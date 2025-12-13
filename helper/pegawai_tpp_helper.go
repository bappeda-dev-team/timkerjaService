package helper

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"sync"
	"timkerjaService/internal"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
)

func MergePenilaianKinerjaParallel(
	ctx context.Context,
	penilaianKinerja []domain.LaporanPenilaian,
	client *internal.KepegawaianClient,
	maxConcurrency int,
) ([]web.LaporanPenilaianKinerjaResponse, error) {

	var (
		responses = make([]web.LaporanPenilaianKinerjaResponse, len(penilaianKinerja))
		sem       = make(chan struct{}, maxConcurrency)
		wg        sync.WaitGroup
	)

	// ==============================
	// 1) GROUPING JENIS NILAI
	// ==============================
	for i, laporan := range penilaianKinerja {
		wg.Add(1)

		go func(i int, laporan domain.LaporanPenilaian) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			groupMap := make(map[string]*web.PenilaianGroupedResponse)

			// Group per pegawai-per-bulan
			for _, p := range laporan.Penilaians {
				key := p.IdPegawai + "_" + p.Tahun + "_" + strconv.Itoa(p.Bulan)

				item, exists := groupMap[key]
				if !exists {
					item = &web.PenilaianGroupedResponse{
						PenilaianGroupedBase: web.PenilaianGroupedBase{
							IdPegawai:       p.IdPegawai,
							NamaPegawai:     p.NamaPegawai,
							SusunanTimId:    p.SusunanTimId,
							LevelJabatanTim: p.LevelJabatanTim,
							NamaJabatanTim:  p.NamaJabatanTim,

							// Akan diisi dari API eksternal
							Pangkat:      "",
							Golongan:     "",
							JenisJabatan: "",

							KodeTim: p.KodeTim,
							Tahun:   p.Tahun,
							Bulan:   p.Bulan,
						},
						Tpp: &web.PenilaianTppExtension{},
					}
					groupMap[key] = item
				}

				switch p.JenisNilai {
				case "KINERJA_BAPPEDA":
					item.KinerjaBappeda = p.NilaiKinerja
				case "KINERJA_TIM":
					item.KinerjaTim = p.NilaiKinerja
				case "KINERJA_PERSON":
					item.KinerjaPerson = p.NilaiKinerja
				}
			}

			// Convert map → slice dan hitung nilai akhir
			grouped := make([]web.PenilaianGroupedResponse, 0, len(groupMap))
			for _, v := range groupMap {
				v.NilaiAkhir = hitungNilaiAkhir(v.PenilaianGroupedBase)
				grouped = append(grouped, *v)
			}

			responses[i] = web.LaporanPenilaianKinerjaResponse{
				NamaTim:           laporan.NamaTim,
				KodeTim:           laporan.KodeTim,
				IsSekretariat:     laporan.IsSekretariat,
				PenilaianKinerjas: grouped,
			}

		}(i, laporan)
	}

	wg.Wait()

	// ==============================
	// 2) AMBIL DETAIL PEGAWAI (BATCH)
	// ==============================
	idPegawaiSet := map[string]struct{}{}
	for _, resp := range responses {
		for _, p := range resp.PenilaianKinerjas {
			if p.IdPegawai != "" {
				idPegawaiSet[p.IdPegawai] = struct{}{}
			}
		}
	}

	// Siapkan list ID
	listIdPegawais := make([]string, 0, len(idPegawaiSet))
	for id := range idPegawaiSet {
		listIdPegawais = append(listIdPegawais, id)
	}

	// ======================
	// 2.5) SORTING: LevelJabatanTim ASC
	// ======================
	for i := range responses {
		sort.Slice(responses[i].PenilaianKinerjas, func(a, b int) bool {
			A := responses[i].PenilaianKinerjas[a]
			B := responses[i].PenilaianKinerjas[b]

			// level jabatan tim kecil > level besar
			if A.LevelJabatanTim != B.LevelJabatanTim {
				return A.LevelJabatanTim < B.LevelJabatanTim
			}

			return A.SusunanTimId < B.SusunanTimId
		})
	}

	// Ambil detail pegawai batch
	detailPegawais, err := client.GetDetailPegawaiBatch(ctx, listIdPegawais)
	if err != nil {
		log.Printf("ERROR KEPEGAWAIAN HOST: %v\n", err)
		return responses, nil // tetap return meski gagal
	}

	// Jadikan map untuk akses cepat
	dpMap := make(map[string]internal.DetailPegawaiResponse)
	for _, dp := range detailPegawais {
		dpMap[dp.NIP] = dp
	}

	// ==============================
	// 3) MERGE DETAIL PEGAWAI KE RESPONSE
	// ==============================
	for i := range responses {
		for j := range responses[i].PenilaianKinerjas {

			item := &responses[i].PenilaianKinerjas[j]
			dp, ok := dpMap[item.IdPegawai]
			if !ok {
				continue
			}

			// Basic biodata
			item.Pangkat = dp.Pangkat
			item.Golongan = dp.Golongan
			item.JenisJabatan = dp.JenisJabatan

			// TPP extension selalu aman
			if item.Tpp == nil {
				item.Tpp = &web.PenilaianTppExtension{}
			}

			// TPP BASIC — gunakan round, bukan ceil
			item.Tpp.TppBasic = int(math.Round(dp.Tpp))

			// Pajak
			item.Tpp.Pajak = dp.Pajak
		}
	}

	return responses, nil
}

