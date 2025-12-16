package web

type PetugasTimCreateRequest struct {
	IdProgramUnggulan int    `json:"id_program_unggulan" validate:"required"`
	KodeTim           string `json:"kode_tim" validate:"required"`
	PegawaiId         string `json:"pegawai_id" validate:"required"`
	Tahun             int    `json:"tahun" validate:"required"`
	Bulan             int    `json:"bulan"`
}
