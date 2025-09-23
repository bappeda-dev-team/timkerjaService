package web

type JabatanTimCreateRequest struct {
	NamaJabatan  string `json:"nama_jabatan"`
	LevelJabatan int    `json:"level_jabatan"`
}
