package domain

import "time"

type JabatanTim struct {
	Id           int       `json:"id"`
	NamaJabatan  string    `json:"nama_jabatan"`
	LevelJabatan int       `json:"level_jabatan"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
