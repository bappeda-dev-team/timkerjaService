package service

import (
	"context"
	"timkerjaService/model/web"
)

type SusunanTimService interface {
	Create(ctx context.Context, susunanTim web.SusunanTimCreateRequest) (web.SusunanTimResponse, error)
	Update(ctx context.Context, susunanTim web.SusunanTimUpdateRequest) (web.SusunanTimResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.SusunanTimResponse, error)
	FindAll(ctx context.Context) ([]web.SusunanTimResponse, error)
	FindAllByBulanTahun(ctx context.Context, bulan int, tahun int) ([]web.SusunanTimResponse, error)
	FindByKodeTim(ctx context.Context, kodeTim string) ([]web.SusunanTimResponse, error)
}
