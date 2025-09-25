package repository

import (
	"context"
	"database/sql"
	"timkerjaService/model/domain"
)

type TimKerjaRepositoryImpl struct {
}

func NewTimKerjaRepositoryImpl() *TimKerjaRepositoryImpl {
	return &TimKerjaRepositoryImpl{}
}

func (repository *TimKerjaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error) {
	query := "INSERT INTO tim_kerja (kode_tim, nama_tim, keterangan, tahun, is_active) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, timKerja.KodeTim, timKerja.NamaTim, timKerja.Keterangan, timKerja.Tahun, timKerja.IsActive)
	if err != nil {
		return domain.TimKerja{}, err
	}

	return timKerja, nil
}

func (repository *TimKerjaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, timKerja domain.TimKerja) (domain.TimKerja, error) {
	query := "UPDATE tim_kerja SET nama_tim = ?, keterangan = ?, tahun = ?, is_active = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, timKerja.NamaTim, timKerja.Keterangan, timKerja.Tahun, timKerja.IsActive, timKerja.Id)
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
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active FROM tim_kerja WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.TimKerja{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive)
		if err != nil {
			return domain.TimKerja{}, err
		}
		return timKerja, nil
	}

	return domain.TimKerja{}, nil
}

func (repository *TimKerjaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TimKerja, error) {
	query := "SELECT id, kode_tim, nama_tim, keterangan, tahun, is_active FROM tim_kerja ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.TimKerja{}, err
	}
	defer rows.Close()

	var timKerjaList []domain.TimKerja
	for rows.Next() {
		var timKerja domain.TimKerja
		err := rows.Scan(&timKerja.Id, &timKerja.KodeTim, &timKerja.NamaTim, &timKerja.Keterangan, &timKerja.Tahun, &timKerja.IsActive)
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
		var levelJabatan sql.NullInt32 // handle jika level jabatan bisa null

		err := rows.Scan(
			&susunanTim.Id,
			&susunanTim.KodeTim,
			&susunanTim.PegawaiId,
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
