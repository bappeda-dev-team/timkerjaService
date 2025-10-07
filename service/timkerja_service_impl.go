package service

import (
	"context"
	"database/sql"
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
		KodeTim:    helper.GenerateKodeTim(0),
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
		Id:         timKerja.Id,
		NamaTim:    timKerja.NamaTim,
		Keterangan: timKerja.Keterangan,
		Tahun:      timKerja.Tahun,
		IsActive:   timKerja.IsActive,
	}

	timKerjaDomain, err = service.TimKerjaRepository.Update(ctx, tx, timKerjaDomain)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	kodeTim, err := service.TimKerjaRepository.FindById(ctx, tx, timKerjaDomain.Id)
	if err != nil {
		return web.TimKerjaResponse{}, err
	}

	return web.TimKerjaResponse{
		KodeTim:    kodeTim.KodeTim,
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
					NamaPegawai:  st.NamaPegawai,
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
			Keterangan:  timKerja.Keterangan,
			SusunanTims: susunanTimResponses,
		})
	}

	return result, nil
}

func (service *TimKerjaServiceImpl) AddProgramUnggulan(ctx context.Context, programUnggulan web.ProgramUnggulanTimKerjaRequest) (web.ProgramUnggulanTimKerjaResponse, error) {
	err := service.Validator.Struct(programUnggulan)
	if err != nil {
		return web.ProgramUnggulanTimKerjaResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.ProgramUnggulanTimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programUnggulanDomain := domain.ProgramUnggulanTimKerja{
		KodeTim:           programUnggulan.KodeTim,
		IdProgramUnggulan: programUnggulan.IdProgramUnggulan,
		Tahun:             programUnggulan.Tahun,
		KodeOpd:           programUnggulan.KodeOpd,
	}

	programUnggulanDomain, err = service.TimKerjaRepository.AddProgramUnggulan(ctx, tx, programUnggulanDomain)
	if err != nil {
		return web.ProgramUnggulanTimKerjaResponse{}, err
	}
	// setelah simpan cek external service
	perencanaanClient := internal.NewPerencanaanClient(
		"https://testapi.kertaskerja.cc",
		&http.Client{Timeout: 25 * time.Second},
	)

	perencanaanResp, err := perencanaanClient.GetProgramUnggulan(ctx, programUnggulanDomain.IdProgramUnggulan)
	if err != nil {
		log.Printf("gagal cek program unggulan ke service eksternal: %v", err)
	}

	// inject namaProgramUnggulan
	namaProgramUnggulan := "NOT_CHECKED"
	if perencanaanResp != nil {
		namaProgramUnggulan = perencanaanResp.NamaTagging
	}

	return web.ProgramUnggulanTimKerjaResponse{
		Id:                programUnggulanDomain.Id,
		KodeTim:           programUnggulanDomain.KodeTim,
		IdProgramUnggulan: programUnggulanDomain.IdProgramUnggulan,
		ProgramUnggulan:   namaProgramUnggulan,
		Tahun:             programUnggulan.Tahun,
		KodeOpd:           programUnggulan.KodeOpd,
	}, nil
}

func (service *TimKerjaServiceImpl) FindAllProgramUnggulanTim(ctx context.Context, kodeTim string) ([]web.ProgramUnggulanTimKerjaResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.ProgramUnggulanTimKerjaResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	programUnggulans, err := service.TimKerjaRepository.FindProgramUnggulanByKodeTim(ctx, tx, kodeTim)
	if err != nil {
		return []web.ProgramUnggulanTimKerjaResponse{}, err
	}

	if len(programUnggulans) == 0 {
		return []web.ProgramUnggulanTimKerjaResponse{}, nil
	}

	// perencanaan service
	perencanaanHost := os.Getenv("PERENCANAAN_HOST")
	perencanaanClient := internal.NewPerencanaanClient(
		perencanaanHost,
		&http.Client{Timeout: 25 * time.Second},
	)

	var perencanaanResponses []domain.ProgramUnggulanTimKerja
	for _, p := range programUnggulans {
		resp := domain.ProgramUnggulanTimKerja{
			Id:                p.Id,
			KodeTim:           p.KodeTim,
			IdProgramUnggulan: p.IdProgramUnggulan,
			Tahun:             p.Tahun,
			KodeOpd:           p.KodeOpd,
		}

		perencanaanResp, err := perencanaanClient.GetProgramUnggulan(ctx, p.IdProgramUnggulan)
		if err != nil {
			log.Printf("gagal cek program unggulan ke service eksternal: %v", err)
			resp.NamaProgramUnggulan = "NOT_CHECKED"
		} else if perencanaanResp != nil {
			resp.NamaProgramUnggulan = perencanaanResp.NamaTagging
		}
		log.Printf("PERENCANAN RESP: %v", perencanaanResp)
		perencanaanResponses = append(perencanaanResponses, resp)
	}

	return helper.ToProgramUnggulanResponses(perencanaanResponses), nil
}
