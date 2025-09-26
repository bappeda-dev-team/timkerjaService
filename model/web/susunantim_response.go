package web

import "time"

type SusunanTimResponse struct {
	Id             int       `json:"id"`
	KodeTim        string    `json:"kode_tim"`
	PegawaiId      string    `json:"nip"`
	NamaPegawai    string    `json:"nama_pegawai"`
	NamaJabatanTim string    `json:"nama_jabatan_tim"`
	IsActive       bool      `json:"is_active"`
	Keterangan     *string   `json:"keterangan"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
