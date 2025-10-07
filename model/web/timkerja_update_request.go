package web

type TimKerjaUpdateRequest struct {
	Id            int    `json:"id"`
	NamaTim       string `json:"nama_tim"`
	Keterangan    string `json:"keterangan"`
	Tahun         string `json:"tahun"`
	IsActive      bool   `json:"is_active"`
	IsSekretariat bool   `json:"is_sekretariat"`
}
