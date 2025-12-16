package domain

import (
	"time"
)

type PetugasTim struct {
	Id                int
	IdProgramUnggulan int
	KodeTim           string
	PegawaiId         string
	NamaPegawai       string
	Tahun             int
	Bulan             int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
