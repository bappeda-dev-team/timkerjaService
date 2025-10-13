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

type ProgramUnggulanResponse struct {
	Id                        int    `json:"id"`
	KodeProgramUnggulan       string `json:"kode_program_unggulan"`
	NamaTagging               string `json:"nama_program_unggulan"`
	KeteranganProgramUnggulan string `json:"rencana_implementasi"`
	Keterangan                string `json:"keterangan"`
	TahunAwal                 string `json:"tahun_awal"`
	TahunAkhir                string `json:"tahun_akhir"`
	IsActive                  bool   `json:"is_active"`
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
