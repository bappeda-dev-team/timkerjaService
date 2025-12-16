package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type PetugasTimRepository interface {
	Create(ctx context.Context, tx *sql.Tx, petugasTimDomain domain.PetugasTim) (domain.PetugasTim, error)
	Delete(ctx context.Context, tx *sql.Tx, idPetugasTim int) error
	FindAllByIdProgramUnggulans(ctx context.Context, tx *sql.Tx, idProgramUnggulans []int) ([]domain.PetugasTim, error)
}
