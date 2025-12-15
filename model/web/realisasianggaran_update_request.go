package web

type RealisasiAnggaranUpdateRequest struct {
	Id                int    `json:"id" validate:"required"`
	KodeTim           string `json:"kode_tim" validate:"required"`
	IdPohon           int    `json:"id_pohon" validate:"required"`
	IdRencanaKinerja  string `json:"id_rencana_kinerja" validate:"required"`
	KodeSubkegiatan   string `json:"kode_subkegiatan" validate:"required"`
	RealisasiAnggaran int    `json:"realisasi_anggaran" validate:"required"`
	KodeOpd           string `json:"kode_opd" validate:"required"`
	RencanaAksi       string `json:"rencana_aksi"`
	FaktorPendorong   string `json:"faktor_pendorong"`
	FaktorPenghambat  string `json:"faktor_penghambat"`
	RekomendasiTl     string `json:"rekomendasi_tl"`
	RisikoHukum       string `json:"risiko_hukum"`
	BuktiDukung       string `json:"bukti_dukung"`
	Bulan             string `json:"bulan" validate:"required"`
	Tahun             string `json:"tahun" validate:"required"`
}
