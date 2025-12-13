package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
	"timkerjaService/model/domain"
)

func TestNewSusunanTimRepositoryImpl(t *testing.T) {
	tests := []struct {
		name string
		want *SusunanTimRepositoryImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSusunanTimRepositoryImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSusunanTimRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSusunanTimRepositoryImpl_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	tx, _ := db.Begin()

	repo := &SusunanTimRepositoryImpl{}

	input := domain.SusunanTim{
		KodeTim:        "TIM-1",
		PegawaiId:      "123",
		NamaPegawai:    "Budi",
		IdJabatanTim:   1,
		NamaJabatanTim: "Ketua",
		IsActive:       true,
	}

	mock.ExpectExec("INSERT INTO susunan_tim").
		WithArgs(
			input.KodeTim,
			input.PegawaiId,
			input.NamaPegawai,
			input.IdJabatanTim,
			input.NamaJabatanTim,
			input.IsActive,
			input.Keterangan,
		).
		WillReturnResult(sqlmock.NewResult(10, 1))

	result, err := repo.Create(context.Background(), tx, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Id != 10 {
		t.Errorf("expected id 10, got %d", result.Id)
	}
}
