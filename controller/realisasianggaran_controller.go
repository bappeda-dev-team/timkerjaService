package controller

import "github.com/labstack/echo/v4"

type RealisasiAnggaranController interface {
	Delete(c echo.Context) error
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
	Upsert(c echo.Context) error
}
