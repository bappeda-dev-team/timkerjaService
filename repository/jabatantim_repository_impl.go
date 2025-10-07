package repository

import (
	"context"
	"database/sql"
	"errors"
	"timkerjaService/model/domain"
)

type JabatanTimRepositoryImpl struct {
}

func NewJabatanTimRepositoryImpl() *JabatanTimRepositoryImpl {
	return &JabatanTimRepositoryImpl{}
}

func (repository *JabatanTimRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, jabatanTim domain.JabatanTim) (domain.JabatanTim, error) {
	query := "INSERT INTO jabatan_tim (nama_jabatan, level_jabatan) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, jabatanTim.NamaJabatan, jabatanTim.LevelJabatan)
	if err != nil {
		return domain.JabatanTim{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.JabatanTim{}, err
	}

	jabatanTim.Id = int(id)

	return jabatanTim, nil
}

func (repository *JabatanTimRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, jabatanTim domain.JabatanTim) (domain.JabatanTim, error) {
	query := "UPDATE jabatan_tim SET nama_jabatan = ?, level_jabatan = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, jabatanTim.NamaJabatan, jabatanTim.LevelJabatan, jabatanTim.Id)
	if err != nil {
		return domain.JabatanTim{}, err
	}

	return jabatanTim, nil
}

func (repository *JabatanTimRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM jabatan_tim WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *JabatanTimRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JabatanTim, error) {
	query := "SELECT id, nama_jabatan, level_jabatan FROM jabatan_tim WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.JabatanTim{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var jabatanTim domain.JabatanTim
		err := rows.Scan(&jabatanTim.Id, &jabatanTim.NamaJabatan, &jabatanTim.LevelJabatan)
		if err != nil {
			return domain.JabatanTim{}, err
		}
		return jabatanTim, nil
	}

	return domain.JabatanTim{}, errors.New("jabatan tim not found")
}

func (repository *JabatanTimRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.JabatanTim, error) {
	query := "SELECT id, nama_jabatan, level_jabatan FROM jabatan_tim ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.JabatanTim{}, err
	}
	defer rows.Close()

	var jabatanTimList []domain.JabatanTim
	for rows.Next() {
		var jabatanTim domain.JabatanTim
		err := rows.Scan(&jabatanTim.Id, &jabatanTim.NamaJabatan, &jabatanTim.LevelJabatan)
		if err != nil {
			return []domain.JabatanTim{}, err
		}

		jabatanTimList = append(jabatanTimList, jabatanTim)
	}

	return jabatanTimList, nil
}
