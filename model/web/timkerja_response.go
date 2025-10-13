package web

import "time"

type TimKerjaResponse struct {
	Id            int       `json:"id"`
	KodeTim       string    `json:"kode_tim"`
	NamaTim       string    `json:"nama_tim"`
	Keterangan    string    `json:"keterangan"`
	Tahun         string    `json:"tahun"`
	IsActive      bool      `json:"is_active"`
	IsSekretariat bool      `json:"is_sekretariat"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProgramUnggulanTimKerjaResponse struct {
	Id                int    `json:"id"`
	KodeTim           string `json:"kode_tim"`
	IdProgramUnggulan int    `json:"id_program_unggulan"`
	ProgramUnggulan   string `json:"program_unggulan"`
	Tahun             string `json:"tahun"`
	KodeOpd           string `json:"kode_opd"`
}

type ProgramUnggulanFullResponse struct {
	Id                int    `json:"id"`
	KodeTim           string `json:"kode_tim"`
	IdProgramUnggulan int    `json:"id_program_unggulan"`
	ProgramUnggulan   string `json:"program_unggulan"`
	PohonKinerja      string `json:"pohon_kinerja"`
	LevelPohon        string `json:"level_pohon"`
	Tahun             string `json:"tahun"`
	KodeOpd           string `json:"kode_opd"`
}

type IndikatorPohon struct {
	IdIndikator string            `json:"id_indikator"`
	IdPokin     string            `json:"id_pokin,omitempty"`
	Indikator   string            `json:"nama_indikator"`
	Kode        string            `json:"kode"`
	Target      []TargetIndikator `json:"targets"`
}

type TargetIndikator struct {
	IdTarget    string `json:"id_target"`
	IndikatorId string `json:"indikator_id"`
	Target      string `json:"target"`
	Satuan      string `json:"satuan"`
	Tahun       int    `json:"tahun,omitempty"`
}

type RencanaKinerjaTimKerjaResponse struct {
	Id              int    `json:"id"`
	KodeTim         string `json:"kode_tim"`
	IdRencanKinerja string    `json:"id_rencana_kinerja"`
	RencanaKinerja  string `json:"rencana_kinerja"`
	Tahun           string `json:"tahun"`
	KodeOpd         string `json:"kode_opd"`
}
