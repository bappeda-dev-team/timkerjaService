package web

type TimKerjaCreateRequest struct {
	KodeTim    string `json:"kode_tim"`
	NamaTim    string `json:"nama_tim"`
	Keterangan string `json:"keterangan"`
	Tahun      string `json:"tahun"`
	IsActive   bool   `json:"is_active"`
}
