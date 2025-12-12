package web

import (
	"time"
)

type PenilaianKinerjaResponse struct {
	Id             int       `json:"id"`
	IdPegawai      string    `json:"id_pegawai" validate:"required"`
	NamaPegawai    string    `json:"nama_pegawai"`
	NamaJabatanTim string    `json:"nama_jabatan_tim"`
	KodeTim        string    `json:"kode_tim" validate:"required"`
	JenisNilai     string    `json:"jenis_nilai" validate:"required"`
	NilaiKinerja   int       `json:"nilai_kinerja" validate:"required"`
	Tahun          string    `json:"tahun" validate:"required"`
	Bulan          int       `json:"bulan" validate:"required"`
	KodeOpd        string    `json:"kode_opd" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      string    `json:"created_by"`
}

type PenilaianGroupedBase struct {
	IdPegawai       string `json:"id_pegawai"`
	NamaPegawai     string `json:"nama_pegawai"`
	SusunanTimId    int    `json:"-"`
	LevelJabatanTim int    `json:"level_jabatan_tim"`
	NamaJabatanTim  string `json:"nama_jabatan_tim"`
	Pangkat         string `json:"pangkat"`
	Golongan        string `json:"golongan"`
	JenisJabatan    string `json:"jenis_jabatan"`
	KodeTim         string `json:"kode_tim"`
	Tahun           string `json:"tahun"`
	Bulan           int    `json:"bulan"`

	KinerjaBappeda int `json:"kinerja_bappeda"`
	KinerjaTim     int `json:"kinerja_tim"`
	KinerjaPerson  int `json:"kinerja_person"`
	NilaiAkhir     int `json:"nilai_akhir"`
}

type PenilaianTppExtension struct {
	TppBasic             int     `json:"tpp_basic"`
	PersentasePenerimaan string  `json:"persentase_penerimaan"`
	JumlahKotor          int     `json:"jumlah_kotor"`
	Pajak                float64 `json:"pajak"`
	JumlahPajak          int     `json:"jumlah_pajak"`
	PotonganBPJS         float64 `json:"potongan_bpjs"`
	JumlahBersih         int     `json:"jumlah_bersih"`
}

type PenilaianGroupedResponse struct {
	PenilaianGroupedBase
	Tpp *PenilaianTppExtension `json:"tpp_pegawai,omitempty"`
}

type LaporanPenilaianKinerjaResponse struct {
	NamaTim           string                     `json:"nama_tim" validate:"required"`
	KodeTim           string                     `json:"kode_tim" validate:"required"`
	IsSekretariat     bool                       `json:"is_sekretariat"`
	PenilaianKinerjas []PenilaianGroupedResponse `json:"penilaian_kinerjas"`
}
