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
	FindAll(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error)
	FindAllWithSusunan(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	FindAllWithSusunanByBulanTahun(ctx context.Context, tx *sql.Tx, bulan int, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	AddProgramUnggulan(ctx context.Context, tx *sql.Tx, programUnggulan domain.ProgramUnggulanTimKerja) (domain.ProgramUnggulanTimKerja, error)
	FindProgramUnggulanByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string, tahun int) ([]domain.ProgramUnggulanTimKerja, error)
	FindAllTimNonSekretariatWithSusunan(ctx context.Context, tx *sql.Tx, bulan int, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	FindAllTimNonSekretariat(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error)
	FindAllTimSekretariatWithSusunan(ctx context.Context, tx *sql.Tx, bulan int, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error)
	FindAllTimSekretariat(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error)
	DeleteProgramUnggulan(ctx context.Context, tx *sql.Tx, id int, kodeTim string) error
	AddRencanaKinerja(ctx context.Context, tx *sql.Tx, rencanaKinerja domain.RencanaKinerjaTimKerja) (domain.RencanaKinerjaTimKerja, error)
	FindRencanaKinerjaByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string, tahun int) ([]domain.RencanaKinerjaTimKerja, error)
	DeleteRencanaKinerja(ctx context.Context, tx *sql.Tx, id int, kodeTim string) error
	SaveRealisasiPokin(ctx context.Context, tx *sql.Tx, realisasi domain.RealisasiPokin) (domain.RealisasiPokin, error)
	UpdateRealisasiPokin(ctx context.Context, tx *sql.Tx, realisasi domain.RealisasiPokin) (domain.RealisasiPokin, error)
	FindAllRealisasiPokinByKodeItemTahun(ctx context.Context, tx *sql.Tx, kodeTim string, tahun string) ([]domain.RealisasiPokin, error)
	FindRealisasiByKodeTimAndPohonIDs(ctx context.Context, tx *sql.Tx, kodeTim string, bulan int, tahun int, pohonIDs []int) (map[int]domain.RealisasiAnggaranRecord, error)
	FindRealisasiByKodeTimAndRekinSekretariatIds(ctx context.Context, tx *sql.Tx, kodeTim string, bulan int, tahun int, rekinSekretIds []int) (map[int]domain.RealisasiAnggaranRecord, error)
}
