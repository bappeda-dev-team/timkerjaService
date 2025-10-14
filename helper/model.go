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

			// === Fetch API eksternal ===

			dataRincian, err := client.GetProgramUnggulan(ctx, r.KodeProgramUnggulan)
			if err != nil {
				log.Printf("⚠️ gagal fetch rincian program unggulan [%v]: %v", r.KodeProgramUnggulan, err)
				resp.ProgramUnggulan = "-"
				responses[i] = resp
				return
			}

			if len(dataRincian.Data) == 0 {
				resp.ProgramUnggulan = "-"
				responses[i] = resp
				return
			}

            programUnggulan, err := client.GetNamaProgramUnggulan(ctx, r.IdProgramUnggulan)
			if err != nil {
				log.Printf("⚠️ gagal fetch program unggulan [%v]: %v", r.IdProgramUnggulan, err)
				resp.ProgramUnggulan = "-"
				responses[i] = resp
				return
			}

			// === Gunakan data pertama sebagai program unggulan utama ===
			resp.ProgramUnggulan = programUnggulan.RencanaImplementasi

			// === Simpan seluruh elemen data API ke dalam Pokin ===
			resp.Pokin = dataRincian.Data

			responses[i] = resp
		}(i, r)
	}

	wg.Wait()
	return responses
}
