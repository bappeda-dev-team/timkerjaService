package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"timkerjaService/helper"
	"timkerjaService/internal"
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
	EventClient          *internal.EventClient
}

func NewSusunanTimServiceImpl(
	susunanTimRepository repository.SusunanTimRepository,
	timKerjaService TimKerjaService,
	db *sql.DB,
	validator *validator.Validate,
	eventClient *internal.EventClient,
) *SusunanTimServiceImpl {
	return &SusunanTimServiceImpl{
		SusunanTimRepository: susunanTimRepository,
		TimKerjaService:      timKerjaService,
		DB:                   db,
		Validator:            validator,
		EventClient:          eventClient,
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
	// pakai audit, commit manual
	// defer helper.CommitOrRollback(tx)
	defer tx.Rollback()

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
	// commit db
	if err := tx.Commit(); err != nil {
		return web.SusunanTimResponse{}, err
	}

	// Audit setelah database berhasil commit.
	event := internal.NewCreateEvent(
		"susunan_tim",
		strconv.Itoa(susunanTimDomain.Id),
		susunanTimDomain,
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event susunan_tim id=%d: %v",
			susunanTimDomain.Id,
			err,
		)
	}

	return web.SusunanTimResponse{
		Id:             susunanTimDomain.Id,
		KodeTim:        susunanTimDomain.KodeTim,
		PegawaiId:      susunanTimDomain.PegawaiId,
		IdJabatanTim:   susunanTim.IdJabatanTim,
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
	// pakai audit, commit manual
	// defer helper.CommitOrRollback(tx)
	defer tx.Rollback()
	before, err := service.SusunanTimRepository.FindById(ctx, tx, susunanTim.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return web.SusunanTimResponse{}, errors.New(
				"id susunan_tim tidak ditemukan",
			)
		}

		return web.SusunanTimResponse{}, err
	}

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
	// commit db
	if err := tx.Commit(); err != nil {
		return web.SusunanTimResponse{}, err
	}
	// Audit setelah database berhasil commit.
	event := internal.NewUpdateEvent(
		"susunan_tim",
		strconv.Itoa(susunanTimDomain.Id),
		susunanTimDomain,
		before,
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event susunan_tim id=%d: %v",
			susunanTimDomain.Id,
			err,
		)
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
	defer tx.Rollback()
	before, err := service.SusunanTimRepository.FindById(ctx, tx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New(
				"id susunan_tim tidak ditemukan",
			)
		}

		return err
	}

	err = service.SusunanTimRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}
	// commit db
	if err := tx.Commit(); err != nil {
		return err
	}
	// Audit setelah database berhasil commit.
	event := internal.NewDeleteEvent(
		"susunan_tim",
		strconv.Itoa(id),
		before,
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event susunan_tim id=%d: %v",
			id,
			err,
		)
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
	if tahun <= 0 || bulan <= 0 || bulan > 14 {
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
	// pakai audit, commit manual
	// defer helper.CommitOrRollback(tx)
	defer tx.Rollback()

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

	// Create tim kerja baru jika tahun dan bulan tidak sama
	if tahunTarget != tahun || bulanTarget != bulan {

		timKerjaTarget, err := service.TimKerjaService.
			FindByKodeTim(ctx, kodeTim)
		if err != nil {
			return fmt.Errorf("tim kerja tidak ditemukan")
		}

		timExists, err := service.TimKerjaService.
			CheckCloned(ctx, timKerjaTarget.Id, bulanTarget, tahunTarget)
		if err != nil {
			return err
		}

		if timExists {
			return errors.New("Tim kerja sudah ada pada periode tersebut")
		}

		cloneTimKerja, err := service.TimKerjaService.
			CreateWithTx(ctx, tx, web.TimKerjaCreateRequest{
				NamaTim:       timKerjaTarget.NamaTim,
				Keterangan:    timKerjaTarget.Keterangan,
				IsActive:      timKerjaTarget.IsActive,
				IsSekretariat: timKerjaTarget.IsSekretariat,
				Bulan:         bulanTarget,
				Tahun:         strconv.Itoa(tahunTarget),
				CloneFrom:     timKerjaTarget.Id,
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
	// commit db
	if err := tx.Commit(); err != nil {
		return err
	}
	// Audit setelah database berhasil commit.
	event := internal.NewCloneEvent(
		"susunan_tim",
		kodeTimTarget,
		cloneSusunanTim,
		map[string]any{
			"source_kode_tim": kodeTim,
			"target_kode_tim": kodeTimTarget,
			"source_bulan":    bulan,
			"source_tahun":    tahun,
			"target_bulan":    bulanTarget,
			"target_tahun":    tahunTarget,
			"total":           len(cloneSusunanTim),
		},
	)
	if err := service.EventClient.CreateEvent(ctx, event); err != nil {
		log.Printf(
			"gagal membuat audit event susunan_tim id=%d: %v",
			kodeTim,
			err,
		)
	}

	return nil
}

func validateClone(bulan, tahun, bulanTarget, tahunTarget int) error {
	if tahun <= 0 || tahunTarget <= 0 {
		return errors.New("tahun tidak valid")
	}
	if bulan < 1 || bulan > 14 || bulanTarget < 1 || bulanTarget > 14 {
		return errors.New("bulan tidak valid")
	}
	if bulan == bulanTarget && tahun == tahunTarget {
		return errors.New("tidak bisa clone ke bulan dan tahun yang sama")
	}
	return nil
}
