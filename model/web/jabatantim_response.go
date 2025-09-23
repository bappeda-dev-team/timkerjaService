package web

import "time"

type JabatanTimResponse struct {
	Id           int       `json:"id"`
	LevelJabatan int       `json:"level_jabatan"`
	NamaJabatan  string    `json:"nama_jabatan"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
