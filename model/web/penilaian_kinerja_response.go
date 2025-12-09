package web

import (
	"time"
)

type PenilaianKinerjaResponse struct {
	Id           int       `json:"id"`
	IdPegawai    string    `json:"id_pegawai" validate:"required"`
	KodeTim      string    `json:"kode_tim" validate:"required"`
	JenisNilai   string    `json:"jenis_nilai" validate:"required"`
	NilaiKinerja int       `json:"nilai_kinerja" validate:"required"`
	Tahun        string    `json:"tahun" validate:"required"`
	Bulan        int       `json:"bulan" validate:"required"`
	KodeOpd      string    `json:"kode_opd" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedBy    string    `json:"created_by"`
}
