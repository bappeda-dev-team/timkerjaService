package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type SusunanTimRepositoryImpl struct {
}

func NewSusunanTimRepositoryImpl() *SusunanTimRepositoryImpl {
	return &SusunanTimRepositoryImpl{}
}

func (repository *SusunanTimRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error) {
	query := "INSERT INTO susunan_tim (kode_tim, pegawai_id, nama_jabatan_tim, is_active, keterangan) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, susunanTim.KodeTim, susunanTim.PegawaiId, susunanTim.NamaJabatanTim, susunanTim.IsActive, susunanTim.Keterangan)
	if err != nil {
		return domain.SusunanTim{}, err
	}

	return susunanTim, nil
}

func (repository *SusunanTimRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, susunanTim domain.SusunanTim) (domain.SusunanTim, error) {
	query := "UPDATE susunan_tim SET kode_tim = ?, pegawai_id = ?, nama_jabatan_tim = ?, is_active = ?, keterangan = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, susunanTim.KodeTim, susunanTim.PegawaiId, susunanTim.NamaJabatanTim, susunanTim.IsActive, susunanTim.Keterangan, susunanTim.Id)
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
	query := "SELECT id, kode_tim, pegawai_id, nama_jabatan_tim, is_active, keterangan FROM susunan_tim WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.SusunanTim{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return domain.SusunanTim{}, err
		}
		return susunanTim, nil
	}

	return domain.SusunanTim{}, nil
}

func (repository *SusunanTimRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.SusunanTim, error) {
	query := "SELECT id, kode_tim, pegawai_id, nama_jabatan_tim , is_active, keterangan FROM susunan_tim ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.SusunanTim{}, err
	}
	defer rows.Close()

	var susunanTimList []domain.SusunanTim
	for rows.Next() {
		var susunanTim domain.SusunanTim
		err := rows.Scan(&susunanTim.Id, &susunanTim.KodeTim, &susunanTim.PegawaiId, &susunanTim.NamaJabatanTim, &susunanTim.IsActive, &susunanTim.Keterangan)
		if err != nil {
			return []domain.SusunanTim{}, err
		}

		susunanTimList = append(susunanTimList, susunanTim)
	}

	return susunanTimList, nil
}
