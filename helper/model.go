package helper

import (
	"context"
	"log"
	"sync"
	"timkerjaService/internal"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
)

func ToJabatanTimResponses(jabatanTimDomains []domain.JabatanTim) []web.JabatanTimResponse {
	jabatanTimResponses := make([]web.JabatanTimResponse, len(jabatanTimDomains))
	for i, jabatanTimDomain := range jabatanTimDomains {
		jabatanTimResponses[i] = web.JabatanTimResponse{
			Id:           jabatanTimDomain.Id,
			NamaJabatan:  jabatanTimDomain.NamaJabatan,
			LevelJabatan: jabatanTimDomain.LevelJabatan,
			CreatedAt:    jabatanTimDomain.CreatedAt,
			UpdatedAt:    jabatanTimDomain.UpdatedAt,
		}
	}
	return jabatanTimResponses
}

func ToSusunanTimResponses(susunanTimDomains []domain.SusunanTim) []web.SusunanTimResponse {
	susunanTimResponses := make([]web.SusunanTimResponse, len(susunanTimDomains))
	for i, susunanTimDomain := range susunanTimDomains {
		susunanTimResponses[i] = web.SusunanTimResponse{
			Id:             susunanTimDomain.Id,
			KodeTim:        susunanTimDomain.KodeTim,
			PegawaiId:      susunanTimDomain.PegawaiId,
			NamaPegawai:    susunanTimDomain.NamaPegawai,
			NamaJabatanTim: susunanTimDomain.NamaJabatanTim,
			IsActive:       susunanTimDomain.IsActive,
			Keterangan:     susunanTimDomain.Keterangan,
			CreatedAt:      susunanTimDomain.CreatedAt,
			UpdatedAt:      susunanTimDomain.UpdatedAt,
		}
	}
	return susunanTimResponses
}

func ToTimKerjaResponses(timKerjaDomains []domain.TimKerja) []web.TimKerjaResponse {
	timKerjaResponses := make([]web.TimKerjaResponse, len(timKerjaDomains))
	for i, timKerjaDomain := range timKerjaDomains {
		timKerjaResponses[i] = web.TimKerjaResponse{
			Id:            timKerjaDomain.Id,
			KodeTim:       timKerjaDomain.KodeTim,
			NamaTim:       timKerjaDomain.NamaTim,
			Keterangan:    timKerjaDomain.Keterangan,
			Tahun:         timKerjaDomain.Tahun,
			IsActive:      timKerjaDomain.IsActive,
			IsSekretariat: timKerjaDomain.IsSekretariat,
			CreatedAt:     timKerjaDomain.CreatedAt,
			UpdatedAt:     timKerjaDomain.UpdatedAt,
		}
	}
	return timKerjaResponses
}

func ToProgramUnggulanResponses(programUnggulans []domain.ProgramUnggulanTimKerja) []web.ProgramUnggulanTimKerjaResponse {
	programUnggulanReponses := make([]web.ProgramUnggulanTimKerjaResponse, len(programUnggulans))
	for i, programUnggulanDomain := range programUnggulans {
		programUnggulanReponses[i] = web.ProgramUnggulanTimKerjaResponse{
			Id:                programUnggulanDomain.Id,
			KodeTim:           programUnggulanDomain.KodeTim,
			IdProgramUnggulan: programUnggulanDomain.IdProgramUnggulan,
			ProgramUnggulan:   programUnggulanDomain.NamaProgramUnggulan,
			Tahun:             programUnggulanDomain.Tahun,
			KodeOpd:           programUnggulanDomain.KodeOpd,
		}
	}
	return programUnggulanReponses
}

