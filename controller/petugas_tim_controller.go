package controller

import "github.com/labstack/echo/v4"

type PetugasTimController interface {
	AddPetugas(c echo.Context) error
	DeletePetugas(c echo.Context) error
}
