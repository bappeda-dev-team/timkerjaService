package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"timkerjaService/model/domain"
)

type TimKerjaRepositoryImpl struct {
}

func NewTimKerjaRepositoryImpl() *TimKerjaRepositoryImpl {
	return &TimKerjaRepositoryImpl{}
}

func (repository *TimKerjaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error) {
	query := "INSERT INTO tim_kerja (kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query,
		timKerja.KodeTim,
		timKerja.NamaTim,
		timKerja.Keterangan,
		timKerja.Tahun,
		timKerja.IsActive,
		timKerja.IsSekretariat)
	if err != nil {
		return domain.TimKerja{}, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.TimKerja{}, fmt.Errorf("gagal ambil last insert id: %w", err)
	}
	timKerja.Id = int(lastID)
	return timKerja, nil
}

func (repository *TimKerjaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error) {
	query := "UPDATE tim_kerja SET nama_tim = ?, keterangan = ?, tahun = ?, is_active = ?, is_sekretariat = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, timKerja.NamaTim, timKerja.Keterangan, timKerja.Tahun, timKerja.IsActive, timKerja.IsSekretariat, timKerja.Id)
	if err != nil {
		return domain.TimKerja{}, err
	}

	return timKerja, nil
}

func (repository *TimKerjaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM tim_kerja WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *TimKerjaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.TimKerja{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive, &timKerja.IsSekretariat)
		if err != nil {
			return domain.TimKerja{}, err
		}
		return timKerja, nil
	}

	return domain.TimKerja{}, nil
}

func (repository *TimKerjaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE tahun = ? ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, tahun)
	if err != nil {
		return []domain.TimKerja{}, err
	}
	defer rows.Close()

	var timKerjaList []domain.TimKerja
	for rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive, &timKerja.IsSekretariat)
		if err != nil {
			return []domain.TimKerja{}, err
		}

		timKerjaList = append(timKerjaList, timKerja)
	}

	return timKerjaList, nil

}

func (repository *TimKerjaRepositoryImpl) FindAllWithSusunan(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAll(ctx, tx, tahun)
	if err != nil {
		return nil, nil, err
	}

	// Get all susunan tim with jabatan details
	query := `
        SELECT 
            st.id, 
            st.kode_tim, 
            st.pegawai_id,
            st.nama_pegawai,
            st.nama_jabatan_tim,
            jt.level_jabatan, -- ambil dari tabel jabatan_tim
            st.keterangan,
            st.is_active
        FROM susunan_tim st
        LEFT JOIN jabatan_tim jt ON st.nama_jabatan_tim = jt.nama_jabatan -- join dengan jabatan_tim untuk dapat level
        ORDER BY st.kode_tim, jt.level_jabatan ASC`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	susunanTimMap := make(map[string][]domain.SusunanTim)

	for rows.Next() {
		var susunanTim domain.SusunanTim
		var levelJabatan sql.NullInt32

		err := rows.Scan(
			&susunanTim.Id,
			&susunanTim.KodeTim,
			&susunanTim.PegawaiId,
			&susunanTim.NamaPegawai,
			&susunanTim.NamaJabatanTim,
			&levelJabatan,
			&susunanTim.Keterangan,
			&susunanTim.IsActive,
		)
		if err != nil {
			return nil, nil, err
		}

		// Handle null values
		if levelJabatan.Valid {
			susunanTim.LevelJabatan = int(levelJabatan.Int32)
		}

		susunanTimMap[susunanTim.KodeTim] = append(susunanTimMap[susunanTim.KodeTim], susunanTim)
	}

	return timKerjaList, susunanTimMap, nil
}

func (repository *TimKerjaRepositoryImpl) AddProgramUnggulan(ctx context.Context, tx *sql.Tx, programUnggulan domain.ProgramUnggulanTimKerja) (domain.ProgramUnggulanTimKerja, error) {
	log.Printf("Program Unggulan Input: %v", programUnggulan)
	query := "INSERT INTO tb_program_unggulan(kode_tim, id_program_unggulan, tahun, kode_opd, kode_program_unggulan) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, programUnggulan.KodeTim, programUnggulan.IdProgramUnggulan, programUnggulan.Tahun, programUnggulan.KodeOpd, programUnggulan.KodeProgramUnggulan)
	if err != nil {
		return domain.ProgramUnggulanTimKerja{}, err
	}

	return programUnggulan, nil
}

