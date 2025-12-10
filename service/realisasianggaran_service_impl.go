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

func (s *RealisasiAnggaranServiceImpl) FindAll(ctx context.Context, kodeSubkegiatan string, kodeTim string, idRencanaKinerja string, bulan string, tahun string) ([]web.RealisasiAnggaranResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	list, err := s.RealisasiAnggaranRepository.FindAll(ctx, tx, kodeSubkegiatan, kodeTim, idRencanaKinerja, bulan, tahun)
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

	// rangebulan := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// for _, bulan := range rangebulan {
	// 	if req.Bulan == bulan {
	// 		return web.RealisasiAnggaranResponse{}, errors.New("bulan tidak valid")
	// 	}
	// }
	if req.Bulan < 1 || req.Bulan > 12 {
		return web.RealisasiAnggaranResponse{}, errors.New("bulan tidak valid")
	}

	d := domain.RealisasiAnggaran{
		KodeSubkegiatan:   req.KodeSubkegiatan,
		RealisasiAnggaran: req.RealisasiAnggaran,
		KodeOpd:           req.KodeOpd,
		RencanaAksi:       req.RencanaAksi,
		FaktorPendorong:   req.FaktorPendorong,
		FaktorPenghambat:  req.FaktorPenghambat,
		RekomendasiTl:     req.RekomendasiTl,
		RisikoHukum:       req.RisikoHukum,
		BuktiDukung:       req.BuktiDukung,
		Bulan:             req.Bulan,
		Tahun:             req.Tahun,
		KodeTim:           req.KodeTim,
		IdPohon:           req.IdPohon,
		IdRencanaKinerja:  req.IdRencanaKinerja,
		IdProgramUnggulan: req.IdProgramUnggulan,
	}

	ra, err := s.RealisasiAnggaranRepository.Upsert(ctx, tx, d)
	if err != nil {
		return web.RealisasiAnggaranResponse{}, err
	}

	// opsional: ambil kembali via FindAll/FindById untuk isi CreatedAt/UpdatedAt
	return helper.ToRealisasiAnggaranResponse(ra), nil
}
