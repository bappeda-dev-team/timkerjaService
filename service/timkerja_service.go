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
}
