package service

import (
	"context"
	"timkerjaService/model/web"
)

type PenilaianKinerjaService interface {
	Create(ctx context.Context, req web.PenilaianKinerjaRequest) (web.PenilaianKinerjaResponse, error)
	Update(ctx context.Context, req web.PenilaianKinerjaRequest, id int) (web.PenilaianKinerjaResponse, error)
}
