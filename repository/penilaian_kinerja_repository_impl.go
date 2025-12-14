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
  st.id AS susunan_tim_id,
  st.pegawai_id,
  st.nama_pegawai,
  jt.level_jabatan,
  st.nama_jabatan_tim,
  st.kode_tim,
  tk.nama_tim,
  tk.is_sekretariat,
  tk.keterangan,

  p.id,
  p.jenis_nilai,
  p.nilai_kinerja,
  p.tahun,
  p.bulan,
  p.kode_opd,
  p.created_at,
  p.updated_at,
  p.created_by

FROM susunan_tim st
LEFT JOIN jabatan_tim jt
  ON jt.id = st.jabatan_tim_id
LEFT JOIN tim_kerja tk
  ON tk.kode_tim = st.kode_tim
LEFT JOIN penilaian_kinerja p
  ON p.id_pegawai = st.pegawai_id
  AND p.tahun = ?
  AND p.bulan = ?
ORDER BY st.kode_tim, st.id;
`

	rows, err := tx.QueryContext(ctx, query, tahun, bulan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groupMap := make(map[string]*domain.LaporanPenilaian)
	order := []string{}

	for rows.Next() {
		var (
			susunanTimID    int
			pegawaiID       string
			namaPegawaiNS   sql.NullString
			levelJabatanNS  sql.NullInt64
			namaJabatanNS   sql.NullString
			kodeTim         string
			namaTimNS       sql.NullString
			isSekretariatNS sql.NullBool
			keteranganTimNS sql.NullString

			idNS           sql.NullInt64
			jenisNilaiNS   sql.NullString
			nilaiKinerjaNS sql.NullInt64
			tahunNS        sql.NullInt64
			bulanNS        sql.NullInt64
			kodeOpdNS      sql.NullString
			createdAtNS    sql.NullTime
			updatedAtNS    sql.NullTime
			createdByNS    sql.NullString
		)

		if err := rows.Scan(
			&susunanTimID,
			&pegawaiID,
			&namaPegawaiNS,
			&levelJabatanNS,
			&namaJabatanNS,
			&kodeTim,
			&namaTimNS,
			&isSekretariatNS,
			&keteranganTimNS,
			&idNS,
			&jenisNilaiNS,
			&nilaiKinerjaNS,
			&tahunNS,
			&bulanNS,
			&kodeOpdNS,
			&createdAtNS,
			&updatedAtNS,
			&createdByNS,
		); err != nil {
			return nil, err
		}

		// init group
		group, exists := groupMap[kodeTim]
		if !exists {
			group = &domain.LaporanPenilaian{
				NamaTim:       stringOrEmpty(namaTimNS),
				KodeTim:       kodeTim,
				IsSekretariat: boolOrFalse(isSekretariatNS),
				Keterangan:    stringOrEmpty(keteranganTimNS),
				Penilaians:    []domain.PenilaianKinerja{},
			}
			groupMap[kodeTim] = group
			order = append(order, kodeTim)
		}

		// Bentuk objek penilaian (bisa kosong/default)
		pen := domain.PenilaianKinerja{
			Id:              intOrZero(idNS),
			IdPegawai:       pegawaiID,
			NamaPegawai:     stringOrEmpty(namaPegawaiNS),
			SusunanTimId:    susunanTimID,
			LevelJabatanTim: intOrZero(levelJabatanNS),
			NamaJabatanTim:  stringOrEmpty(namaJabatanNS),
			KodeTim:         kodeTim,
			JenisNilai:      stringOrEmpty(jenisNilaiNS),
			NilaiKinerja:    intOrZero(nilaiKinerjaNS),
			Tahun:           strconv.Itoa(tahun),
			Bulan:           bulan,
			KodeOpd:         stringOrEmpty(kodeOpdNS),
			CreatedAt:       timeOrZero(createdAtNS),
			UpdatedAt:       timeOrZero(updatedAtNS),
			CreatedBy:       stringOrEmpty(createdByNS),
		}

		group.Penilaians = append(group.Penilaians, pen)
	}

	// convert ke slice sesuai order
	result := make([]domain.LaporanPenilaian, 0, len(order))
	for _, k := range order {
		result = append(result, *groupMap[k])
	}

	return result, nil
}

func stringOrEmpty(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func intOrZero(ns sql.NullInt64) int {
	if ns.Valid {
		return int(ns.Int64)
	}
	return 0
}

func boolOrFalse(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func timeOrZero(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}
