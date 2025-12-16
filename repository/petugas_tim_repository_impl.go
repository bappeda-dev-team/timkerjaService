package repository

import (
	"context"
	"database/sql"
	"fmt"
	"timkerjaService/model/domain"
)

type PetugasTimRepositoryImpl struct {
}

func NewPetugasTimRepositoryImpl() *PetugasTimRepositoryImpl {
	return &PetugasTimRepositoryImpl{}
}

func (repo *PetugasTimRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, petugasTimDomain domain.PetugasTim) (domain.PetugasTim, error) {
	query := `INSERT INTO petugas_tim (id_program_unggulan, kode_tim, pegawai_id, tahun, bulan)
                  VALUES (?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, query,
		petugasTimDomain.IdProgramUnggulan,
		petugasTimDomain.KodeTim,
		petugasTimDomain.PegawaiId,
		petugasTimDomain.Tahun,
		petugasTimDomain.Bulan)
	if err != nil {
		return domain.PetugasTim{}, err
	}
	idHasil, err := result.LastInsertId()
	if err != nil {
		return domain.PetugasTim{}, fmt.Errorf("gagal menyimpan: %v", err)
	}
	petugasTimDomain.Id = int(idHasil)
	return petugasTimDomain, nil
}

func (repo *PetugasTimRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idPetugasTim int) error {
	query := `DELETE FROM petugas_tim pt WHERE pt.id = ?`
	_, err := tx.ExecContext(ctx, query, idPetugasTim)
	if err != nil {
		return err
	}
	return nil
}
