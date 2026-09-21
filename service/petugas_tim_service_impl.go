package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"timkerjaService/helper"
	"timkerjaService/internal"
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
	EventClient          *internal.EventClient
}

func NewPetugasTimServiceImpl(
	petugasTimRepository repository.PetugasTimRepository,
	susunanTimRepository repository.SusunanTimRepository,
	db *sql.DB,
	validator *validator.Validate,
	eventClient *internal.EventClient,
) *PetugasTimServiceImpl {
	return &PetugasTimServiceImpl{
		PetugasTimRepository: petugasTimRepository,
		SusunanTimRepository: susunanTimRepository,
		DB:                   db,
		Validator:            validator,
		EventClient:          eventClient,
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
	// pakai audit, commit manual
	// defer helper.CommitOrRollback(tx)
	defer tx.Rollback()

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
	// commit db
	if err := tx.Commit(); err != nil {
		return web.PetugasTimResponse{}, err
	}

	// Audit setelah database berhasil commit.
	event := internal.NewCreateEvent(
		"petugas_tim",
		strconv.Itoa(result.Id),
		result,
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event petugas_tim id=%d: %v",
			result.Id,
			err,
		)
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
	// pakai audit, commit manual
	// defer helper.CommitOrRollback(tx)
	defer tx.Rollback()
	before, err := service.PetugasTimRepository.FindById(ctx, tx, idPetugasTim)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New(
				"id petugas_tim tidak ditemukan",
			)
		}

		return err
	}

	err = service.PetugasTimRepository.Delete(ctx, tx, idPetugasTim)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	// Audit setelah database berhasil commit.
	event := internal.NewDeleteEvent(
		"petugas_tim",
		strconv.Itoa(before.Id),
		before,
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event petugas_tim id=%d: %v",
			before.Id,
			err,
		)
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
