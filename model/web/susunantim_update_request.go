package web

type SusunanTimUpdateRequest struct {
	Id             int    `json:"id"`
	KodeTim        string `json:"kode_tim"`
	PegawaiId      string `json:"nip"`
	NamaJabatanTim string `json:"nama_jabatan_tim"`
	IsActive       bool   `json:"is_active"`
	Keterangan     string `json:"keterangan"`
}
