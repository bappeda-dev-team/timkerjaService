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
		KodeTim:        susunanTim.KodeTim,
		PegawaiId:      susunanTim.PegawaiId,
		NamaPegawai:    susunanTim.NamaPegawai,
		NamaJabatanTim: susunanTim.NamaJabatanTim,
		IsActive:       susunanTim.IsActive,
		Keterangan:     &susunanTim.Keterangan,
	}

	susunanTimDomain, err = service.SusunanTimRepository.Create(ctx, tx, susunanTimDomain)
	if err != nil {
		return web.SusunanTimResponse{}, err
	}

	return web.SusunanTimResponse{
		Id:             susunanTimDomain.Id,
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		NamaPegawai:    susunanTimDomain.NamaPegawai,
		NamaJabatanTim: susunanTimDomain.NamaJabatanTim,
		IsActive:       susunanTimDomain.IsActive,
		Keterangan:     susunanTimDomain.Keterangan,
	}, nil
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
		NamaPegawai:    susunanTim.NamaPegawai,
		NamaJabatanTim: susunanTim.NamaJabatanTim,
		IsActive:       susunanTim.IsActive,
		Keterangan:     &susunanTim.Keterangan,
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
		return web.SusunanTimResponse{}, err
	}

	return web.SusunanTimResponse{
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		NamaPegawai:    susunanTimDomain.NamaPegawai,
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
