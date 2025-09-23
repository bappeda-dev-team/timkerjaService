package domain

import "time"

type SusunanTim struct {
	Id             int
	KodeTim        string
	PegawaiId      string
	NamaJabatanTim string
	IsActive       bool
	Keterangan     *string
	LevelJabatan   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
