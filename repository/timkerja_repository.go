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
	AddProgramUnggulan(ctx context.Context, tx *sql.Tx, programUnggulan domain.ProgramUnggulanTimKerja) (domain.ProgramUnggulanTimKerja, error)
	FindProgramUnggulanByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.ProgramUnggulanTimKerja, error)
	FindAllTimNonSekretariatWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	FindAllTimNonSekretariat(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error)
	FindAllTimSekretariatWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	FindAllTimSekretariat(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error)
}
