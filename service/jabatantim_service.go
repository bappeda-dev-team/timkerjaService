package service

import (
	"context"
	"timkerjaService/model/web"
)

type JabatanTimService interface {
	Create(ctx context.Context, jabatanTim web.JabatanTimCreateRequest) (web.JabatanTimResponse, error)
	Update(ctx context.Context, jabatanTim web.JabatanTimUpdateRequest) (web.JabatanTimResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.JabatanTimResponse, error)
	FindAll(ctx context.Context) ([]web.JabatanTimResponse, error)
}