func ToRencanaKinerjaTimResponses(rencanaKinerjas []domain.RencanaKinerjaTimKerja) []web.RencanaKinerjaTimKerjaResponse {
	rencanaKinerjaReponses := make([]web.RencanaKinerjaTimKerjaResponse, len(rencanaKinerjas))
	for i, rencanaKinerjaDomain := range rencanaKinerjas {
		rencanaKinerjaReponses[i] = web.RencanaKinerjaTimKerjaResponse{
			Id:               rencanaKinerjaDomain.Id,
			KodeTim:          rencanaKinerjaDomain.KodeTim,
			IdRencanaKinerja: rencanaKinerjaDomain.IdRencanaKinerja,
			IdPegawai:        rencanaKinerjaDomain.IdPegawai,
			RencanaKinerja:   rencanaKinerjaDomain.RencanaKinerja,
			Tahun:            rencanaKinerjaDomain.Tahun,
			KodeOpd:          rencanaKinerjaDomain.KodeOpd,
		}
	}
	return rencanaKinerjaReponses
}

func ToRealisasiPokinResponses(realisasis []domain.RealisasiPokin) []web.RealisasiResponse {
	realisasiResponses := make([]web.RealisasiResponse, len(realisasis))
	for i, realisasi := range realisasis {
		realisasiResponses[i] = web.RealisasiResponse{
			Id:               realisasi.Id,
			IdPokin:          realisasi.IdPokin,
			KodeTim:          realisasi.KodeTim,
			JenisPohon:       realisasi.JenisPohon,
			JenisItem:        realisasi.JenisItem,
			KodeItem:         realisasi.KodeItem,
			NamaItem:         realisasi.NamaItem,
			Pagu:             realisasi.Pagu,
			Realisasi:        realisasi.Realisasi,
			FaktorPendorong:  realisasi.FaktorPendorong,
			FaktorPenghambat: realisasi.FaktorPenghambat,
			Rtl:              realisasi.Rtl,
			UrlBuktiDukung:   realisasi.UrlBuktiDukung,
			Tahun:            realisasi.Tahun,
			KodeOpd:          realisasi.KodeOpd,
		}
	}
	return realisasiResponses
}

// internal
func MergeRencanaKinerjaWithRekinParallel(
	ctx context.Context,
	rencanas []domain.RencanaKinerjaTimKerja,
	client *internal.PerencanaanClient,
	maxConcurrency int,
) []web.RencanaKinerjaTimKerjaResponse {
	responses := make([]web.RencanaKinerjaTimKerjaResponse, len(rencanas))
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for i, r := range rencanas {
		wg.Add(1)
		go func(i int, r domain.RencanaKinerjaTimKerja) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			resp := web.RencanaKinerjaTimKerjaResponse{
				Id:               r.Id,
				KodeTim:          r.KodeTim,
				IdRencanaKinerja: r.IdRencanaKinerja,
				IdPegawai:        r.IdPegawai,
				Tahun:            r.Tahun,
				KodeOpd:          r.KodeOpd,
				PaguAnggaran:     0,
			}

			// === Fetch API eksternal ===
			dataRincian, err := client.GetDataRincianKerja(ctx, r.IdRencanaKinerja, r.IdPegawai)
			if err != nil {
				log.Printf("⚠️ gagal fetch rincian kerja [%v]: %v", r.IdRencanaKinerja, err)
				resp.RencanaKinerja = "NOT_CHECKED"
				responses[i] = resp
				return
			}
			if dataRincian == nil {
				resp.RencanaKinerja = "NOT_FOUND"
				responses[i] = resp
				return
			}

			// === Map hasil dari API ===
			resp.RencanaKinerja = dataRincian.RencanaKinerja.NamaRencanaKinerja
			resp.NamaPegawai = dataRincian.RencanaKinerja.NamaPegawai
			resp.PaguAnggaran = dataRincian.RencanaKinerja.Pagu
			resp.RencanaAksi = dataRincian.RencanaAksi.RencanaAksi

			// tambah disini kebutuhan tambahan
			//
			resp.Indikator = dataRincian.RencanaKinerja.Indikator

			resp.SubKegiatan = dataRincian.SubKegiatan

			responses[i] = resp
		}(i, r)
	}

	wg.Wait()
	return responses
}

