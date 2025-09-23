package web

type JabatanTimUpdateRequest struct {
	Id           int    `json:"id"`
	NamaJabatan  string `json:"nama_jabatan"`
	LevelJabatan int    `json:"level_jabatan"`
}
