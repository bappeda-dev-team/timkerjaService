package repository

import (
	"context"
	"database/sql"
	"fmt"
	"timkerjaService/model/domain"
)

type RealisasiAnggaranRepositoryImpl struct {
}

func NewRealisasiAnggaranRepositoryImpl() *RealisasiAnggaranRepositoryImpl {
	return &RealisasiAnggaranRepositoryImpl{}
}

// func (r *RealisasiAnggaranRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, ra domain.RealisasiAnggaran) (domain.RealisasiAnggaran, error) {
// 	query := `
// 		INSERT INTO realisasi_anggaran (
// 			kode_subkegiatan,
// 			realisasi_anggaran,
// 			kode_opd,
// 			rencana_aksi,
// 			faktor_pendorong,
// 			faktor_penghambat,
// 			rekomendasi_tl,
// 			bukti_dukung,
// 			bulan,
// 			tahun
// 		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`
// 	res, err := tx.ExecContext(
// 		ctx,
// 		query,
// 		ra.KodeSubkegiatan,
// 		ra.RealisasiAnggaran,
// 		ra.KodeOpd,
// 		ra.RencanaAksi,
// 		ra.FaktorPendorong,
// 		ra.FaktorPenghambat,
// 		ra.RekomendasiTl,
// 		ra.BuktiDukung,
// 		ra.Bulan,
// 		ra.Tahun,
// 	)
// 	if err != nil {
// 		return domain.RealisasiAnggaran{}, fmt.Errorf("gagal insert realisasi_anggaran: %w", err)
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return domain.RealisasiAnggaran{}, fmt.Errorf("gagal ambil last insert id: %w", err)
// 	}

// 	ra.Id = int(id)
// 	return ra, nil
// }

// func (r *RealisasiAnggaranRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ra domain.RealisasiAnggaran) (domain.RealisasiAnggaran, error) {
// 	query := `
// 		UPDATE realisasi_anggaran
// 		SET
// 			kode_subkegiatan = ?,
// 			realisasi_anggaran = ?,
// 			kode_opd = ?,
// 			rencana_aksi = ?,
// 			faktor_pendorong = ?,
// 			faktor_penghambat = ?,
// 			rekomendasi_tl = ?,
// 			bukti_dukung = ?,
// 			bulan = ?,
// 			tahun = ?
// 		WHERE id = ?
// 	`
// 	_, err := tx.ExecContext(
// 		ctx,
// 		query,
// 		ra.KodeSubkegiatan,
// 		ra.RealisasiAnggaran,
// 		ra.KodeOpd,
// 		ra.RencanaAksi,
// 		ra.FaktorPendorong,
// 		ra.FaktorPenghambat,
// 		ra.RekomendasiTl,
// 		ra.BuktiDukung,
// 		ra.Bulan,
// 		ra.Tahun,
// 		ra.Id,
// 	)
// 	if err != nil {
// 		return domain.RealisasiAnggaran{}, fmt.Errorf("gagal update realisasi_anggaran id=%d: %w", ra.Id, err)
// 	}
// 	return ra, nil
// }

func (r *RealisasiAnggaranRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := `DELETE FROM realisasi_anggaran WHERE id = ?`
	if _, err := tx.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("gagal delete realisasi_anggaran id=%d: %w", id, err)
	}
	return nil
}

func (r *RealisasiAnggaranRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.RealisasiAnggaran, error) {
	query := `
		SELECT
			id,
			id_program_unggulan,
			kode_tim,
			id_rencana_kinerja,
			kode_subkegiatan,
			realisasi_anggaran,
			kode_opd,
			rencana_aksi,
			faktor_pendorong,
			faktor_penghambat,
			rekomendasi_tl,
			bukti_dukung,
			bulan,
			tahun,
			created_at,
			updated_at
		FROM realisasi_anggaran
		WHERE id = ?
	`
	row := tx.QueryRowContext(ctx, query, id)

	var ra domain.RealisasiAnggaran
	if err := row.Scan(
		&ra.Id,
		&ra.IdProgramUnggulan,
		&ra.KodeTim,
		&ra.IdRencanaKinerja,
		&ra.KodeSubkegiatan,
		&ra.RealisasiAnggaran,
		&ra.KodeOpd,
		&ra.RencanaAksi,
		&ra.FaktorPendorong,
		&ra.FaktorPenghambat,
		&ra.RekomendasiTl,
		&ra.BuktiDukung,
		&ra.Bulan,
		&ra.Tahun,
		&ra.CreatedAt,
		&ra.UpdatedAt,
	); err != nil {
		return domain.RealisasiAnggaran{}, err
	}

	return ra, nil
}

