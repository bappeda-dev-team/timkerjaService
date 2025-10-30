//go:build wireinject
// +build wireinject

package main

import (
	"timkerjaService/app"

	"timkerjaService/controller"
	"timkerjaService/repository"
	"timkerjaService/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var timKerjaSet = wire.NewSet(
	repository.NewTimKerjaRepositoryImpl,
	wire.Bind(new(repository.TimKerjaRepository), new(*repository.TimKerjaRepositoryImpl)),
	service.NewTimKerjaServiceImpl,
	wire.Bind(new(service.TimKerjaService), new(*service.TimKerjaServiceImpl)),
	controller.NewTimKerjaControllerImpl,
	wire.Bind(new(controller.TimKerjaController), new(*controller.TimKerjaControllerImpl)),
)

var susunanTimSet = wire.NewSet(
	repository.NewSusunanTimRepositoryImpl,
	wire.Bind(new(repository.SusunanTimRepository), new(*repository.SusunanTimRepositoryImpl)),
	service.NewSusunanTimServiceImpl,
	wire.Bind(new(service.SusunanTimService), new(*service.SusunanTimServiceImpl)),
	controller.NewSusunanTimControllerImpl,
	wire.Bind(new(controller.SusunanTimController), new(*controller.SusunanTimControllerImpl)),
)

var jabatanTimSet = wire.NewSet(
	repository.NewJabatanTimRepositoryImpl,
	wire.Bind(new(repository.JabatanTimRepository), new(*repository.JabatanTimRepositoryImpl)),
	service.NewJabatanTimServiceImpl,
	wire.Bind(new(service.JabatanTimService), new(*service.JabatanTimServiceImpl)),
	controller.NewJabatanTimControllerImpl,
	wire.Bind(new(controller.JabatanTimController), new(*controller.JabatanTimControllerImpl)),
)

var realisasiAnggaranSet = wire.NewSet(
	repository.NewRealisasiAnggaranRepositoryImpl,
	wire.Bind(new(repository.RealisasiAnggaranRepository), new(*repository.RealisasiAnggaranRepositoryImpl)),
	service.NewRealisasiAnggaranServiceImpl,
	wire.Bind(new(service.RealisasiAnggaranService), new(*service.RealisasiAnggaranServiceImpl)),
	controller.NewRealisasiAnggaranControllerImpl,
	wire.Bind(new(controller.RealisasiAnggaranController), new(*controller.RealisasiAnggaranControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		timKerjaSet,
		susunanTimSet,
		jabatanTimSet,
		realisasiAnggaranSet,
		app.NewRouter,
	)
	return nil
}
