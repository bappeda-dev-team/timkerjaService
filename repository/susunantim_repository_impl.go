package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"timkerjaService/model/domain"
)

type SusunanTimRepositoryImpl struct {
}

func NewSusunanTimRepositoryImpl() *SusunanTimRepositoryImpl {
	return &SusunanTimRepositoryImpl{}
}

func (repository *SusunanTimRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error) {
	query := "INSERT INTO susunan_tim (kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim, is_active, keterangan, bulan, tahun) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, susunanTim.KodeTim, susunanTim.PegawaiId, susunanTim.NamaPegawai, susunanTim.IdJabatanTim, susunanTim.NamaJabatanTim, susunanTim.IsActive, susunanTim.Keterangan, susunanTim.Bulan, susunanTim.Tahun)
	if err != nil {
		return domain.SusunanTim{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.SusunanTim{}, err
	}

	susunanTim.Id = int(id)

	return susunanTim, nil
}

func (repository *SusunanTimRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error) {
	query := "UPDATE susunan_tim SET kode_tim = ?, pegawai_id = ?, nama_pegawai = ?, jabatan_tim_id = ?, nama_jabatan_tim = ?, is_active = ?, keterangan = ?, bulan = ?, tahun = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, susunanTim.KodeTim, susunanTim.PegawaiId, susunanTim.NamaPegawai, susunanTim.IdJabatanTim, susunanTim.NamaJabatanTim, susunanTim.IsActive, susunanTim.Keterangan, susunanTim.Bulan, susunanTim.Tahun, susunanTim.Id)
	if err != nil {
		return domain.SusunanTim{}, err
	}

	return susunanTim, nil
}

func (repository *SusunanTimRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM susunan_tim WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *SusunanTimRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim, is_active, keterangan FROM susunan_tim WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.SusunanTim{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return domain.SusunanTim{}, err
		}
		return susunanTim, nil
	}

	return domain.SusunanTim{}, errors.New("susunan tim not found")
}

func (repository *SusunanTimRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim , is_active, keterangan FROM susunan_tim ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.SusunanTim{}, err
	}
	defer rows.Close()

	var susunanTimList []domain.SusunanTim
	for rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return []domain.SusunanTim{}, err
		}

		susunanTimList = append(susunanTimList, susunanTim)
	}

	return susunanTimList, nil
}

func (repository *SusunanTimRepositoryImpl) FindAllByBulanTahun(ctx context.Context, tx *sql.Tx, bulan int, tahun int) ([]domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim , is_active, keterangan FROM susunan_tim WHERE bulan = ? AND tahun = ? ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, bulan, tahun)
	if err != nil {
		return []domain.SusunanTim{}, err
	}
	defer rows.Close()

	var susunanTimList []domain.SusunanTim
	for rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return []domain.SusunanTim{}, err
		}

		susunanTimList = append(susunanTimList, susunanTim)
	}

	return susunanTimList, nil
}

func (repository *SusunanTimRepositoryImpl) FindByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim , is_active, keterangan FROM susunan_tim WHERE kode_tim = ? ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeTim)
	if err != nil {
		return []domain.SusunanTim{}, err
	}
	defer rows.Close()

	var susunanTimList []domain.SusunanTim
	for rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return []domain.SusunanTim{}, err
		}

		susunanTimList = append(susunanTimList, susunanTim)
	}

	return susunanTimList, nil
}

func (repository *SusunanTimRepositoryImpl) FindByKodeTimBulanTahun(ctx context.Context, tx *sql.Tx, kodeTim string, bulan int, tahun int) ([]domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim , is_active, keterangan, bulan, tahun FROM susunan_tim WHERE kode_tim = ? AND bulan = ? AND tahun = ? ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, kodeTim, bulan, tahun)
	if err != nil {
		return []domain.SusunanTim{}, err
	}
	defer rows.Close()

	var susunanTimList []domain.SusunanTim
	for rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan, &susunanTim.Bulan, &susunanTim.Tahun)
		if err != nil {
			return []domain.SusunanTim{}, err
		}

		susunanTimList = append(susunanTimList, susunanTim)
	}

	return susunanTimList, nil
}

func (repository *SusunanTimRepositoryImpl) FindByIdPegawai(ctx context.Context, tx *sql.Tx, idPegawai string) (domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_pegawai, jabatan_tim_id, nama_jabatan_tim, is_active, keterangan FROM susunan_tim WHERE pegawai_id = ?"
	rows, err := tx.QueryContext(ctx, query, idPegawai)
	if err != nil {
		return domain.SusunanTim{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaPegawai, &susunanTim.IdJabatanTim, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return domain.SusunanTim{}, err
		}
		return susunanTim, nil
	}

	return domain.SusunanTim{}, errors.New("Pegawai tidak ditemukan")
}

func (repository *SusunanTimRepositoryImpl) SaveAll(ctx context.Context, tx *sql.Tx, susunanTims []domain.SusunanTim) error {
	if len(susunanTims) == 0 {
		return errors.New("Susunan Tim tidak boleh kosong")
	}

	var (
		valueStrings []string
		valueArgs    []any
	)

	for _, v := range susunanTims {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs,
			v.KodeTim,
			v.Bulan,
			v.Tahun,
			v.PegawaiId,
			v.NamaPegawai,
			v.IdJabatanTim,
			v.NamaJabatanTim,
			v.IsActive,
			v.Keterangan,
		)
	}
	query := `
		INSERT INTO susunan_tim (
			kode_tim, bulan, tahun,
			pegawai_id, nama_pegawai,
			jabatan_tim_id, nama_jabatan_tim,
			is_active, keterangan
		)
		VALUES ` + strings.Join(valueStrings, ",")

	_, err := tx.ExecContext(ctx, query, valueArgs...)
	return err
}

func (repository *SusunanTimRepositoryImpl) ExistsByKodeTimBulanTahun(
	ctx context.Context,
	tx *sql.Tx,
	kodeTim string,
	bulan int,
	tahun int,
) (bool, error) {

	query := `
		SELECT 1
		FROM susunan_tim
		WHERE kode_tim = ?
		  AND bulan = ?
		  AND tahun = ?
		LIMIT 1
	`

	var dummy int
	err := tx.QueryRowContext(ctx, query, kodeTim, bulan, tahun).Scan(&dummy)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