func (r *RealisasiAnggaranRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeSubkegiatan string, kodeTim string, idRencanaKinerja string, bulan string, tahun string) ([]domain.RealisasiAnggaran, error) {
	query := `
		SELECT
			id,
			kode_tim,
			id_rencana_kinerja,
			kode_subkegiatan,
			realisasi_anggaran,
			kode_opd,
			rencana_aksi,
			faktor_pendorong,
			faktor_penghambat,
			rekomendasi_tl,
			bukti_dukung,
			bulan,
			tahun,
			created_at,
			updated_at
		FROM realisasi_anggaran
		WHERE kode_subkegiatan = ? AND kode_tim = ? AND id_rencana_kinerja = ? AND bulan = ? AND tahun = ?
		ORDER BY id ASC
	`

	rows, err := tx.QueryContext(ctx, query, kodeSubkegiatan, kodeTim, idRencanaKinerja, bulan, tahun)
	if err != nil {
		return nil, fmt.Errorf("gagal query realisasi_anggaran: %w", err)
	}
	defer rows.Close()

	var list []domain.RealisasiAnggaran
	for rows.Next() {
		var ra domain.RealisasiAnggaran
		if err := rows.Scan(
			&ra.Id,
			&ra.KodeTim,
			&ra.IdRencanaKinerja,
			&ra.KodeSubkegiatan,
			&ra.RealisasiAnggaran,
			&ra.KodeOpd,
			&ra.RencanaAksi,
			&ra.FaktorPendorong,
			&ra.FaktorPenghambat,
			&ra.RekomendasiTl,
			&ra.BuktiDukung,
			&ra.Bulan,
			&ra.Tahun,
			&ra.CreatedAt,
			&ra.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("gagal scan realisasi_anggaran: %w", err)
		}
		list = append(list, ra)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterasi hasil realisasi_anggaran error: %w", err)
	}

	return list, nil
}

func (r *RealisasiAnggaranRepositoryImpl) Upsert(ctx context.Context, tx *sql.Tx, ra domain.RealisasiAnggaran) (domain.RealisasiAnggaran, error) {
	// duplicate key : id_pohon, bulan, tahun
	query := `
INSERT INTO realisasi_anggaran (
	kode_tim, id_rencana_kinerja, id_pohon, id_program_unggulan, kode_subkegiatan, realisasi_anggaran, kode_opd, rencana_aksi, faktor_pendorong,
	faktor_penghambat, risiko_hukum, rekomendasi_tl, bukti_dukung, bulan, tahun
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
	realisasi_anggaran = VALUES(realisasi_anggaran),
	kode_opd           = VALUES(kode_opd),
	rencana_aksi       = VALUES(rencana_aksi),
	faktor_pendorong   = VALUES(faktor_pendorong),
	faktor_penghambat  = VALUES(faktor_penghambat),
	risiko_hukum       = VALUES(risiko_hukum),
	rekomendasi_tl     = VALUES(rekomendasi_tl),
	bukti_dukung       = VALUES(bukti_dukung),
    id_rencana_kinerja = VALUES(id_rencana_kinerja),
    id_program_unggulan = VALUES(id_program_unggulan),
    kode_tim           = VALUES(kode_tim),
	updated_at         = NOW()
`
	_, err := tx.ExecContext(ctx, query,
		ra.KodeTim, ra.IdRencanaKinerja, ra.IdPohon, ra.IdProgramUnggulan, ra.KodeSubkegiatan, ra.RealisasiAnggaran, ra.KodeOpd, ra.RencanaAksi, ra.FaktorPendorong,
		ra.FaktorPenghambat, ra.RisikoHukum, ra.RekomendasiTl, ra.BuktiDukung, ra.Bulan, ra.Tahun,
	)
	if err != nil {
		return domain.RealisasiAnggaran{}, fmt.Errorf("upsert realisasi_anggaran: %w", err)
	}
	return ra, nil
}
