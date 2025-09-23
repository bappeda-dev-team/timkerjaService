package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type JabatanTimRepository interface {
	Create(ctx context.Context, tx *sql.Tx, jabatanTim domain.JabatanTim) (domain.JabatanTim, error)
	Update(ctx context.Context, tx *sql.Tx, jabatanTim domain.JabatanTim) (domain.JabatanTim, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JabatanTim, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.JabatanTim, error)
}
