package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
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
	result, err := s.PenilaianKinerjaRepository.FindByTahunBulan(ctx, tx, tahun, bulan)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// ==========================
	// GROUPING JENIS NILAI
	// ==========================

	responses := make([]web.LaporanPenilaianKinerjaResponse, 0)

	for _, laporan := range result {

		groupMap := make(map[string]*web.PenilaianGroupedResponse)

		for _, p := range laporan.Penilaians {

			// unique key per pegawai per bulan
			key := p.IdPegawai + "_" + p.Tahun + "_" + strconv.Itoa(p.Bulan)

			item, exists := groupMap[key]
			if !exists {
				item = &web.PenilaianGroupedResponse{
					IdPegawai:      p.IdPegawai,
					NamaPegawai:    p.NamaPegawai,
					NamaJabatanTim: p.NamaJabatanTim,
					KodeTim:        p.KodeTim,
					Tahun:          p.Tahun,
					Bulan:          p.Bulan,
				}
				groupMap[key] = item
			}

			// Masukkan ke field sesuai jenis nilai
			switch p.JenisNilai {
			case "KINERJA_BAPPEDA":
				item.KinerjaBappeda = p.NilaiKinerja
			case "KINERJA_TIM":
				item.KinerjaTim = p.NilaiKinerja
			case "KINERJA_PERSON":
				item.KinerjaPerson = p.NilaiKinerja
			}
		}

		// convert map → slice
		groupedList := make([]web.PenilaianGroupedResponse, 0, len(groupMap))
		for _, v := range groupMap {
			v.NilaiAkhir = hitungNilaiAkhir(*v)
			groupedList = append(groupedList, *v)
		}

		responses = append(responses, web.LaporanPenilaianKinerjaResponse{
			NamaTim:           laporan.NamaTim,
			KodeTim:           laporan.KodeTim,
			IsSekretariat:     laporan.IsSekretariat,
			PenilaianKinerjas: groupedList,
		})
	}

	return responses, nil
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

func hitungNilaiAkhir(item web.PenilaianGroupedResponse) float32 {
	xs := []float32{}

	if item.KinerjaBappeda > 0 {
		xs = append(xs, float32(item.KinerjaBappeda))
	}
	if item.KinerjaTim > 0 {
		xs = append(xs, float32(item.KinerjaTim))
	}
	if item.KinerjaPerson > 0 {
		xs = append(xs, float32(item.KinerjaPerson))
	}

	if len(xs) == 0 {
		return 0
	}

	return average(xs)
}

func average(xs []float32) float32 {
	var total float32
	for _, v := range xs {
		total += v
	}
	return total / float32(len(xs))
}
