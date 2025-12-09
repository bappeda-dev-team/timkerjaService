package domain

import (
	"time"
	"timkerjaService/internal"
)

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
	Pokin               []internal.TaggingPohonKinerjaItem
}

type RencanaKinerjaTimKerja struct {
	Id               int
	KodeTim          string
	IdRencanaKinerja string
	IdPegawai        string
	RencanaKinerja   string
	Tahun            string
	KodeOpd          string
}

type RealisasiPokin struct {
	Id               int
	IdPokin          int
	KodeTim          string
	JenisPohon       string
	JenisItem        string
	KodeItem         string
	NamaItem         string
	Pagu             int
	Realisasi        int
	FaktorPendorong  string
	FaktorPenghambat string
	Rtl              string
	UrlBuktiDukung   string
	Tahun            string
	KodeOpd          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
