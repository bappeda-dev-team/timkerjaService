package web

type RealisasiAnggaranCreateRequest struct {
	KodeSubkegiatan   string `json:"kode_subkegiatan" validate:"required"`
	RealisasiAnggaran int    `json:"realisasi_anggaran" validate:"required"`
	KodeOpd           string `json:"kode_opd" validate:"required"`
	RencanaAksi       string `json:"rencana_aksi"`
	FaktorPendorong   string `json:"faktor_pendorong"`
	FaktorPenghambat  string `json:"faktor_penghambat"`
	RekomendasiTl     string `json:"rekomendasi_tl"`
	BuktiDukung       string `json:"bukti_dukung"`
	Bulan             string `json:"bulan" validate:"required"`
	Tahun             string `json:"tahun" validate:"required"`
}
