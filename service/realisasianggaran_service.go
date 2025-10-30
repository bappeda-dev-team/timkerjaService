package service

import (
	"context"
	"timkerjaService/model/web"
)

type RealisasiAnggaranService interface {
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.RealisasiAnggaranResponse, error)
	FindAll(ctx context.Context, kodeSubkegiatan string, bulan string, tahun string) ([]web.RealisasiAnggaranResponse, error)
	Upsert(ctx context.Context, req web.RealisasiAnggaranCreateRequest) (web.RealisasiAnggaranResponse, error)
}
