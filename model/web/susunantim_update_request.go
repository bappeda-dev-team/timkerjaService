package web

type SusunanTimUpdateRequest struct {
	Id             int    `json:"id"`
	KodeTim        string `json:"kode_tim" validate:"required"`
	PegawaiId      string `json:"nip" validate:"required"`
	IdJabatanTim   int    `json:"id_jabatan_tim" validate:"required"`
	NamaPegawai    string `json:"nama_pegawai" validate:"required"`
	NamaJabatanTim string `json:"nama_jabatan_tim" validate:"required"`
	IsActive       bool   `json:"is_active"`
	Keterangan     string `json:"keterangan"`
}