func (repository *TimKerjaRepositoryImpl) FindProgramUnggulanByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string, tahun int) ([]domain.ProgramUnggulanTimKerja, error) {
	query := "SELECT pu.id, pu.kode_tim, pu.id_program_unggulan, pu.kode_program_unggulan, pu.tahun, pu.kode_opd FROM tb_program_unggulan pu WHERE pu.kode_tim = ? AND pu.tahun = ?"
	rows, err := tx.QueryContext(ctx, query, kodeTim, tahun)
	if err != nil {
		return []domain.ProgramUnggulanTimKerja{}, err
	}
	defer rows.Close()

	var listProgramUnggulans []domain.ProgramUnggulanTimKerja

	for rows.Next() {
		var programUnggulan domain.ProgramUnggulanTimKerja

		err := rows.Scan(
			&programUnggulan.Id,
			&programUnggulan.KodeTim,
			&programUnggulan.IdProgramUnggulan,
			&programUnggulan.KodeProgramUnggulan,
			&programUnggulan.Tahun,
			&programUnggulan.KodeOpd,
		)
		if err != nil {
			return []domain.ProgramUnggulanTimKerja{}, err
		}

		listProgramUnggulans = append(listProgramUnggulans, programUnggulan)
	}

	return listProgramUnggulans, nil
}

func (repository *TimKerjaRepositoryImpl) FindAllTimNonSekretariat(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE NOT is_sekretariat AND tahun = ? AND is_active ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, tahun)
	if err != nil {
		return []domain.TimKerja{}, err
	}
	defer rows.Close()

	var timKerjaList []domain.TimKerja
	for rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive, &timKerja.IsSekretariat)
		if err != nil {
			return []domain.TimKerja{}, err
		}

		timKerjaList = append(timKerjaList, timKerja)
	}

	return timKerjaList, nil

}

func (repository *TimKerjaRepositoryImpl) FindAllTimNonSekretariatWithSusunan(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAllTimNonSekretariat(ctx, tx, tahun)
	if err != nil {
		return nil, nil, err
	}

	// Get all susunan tim with jabatan details
	query := `
        SELECT
            st.id,
            st.kode_tim,
            st.pegawai_id,
            st.nama_pegawai,
            st.nama_jabatan_tim,
            jt.level_jabatan, -- ambil dari tabel jabatan_tim
            st.keterangan,
            st.is_active
        FROM susunan_tim st
        LEFT JOIN jabatan_tim jt ON st.nama_jabatan_tim = jt.nama_jabatan -- join dengan jabatan_tim untuk dapat level
        ORDER BY st.kode_tim, jt.level_jabatan ASC`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	susunanTimMap := make(map[string][]domain.SusunanTim)

	for rows.Next() {
		var susunanTim domain.SusunanTim
		var levelJabatan sql.NullInt32

		err := rows.Scan(
			&susunanTim.Id,
			&susunanTim.KodeTim,
			&susunanTim.PegawaiId,
			&susunanTim.NamaPegawai,
			&susunanTim.NamaJabatanTim,
			&levelJabatan,
			&susunanTim.Keterangan,
			&susunanTim.IsActive,
		)
		if err != nil {
			return nil, nil, err
		}

		// Handle null values
		if levelJabatan.Valid {
			susunanTim.LevelJabatan = int(levelJabatan.Int32)
		}

		susunanTimMap[susunanTim.KodeTim] = append(susunanTimMap[susunanTim.KodeTim], susunanTim)
	}

	return timKerjaList, susunanTimMap, nil
}

func (repository *TimKerjaRepositoryImpl) FindAllTimSekretariat(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE is_sekretariat AND tahun = ? AND is_active ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query, tahun)
	if err != nil {
		return []domain.TimKerja{}, err
	}
	defer rows.Close()

	var timKerjaList []domain.TimKerja
	for rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive, &timKerja.IsSekretariat)
		if err != nil {
			return []domain.TimKerja{}, err
		}

		timKerjaList = append(timKerjaList, timKerja)
	}

	return timKerjaList, nil

}

