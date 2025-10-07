package domain

import "time"

type TimKerja struct {
	Id            int
	KodeTim       string
	NamaTim       string
	Keterangan    string
	Tahun         string
	IsActive      bool
	IsSekretariat bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProgramUnggulanTimKerja struct {
	Id                  int
	KodeTim             string
	KodeProgramUnggulan string
	IdProgramUnggulan   int
	NamaProgramUnggulan string
	Tahun               string
	KodeOpd             string
}
