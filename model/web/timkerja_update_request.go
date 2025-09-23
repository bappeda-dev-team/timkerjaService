package web

type TimKerjaUpdateRequest struct {
	Id         int    `json:"id"`
	KodeTim    string `json:"kode_tim"`
	NamaTim    string `json:"nama_tim"`
	Keterangan string `json:"keterangan"`
	Tahun      string `json:"tahun"`
	IsActive   bool   `json:"is_active"`
}
