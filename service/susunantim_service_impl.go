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

type SusunanTimServiceImpl struct {
	SusunanTimRepository repository.SusunanTimRepository
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewSusunanTimServiceImpl(susunanTimRepository repository.SusunanTimRepository, db *sql.DB, validator *validator.Validate) *SusunanTimServiceImpl {
	return &SusunanTimServiceImpl{
		SusunanTimRepository: susunanTimRepository,
		DB:                   db,
		Validator:            validator,
	}
}

func (service *SusunanTimServiceImpl) Create(ctx context.Context, susunanTim web.SusunanTimCreateRequest) (web.SusunanTimResponse, error) {
	if err := service.Validator.Struct(susunanTim); err != nil {

		validationErrors := err.(validator.ValidationErrors)

		fieldErrors := make(map[string]string)
		for _, fe := range validationErrors {
			fieldErrors[fe.Field()] = fe.Tag()
		}

		return web.SusunanTimResponse{}, &web.ValidationError{
			Message: "INVALID_FIELD",
			Fields:  fieldErrors,
		}
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}
	defer helper.NewCommitOrRollback(tx, &err)

	susunanTimDomain := domain.SusunanTim{
		KodeTim:        susunanTim.KodeTim,
		PegawaiId:      susunanTim.PegawaiId,
		NamaPegawai:    susunanTim.NamaPegawai,
		IdJabatanTim:   susunanTim.IdJabatanTim,
		NamaJabatanTim: susunanTim.NamaJabatanTim,
		IsActive:       susunanTim.IsActive,
		Keterangan:     &susunanTim.Keterangan,
		Bulan:          susunanTim.Bulan,
		Tahun:          susunanTim.Tahun,
	}

	susunanTimDomain, err = service.SusunanTimRepository.Create(ctx, tx, susunanTimDomain)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}

	res := web.SusunanTimResponse{
		Id:             susunanTimDomain.Id,
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		IdJabatanTim:   susunanTim.IdJabatanTim,
		NamaPegawai:    susunanTimDomain.NamaPegawai,
		NamaJabatanTim: susunanTimDomain.NamaJabatanTim,
		IsActive:       susunanTimDomain.IsActive,
		Keterangan:     susunanTimDomain.Keterangan,
	}

	return res, nil
}

func (service *SusunanTimServiceImpl) Update(ctx context.Context, susunanTim web.SusunanTimUpdateRequest) (web.SusunanTimResponse, error) {
	err := service.Validator.Struct(susunanTim)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	susunanTimDomain := domain.SusunanTim{
		Id:             susunanTim.Id,
		KodeTim:        susunanTim.KodeTim,
		PegawaiId:      susunanTim.PegawaiId,
		IdJabatanTim:   susunanTim.IdJabatanTim,
		NamaPegawai:    susunanTim.NamaPegawai,
		NamaJabatanTim: susunanTim.NamaJabatanTim,
		IsActive:       susunanTim.IsActive,
		Keterangan:     &susunanTim.Keterangan,
		Bulan:          susunanTim.Bulan,
		Tahun:          susunanTim.Tahun,
	}

	susunanTimDomain, err = service.SusunanTimRepository.Update(ctx, tx, susunanTimDomain)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}

	return web.SusunanTimResponse{
		Id:             susunanTimDomain.Id,
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		NamaPegawai:    susunanTimDomain.NamaPegawai,
		IdJabatanTim:   susunanTim.IdJabatanTim,
		NamaJabatanTim: susunanTimDomain.NamaJabatanTim,
		IsActive:       susunanTimDomain.IsActive,
		Keterangan:     susunanTimDomain.Keterangan,
	}, nil
}

func (service *SusunanTimServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.SusunanTimRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *SusunanTimServiceImpl) FindById(ctx context.Context, id int) (web.SusunanTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	susunanTimDomain, err := service.SusunanTimRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.SusunanTimResponse{}, errors.New("susunan tim not found")
	}

	return web.SusunanTimResponse{
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		NamaPegawai:    susunanTimDomain.NamaPegawai,
		IdJabatanTim:   susunanTimDomain.IdJabatanTim,
		NamaJabatanTim: susunanTimDomain.NamaJabatanTim,
		IsActive:       susunanTimDomain.IsActive,
		Keterangan:     susunanTimDomain.Keterangan,
	}, nil
}

func (service *SusunanTimServiceImpl) FindAll(ctx context.Context) ([]web.SusunanTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.SusunanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	susunanTimDomains, err := service.SusunanTimRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.SusunanTimResponse{}, err
	}

	return helper.ToSusunanTimResponses(susunanTimDomains), nil
}

func (service *SusunanTimServiceImpl) FindByKodeTim(ctx context.Context, kodeTim string) ([]web.SusunanTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.SusunanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	susunanTimDomains, err := service.SusunanTimRepository.FindByKodeTim(ctx, tx, kodeTim)
	if err != nil {
		return []web.SusunanTimResponse{}, err
	}

	return helper.ToSusunanTimResponses(susunanTimDomains), nil
}
