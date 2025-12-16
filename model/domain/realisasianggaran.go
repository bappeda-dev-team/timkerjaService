package domain

import "time"

type RealisasiAnggaran struct {
	Id                          int
	IdRencanaKinerjaSekretariat int
	IdProgramUnggulan           int
	KodeSubkegiatan             string
	RealisasiAnggaran           int
	KodeTim                     string
	IdPohon                     int
	IdRencanaKinerja            string
	KodeOpd                     string
	RencanaAksi                 string
	FaktorPendorong             string
	FaktorPenghambat            string
	RisikoHukum                 string
	RekomendasiTl               string
	BuktiDukung                 string
	Bulan                       int
	Tahun                       string
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

type RealisasiAnggaranRecord struct {
	IdPohon                     int
	IdRencanaKinerjaSekretariat int
	IdProgramUnggulan           int
	RealisasiAnggaran           int
	RencanaAksi                 string
	FaktorPendorong             string
	FaktorPenghambat            string
	RisikoHukum                 string
	RekomendasiTl               string
}
