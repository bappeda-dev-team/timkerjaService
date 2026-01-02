package web

type CloneSusunanTimRequest struct {
	KodeTim     string `json:"kodeTim" validate:"required"`
	Bulan       int    `json:"bulan" validate:"required"`
	Tahun       int    `json:"tahun" validate:"required"`
	BulanTarget int    `json:"bulanTarget" validate:"required"`
	TahunTarget int    `json:"tahunTarget" validate:"required"`
}
