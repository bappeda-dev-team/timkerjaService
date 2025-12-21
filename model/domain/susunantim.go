package domain

import "time"

type SusunanTim struct {
	Id             int
	KodeTim        string
	PegawaiId      string
	IdJabatanTim   int
	NamaJabatanTim string
	NamaPegawai    string
	IsActive       bool
	Keterangan     *string
	LevelJabatan   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Bulan          int
	Tahun          int
}
