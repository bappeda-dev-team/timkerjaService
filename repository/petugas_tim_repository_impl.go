package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

func (repo *PetugasTimRepositoryImpl) FindAllByIdProgramUnggulans(
	ctx context.Context,
	tx *sql.Tx,
	idProgramUnggulans []int,
	bulan int,
	tahun int,
) ([]domain.PetugasTim, error) {

	if len(idProgramUnggulans) == 0 {
		return []domain.PetugasTim{}, nil
	}

	placeholders := make([]string, len(idProgramUnggulans))
	args := make([]any, 0, len(idProgramUnggulans))

	args = append(args, bulan, tahun)

	for i, id := range idProgramUnggulans {
		placeholders[i] = "?"
		args = append(args, id)
	}

	query := fmt.Sprintf(`
SELECT
    pt.id,
    pt.id_program_unggulan,
    pt.kode_tim,
    pt.pegawai_id,
    st.nama_pegawai
FROM petugas_tim pt
JOIN susunan_tim st
  ON st.pegawai_id = pt.pegawai_id
 AND st.kode_tim   = pt.kode_tim
WHERE pt.bulan = ?
  AND pt.tahun = ?
  AND pt.id_program_unggulan IN (%s)
ORDER BY pt.pegawai_id, pt.id
`, strings.Join(placeholders, ","))

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]domain.PetugasTim, 0)

	type petugasKey struct {
		PegawaiId         string
		IdProgramUnggulan int
		KodeTim           string
	}
	seenPegawai := make(map[petugasKey]struct{})

	for rows.Next() {
		var pet domain.PetugasTim

		if err := rows.Scan(
			&pet.Id,
			&pet.IdProgramUnggulan,
			&pet.KodeTim,
			&pet.PegawaiId,
			&pet.NamaPegawai,
		); err != nil {
			return nil, err
		}

		key := petugasKey{
			PegawaiId:         pet.PegawaiId,
			IdProgramUnggulan: pet.IdProgramUnggulan,
			KodeTim:           pet.KodeTim,
		}

		if _, exists := seenPegawai[key]; exists {
			continue
		}

		seenPegawai[key] = struct{}{}
		results = append(results, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
