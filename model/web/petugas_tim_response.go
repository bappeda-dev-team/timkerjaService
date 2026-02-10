package web

type PetugasTimResponse struct {
	Id          int    `json:"id"`
	PegawaiId   string `json:"pegawai_id"` // nip
	NamaPegawai string `json:"nama_pegawai"`
	NamaTim     string `json:"nama_tim"`
	KodeTim     string `json:"kode_tim"`
}
