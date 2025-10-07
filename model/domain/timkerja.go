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

type ProgramUnggulanTimKerja struct {
	Id                  int
	KodeTim             string
	IdProgramUnggulan   int
	NamaProgramUnggulan string
	Tahun               string
	KodeOpd             string
}
