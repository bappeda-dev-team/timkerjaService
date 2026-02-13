package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"timkerjaService/helper"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
	"timkerjaService/repository"

	"github.com/go-playground/validator/v10"
)

type SusunanTimServiceImpl struct {
	SusunanTimRepository repository.SusunanTimRepository
	TimKerjaService      TimKerjaService
	DB                   *sql.DB
	Validator            *validator.Validate
}

func NewSusunanTimServiceImpl(
	susunanTimRepository repository.SusunanTimRepository,
	timKerjaService TimKerjaService,
	db *sql.DB,
	validator *validator.Validate) *SusunanTimServiceImpl {
	return &SusunanTimServiceImpl{
		SusunanTimRepository: susunanTimRepository,
		TimKerjaService:      timKerjaService,
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

func (service *SusunanTimServiceImpl) FindAllByBulanTahun(ctx context.Context, bulan int, tahun int) ([]web.SusunanTimResponse, error) {
	if tahun <= 0 || bulan <= 0 || bulan > 12 {
		return nil, errors.New("tahun atau bulan tidak valid")
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.SusunanTimResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	susunanTimDomains, err := service.SusunanTimRepository.FindAllByBulanTahun(ctx, tx, bulan, tahun)
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

func (service *SusunanTimServiceImpl) CloneByKodeTim(ctx context.Context, bulan int, tahun int, kodeTim string, bulanTarget int, tahunTarget int) error {
	// guard bulan tahun
	if err := validateClone(bulan, tahun, bulanTarget, tahunTarget); err != nil {
		return err
	}

	// tx db
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	// cek untuk memastikan susunan tim belum ada di bulan tahun target
	exists, err := service.SusunanTimRepository.
		ExistsByKodeTimBulanTahun(ctx, tx, kodeTim, bulanTarget, tahunTarget)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("susunan tim target sudah ada")
	}
	susunanTims, err := service.SusunanTimRepository.
		FindByKodeTimBulanTahun(ctx, tx, kodeTim, bulan, tahun)
	if err != nil {
		return fmt.Errorf("find susunan tim gagal: %w", err)
	}
	if len(susunanTims) <= 0 {
		return errors.New("Susunan Tim Tidak ditemukan")
	}

	kodeTimTarget := kodeTim

	// Create tim kerja baru jika tahun tidak sama
	if tahunTarget != tahun {
		timKerjaTarget, err := service.TimKerjaService.
			FindByKodeTim(ctx, kodeTim)
		if err != nil {
			return fmt.Errorf("tim kerja tidak ditemukan")
		}

		cloneTimKerja, err := service.TimKerjaService.
			CreateWithTx(ctx, tx, web.TimKerjaCreateRequest{
				NamaTim:       timKerjaTarget.NamaTim,
				Keterangan:    timKerjaTarget.Keterangan,
				IsActive:      timKerjaTarget.IsActive,
				IsSekretariat: timKerjaTarget.IsSekretariat,
				Tahun:         strconv.Itoa(tahunTarget),
			})
		if err != nil {
			return fmt.Errorf("tim kerja gagal di clone: %w", err)
		}

		kodeTimTarget = cloneTimKerja.KodeTim
	}

	cloneSusunanTim := make([]domain.SusunanTim, 0, len(susunanTims))
	for _, st := range susunanTims {
		newSusunanTim := domain.SusunanTim{
			KodeTim:        kodeTimTarget,
			Bulan:          bulanTarget,
			Tahun:          tahunTarget,
			PegawaiId:      st.PegawaiId,
			NamaPegawai:    st.NamaPegawai,
			IdJabatanTim:   st.IdJabatanTim,
			NamaJabatanTim: st.NamaJabatanTim,
			IsActive:       st.IsActive,
			Keterangan:     st.Keterangan,
		}

		cloneSusunanTim = append(cloneSusunanTim, newSusunanTim)
	}
	err = service.SusunanTimRepository.SaveAll(ctx, tx, cloneSusunanTim)
	if err != nil {
		return fmt.Errorf("save clone susunan tim gagal: %w", err)
	}

	return nil
}

func validateClone(bulan, tahun, bulanTarget, tahunTarget int) error {
	if tahun <= 0 || tahunTarget <= 0 {
		return errors.New("tahun tidak valid")
	}
	if bulan < 1 || bulan > 12 || bulanTarget < 1 || bulanTarget > 12 {
		return errors.New("bulan tidak valid")
	}
	if bulan == bulanTarget && tahun == tahunTarget {
		return errors.New("tidak bisa clone ke bulan dan tahun yang sama")
	}
	return nil
}
