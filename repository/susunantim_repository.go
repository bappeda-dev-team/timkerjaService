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
	FindAllByBulanTahun(ctx context.Context, tx *sql.Tx, bulan int, tahun int) ([]domain.SusunanTim, error)
	FindByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.SusunanTim, error)
	FindByIdPegawai(ctx context.Context, tx *sql.Tx, idPegawai string) (domain.SusunanTim, error)
	FindByKodeTimBulanTahun(ctx context.Context, tx *sql.Tx, kodeTim string, bulan int, tahun int) ([]domain.SusunanTim, error)
	SaveAll(ctx context.Context, tx *sql.Tx, susunanTims []domain.SusunanTim) error
	ExistsByKodeTimBulanTahun(ctx context.Context, tx *sql.Tx, kodeTim string, bulan int, tahun int) (bool, error)
}
