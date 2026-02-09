package web

import "time"

type RealisasiAnggaranResponse struct {
	Id                          int       `json:"id"`
	IdProgramUnggulan           int       `json:"id_program_unggulan"`
	IdRencanaKinerjaSekretariat int       `json:"id_rencana_kinerja_sekretariat"`
	KodeTim                     string    `json:"kode_tim"`
	IdPohon                     int       `json:"id_pohon" validate:"required"`
	IdRencanaKinerja            string    `json:"id_rencana_kinerja"`
	KodeSubkegiatan             string    `json:"kode_subkegiatan"`
	RealisasiAnggaran           int       `json:"realisasi_anggaran"`
	KodeOpd                     string    `json:"kode_opd"`
	RencanaAksi                 string    `json:"rencana_aksi"`
	FaktorPendorong             string    `json:"faktor_pendorong"`
	FaktorPenghambat            string    `json:"faktor_penghambat"`
	RekomendasiTl               string    `json:"rekomendasi_tl"`
	RisikoHukum                 string    `json:"risiko_hukum"`
	BuktiDukung                 string    `json:"bukti_dukung"`
	Bulan                       int       `json:"bulan"`
	Tahun                       string    `json:"tahun"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	CatatanRealisasiAnggaran    string    `json:"catatan_realisasi_anggaran"`
	CatatanPenataUsahaKeuangan  string    `json:"catatan_penata_usaha_keuangan"`
	CatatanPelaporanKeuangan    string    `json:"catatan_pelaporan_keuangan"`
	CatatanPelaporanAset        string    `json:"catatan_pelaporan_aset"`
}
