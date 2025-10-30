package service

import (
	"context"
	"database/sql"
	"errors"
	"timkerjaService/helper"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
	"timkerjaService/repository"
)

type RealisasiAnggaranServiceImpl struct {
	DB                          *sql.DB
	RealisasiAnggaranRepository repository.RealisasiAnggaranRepository
}

func NewRealisasiAnggaranServiceImpl(db *sql.DB, repo repository.RealisasiAnggaranRepository) *RealisasiAnggaranServiceImpl {
	return &RealisasiAnggaranServiceImpl{
		DB:                          db,
		RealisasiAnggaranRepository: repo,
	}
}

func (s *RealisasiAnggaranServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	return s.RealisasiAnggaranRepository.Delete(ctx, tx, id)
}

func (s *RealisasiAnggaranServiceImpl) FindById(ctx context.Context, id int) (web.RealisasiAnggaranResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RealisasiAnggaranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	d, err := s.RealisasiAnggaranRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.RealisasiAnggaranResponse{}, errors.New("realisasi anggaran not found")
	}

	return helper.ToRealisasiAnggaranResponse(d), nil
}

func (s *RealisasiAnggaranServiceImpl) FindAll(ctx context.Context, kodeSubkegiatan string, bulan string, tahun string) ([]web.RealisasiAnggaranResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	list, err := s.RealisasiAnggaranRepository.FindAll(ctx, tx, kodeSubkegiatan, bulan, tahun)
	if err != nil {
		return nil, err
	}

	return helper.ToRealisasiAnggaranResponses(list), nil
}

func (s *RealisasiAnggaranServiceImpl) Upsert(ctx context.Context, req web.RealisasiAnggaranCreateRequest) (web.RealisasiAnggaranResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.RealisasiAnggaranResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	d := domain.RealisasiAnggaran{
		KodeSubkegiatan:   req.KodeSubkegiatan,
		RealisasiAnggaran: req.RealisasiAnggaran,
		KodeOpd:           req.KodeOpd,
		RencanaAksi:       req.RencanaAksi,
		FaktorPendorong:   req.FaktorPendorong,
		FaktorPenghambat:  req.FaktorPenghambat,
		RekomendasiTl:     req.RekomendasiTl,
		BuktiDukung:       req.BuktiDukung,
		Bulan:             req.Bulan,
		Tahun:             req.Tahun,
	}

	ra, err := s.RealisasiAnggaranRepository.Upsert(ctx, tx, d)
	if err != nil {
		return web.RealisasiAnggaranResponse{}, err
	}

	// opsional: ambil kembali via FindAll/FindById untuk isi CreatedAt/UpdatedAt
	return helper.ToRealisasiAnggaranResponse(ra), nil
}
