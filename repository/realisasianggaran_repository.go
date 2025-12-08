package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type RealisasiAnggaranRepository interface {
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.RealisasiAnggaran, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeSubkegiatan string, kodeTim string, idRencanaKinerja string, bulan string, tahun string) ([]domain.RealisasiAnggaran, error)
	Upsert(ctx context.Context, tx *sql.Tx, ra domain.RealisasiAnggaran) (domain.RealisasiAnggaran, error)
}
