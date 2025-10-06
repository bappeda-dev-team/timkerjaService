package web

type TimKerjaCreateRequest struct {
	NamaTim    string `json:"nama_tim"`
	Keterangan string `json:"keterangan"`
	Tahun      string `json:"tahun"`
	IsActive   bool   `json:"is_active"`
}

type ProgramUnggulanTimKerjaRequest struct {
	KodeTim           string
	IdProgramUnggulan int    `json:"id_program_unggulan"`
	Tahun             string `json:"tahun"`
	KodeOpd           string `json:"kode_opd"`
	RealisasiAnggaran int    `json:"realisasi_anggaran"`
	FaktorPendorong   string `json:"faktor_pendorong"`
	FaktorPenghambat  string `json:"faktor_penghambat"`
	FileUrl           string `json:"file_url"`
}
