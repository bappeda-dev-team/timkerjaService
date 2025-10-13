package helper

import (
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
			Id:              rencanaKinerjaDomain.Id,
			KodeTim:         rencanaKinerjaDomain.KodeTim,
			IdRencanKinerja: rencanaKinerjaDomain.IdRencanaKinerja,
			IdPegawai:       rencanaKinerjaDomain.IdPegawai,
			RencanaKinerja:  rencanaKinerjaDomain.RencanaKinerja,
			Tahun:           rencanaKinerjaDomain.Tahun,
			KodeOpd:         rencanaKinerjaDomain.KodeOpd,
		}
	}
	return rencanaKinerjaReponses
}
