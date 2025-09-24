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

	e.GET("/api/v1/timkerja/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/v1/timkerja/swagger/doc.json", echoSwagger.WrapHandler)

	e.POST("/api/v1/timkerja/timkerja", timKerjaController.Create)
	e.PUT("/api/v1/timkerja/timkerja/:id", timKerjaController.Update)
	e.DELETE("/api/v1/timkerja/timkerja/:id", timKerjaController.Delete)
	e.GET("/api/v1/timkerja/timkerja/:id", timKerjaController.FindById)
	e.GET("/api/v1/timkerja/only_timkerja", timKerjaController.FindAll)
	e.GET("/api/v1/timkerja/timkerja", timKerjaController.FindAllTm)

	e.POST("/api/v1/timkerja/susunantim", susunanTimController.Create)
	e.PUT("/api/v1/timkerja/susunantim/:id", susunanTimController.Update)
	e.DELETE("/api/v1/timkerja/susunantim/:id", susunanTimController.Delete)
	e.GET("/api/v1/timkerja/susunantim/:id", susunanTimController.FindById)
	e.GET("/api/v1/timkerja/susunantim", susunanTimController.FindAll)

	e.POST("/api/v1/timkerja/jabatantim", jabatanTimController.Create)
	e.PUT("/api/v1/timkerja/jabatantim/:id", jabatanTimController.Update)
	e.DELETE("/api/v1/timkerja/jabatantim/:id", jabatanTimController.Delete)
	e.GET("/api/v1/timkerja/jabatantim/:id", jabatanTimController.FindById)
	e.GET("/api/v1/timkerja/jabatantim", jabatanTimController.FindAll)

	return e
}
