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
			Id:         timKerjaDomain.Id,
			KodeTim:    timKerjaDomain.KodeTim,
			NamaTim:    timKerjaDomain.NamaTim,
			Keterangan: timKerjaDomain.Keterangan,
			Tahun:      timKerjaDomain.Tahun,
			IsActive:   timKerjaDomain.IsActive,
			CreatedAt:  timKerjaDomain.CreatedAt,
			UpdatedAt:  timKerjaDomain.UpdatedAt,
		}
	}
	return timKerjaResponses
}
