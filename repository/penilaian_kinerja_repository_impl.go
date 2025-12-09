package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"
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
	INSERT INTO penilaian_kinerja
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
	UPDATE penilaian_kinerja
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
	FROM penilaian_kinerja
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
func (repo *PenilaianKinerjaRepositoryImpl) FindByTahunBulan(
	ctx context.Context,
	tx *sql.Tx,
	tahun int,
	bulan int,
) ([]domain.LaporanPenilaian, error) {

	query := `
SELECT
  p.id,
  p.id_pegawai,
  st.nama_pegawai,
  st.nama_jabatan_tim,
  p.kode_tim,
  p.jenis_nilai,
  p.nilai_kinerja,
  p.tahun,
  p.bulan,
  p.kode_opd,
  p.created_at,
  p.updated_at,
  p.created_by,
  tk.nama_tim
FROM penilaian_kinerja p
LEFT JOIN susunan_tim st
  ON st.pegawai_id = p.id_pegawai
  AND st.kode_tim = p.kode_tim
LEFT JOIN tim_kerja tk
  ON tk.kode_tim = p.kode_tim
WHERE p.tahun = ? AND p.bulan = ?
ORDER BY tk.nama_tim, p.kode_tim, p.id;
`

	rows, err := tx.QueryContext(ctx, query, tahun, bulan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// map untuk grouping, dan slice untuk menjaga urutan tim sesuai first-seen
	groupMap := make(map[string]*domain.LaporanPenilaian)
	order := make([]string, 0, 16)

	for rows.Next() {
		var (
			id              int
			idPegawai       string
			namaPegawaiNS   sql.NullString
			namaJabatanNS   sql.NullString
			kodeTim         string
			jenisNilai      string
			nilaiKinerjaInt sql.NullInt64
			tahunStr        sql.NullString
			bulanInt        sql.NullInt64
			kodeOpdNS       sql.NullString
			createdAtNS     sql.NullTime
			updatedAtNS     sql.NullTime
			createdByNS     sql.NullString
			namaTimNS       sql.NullString
		)

		if err := rows.Scan(
			&id,
			&idPegawai,
			&namaPegawaiNS,
			&namaJabatanNS,
			&kodeTim,
			&jenisNilai,
			&nilaiKinerjaInt,
			&tahunStr,
			&bulanInt,
			&kodeOpdNS,
			&createdAtNS,
			&updatedAtNS,
			&createdByNS,
			&namaTimNS,
		); err != nil {
			return nil, err
		}

		// konversi sql.Null* ke tipe domain
		nilaiKinerja := 0
		if nilaiKinerjaInt.Valid {
			nilaiKinerja = int(nilaiKinerjaInt.Int64)
		}

		tahunVal := ""
		if tahunStr.Valid {
			tahunVal = tahunStr.String
		} else {
			// jika di DB disimpan sebagai int, bisa juga konversi dari param tahun
			tahunVal = strconv.Itoa(tahun)
		}

		bulanVal := 0
		if bulanInt.Valid {
			bulanVal = int(bulanInt.Int64)
		} else {
			bulanVal = bulan
		}

		kodeOpd := ""
		if kodeOpdNS.Valid {
			kodeOpd = kodeOpdNS.String
		}

		namaPegawai := ""
		if namaPegawaiNS.Valid {
			namaPegawai = namaPegawaiNS.String
		}

		namaJabatanTim := ""
		if namaJabatanNS.Valid {
			namaJabatanTim = namaJabatanNS.String
		}

		createdAt := time.Time{}
		if createdAtNS.Valid {
			createdAt = createdAtNS.Time
		}

		updatedAt := time.Time{}
		if updatedAtNS.Valid {
			updatedAt = updatedAtNS.Time
		}

		createdBy := ""
		if createdByNS.Valid {
			createdBy = createdByNS.String
		}

		namaTim := ""
		if namaTimNS.Valid {
			namaTim = namaTimNS.String
		}

		pen := domain.PenilaianKinerja{
			Id:              id,
			IdPegawai:       idPegawai,
			NamaPegawai:     namaPegawai,
			NamaJabatanTim:  namaJabatanTim,
			KodeTim:         kodeTim,
			JenisNilai:      jenisNilai,
			NilaiKinerja:    nilaiKinerja,
			Tahun:           tahunVal,
			Bulan:           bulanVal,
			KodeOpd:         kodeOpd,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
			CreatedBy:       createdBy,
		}

		// grouping by kode_tim
		group, ok := groupMap[kodeTim]
		if !ok {
			// buat new group
			group = &domain.LaporanPenilaian{
				NamaTim:    namaTim,
				KodeTim:    kodeTim,
				Penilaians: make([]domain.PenilaianKinerja, 0, 8),
			}
			groupMap[kodeTim] = group
			order = append(order, kodeTim)
		}

		group.Penilaians = append(group.Penilaians, pen)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// convert map -> slice mengikuti order
	result := make([]domain.LaporanPenilaian, 0, len(order))
	for _, kode := range order {
		if g, ok := groupMap[kode]; ok {
			result = append(result, *g)
		}
	}

	return result, nil
}
