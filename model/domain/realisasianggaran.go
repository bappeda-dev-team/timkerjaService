package domain

import "time"

type RealisasiAnggaran struct {
	Id                int
	KodeSubkegiatan   string
	RealisasiAnggaran int
	KodeOpd           string
	RencanaAksi       string
	FaktorPendorong   string
	FaktorPenghambat  string
	RekomendasiTl     string
	BuktiDukung       string
	Bulan             string
	Tahun             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
