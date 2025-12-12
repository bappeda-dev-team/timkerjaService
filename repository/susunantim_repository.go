package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type SusunanTimRepository interface {
	Create(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error)
	Update(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.SusunanTim, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.SusunanTim, error)
	FindByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.SusunanTim, error)
}
