package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"timkerjaService/model/domain"
	"timkerjaService/model/web"
	"timkerjaService/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
)

// REPOSITORY MOCK
type MockSusunanTimRepository struct {
	CreateFn func(ctx context.Context, tx *sql.Tx, data domain.SusunanTim) (domain.SusunanTim, error)
}

func (m *MockSusunanTimRepository) Create(
	ctx context.Context,
	tx *sql.Tx,
	data domain.SusunanTim,
) (domain.SusunanTim, error) {
	if m.CreateFn == nil {
		panic("CreateFn is nil")
	}
	return m.CreateFn(ctx, tx, data)
}

func (m *MockSusunanTimRepository) Update(
	ctx context.Context,
	tx *sql.Tx,
	data domain.SusunanTim,
) (domain.SusunanTim, error) {
	panic("Update not implemented")
}

func (m *MockSusunanTimRepository) Delete(
	ctx context.Context,
	tx *sql.Tx,
	id int,
) error {
	panic("Delete not implemented")
}

func (m *MockSusunanTimRepository) FindById(
	ctx context.Context,
	tx *sql.Tx,
	id int,
) (domain.SusunanTim, error) {
	panic("FindById not implemented")
}

func (m *MockSusunanTimRepository) FindAll(
	ctx context.Context,
	tx *sql.Tx,
) ([]domain.SusunanTim, error) {
	panic("FindAll not implemented")
}

func (m *MockSusunanTimRepository) FindByKodeTim(
	ctx context.Context,
	tx *sql.Tx,
	kodeTim string,
) ([]domain.SusunanTim, error) {
	panic("FindByKodeTim not implemented")
}

func TestSusunanTimServiceImpl_Create(t *testing.T) {
	ctx := context.Background()

	// === request yang VALID ===
	validReq := web.SusunanTimCreateRequest{
		KodeTim:        "TEST-KODE",
		PegawaiId:      "123456789012345678",
		NamaPegawai:    "TEST PEGAWAI",
		IdJabatanTim:   12,
		NamaJabatanTim: "KETUA",
		IsActive:       true,
		Keterangan:     "NO KETERANGAN",
	}

	tests := []struct {
		name       string
		req        web.SusunanTimCreateRequest
		repository repository.SusunanTimRepository
		wantErr    bool
	}{
		{
			name: "SUCCESS",
			req:  validReq,
			repository: &MockSusunanTimRepository{
				CreateFn: func(ctx context.Context, tx *sql.Tx, data domain.SusunanTim) (domain.SusunanTim, error) {
					data.Id = 10
					return data, nil
				},
			},
			wantErr: false,
		},
		{
			name:       "VALIDATION ERROR",
			req:        web.SusunanTimCreateRequest{}, // invalid
			repository: &MockSusunanTimRepository{},
			wantErr:    true,
		},
		{
			name: "REPOSITORY ERROR",
			req:  validReq, // HARUS valid
			repository: &MockSusunanTimRepository{
				CreateFn: func(ctx context.Context, tx *sql.Tx, data domain.SusunanTim) (domain.SusunanTim, error) {
					return domain.SusunanTim{}, errors.New("db error")
				},
			},
			wantErr: true,
		},
	}

	// EXPECTATION
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create sqlmock: %v", err)
			}
			defer db.Close()

			// === SET EXPECTATION TRANSACTION ===
			switch tt.name {
			case "SUCCESS":
				mock.ExpectBegin()
				mock.ExpectCommit()
			case "REPOSITORY ERROR":
				mock.ExpectBegin()
				mock.ExpectRollback()
				// VALIDATION ERROR → TIDAK ADA TX
			}

			service := &SusunanTimServiceImpl{
				SusunanTimRepository: tt.repository,
				DB:                   db,
				Validator:            validator.New(),
			}

			result, err := service.Create(ctx, tt.req)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result.Id != 10 {
				t.Errorf("expected Id=10, got %d", result.Id)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unmet sqlmock expectations: %v", err)
			}
		})
	}
}
