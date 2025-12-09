package domain

import (
	"time"
)

type PenilaianKinerja struct {
	Id             int
	IdPegawai      string
	NamaPegawai    string
	NamaJabatanTim string
	KodeTim        string
	JenisNilai     string // Konker, Bappeda, Kerja TIM
	NilaiKinerja   int
	Tahun          string
	Bulan          int
	KodeOpd        string
	CreatedAt      time.Time // injected
	UpdatedAt      time.Time // injected
	CreatedBy      string    // injected
}

type LaporanPenilaian struct {
	NamaTim    string
	KodeTim    string
	Penilaians []PenilaianKinerja
}
