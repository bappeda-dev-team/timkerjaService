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
	bulan int,
	tahun int,
	kodeOpd string,
) ([]web.LaporanPenilaianKinerjaResponse, error) {

	var (
		responses = make([]web.LaporanPenilaianKinerjaResponse, len(penilaianKinerja))
		sem       = make(chan struct{}, maxConcurrency)
		wg        sync.WaitGroup
	)
	// ============
	// 0 SET NILAI OPD = sama semua
	// SET NILAI TIM = sama semua by tim
	// ===========

	// get nilai opd (KINERJA_BAPPEDA)
	// get Maximum nilai
	kinerjaOpd := 0
	kinerjaPerTim := make(map[string]int)

	for _, laporan := range penilaianKinerja {
		for _, pp := range laporan.Penilaians {
			if pp.JenisNilai == "KINERJA_BAPPEDA" {
				if pp.NilaiKinerja > kinerjaOpd {
					kinerjaOpd = pp.NilaiKinerja
				}
			}
			if pp.JenisNilai == "KINERJA_TIM" {
				if pp.NilaiKinerja > kinerjaPerTim[pp.KodeTim] {
					kinerjaPerTim[pp.KodeTim] = pp.NilaiKinerja
				}
			}
		}
	}

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
				case "KINERJA_KEHADIRAN":
					item.KinerjaKehadiran = p.NilaiKinerja
				}
			}

			// Convert map → slice dan hitung nilai akhir
			grouped := make([]web.PenilaianGroupedResponse, 0, len(groupMap))
			for _, v := range groupMap {
				v.KinerjaBappeda = kinerjaOpd
				if val, ok := kinerjaPerTim[v.KodeTim]; ok {
					v.KinerjaTim = val
				}
				v.NilaiAkhir = hitungNilaiAkhir(v.PenilaianGroupedBase)
				grouped = append(grouped, *v)
			}

			responses[i] = web.LaporanPenilaianKinerjaResponse{
				NamaTim:           laporan.NamaTim,
				KodeTim:           laporan.KodeTim,
				IsSekretariat:     laporan.IsSekretariat,
				Keterangan:        laporan.Keterangan,
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

	log.Printf("KODE OPD: %s", kodeOpd)
	// Ambil detail pegawai batch
	detailPegawais, err := client.GetDetailPegawaiBatch(ctx, listIdPegawais, bulan, tahun, kodeOpd)
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

			// BPJS
			item.Tpp.PotonganBPJS1 = dp.Bpjs1
			item.Tpp.PotonganBPJS4 = dp.Bpjs4
		}
	}

	return responses, nil
}

func hitungNilaiAkhir(item web.PenilaianGroupedBase) int {

	// yang 0 tidak masuk slice
	// xs := []float64{}
	// if item.KinerjaBappeda > 0 {
	// 	xs = append(xs, float64(item.KinerjaBappeda))
	// }
	// if item.KinerjaTim > 0 {
	// 	xs = append(xs, float64(item.KinerjaTim))
	// }
	// if item.KinerjaPerson > 0 {
	// 	xs = append(xs, float64(item.KinerjaPerson))
	// }
	// if len(xs) == 0 {
	// 	return 0
	// }
	// yang 0 tetap masuk ke slice
	xs := []float64{
		float64(item.KinerjaBappeda),
		float64(item.KinerjaTim),
		float64(item.KinerjaPerson),
	}

	avgNilai := average(xs)

	// percentage
	hasilAkhir := avgNilai * float64(item.KinerjaKehadiran) / 100

	return int(math.Ceil(hasilAkhir))
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
			Keterangan:        lap.Keterangan,
			PenilaianKinerjas: make([]web.PenilaianGroupedResponse, len(lap.PenilaianKinerjas)),
		}

		for j, p := range lap.PenilaianKinerjas {

			// copy original including TPP from API
			item := p // copy struct

			// Pastikan extension TPP tidak nil
			if item.Tpp == nil {
				item.Tpp = &web.PenilaianTppExtension{}
			}

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
	// tpp.PotonganBPJS = float64(tpp.JumlahKotor) * tpp.PotonganBPJS
	// bpjs 1
	potonganBpjs1 := float64(tpp.JumlahKotor) * tpp.PotonganBPJS1
	// bpjs 4
	potonganBpjs4 := float64(tpp.JumlahKotor) * tpp.PotonganBPJS4

	// bpjsAmount := int(tpp.PotonganBPJS)
	bpjs1Amount := int(potonganBpjs1)
	bpjs4Amount := int(potonganBpjs4)

	tpp.Bpjs1 = bpjs1Amount
	tpp.Bpjs4 = bpjs4Amount

	// 4. Jumlah Bersih
	tpp.JumlahBersih = tpp.JumlahKotor - tpp.JumlahPajak - bpjs1Amount - bpjs4Amount
}

func ConvertToAllLaporan(
	src []web.LaporanPenilaianKinerjaResponse,
) []web.PenilaianGroupedResponse {

	// slice kosong, tapi siap di-append
	out := make([]web.PenilaianGroupedResponse, 0)

	for _, lap := range src {
		for _, p := range lap.PenilaianKinerjas {

			// copy struct (aman, karena kita mau hasil terpisah)
			item := p

			// pastikan TPP tidak nil
			if item.Tpp == nil {
				item.Tpp = &web.PenilaianTppExtension{}
			}

			// inject context dari parent
			item.NamaTim = lap.NamaTim

			// konfigurasi tambahan
			// item.Tpp.PotonganBPJS = 0.01

			// hitung TPP (pakai pointer ke item)
			HitungTPP(&item)

			// validasi hasil
			// ADA PERUBAHAN, DULU HANYA YANG PUNYA TPP YANG MASUK
			// SEKARANG DI BKAD DAN PERUBAHAN TERBARU
			// SEMUA PEGAWAI MUNCUL
			// TODO FIX LOGIC TPP JUMLAH KOTOR
			// if item.NamaPegawai != "" && item.Tpp.JumlahKotor != 0 {
			if item.NamaPegawai != "" {
				out = append(out, item)
			}
		}
	}

	return out
}
