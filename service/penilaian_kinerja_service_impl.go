package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	"timkerjaService/helper"
	"timkerjaService/internal"
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

func (s *PenilaianKinerjaServiceImpl) All(
	ctx context.Context,
	tahun int,
	bulan int,
) ([]web.LaporanPenilaianKinerjaResponse, error) {

	if tahun <= 0 || bulan <= 0 || bulan > 12 {
		return nil, errors.New("tahun atau bulan tidak valid")
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Ambil data RAW dari repository (belum grouped jenis nilai)
	penilaianKinerja, err := s.PenilaianKinerjaRepository.FindByTahunBulan(ctx, tx, tahun, bulan)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	kepegawaianHost := os.Getenv("PERENCANAAN_HOST")
	if kepegawaianHost == "" {
		log.Println("PERENCANAAN_HOST belum diatur — skip merge eksternal")
	}

	kepegawaianClient := internal.NewKepegawaianClient(
		kepegawaianHost,
		&http.Client{Timeout: 25 * time.Second},
	)

	// gabung dengan api internal tim
	merged, err := helper.MergePenilaianKinerjaParallel(ctx, penilaianKinerja, kepegawaianClient, 5)
	if err != nil {
		return nil, err
	}

	return merged, nil
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

func (s *PenilaianKinerjaServiceImpl) TppPegawaiAll(
	ctx context.Context,
	tahun int,
	bulan int,
) ([]web.LaporanPenilaianKinerjaResponse, error) {

	if tahun <= 0 || bulan <= 0 || bulan > 12 {
		return nil, errors.New("tahun atau bulan tidak valid")
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Ambil data RAW dari repository (belum grouped jenis nilai)
	penilaianKinerja, err := s.PenilaianKinerjaRepository.FindByTahunBulan(ctx, tx, tahun, bulan)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	kepegawaianHost := os.Getenv("PERENCANAAN_HOST")
	if kepegawaianHost == "" {
		log.Println("PERENCANAAN_HOST belum diatur — skip merge eksternal")
	}

	kepegawaianClient := internal.NewKepegawaianClient(
		kepegawaianHost,
		&http.Client{Timeout: 25 * time.Second},
	)

	// gabung dengan api internal tim
	merged, err := helper.MergePenilaianKinerjaParallel(ctx, penilaianKinerja, kepegawaianClient, 5)
	if err != nil {
		return nil, err
	}
	// TODO: merge baris diatas dengan method All

	// convert tambah tpp dan perhitungan
	result := helper.ConvertToTppPegawaiResponse(merged)

	return result, nil
}
