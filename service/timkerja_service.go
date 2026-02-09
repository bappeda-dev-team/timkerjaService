package service

import (
	"context"
	"timkerjaService/model/web"
)

type TimKerjaService interface {
	Create(ctx context.Context, timKerja web.TimKerjaCreateRequest) (web.TimKerjaResponse, error)
	Update(ctx context.Context, timKerja web.TimKerjaUpdateRequest) (web.TimKerjaResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.TimKerjaResponse, error)
	FindByKodeTim(ctx context.Context, kodeTim string) (web.TimKerjaResponse, error)
	FindAll(ctx context.Context, tahun int) ([]web.TimKerjaResponse, error)
	FindAllTm(ctx context.Context, tahun int) ([]web.TimKerjaDetailResponse, error)
	FindAllTmByBulanTahun(ctx context.Context, bulan int, tahun int) ([]web.TimKerjaDetailResponse, error)
	AddProgramUnggulan(ctx context.Context, timKerja web.ProgramUnggulanTimKerjaRequest) (web.ProgramUnggulanTimKerjaResponse, error)
	// TODO: ubah tahun ke string
	FindAllProgramUnggulanTim(ctx context.Context, kodeTim string, bulan int, tahun int) ([]web.ProgramUnggulanTimKerjaResponse, error)
	FindAllTimNonSekretariat(ctx context.Context, bulan int, tahun int) ([]web.TimKerjaDetailResponse, error)
	FindAllTimSekretariat(ctx context.Context, bulan int, tahun int) ([]web.TimKerjaDetailResponse, error)
	DeleteProgramUnggulan(ctx context.Context, id int, kodeTim string) error
	AddRencanaKinerja(ctx context.Context, timkerja web.RencanaKinerjaRequest) (web.RencanaKinerjaTimKerjaResponse, error)
	FindAllRencanaKinerjaTim(ctx context.Context, kodeTim string, bulan int, tahun int) ([]web.RencanaKinerjaTimKerjaResponse, error)
	DeleteRencanaKinerjaTim(ctx context.Context, id int, kodeTim string) error
	SaveRealisasiPokin(ctx context.Context, realisasi web.RealisasiRequest) (web.RealisasiResponse, error)
	GetRealisasiPokin(ctx context.Context, kodeItem string, tahun string) ([]web.RealisasiResponse, error)
	FindAllProgramUnggulanOpd(ctx context.Context, kodeOpd string, bulan int, tahun string) ([]web.ProgramUnggulanTimKerjaResponse, error)
}
