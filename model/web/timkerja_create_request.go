package web

type TimKerjaCreateRequest struct {
	NamaTim       string `json:"nama_tim"`
	Keterangan    string `json:"keterangan"`
	Tahun         string `json:"tahun"`
	IsActive      bool   `json:"is_active"`
	IsSekretariat bool   `json:"is_sekretariat"`
}

type ProgramUnggulanTimKerjaRequest struct {
	KodeTim             string
	KodeProgramUnggulan string
	IdProgramUnggulan   int    `json:"id_program_unggulan" validate:"required"`
	Tahun               string `json:"tahun" validate:"required"`
	KodeOpd             string `json:"kode_opd" validate:"required"`
	RealisasiAnggaran   int    `json:"realisasi_anggaran"`
	FaktorPendorong     string `json:"faktor_pendorong"`
	FaktorPenghambat    string `json:"faktor_penghambat"`
	FileUrl             string `json:"file_url"`
}

type RencanaKinerjaRequest struct {
	KodeTim           string
	IdRencanaKinerja  string `json:"id_rencana_kinerja" validate:"required"`
	Tahun             string `json:"tahun" validate:"required"`
	KodeOpd           string `json:"kode_opd" validate:"required"`
	RealisasiAnggaran int    `json:"realisasi_anggaran"`
	FaktorPendorong   string `json:"faktor_pendorong"`
	FaktorPenghambat  string `json:"faktor_penghambat"`
	FileUrl           string `json:"file_url"`
}

type RealisasiRequest struct {
	IdPokin          int    `json:"id_pokin" validate:"required"`
	KodeTim          string `json:"kode_tim" validate:"required"`
	JenisPohon       string `json:"jenis_pohon" validate:"required"`
	JenisItem        string `json:"jenis_item" validate:"required"`
	KodeItem         string `json:"kode_item" validate:"required"`
	NamaItem         string `json:"nama_item" validate:"required"`
	Pagu             int    `json:"pagu" validate:"gte=0"`
	Realisasi        int    `json:"realisasi" validate:"gte=0"`
	FaktorPendorong  string `json:"faktor_pendorong"`
	FaktorPenghambat string `json:"faktor_penghambat"`
	Rtl              string `json:"rtl"`
	UrlBuktiDukung   string `json:"url_bukti_dukung"`
	Tahun            string `json:"tahun" validate:"required"`
	KodeOpd          string `json:"kode_opd" validate:"required"`
}
