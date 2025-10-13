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
	FindAll(ctx context.Context) ([]web.TimKerjaResponse, error)
	FindAllTm(ctx context.Context) ([]web.TimKerjaDetailResponse, error)
	AddProgramUnggulan(ctx context.Context, timKerja web.ProgramUnggulanTimKerjaRequest) (web.ProgramUnggulanTimKerjaResponse, error)
	FindAllProgramUnggulanTim(ctx context.Context, kodeTim string) ([]web.ProgramUnggulanTimKerjaResponse, error)
	FindAllTimNonSekretariat(ctx context.Context) ([]web.TimKerjaDetailResponse, error)
	FindAllTimSekretariat(ctx context.Context) ([]web.TimKerjaDetailResponse, error)
	DeleteProgramUnggulan(ctx context.Context, id int, kodeTim string) error
	AddRencanaKinerja(ctx context.Context, timkerja web.RencanaKinerjaRequest) (web.RencanaKinerjaTimKerjaResponse, error)
	FindAllRencanaKinerjaTim(ctx context.Context, kodeTim string) ([]web.RencanaKinerjaTimKerjaResponse, error)
	DeleteRencanaKinerjaTim(ctx context.Context, id int, kodeTim string) error
}
