package repository

import (
	"context"
	"database/sql"
	"errors"
	"timkerjaService/model/domain"
)

type PenilaianKinerjaRepositoryImpl struct {
}

func NewPenilaianKinerjaRepositoryImpl() *PenilaianKinerjaRepositoryImpl {
	return &PenilaianKinerjaRepositoryImpl{}
}

func (repo *PenilaianKinerjaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, penilaian domain.PenilaianKinerja) (domain.PenilaianKinerja, error) {
	return domain.PenilaianKinerja{}, errors.New("Create Unimplemented :(")
}

func (repo *PenilaianKinerjaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, penilaian domain.PenilaianKinerja, id int) (domain.PenilaianKinerja, error) {
	return domain.PenilaianKinerja{}, errors.New("Update Unimplemented :(")
}

func (repo *PenilaianKinerjaRepositoryImpl) ExistById(ctx context.Context, tx *sql.Tx, id int) (bool, error) {
	return false, errors.New("Exist by id Not implemented")
}
