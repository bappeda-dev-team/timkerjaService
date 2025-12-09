package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type PenilaianKinerjaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, penilaian domain.PenilaianKinerja) (domain.PenilaianKinerja, error)
	Update(ctx context.Context, tx *sql.Tx, penilaian domain.PenilaianKinerja, id int) (domain.PenilaianKinerja, error)
	ExistById(ctx context.Context, tx *sql.Tx, id int) (bool, error)
	FindByTahunBulan(ctx context.Context, tx *sql.Tx, tahun int, bulan int) ([]domain.LaporanPenilaian, error)
}
