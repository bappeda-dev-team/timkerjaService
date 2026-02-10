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

type PetugasTimServiceImpl struct {
	PetugasTimRepository repository.PetugasTimRepository
	SusunanTimRepository repository.SusunanTimRepository
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewPetugasTimServiceImpl(petugasTimRepository repository.PetugasTimRepository, susunanTimRepository repository.SusunanTimRepository, db *sql.DB, validator *validator.Validate) *PetugasTimServiceImpl {
	return &PetugasTimServiceImpl{
		PetugasTimRepository: petugasTimRepository,
		SusunanTimRepository: susunanTimRepository,
		DB:                   db,
		Validator:            validator,
	}
}

func (service *PetugasTimServiceImpl) Create(ctx context.Context, petugasTimReq web.PetugasTimCreateRequest) (web.PetugasTimResponse, error) {
	err := service.Validator.Struct(petugasTimReq)
	if err != nil {
		return web.PetugasTimResponse{}, err
	}
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.PetugasTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	petugasTimDomain := domain.PetugasTim{
		IdProgramUnggulan: petugasTimReq.IdProgramUnggulan,
		PegawaiId:         petugasTimReq.PegawaiId,
		KodeTim:           petugasTimReq.KodeTim,
		Tahun:             petugasTimReq.Tahun,
		Bulan:             petugasTimReq.Bulan,
	}

	result, err := service.PetugasTimRepository.Create(ctx, tx, petugasTimDomain)
	if err != nil {
		return web.PetugasTimResponse{}, err
	}

	currentPegawaiId := result.PegawaiId
	namaPegawai, err := service.SusunanTimRepository.FindByIdPegawai(ctx, tx, currentPegawaiId)
	if err != nil {
		return web.PetugasTimResponse{}, err
	}

	return web.PetugasTimResponse{
		Id:          result.Id,
		PegawaiId:   result.PegawaiId,
		NamaPegawai: namaPegawai.NamaPegawai,
	}, nil
}

func (service *PetugasTimServiceImpl) Delete(ctx context.Context, idPetugasTim int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.PetugasTimRepository.Delete(ctx, tx, idPetugasTim)
	if err != nil {
		return err
	}

	return nil
}

func (service *PetugasTimServiceImpl) FindAllByIdProgramUnggulans(ctx context.Context, idProgramUnggulans []int, bulan int, tahun int) (map[int][]web.PetugasTimResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	petugasTims, err := service.PetugasTimRepository.FindAllByIdProgramUnggulans(ctx, tx, idProgramUnggulans, bulan, tahun)
	if err != nil {
		return nil, err
	}
	results := make(map[int][]web.PetugasTimResponse)
	for _, pt := range petugasTims {

		resp := web.PetugasTimResponse{
			Id:          pt.Id,
			PegawaiId:   pt.PegawaiId,
			NamaPegawai: pt.NamaPegawai,
			KodeTim:     pt.KodeTim,
			NamaTim:     pt.NamaTim,
		}
		results[pt.IdProgramUnggulan] = append(results[pt.IdProgramUnggulan], resp)
	}

	return results, nil
}
