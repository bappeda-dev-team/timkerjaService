package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type TimKerjaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error)
	Update(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TimKerja, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error)
	FindAllWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
}
