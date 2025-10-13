package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (repository *TimKerjaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
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

func (repository *TimKerjaRepositoryImpl) FindAllWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAll(ctx, tx)
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

func (repository *TimKerjaRepositoryImpl) FindProgramUnggulanByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.ProgramUnggulanTimKerja, error) {
	query := "SELECT pu.id, pu.kode_tim, pu.id_program_unggulan, pu.tahun, pu.kode_opd FROM tb_program_unggulan pu WHERE pu.kode_tim = ?"
	rows, err := tx.QueryContext(ctx, query, kodeTim)
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

func (repository *TimKerjaRepositoryImpl) FindAllTimNonSekretariat(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE NOT is_sekretariat ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
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

func (repository *TimKerjaRepositoryImpl) FindAllTimNonSekretariatWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAllTimNonSekretariat(ctx, tx)
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

func (repository *TimKerjaRepositoryImpl) FindAllTimSekretariat(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active, is_sekretariat FROM tim_kerja WHERE is_sekretariat ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
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

func (repository *TimKerjaRepositoryImpl) FindAllTimSekretariatWithSusunan(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, map[string][]domain.SusunanTim, error) {
	timKerjaList, err := repository.FindAllTimSekretariat(ctx, tx)
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
	// guard biar tidak sembarangan
	timSekretariat, err := repository.FindAllTimSekretariat(ctx, tx)
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

	query := "INSERT INTO rencana_kinerja_sekretariat(kode_tim, id_rencana_kinerja, tahun, kode_opd) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, rencanaKinerja.KodeTim, rencanaKinerja.IdRencanaKinerja, rencanaKinerja.Tahun, rencanaKinerja.KodeOpd)
	if err != nil {
		return domain.RencanaKinerjaTimKerja{}, err
	}

	return rencanaKinerja, nil
}

func (repository *TimKerjaRepositoryImpl) FindRencanaKinerjaByKodeTim(ctx context.Context, tx *sql.Tx, kodeTim string) ([]domain.RencanaKinerjaTimKerja, error) {
	// guard biar tidak sembarangan
	timSekretariat, err := repository.FindAllTimSekretariat(ctx, tx)
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

	query := "SELECT rekin.id, rekin.kode_tim, rekin.id_rencana_kinerja, rekin.tahun, rekin.kode_opd FROM rencana_kinerja_sekretariat rekin WHERE rekin.kode_tim = ?"
	rows, err := tx.QueryContext(ctx, query, kodeTim)
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
