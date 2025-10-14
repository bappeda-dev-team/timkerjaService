package web

import (
	"time"
	"timkerjaService/internal"
)

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
	Id                  int                               `json:"id"`
	KodeTim             string                            `json:"kode_tim"`
	IdProgramUnggulan   int                               `json:"id_program_unggulan"`
	KodeProgramUnggulan string                            `json:"kode_program_unggulan"`
	ProgramUnggulan     string                            `json:"program_unggulan"`
	Tahun               string                            `json:"tahun"`
	KodeOpd             string                            `json:"kode_opd"`
	Pokin               []internal.TaggingPohonKinerjaItem `json:"pohon_kinerja"`
}

type ProgramUnggulanFullResponse struct {
	Id                int    `json:"id"`
	KodeTim           string `json:"kode_tim"`
	IdProgramUnggulan int    `json:"id_program_unggulan"`
	ProgramUnggulan   string `json:"program_unggulan"`
	IdPokin           string `json:"id_pokin"`
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
	Id               int                            `json:"id"`
	KodeTim          string                         `json:"kode_tim"`
	IdRencanaKinerja string                         `json:"id_rencana_kinerja"`
	IdPegawai        string                         `json:"id_pegawai"`
	RencanaKinerja   string                         `json:"rencana_kinerja"`
	Tahun            string                         `json:"tahun"`
	KodeOpd          string                         `json:"kode_opd"`
	IdPohon          int                            `json:"id_pohon,omitempty"`
	NamaPohon        string                         `json:"nama_pohon,omitempty"`
	Indikator        []internal.IndikatorResponse   `json:"indikators,omitempty"`
	SubKegiatan      []internal.SubKegiatanResponse `json:"subkegiatan,omitempty"`
}

type RealisasiResponse struct {
	Id               int    `json:"id"`
	IdPokin          int    `json:"id_pokin"`
	KodeTim          string `json:"kode_tim"`
	JenisPohon       string `json:"jenis_pohon"`
	JenisItem        string `json:"jenis_item"`
	KodeItem         string `json:"kode_item"`
	NamaItem         string `json:"nama_item"`
	Pagu             int    `json:"pagu"`
	Realisasi        int    `json:"realisasi"`
	FaktorPendorong  string `json:"faktor_pendorong"`
	FaktorPenghambat string `json:"faktor_penghambat"`
	Rtl              string `json:"rtl"`
	UrlBuktiDukung   string `json:"url_bukti_dukung"`
	Tahun            string `json:"tahun"`
	KodeOpd          string `json:"kode_opd"`
}
