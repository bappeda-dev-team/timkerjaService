package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/helper"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type SusunanTimControllerImpl struct {
	SusunanTimService service.SusunanTimService
}

func NewSusunanTimControllerImpl(susunanTimService service.SusunanTimService) *SusunanTimControllerImpl {
	return &SusunanTimControllerImpl{
		SusunanTimService: susunanTimService,
	}
}

// @Summary Create Susunan Tim
// @Description Create new Susunan Tim
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param data body web.SusunanTimCreateRequest true "Susunan Tim Create Request"
// @Success 201 {object} web.WebResponse{data=web.SusunanTimResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim [post]
func (controller *SusunanTimControllerImpl) Create(c echo.Context) error {
	SusunanTimCreateRequest := web.SusunanTimCreateRequest{}
	err := c.Bind(&SusunanTimCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	SusunanTimResponse, err := controller.SusunanTimService.Create(c.Request().Context(), SusunanTimCreateRequest)
	if err != nil {
		if ve, ok := err.(*web.ValidationError); ok {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   ve,
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   SusunanTimResponse,
	})
}

// @Summary Update Susunan Tim
// @Description Update Susunan Tim
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param id path int true "Susunan Tim ID"
// @Param data body web.SusunanTimUpdateRequest true "Susunan Tim Update Request"
// @Success 200 {object} web.WebResponse{data=web.SusunanTimResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim/{id} [put]
func (controller *SusunanTimControllerImpl) Update(c echo.Context) error {
	SusunanTimUpdateRequest := web.SusunanTimUpdateRequest{}
	err := c.Bind(&SusunanTimUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	SusunanTimUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	SusunanTimResponse, err := controller.SusunanTimService.Update(c.Request().Context(), SusunanTimUpdateRequest)
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
		Data:   SusunanTimResponse,
	})
}

// @Summary Delete Susunan Tim
// @Description Delete Susunan Tim
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param id path int true "Susunan Tim ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim/{id} [delete]
func (controller *SusunanTimControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	err = controller.SusunanTimService.Delete(c.Request().Context(), idInt)
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

// @Summary Get Susunan Tim by ID
// @Description Get Susunan Tim detail by ID
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param id path int true "Susunan Tim ID"
// @Success 200 {object} web.WebResponse{data=web.SusunanTimResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim/{id} [get]
func (controller *SusunanTimControllerImpl) FindById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	SusunanTimResponse, err := controller.SusunanTimService.FindById(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   SusunanTimResponse,
	})
}

// @Summary Get All Susunan Tim
// @Description Get All Susunan Tim
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param tahun query int true "Tahun penilaian (ex: 2025)"
// @Param bulan query int true "Bulan penilaian (ex: 1)"
// @Success 200 {object} web.WebResponse{data=[]web.SusunanTimResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim [get]
func (controller *SusunanTimControllerImpl) FindAll(c echo.Context) error {
	// TODO: ubah ke current month
	bulan, err := helper.GetQueryIntWithDefault(c, "bulan", 12)
	// TODO: ubah ke string agar kompatible dengan DB
	tahun, err := helper.GetQueryIntWithDefault(c, "tahun", 2025)

	SusunanTimResponses, err := controller.SusunanTimService.FindAllByBulanTahun(c.Request().Context(), bulan, tahun)
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
		Data:   SusunanTimResponses,
	})
}

// @Summary Susunan Tim by Kode Tim
// @Description Get Pelaksana Tim by Kode Tim
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.SusunanTimResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /susunantim/{kodeTim}/pelaksana [get]
func (controller *SusunanTimControllerImpl) FindByKodeTim(c echo.Context) error {
	kodeTim := c.Param("kodeTim")
	if kodeTim == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   "KODE TIM TIDAK BOLEH KOSONG",
		})
	}

	SusunanTimResponses, err := controller.SusunanTimService.FindByKodeTim(c.Request().Context(), kodeTim)
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
		Data:   SusunanTimResponses,
	})
}