func (repository *TimKerjaRepositoryImpl) FindAllTimSekretariatWithSusunan(ctx context.Context, tx *sql.Tx, tahun int) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAllTimSekretariat(ctx, tx, tahun)
	if err != nil {
		return nil, nil, err
	}

	// Get all susunan tim with jabatan details
	query := `
        SELECT
            st.id,
            st.kode_tim,
            st.pegawai_id,
            st.nama_pegawai,
            st.nama_jabatan_tim,
            jt.level_jabatan, -- ambil dari tabel jabatan_tim
            st.keterangan,
            st.is_active
        FROM susunan_tim st
        LEFT JOIN jabatan_tim jt ON st.nama_jabatan_tim = jt.nama_jabatan -- join dengan jabatan_tim untuk dapat level
        ORDER BY st.kode_tim, jt.level_jabatan ASC`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	susunanTimMap := make(map[string][]domain.SusunanTim)

	for rows.Next() {
		var susunanTim domain.SusunanTim
		var levelJabatan sql.NullInt32

		err := rows.Scan(
			&susunanTim.Id,
			&susunanTim.KodeTim,
			&susunanTim.PegawaiId,
			&susunanTim.NamaPegawai,
			&susunanTim.NamaJabatanTim,
			&levelJabatan,
			&susunanTim.Keterangan,
			&susunanTim.IsActive,
		)
		if err != nil {
			return nil, nil, err
		}

		// Handle null values
		if levelJabatan.Valid {
			susunanTim.LevelJabatan = int(levelJabatan.Int32)
		}

		susunanTimMap[susunanTim.KodeTim] = append(susunanTimMap[susunanTim.KodeTim], susunanTim)
	}

	return timKerjaList, susunanTimMap, nil
}

func (repository *TimKerjaRepositoryImpl) DeleteProgramUnggulan(ctx context.Context, tx *sql.Tx, id int, kodeTim string) error {
	query := "DELETE FROM tb_program_unggulan WHERE id = ? AND kode_tim = ?"
	_, err := tx.ExecContext(ctx, query, id, kodeTim)
	if err != nil {
		return err
	}
	return nil
}

func (repository *TimKerjaRepositoryImpl) AddRencanaKinerja(ctx context.Context, tx *sql.Tx, rencanaKinerja domain.RencanaKinerjaTimKerja) (domain.RencanaKinerjaTimKerja, error) {
	tahun, err := strconv.Atoi(rencanaKinerja.Tahun)
	if err != nil {
		return domain.RencanaKinerjaTimKerja{}, fmt.Errorf("gagal ambil tim sekretariat: %w", err)
	}
	// guard biar tidak sembarangan
	timSekretariat, err := repository.FindAllTimSekretariat(ctx, tx, tahun)
	if err != nil {
		return domain.RencanaKinerjaTimKerja{}, fmt.Errorf("gagal ambil tim sekretariat: %w", err)
	}
	// Buat lookup map agar pengecekan KodeTim O(1)
	timMap := make(map[string]struct{}, len(timSekretariat))
	for _, tim := range timSekretariat {
		timMap[tim.KodeTim] = struct{}{}
	}

	if _, ok := timMap[rencanaKinerja.KodeTim]; !ok {
		return domain.RencanaKinerjaTimKerja{}, fmt.Errorf(
			"kode tim '%s' tidak termasuk dalam tim sekretariat", rencanaKinerja.KodeTim,
		)
	}

	query := "INSERT INTO rencana_kinerja_sekretariat(kode_tim, id_rencana_kinerja, id_pegawai, tahun, kode_opd) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, rencanaKinerja.KodeTim, rencanaKinerja.IdRencanaKinerja, rencanaKinerja.IdPegawai, rencanaKinerja.Tahun, rencanaKinerja.KodeOpd)
	if err != nil {
		return domain.RencanaKinerjaTimKerja{}, err
	}

	return rencanaKinerja, nil
}

func (repository *TimKerjaRepositoryImpl) FindRencanaKinerjaByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string, tahun int) ([]domain.RencanaKinerjaTimKerja, error) {
	// guard biar tidak sembarangan
	timSekretariat, err := repository.FindAllTimSekretariat(ctx, tx, tahun)
	if err != nil {
		return []domain.RencanaKinerjaTimKerja{}, fmt.Errorf("gagal ambil tim sekretariat: %w", err)
	}
	// Buat lookup map agar pengecekan KodeTim O(1)
	timMap := make(map[string]struct{}, len(timSekretariat))
	for _, tim := range timSekretariat {
		timMap[tim.KodeTim] = struct{}{}
	}

	if _, ok := timMap[kodeTim]; !ok {
		return []domain.RencanaKinerjaTimKerja{}, fmt.Errorf(
			"kode tim '%s' tidak termasuk dalam tim sekretariat", kodeTim,
		)
	}

	query := "SELECT rekin.id, rekin.kode_tim, rekin.id_rencana_kinerja, rekin.id_pegawai, rekin.tahun, rekin.kode_opd FROM rencana_kinerja_sekretariat rekin WHERE rekin.kode_tim = ? AND rekin.tahun = ?"
	rows, err := tx.QueryContext(ctx, query, kodeTim, tahun)
	if err != nil {
		return []domain.RencanaKinerjaTimKerja{}, err
	}
	defer rows.Close()

	var listRencanaKinerja []domain.RencanaKinerjaTimKerja

	for rows.Next() {
		var rencanaKinerja domain.RencanaKinerjaTimKerja

		err := rows.Scan(
			&rencanaKinerja.Id,
			&rencanaKinerja.KodeTim,
			&rencanaKinerja.IdRencanaKinerja,
			&rencanaKinerja.IdPegawai,
			&rencanaKinerja.Tahun,
			&rencanaKinerja.KodeOpd,
		)
		if err != nil {
			return []domain.RencanaKinerjaTimKerja{}, err
		}

		listRencanaKinerja = append(listRencanaKinerja, rencanaKinerja)
	}

	return listRencanaKinerja, nil
}

