package web

import "time"

type RealisasiAnggaranResponse struct {
	Id                int       `json:"id"`
	IdProgramUnggulan int       `json:"id_program_unggulan"`
	KodeTim           string    `json:"kode_tim"`
	IdRencanaKinerja  string    `json:"id_rencana_kinerja"`
	KodeSubkegiatan   string    `json:"kode_subkegiatan"`
	RealisasiAnggaran int       `json:"realisasi_anggaran"`
	KodeOpd           string    `json:"kode_opd"`
	RencanaAksi       string    `json:"rencana_aksi"`
	FaktorPendorong   string    `json:"faktor_pendorong"`
	FaktorPenghambat  string    `json:"faktor_penghambat"`
	RekomendasiTl     string    `json:"rekomendasi_tl"`
	BuktiDukung       string    `json:"bukti_dukung"`
	Bulan             int       `json:"bulan"`
	Tahun             string    `json:"tahun"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