func MergeProgramUnggulanFromApiParallel(
	ctx context.Context,
	programUnggulans []domain.ProgramUnggulanTimKerja,
	client *internal.PerencanaanClient,
	maxConcurrency int,
) []web.ProgramUnggulanTimKerjaResponse {
	responses := make([]web.ProgramUnggulanTimKerjaResponse, len(programUnggulans))
	var idBatch []int
	for _, r := range programUnggulans {
		if r.IdProgramUnggulan != 0 {
			idBatch = append(idBatch, r.IdProgramUnggulan)
		}
	}
	// nama program unggulans
	batchResp, err := client.GetNamaProgramUnggulanBatch(ctx, idBatch)
	if err != nil {
		log.Printf("gagal fetch batch program unggulan: %v", err)
	}

	programUnggulanMap := make(map[int]internal.ProgramUnggulanResponse)
	for _, item := range batchResp {
		programUnggulanMap[item.Id] = internal.ProgramUnggulanResponse{
			Id:                  item.Id,
			RencanaImplementasi: item.RencanaImplementasi,
		}
	}

	// rincian program unggulans
	var kodeBatch []string
	for _, r := range programUnggulans {
		if r.IdProgramUnggulan != 0 {
			kodeBatch = append(kodeBatch, r.KodeProgramUnggulan)
		}
	}
	rincianBatchResp, err := client.GetRincianProgramUnggulans(ctx, kodeBatch)
	if err != nil {
		log.Printf("gagal fetch batch program unggulan: %v", err)
	}

	rincianMap := make(map[string][]internal.TaggingPohonKinerjaItem)
	for _, item := range rincianBatchResp {
		kode := item.KodeProgramUnggulan
		rincianMap[kode] = append(rincianMap[kode], item)
	}

	// gabung
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for i, r := range programUnggulans {
		wg.Add(1)
		go func(i int, r domain.ProgramUnggulanTimKerja) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			resp := web.ProgramUnggulanTimKerjaResponse{
				Id:                  r.Id,
				KodeTim:             r.KodeTim,
				IdProgramUnggulan:   r.IdProgramUnggulan,
				KodeProgramUnggulan: r.KodeProgramUnggulan,
				Tahun:               r.Tahun,
				KodeOpd:             r.KodeOpd,
			}
			// === Ambil data dari hasil batch ===
			if pu, ok := programUnggulanMap[r.IdProgramUnggulan]; ok {
				resp.ProgramUnggulan = pu.RencanaImplementasi
			} else {
				resp.ProgramUnggulan = "-"
			}

			// === Ambil rincian batch ===

			if items, ok := rincianMap[r.KodeProgramUnggulan]; ok {
				resp.Pokin = items
			} else {
				resp.Pokin = []internal.TaggingPohonKinerjaItem{}
			}

			responses[i] = resp
		}(i, r)
	}

	wg.Wait()
	return responses
}

func ToRealisasiAnggaranResponse(d domain.RealisasiAnggaran) web.RealisasiAnggaranResponse {
	return web.RealisasiAnggaranResponse{
		Id:                d.Id,
		KodeTim:           d.KodeTim,
		IdProgramUnggulan: d.IdProgramUnggulan,
		IdPohon:           d.IdPohon,
		IdRencanaKinerja:  d.IdRencanaKinerja,
		KodeSubkegiatan:   d.KodeSubkegiatan,
		RealisasiAnggaran: d.RealisasiAnggaran,
		KodeOpd:           d.KodeOpd,
		RencanaAksi:       d.RencanaAksi,
		FaktorPendorong:   d.FaktorPendorong,
		FaktorPenghambat:  d.FaktorPenghambat,
		RekomendasiTl:     d.RekomendasiTl,
		BuktiDukung:       d.BuktiDukung,
		Bulan:             d.Bulan,
		Tahun:             d.Tahun,
		CreatedAt:         d.CreatedAt,
		UpdatedAt:         d.UpdatedAt,
	}
}

func ToRealisasiAnggaranResponses(ds []domain.RealisasiAnggaran) []web.RealisasiAnggaranResponse {
	out := make([]web.RealisasiAnggaranResponse, len(ds))
	for i, d := range ds {
		out[i] = ToRealisasiAnggaranResponse(d)
	}
	return out
}
