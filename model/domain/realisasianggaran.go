package domain

import "time"

type RealisasiAnggaran struct {
	Id                int
	IdProgramUnggulan *int
	KodeSubkegiatan   string
	RealisasiAnggaran int
	KodeTim           string
	IdRencanaKinerja  string
	KodeOpd           string
	RencanaAksi       string
	FaktorPendorong   string
	FaktorPenghambat  string
	RekomendasiTl     string
	BuktiDukung       string
	Bulan             int
	Tahun             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
