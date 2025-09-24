package app

import (
	"timkerjaService/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(timKerjaController controller.TimKerjaController, susunanTimController controller.SusunanTimController, jabatanTimController controller.JabatanTimController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/v1/timkerja/swagger/doc.json", echoSwagger.WrapHandler)

	e.POST("/timkerja", timKerjaController.Create)
	e.PUT("/timkerja/:id", timKerjaController.Update)
	e.DELETE("/timkerja/:id", timKerjaController.Delete)
	e.GET("/timkerja/:id", timKerjaController.FindById)
	e.GET("/only_timkerja", timKerjaController.FindAll)
	e.GET("/timkerja", timKerjaController.FindAllTm)

	e.POST("/susunantim", susunanTimController.Create)
	e.PUT("/susunantim/:id", susunanTimController.Update)
	e.DELETE("/susunantim/:id", susunanTimController.Delete)
	e.GET("/susunantim/:id", susunanTimController.FindById)
	e.GET("/susunantim", susunanTimController.FindAll)

	e.POST("/jabatantim", jabatanTimController.Create)
	e.PUT("/jabatantim/:id", jabatanTimController.Update)
	e.DELETE("/jabatantim/:id", jabatanTimController.Delete)
	e.GET("/jabatantim/:id", jabatanTimController.FindById)
	e.GET("/jabatantim", jabatanTimController.FindAll)

	return e
}
