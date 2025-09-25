package web

type TimKerjaCreateRequest struct {
	NamaTim    string `json:"nama_tim"`
	Keterangan string `json:"keterangan"`
	Tahun      string `json:"tahun"`
	IsActive   bool   `json:"is_active"`
}
