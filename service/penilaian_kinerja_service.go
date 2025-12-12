package service

import (
	"context"
	"timkerjaService/model/web"
)

type PenilaianKinerjaService interface {
	All(ctx context.Context, tahun int, bulan int) ([]web.LaporanPenilaianKinerjaResponse, error)
	Create(ctx context.Context, req web.PenilaianKinerjaRequest) (web.PenilaianKinerjaResponse, error)
	Update(ctx context.Context, req web.PenilaianKinerjaRequest, id int) (web.PenilaianKinerjaResponse, error)
	// PenilaianGrouped with TPP
	TppPegawaiAll(ctx context.Context, tahun int, bulan int) ([]web.LaporanPenilaianKinerjaResponse, error)
}
