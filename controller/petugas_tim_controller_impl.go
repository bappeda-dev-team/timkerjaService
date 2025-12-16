package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type PetugasTimControllerImpl struct {
	PetugasTimService service.PetugasTimService
}

func NewPetugasTimControllerImpl(petugasTimService service.PetugasTimService) *PetugasTimControllerImpl {
	return &PetugasTimControllerImpl{
		PetugasTimService: petugasTimService,
	}
}

// @Summary Add Petugas
// @Description Add petugas to the program unggulan
// @Tags Petugas
// @Accept json
// @Produce json
// @Param data body web.PetugasTimCreateRequest true "Petugas Tim Create Request"
// @Success 201 {object} web.WebResponse{data=web.PetugasTimResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /petugas_tim [post]
func (controller *PetugasTimControllerImpl) AddPetugas(c echo.Context) error {
	petugasTimCreateRequest := web.PetugasTimCreateRequest{}
	err := c.Bind(&petugasTimCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := controller.PetugasTimService.Create(c.Request().Context(), petugasTimCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

// @Summary Delete Petugas
// @Description Remove petugas from the program unggulan assigned
// @Tags Petugas
// @Accept json
// @Produce json
// @Param idPetugasTim path int true "Petugas Tim ID"
// @Success 200 {object} web.WebResponse "DELETED"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /petugas_tim/:idPetugasTim [delete]
func (controller *PetugasTimControllerImpl) DeletePetugas(c echo.Context) error {
	idPetugasTim := c.Param("idPetugasTim")
	if idPetugasTim == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	intIdPetugas, err := strconv.Atoi(idPetugasTim)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.PetugasTimService.Delete(c.Request().Context(), intIdPetugas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}
