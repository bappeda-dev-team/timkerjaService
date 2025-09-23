package service

import (
	"context"
	"database/sql"
	"timkerjaService/helper"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
	"timkerjaService/repository"

	"github.com/go-playground/validator/v10"
)

type TimKerjaServiceImpl struct {
	TimKerjaRepository repository.TimKerjaRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewTimKerjaServiceImpl(timKerjaRepository repository.TimKerjaRepository, db *sql.DB, validator *validator.Validate) *TimKerjaServiceImpl {
	return &TimKerjaServiceImpl{
		TimKerjaRepository: timKerjaRepository,
		DB:                 db,
		Validator:          validator,
	}
}

func (service *TimKerjaServiceImpl) Create(ctx context.Context, timKerja web.TimKerjaCreateRequest) (web.TimKerjaResponse, error) {
	err := service.Validator.Struct(timKerja)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	timKerjaDomain := domain.TimKerja{
		KodeTim:    timKerja.KodeTim,
		NamaTim:    timKerja.NamaTim,
		Keterangan: timKerja.Keterangan,
		Tahun:      timKerja.Tahun,
		IsActive:   timKerja.IsActive,
	}

	timKerjaDomain, err = service.TimKerjaRepository.Create(ctx, tx, timKerjaDomain)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	return web.TimKerjaResponse{
		KodeTim:    timKerjaDomain.KodeTim,
		NamaTim:    timKerjaDomain.NamaTim,
		Keterangan: timKerjaDomain.Keterangan,
		Tahun:      timKerjaDomain.Tahun,
		IsActive:   timKerjaDomain.IsActive,
	}, nil
}

func (service *TimKerjaServiceImpl) Update(ctx context.Context, timKerja web.TimKerjaUpdateRequest) (web.TimKerjaResponse, error) {
	err := service.Validator.Struct(timKerja)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	timKerjaDomain := domain.TimKerja{
		Id: timKerja.Id,
	}

	timKerjaDomain, err = service.TimKerjaRepository.Update(ctx, tx, timKerjaDomain)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	return web.TimKerjaResponse{
		KodeTim:    timKerjaDomain.KodeTim,
		NamaTim:    timKerjaDomain.NamaTim,
		Keterangan: timKerjaDomain.Keterangan,
		Tahun:      timKerjaDomain.Tahun,
		IsActive:   timKerjaDomain.IsActive,
	}, nil
}

func (service *TimKerjaServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.TimKerjaRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *TimKerjaServiceImpl) FindById(ctx context.Context, id int) (web.TimKerjaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	timKerjaDomain, err := service.TimKerjaRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	return web.TimKerjaResponse{
		KodeTim:    timKerjaDomain.KodeTim,
		NamaTim:    timKerjaDomain.NamaTim,
		Keterangan: timKerjaDomain.Keterangan,
		Tahun:      timKerjaDomain.Tahun,
		IsActive:   timKerjaDomain.IsActive,
	}, nil
}

func (service *TimKerjaServiceImpl) FindAll(ctx context.Context) ([]web.TimKerjaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.TimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	timKerjaDomains, err := service.TimKerjaRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.TimKerjaResponse{}, err
	}
	return helper.ToTimKerjaResponses(timKerjaDomains), nil
}

func (service *TimKerjaServiceImpl) FindAllTm(ctx context.Context) ([]web.TimKerjaDetailResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	timKerjaList, susunanTimMap, err := service.TimKerjaRepository.FindAllWithSusunan(ctx, tx)
	if err != nil {
		return nil, err
	}

	var result []web.TimKerjaDetailResponse

	for _, timKerja := range timKerjaList {
		var susunanTimResponses []web.SusunanTimDetailResponse

		// Get susunan tim for this kode_tim
		if susunanTims, exists := susunanTimMap[timKerja.KodeTim]; exists {
			for _, st := range susunanTims {
				susunanTimResponses = append(susunanTimResponses, web.SusunanTimDetailResponse{
					Id:           st.Id,
					PegawaiId:    st.PegawaiId,
					NamaJabatan:  st.NamaJabatanTim,
					LevelJabatan: st.LevelJabatan,
					Keterangan:   st.Keterangan,
					IsActive:     st.IsActive,
				})
			}
		}

		result = append(result, web.TimKerjaDetailResponse{
			Id:          timKerja.Id,
			KodeTim:     timKerja.KodeTim,
			NamaTim:     timKerja.NamaTim,
			SusunanTims: susunanTimResponses,
		})
	}

	return result, nil
}
