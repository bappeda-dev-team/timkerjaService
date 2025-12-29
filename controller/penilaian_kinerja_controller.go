package controller

import "github.com/labstack/echo/v4"

type PenilaianKinerjaController interface {
	All(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	LaporanTpp(c echo.Context) error
	LaporanTppAll(c echo.Context) error
	// Delete(c echo.Context) error
}
