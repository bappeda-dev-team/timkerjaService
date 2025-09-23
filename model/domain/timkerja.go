package domain

import "time"

type TimKerja struct {
	Id         int
	KodeTim    string
	NamaTim    string
	Keterangan string
	Tahun      string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
