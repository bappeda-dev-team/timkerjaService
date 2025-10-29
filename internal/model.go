package internal

type DataRincianKerjaWrapper struct {
	Code           int                `json:"code"`
	Status         string             `json:"status"`
	RencanaKinerja []DataRincianKerja `json:"rencana_kinerja"`
}

type DataRincianKerja struct {
	RencanaKinerja RencanaKinerjaResponse `json:"rencana_kinerja"`
	SubKegiatan    []SubKegiatanResponse  `json:"subkegiatan"`
}

type RencanaKinerjaResponse struct {
	IdRencanaKinerja     string              `json:"id_rencana_kinerja"`
	IdPohon              int                 `json:"id_pohon"`
	NamaPohon            string              `json:"nama_pohon"`
	NamaRencanaKinerja   string              `json:"nama_rencana_kinerja"`
	Tahun                string              `json:"tahun"`
	StatusRencanaKinerja string              `json:"status_rencana_kinerja"`
	OperasionalDaerah    OperasionalDaerah   `json:"operasional_daerah"`
	PegawaiId            string              `json:"pegawai_id"`
	NamaPegawai          string              `json:"nama_pegawai"`
	Indikator            []IndikatorResponse `json:"indikator"`
}

type OperasionalDaerah struct {
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
}

type IndikatorResponse struct {
	IdIndikator      string           `json:"id_indikator"`
	RencanaKinerjaId string           `json:"rencana_kinerja_id"`
	NamaIndikator    string           `json:"nama_indikator"`
	Targets          []TargetResponse `json:"targets"`
}

type TargetResponse struct {
	IdTarget    string `json:"id_target"`
	IndikatorId string `json:"indikator_id"`
	Target      string `json:"target"`
	Satuan      string `json:"satuan"`
}

type SubKegiatanResponse struct {
	SubKegiatanTerpilihId string                    `json:"subkegiatanterpilih_id"`
	Id                    string                    `json:"id"`
	RekinId               string                    `json:"rekin_id"`
	KodeSubKegiatan       string                    `json:"kode_subkegiatan"`
	NamaSubKegiatan       string                    `json:"nama_sub_kegiatan"`
	PaguSubKegiatan       []PaguSubKegiatanResponse `json:"pagu,omitempty"`
}

// DATA HELPER
type IndikatorSubKegiatanResponse struct {
	Id            string `json:"id"`
	SubKegiatanId string `json:"sub_kegiatan_id"`
	NamaIndikator string `json:"indikator"`
}

type PaguSubKegiatanResponse struct {
	Id            string `json:"id"`
	SubKegiatanId string `json:"sub_kegiatan_id"`
	JenisPagu     string `json:"jenis"`
	PaguAnggaran  int    `json:"pagu_anggaran"`
	Tahun         string `json:"tahun"`
}

// deprecated
type ProgramUnggulanResponse struct {
	Id                  int    `json:"id"`
	KodeProgramUnggulan string `json:"kode_program_unggulan"`
	NamaProgramUnggulan string `json:"nama_program_unggulan"`
	RencanaImplementasi string `json:"rencana_implementasi"`
	Keterangan          string `json:"keterangan"`
	TahunAwal           string `json:"tahun_awal"`
	TahunAkhir          string `json:"tahun_akhir"`
	IsActive            bool   `json:"is_active"`
}

type LaporanTaggingPohonKinerjaResponse struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    []TaggingPohonKinerjaItem `json:"data"`
}

type TaggingPohonKinerjaItem struct {
	KodeProgramUnggulan string             `json:"kode_program_unggulan"`
	NamaProgramUnggulan string             `json:"nama_program_unggulan"`
	RencanaImplementasi string             `json:"rencana_implementasi"`
	IdTagging           int                `json:"id_tagging"`
	IdPohon             int                `json:"id_pohon"`
	Tahun               int                `json:"tahun"`
	NamaPohon           string             `json:"nama_pohon"`
	KodeOpd             string             `json:"kode_opd"`
	NamaOpd             string             `json:"nama_opd"`
	JenisPohon          string             `json:"jenis_pohon"`
	KeteranganTagging   string             `json:"keterangan_tagging"`
	Status              string             `json:"status"`
	Pelaksanas          []PelaksanaPokin   `json:"pelaksanas"`
	Keterangan          string             `json:"keterangan"`
	Indikator           []TaggingIndikator `json:"indikator"`
}

