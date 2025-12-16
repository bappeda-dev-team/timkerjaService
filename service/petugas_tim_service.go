package service

import (
	"context"
	"timkerjaService/model/web"
)

type PetugasTimService interface {
	Create(ctx context.Context, petugasTimReq web.PetugasTimCreateRequest) (web.PetugasTimResponse, error)
	Delete(ctx context.Context, idPetugasTim int) error
}