func (repository *TimKerjaRepositoryImpl) DeleteRencanaKinerja(ctx context.Context, tx *sql.Tx, id int, kodeTim string) error {
	query := "DELETE FROM rencana_kinerja_sekretariat WHERE id = ? AND kode_tim = ?"
	_, err := tx.ExecContext(ctx, query, id, kodeTim)
	if err != nil {
		return err
	}
	return nil
}

func (repository *TimKerjaRepositoryImpl) SaveRealisasiPokin(ctx context.Context, tx *sql.Tx, realisasi domain.RealisasiPokin) (domain.RealisasiPokin, error) {
	query := `INSERT INTO realisasi_pokin(
                id_pokin, kode_tim, jenis_pohon,
				jenis_item, kode_item, nama_item, pagu, realisasi,
                faktor_pendorong, faktor_penghambat, rtl, url_bukti_dukung,
                tahun, kode_opd) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, realisasi.IdPokin, realisasi.KodeItem, realisasi.JenisPohon,
		realisasi.JenisItem, realisasi.KodeItem, realisasi.NamaItem, realisasi.Pagu, realisasi.Realisasi,
		realisasi.FaktorPendorong, realisasi.FaktorPenghambat, realisasi.Rtl, realisasi.UrlBuktiDukung,
		realisasi.Tahun, realisasi.KodeOpd)
	if err != nil {
		return domain.RealisasiPokin{}, err
	}

	return realisasi, nil
}

func (repository *TimKerjaRepositoryImpl) UpdateRealisasiPokin(ctx context.Context, tx *sql.Tx, realisasi domain.RealisasiPokin) (domain.RealisasiPokin, error) {
	query := `
		UPDATE realisasi_pokin
		SET
            kode_item = ?,
			nama_item = ?,
			pagu = ?,
			realisasi = ?,
			faktor_pendorong = ?,
			faktor_penghambat = ?,
			rtl = ?,
			url_bukti_dukung = ?
		WHERE id_pokin = ?
			AND kode_tim = ?
	`

	result, err := tx.ExecContext(ctx, query,
		realisasi.KodeItem,
		realisasi.NamaItem,
		realisasi.Pagu,
		realisasi.Realisasi,
		realisasi.FaktorPendorong,
		realisasi.FaktorPenghambat,
		realisasi.Rtl,
		realisasi.UrlBuktiDukung,
		realisasi.IdPokin,
		realisasi.KodeTim,
	)
	if err != nil {
		return domain.RealisasiPokin{}, fmt.Errorf("gagal update realisasi pokin: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.RealisasiPokin{}, fmt.Errorf("gagal membaca hasil update: %w", err)
	}
	if rowsAffected == 0 {
		return domain.RealisasiPokin{}, fmt.Errorf("tidak ada data yang diperbarui (id_pokin=%v, kode_item=%s, kode_tim=%s)",
			realisasi.IdPokin, realisasi.KodeItem, realisasi.KodeTim)
	}

	return realisasi, nil
}

func (repo *TimKerjaRepositoryImpl) FindAllRealisasiPokinByKodeItemTahun(ctx context.Context, tx *sql.Tx, kodeTim string, tahun string) ([]domain.RealisasiPokin, error) {
	query := `
		SELECT
			id,
			id_pokin,
			kode_tim,
			jenis_pohon,
			jenis_item,
			kode_item,
			nama_item,
			pagu,
			realisasi,
			faktor_pendorong,
			faktor_penghambat,
			rtl,
			url_bukti_dukung,
			tahun,
			kode_opd,
			created_at,
			updated_at
		FROM realisasi_pokin
		WHERE kode_tim = ? AND tahun = ?
		ORDER BY id ASC
	`

	rows, err := tx.QueryContext(ctx, query, kodeTim, tahun)
	if err != nil {
		return nil, fmt.Errorf("gagal query realisasi_pokin: %w", err)
	}
	defer rows.Close()

	var list []domain.RealisasiPokin

	for rows.Next() {
		var r domain.RealisasiPokin

		err := rows.Scan(
			&r.Id,
			&r.IdPokin,
			&r.KodeTim,
			&r.JenisPohon,
			&r.JenisItem,
			&r.KodeItem,
			&r.NamaItem,
			&r.Pagu,
			&r.Realisasi,
			&r.FaktorPendorong,
			&r.FaktorPenghambat,
			&r.Rtl,
			&r.UrlBuktiDukung,
			&r.Tahun,
			&r.KodeOpd,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("gagal memindai hasil realisasi_pokin: %w", err)
		}

		list = append(list, r)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi hasil realisasi_pokin: %w", err)
	}

	return list, nil
}

func (r *TimKerjaRepositoryImpl) FindRealisasiByKodeTimAndPohonIDs(
	ctx context.Context,
	tx *sql.Tx,
	kodeTim string,
	bulan int,
	tahun int,
	pohonIDs []int,
) (map[int]domain.RealisasiAnggaranRecord, error) {

	if len(pohonIDs) == 0 {
		return map[int]domain.RealisasiAnggaranRecord{}, nil
	}

	// --- 1) Build placeholder MySQL: ?, ?, ?, ...
	placeholders := make([]string, len(pohonIDs))
	args := make([]any, 0, len(pohonIDs)+3)

	for i, id := range pohonIDs {
		placeholders[i] = "?"
		args = append(args, id)
	}

	args = append(args, kodeTim)
	args = append(args, bulan)
	args = append(args, tahun)

	query := fmt.Sprintf(`
        SELECT
            id_pohon, realisasi_anggaran, rencana_aksi, faktor_pendorong,
            faktor_penghambat, risiko_hukum, rekomendasi_tl
        FROM realisasi_anggaran
        WHERE id_pohon IN (%s)
          AND kode_tim = ?
          AND bulan = ?
          AND tahun = ?
    `, strings.Join(placeholders, ","))

	// --- 2) Execute query
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// --- 3) Parse result
	result := make(map[int]domain.RealisasiAnggaranRecord)
	for rows.Next() {
		var rec domain.RealisasiAnggaranRecord
		if err := rows.Scan(
			&rec.IdPohon,
			&rec.RealisasiAnggaran,
			&rec.RencanaAksi,
			&rec.FaktorPendorong,
			&rec.FaktorPenghambat,
			&rec.RisikoHukum,
			&rec.RekomendasiTl,
		); err != nil {
			return nil, err
		}

		result[rec.IdPohon] = rec
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TimKerjaRepositoryImpl) FindRealisasiByKodeTimAndRekinSekretariatIds(
	ctx context.Context,
	tx *sql.Tx,
	kodeTim string,
	bulan int,
	tahun int,
	rekinSekretIds []int,
) (map[int]domain.RealisasiAnggaranRecord, error) {

	if len(rekinSekretIds) == 0 {
		return map[int]domain.RealisasiAnggaranRecord{}, nil
	}

	// --- 1) Build placeholder MySQL: ?, ?, ?, ...
	placeholders := make([]string, len(rekinSekretIds))
	args := make([]any, 0, len(rekinSekretIds)+3)

	for i, id := range rekinSekretIds {
		placeholders[i] = "?"
		args = append(args, id)
	}

	args = append(args, kodeTim)
	args = append(args, bulan)
	args = append(args, tahun)

	query := fmt.Sprintf(`
        SELECT
            id_rencana_kinerja_sekretariat, realisasi_anggaran, faktor_pendorong,
            faktor_penghambat, risiko_hukum, rekomendasi_tl
        FROM realisasi_anggaran
        WHERE id_program_unggulan = 0
          AND id_rencana_kinerja_sekretariat IN (%s)
          AND kode_tim = ?
          AND bulan = ?
          AND tahun = ?
    `, strings.Join(placeholders, ","))

	// --- 2) Execute query
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// --- 3) Parse result
	result := make(map[int]domain.RealisasiAnggaranRecord)
	for rows.Next() {
		var rec domain.RealisasiAnggaranRecord
		if err := rows.Scan(
			&rec.IdRencanaKinerjaSekretariat,
			&rec.RealisasiAnggaran,
			&rec.FaktorPendorong,
			&rec.FaktorPenghambat,
			&rec.RisikoHukum,
			&rec.RekomendasiTl,
		); err != nil {
			return nil, err
		}

		result[rec.IdRencanaKinerjaSekretariat] = rec
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
