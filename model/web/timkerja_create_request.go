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
