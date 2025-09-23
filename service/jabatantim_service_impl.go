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

type JabatanTimServiceImpl struct {
	JabatanTimRepository repository.JabatanTimRepository
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewJabatanTimServiceImpl(jabatanTimRepository repository.JabatanTimRepository, db *sql.DB, validator *validator.Validate) *JabatanTimServiceImpl {
	return &JabatanTimServiceImpl{
		JabatanTimRepository: jabatanTimRepository,
		DB:                   db,
		Validator:            validator,
	}
}

func (service *JabatanTimServiceImpl) Create(ctx context.Context, jabatanTim web.JabatanTimCreateRequest) (web.JabatanTimResponse, error) {
	err := service.Validator.Struct(jabatanTim)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanTimDomain := domain.JabatanTim{
		NamaJabatan:  jabatanTim.NamaJabatan,
		LevelJabatan: jabatanTim.LevelJabatan,
	}

	jabatanTimDomain, err = service.JabatanTimRepository.Create(ctx, tx, jabatanTimDomain)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}

	return web.JabatanTimResponse{
		NamaJabatan:  jabatanTimDomain.NamaJabatan,
		LevelJabatan: jabatanTimDomain.LevelJabatan,
	}, nil
}

func (service *JabatanTimServiceImpl) Update(ctx context.Context, jabatanTim web.JabatanTimUpdateRequest) (web.JabatanTimResponse, error) {
	err := service.Validator.Struct(jabatanTim)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanTimDomain := domain.JabatanTim{
		Id:           jabatanTim.Id,
		NamaJabatan:  jabatanTim.NamaJabatan,
		LevelJabatan: jabatanTim.LevelJabatan,
	}

	jabatanTimDomain, err = service.JabatanTimRepository.Update(ctx, tx, jabatanTimDomain)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}

	return web.JabatanTimResponse{
		Id:           jabatanTimDomain.Id,
		NamaJabatan:  jabatanTimDomain.NamaJabatan,
		LevelJabatan: jabatanTimDomain.LevelJabatan,
	}, nil
}

func (service *JabatanTimServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.JabatanTimRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *JabatanTimServiceImpl) FindById(ctx context.Context, id int) (web.JabatanTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanTimDomain, err := service.JabatanTimRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.JabatanTimResponse{}, err
	}

	return web.JabatanTimResponse{
		Id:           jabatanTimDomain.Id,
		NamaJabatan:  jabatanTimDomain.NamaJabatan,
		LevelJabatan: jabatanTimDomain.LevelJabatan,
	}, nil

}

func (service *JabatanTimServiceImpl) FindAll(ctx context.Context) ([]web.JabatanTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.JabatanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jabatanTimDomains, err := service.JabatanTimRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.JabatanTimResponse{}, err
	}

	return helper.ToJabatanTimResponses(jabatanTimDomains), nil
}
