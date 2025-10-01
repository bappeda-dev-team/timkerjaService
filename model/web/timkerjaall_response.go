package web

type TimKerjaDetailResponse struct {
	Id          int                        `json:"id"`
	KodeTim     string                     `json:"kode_tim"`
	NamaTim     string                     `json:"nama_tim"`
	Keterangan  string                     `json:"keterangan"`
	SusunanTims []SusunanTimDetailResponse `json:"susunan_tims"`
}

type SusunanTimDetailResponse struct {
	Id           int     `json:"id"`
	PegawaiId    string  `json:"nip"`
	NamaPegawai  string  `json:"nama_pegawai"`
	NamaJabatan  string  `json:"nama_jabatan"`
	LevelJabatan int     `json:"level_jabatan"`
	Keterangan   *string `json:"keterangan"`
	IsActive     bool    `json:"is_active"`
}
