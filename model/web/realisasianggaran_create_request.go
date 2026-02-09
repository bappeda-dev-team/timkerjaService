package web

type RealisasiAnggaranCreateRequest struct {
	KodeSubkegiatan             string `json:"kode_subkegiatan" validate:"required"`
	KodeTim                     string `json:"kode_tim" validate:"required"`
	IdRencanaKinerjaSekretariat int    `json:"id_rencana_kinerja_sekretariat"`
	IdProgramUnggulan           int    `json:"id_program_unggulan"`
	IdPohon                     int    `json:"id_pohon" validate:"required"`
	IdRencanaKinerja            string `json:"id_rencana_kinerja" validate:"required"`
	RealisasiAnggaran           int    `json:"realisasi_anggaran" validate:"required"`
	KodeOpd                     string `json:"kode_opd" validate:"required"`
	RencanaAksi                 string `json:"rencana_aksi"`
	FaktorPendorong             string `json:"faktor_pendorong"`
	FaktorPenghambat            string `json:"faktor_penghambat"`
	RekomendasiTl               string `json:"rekomendasi_tl"`
	RisikoHukum                 string `json:"risiko_hukum"`
	BuktiDukung                 string `json:"bukti_dukung"`
	Bulan                       int    `json:"bulan" validate:"required"`
	Tahun                       string `json:"tahun" validate:"required"`
	CatatanRealisasiAnggaran    string `json:"catatan_realisasi_anggaran"`
	CatatanPenataUsahaKeuangan  string `json:"catatan_penata_usaha_keuangan"`
	CatatanPelaporanKeuangan    string `json:"catatan_pelaporan_keuangan"`
	CatatanPelaporanAset        string `json:"catatan_pelaporan_aset"`
}