type PelaksanaPokin struct {
	NamaPelaksana   string              `json:"nama_pelaksana"`
	NIPPelaksana    string              `json:"nip_pelaksana"`
	RencanaKinerjas []RencanaKinerjaAsn `json:"rencana_kinerjas"`
}

type RencanaKinerjaAsn struct {
	IdRekin            string           `json:"id_rekin"`
	RencanaKinerja     string           `json:"rencana_kinerja"`
	NamaPelaksana      string           `json:"nama_pelaksana"`
	NIPPelaksana       string           `json:"nip_pelaksana"`
	KodeBidangUrusan   string           `json:"kode_bidang_urusan,omitempty"`
	NamaBidangUrusan   string           `json:"nama_bidang_urusan,omitempty"`
	KodeProgram        string           `json:"kode_program,omitempty"`
	NamaProgram        string           `json:"nama_program,omitempty"`
	KodeSubkegiatan    string           `json:"kode_subkegiatan,omitempty"`
	NamaSubkegiatan    string           `json:"nama_subkegiatan,omitempty"`
	Pagu               int              `json:"pagu"`
	Catatan            string           `json:"keterangan"`
	TahapanPelaksanaan WaktuPelaksanaan `json:"tahapan_pelaksanaan"`
}

type WaktuPelaksanaan struct {
	Tw1 int `json:"tw_1"`
	Tw2 int `json:"tw_2"`
	Tw3 int `json:"tw_3"`
	Tw4 int `json:"tw_4"`
}

type TaggingIndikator struct {
	IdIndikator   string              `json:"id_indikator"`
	NamaIndikator string              `json:"nama_indikator"`
	Kode          string              `json:"kode"`
	Targets       []TaggingTargetItem `json:"targets"`
}

type TaggingTargetItem struct {
	IdTarget    string `json:"id_target"`
	IndikatorId string `json:"indikator_id"`
	Target      string `json:"target"`
	Satuan      string `json:"satuan"`
	Tahun       int    `json:"tahun"`
}

type RencanaAksiResponse struct {
	Id                     string                           `json:"id"`
	RencanaKinerjaId       string                           `json:"rekin_id"`
	KodeOpd                string                           `json:"kode_opd,omitempty"`
	Urutan                 int                              `json:"urutan"`
	NamaRencanaAksi        string                           `json:"nama_rencana_aksi"`
	PelaksanaanRencanaAksi []PelaksanaanRencanaAksiResponse `json:"pelaksanaan"`
	JumlahBobot            int                              `json:"jumlah_bobot,omitempty"`
	TotalBobotRencanaAksi  int                              `json:"total_bobot_rencana_aksi,omitempty"`
}

type BobotBulanan struct {
	Bulan      int `json:"bulan"`
	TotalBobot int `json:"total_bobot"`
}

type RencanaAksiTableResponse struct {
	RencanaAksi      []RencanaAksiResponse `json:"rencana_aksi"`
	TotalPerBulan    []BobotBulanan        `json:"total_per_bulan"`
	TotalKeseluruhan int                   `json:"total_keseluruhan"`
	WaktuDibutuhkan  int                   `json:"waktu_dibutuhkan"`
}

type PelaksanaanRencanaAksiResponse struct {
	Id            string `json:"id"`
	RencanaAksiId string `json:"rencana_aksi_id"`
	Bulan         int    `json:"bulan"`
	Bobot         int    `json:"bobot"`
	BobotAvail    int    `json:"bobot_tersedia,omitempty"`
}

type FindByIdTerkaitRequest struct {
	Ids []int `json:"id_prorgramunggulan" validate:"required,min=1"`
}

type FindByKodeProgramUnggulansRequest struct {
	KodeProgramUnggulan []string `json:"kode_program_unggulan" validate:"required,min=1"`
}
