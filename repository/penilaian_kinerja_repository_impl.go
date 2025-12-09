package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type PenilaianKinerjaRepositoryImpl struct {
}

func NewPenilaianKinerjaRepositoryImpl() *PenilaianKinerjaRepositoryImpl {
	return &PenilaianKinerjaRepositoryImpl{}
}

func (repo *PenilaianKinerjaRepositoryImpl) Create(
	ctx context.Context,
	tx *sql.Tx,
	penilaian domain.PenilaianKinerja,
) (domain.PenilaianKinerja, error) {

	query := `
	INSERT INTO tim_kerja_service.penilaian_kinerja
		(id_pegawai, kode_tim, jenis_nilai, nilai_kinerja, tahun, bulan, kode_opd)
	VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		penilaian.IdPegawai,
		penilaian.KodeTim,
		penilaian.JenisNilai,
		penilaian.NilaiKinerja,
		penilaian.Tahun,
		penilaian.Bulan,
		penilaian.KodeOpd,
	)
	if err != nil {
		return domain.PenilaianKinerja{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.PenilaianKinerja{}, err
	}

	penilaian.Id = int(id)

	return penilaian, nil
}

func (repo *PenilaianKinerjaRepositoryImpl) Update(
	ctx context.Context,
	tx *sql.Tx,
	penilaian domain.PenilaianKinerja,
	id int,
) (domain.PenilaianKinerja, error) {

	query := `
	UPDATE tim_kerja_service.penilaian_kinerja
	SET
		id_pegawai    = ?,
		kode_tim      = ?,
		jenis_nilai   = ?,
		nilai_kinerja = ?,
		tahun         = ?,
		bulan         = ?,
		kode_opd      = ?
	WHERE id = ?;
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		penilaian.IdPegawai,
		penilaian.KodeTim,
		penilaian.JenisNilai,
		penilaian.NilaiKinerja,
		penilaian.Tahun,
		penilaian.Bulan,
		penilaian.KodeOpd,
		id,
	)
	if err != nil {
		return domain.PenilaianKinerja{}, err
	}

	// pastikan ID tetap terset
	penilaian.Id = id

	return penilaian, nil
}

func (repo *PenilaianKinerjaRepositoryImpl) ExistById(
	ctx context.Context,
	tx *sql.Tx,
	id int,
) (bool, error) {

	query := `
	SELECT 1
	FROM tim_kerja_service.penilaian_kinerja
	WHERE id = ?
	LIMIT 1;
	`

	var exists int
	err := tx.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
