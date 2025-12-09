package service

import (
	"context"
	"database/sql"
	"errors"
	"timkerjaService/helper"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
	"timkerjaService/repository"

	"github.com/go-playground/validator/v10"
)

type PenilaianKinerjaServiceImpl struct {
	DB                         *sql.DB
	PenilaianKinerjaRepository repository.PenilaianKinerjaRepository
	Validator                  *validator.Validate
}

func NewPenilaianKinerjaServiceImpl(db *sql.DB, penilaianRepo repository.PenilaianKinerjaRepository, validator *validator.Validate) *PenilaianKinerjaServiceImpl {
	return &PenilaianKinerjaServiceImpl{
		DB:                         db,
		PenilaianKinerjaRepository: penilaianRepo,
		Validator:                  validator,
	}
}

func (s *PenilaianKinerjaServiceImpl) Create(ctx context.Context, req web.PenilaianKinerjaRequest) (web.PenilaianKinerjaResponse, error) {
	err := s.Validator.Struct(req)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	domain := domain.PenilaianKinerja{
		IdPegawai:    req.IdPegawai,
		KodeTim:      req.KodeTim,
		JenisNilai:   req.JenisNilai,
		NilaiKinerja: req.NilaiKinerja,
		Tahun:        req.Tahun,
		Bulan:        req.Bulan,
		KodeOpd:      req.KodeOpd,
		CreatedBy:    "admin_test", // TODO: get from context
	}

	res, err := s.PenilaianKinerjaRepository.Create(ctx, tx, domain)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}

	return web.PenilaianKinerjaResponse{
		Id:           res.Id,
		IdPegawai:    res.IdPegawai,
		KodeTim:      res.KodeTim,
		JenisNilai:   res.JenisNilai,
		NilaiKinerja: res.NilaiKinerja,
		Tahun:        res.Tahun,
		Bulan:        res.Bulan,
		KodeOpd:      res.KodeOpd,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
		CreatedBy:    res.CreatedBy,
	}, nil
}

func (s *PenilaianKinerjaServiceImpl) Update(ctx context.Context, req web.PenilaianKinerjaRequest, id int) (web.PenilaianKinerjaResponse, error) {
	err := s.Validator.Struct(req)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	exist, err := s.PenilaianKinerjaRepository.ExistById(ctx, tx, id)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}
	if exist == false {
		return web.PenilaianKinerjaResponse{}, errors.New("id penilaian tidak ditemukan")
	}

	domain := domain.PenilaianKinerja{
		IdPegawai:    req.IdPegawai,
		KodeTim:      req.KodeTim,
		JenisNilai:   req.JenisNilai,
		NilaiKinerja: req.NilaiKinerja,
		Tahun:        req.Tahun,
		Bulan:        req.Bulan,
		KodeOpd:      req.KodeOpd,
		CreatedBy:    "admin_test", // TODO: get from context
	}

	res, err := s.PenilaianKinerjaRepository.Update(ctx, tx, domain, id)
	if err != nil {
		return web.PenilaianKinerjaResponse{}, err
	}

	return web.PenilaianKinerjaResponse{
		Id:           res.Id,
		IdPegawai:    res.IdPegawai,
		KodeTim:      res.KodeTim,
		JenisNilai:   res.JenisNilai,
		NilaiKinerja: res.NilaiKinerja,
		Tahun:        res.Tahun,
		Bulan:        res.Bulan,
		KodeOpd:      res.KodeOpd,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
		CreatedBy:    res.CreatedBy,
	}, nil
}