func hitungNilaiAkhir(item web.PenilaianGroupedBase) int {
	xs := []float64{}

	if item.KinerjaBappeda > 0 {
		xs = append(xs, float64(item.KinerjaBappeda))
	}
	if item.KinerjaTim > 0 {
		xs = append(xs, float64(item.KinerjaTim))
	}
	if item.KinerjaPerson > 0 {
		xs = append(xs, float64(item.KinerjaPerson))
	}

	if len(xs) == 0 {
		return 0
	}

	return int(math.Ceil(average(xs)))
}

func average(xs []float64) float64 {
	var total float64
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func ConvertToTppPegawaiResponse(
	src []web.LaporanPenilaianKinerjaResponse,
) []web.LaporanPenilaianKinerjaResponse {

	out := make([]web.LaporanPenilaianKinerjaResponse, len(src))

	for i, lap := range src {
		target := web.LaporanPenilaianKinerjaResponse{
			NamaTim:           lap.NamaTim,
			KodeTim:           lap.KodeTim,
			IsSekretariat:     lap.IsSekretariat,
			PenilaianKinerjas: make([]web.PenilaianGroupedResponse, len(lap.PenilaianKinerjas)),
		}

		for j, p := range lap.PenilaianKinerjas {

			// copy original including TPP from API
			item := p // copy struct

			// Pastikan extension TPP tidak nil
			if item.Tpp == nil {
				item.Tpp = &web.PenilaianTppExtension{}
			}

			// Set konfigurasi tambahan di sini
			item.Tpp.PotonganBPJS = 0.01

			// Hitung TPP dengan pointer, agar perubahan tersimpan
			HitungTPP(&item)

			target.PenilaianKinerjas[j] = item
		}

		out[i] = target
	}

	return out
}

func HitungTPP(p *web.PenilaianGroupedResponse) {
	if p.Tpp == nil {
		p.Tpp = &web.PenilaianTppExtension{}
	}

	tpp := p.Tpp

	tpp.PersentasePenerimaan = fmt.Sprintf("%d%%", p.NilaiAkhir)

	// 1. Hitung TPP Kotor = TppBasic * (NilaiAkhir / 100)
	tpp.JumlahKotor = int(float64(tpp.TppBasic) * (float64(p.NilaiAkhir) / 100.0))

	// 2. Pajak = persen pajak * jumlah kotor
	tpp.JumlahPajak = int(float64(tpp.JumlahKotor) * tpp.Pajak)

	// 3. BPJS = persen BPJS * jumlah kotor
	tpp.PotonganBPJS = float64(tpp.JumlahKotor) * tpp.PotonganBPJS

	bpjsAmount := int(tpp.PotonganBPJS)

	// 4. Jumlah Bersih
	tpp.JumlahBersih = tpp.JumlahKotor - tpp.JumlahPajak - bpjsAmount
}
