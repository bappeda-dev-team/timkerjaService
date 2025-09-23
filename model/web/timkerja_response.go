package web

import "time"

type TimKerjaResponse struct {
	Id         int       `json:"id"`
	KodeTim    string    `json:"kode_tim"`
	NamaTim    string    `json:"nama_tim"`
	Keterangan string    `json:"keterangan"`
	Tahun      string    `json:"tahun"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
