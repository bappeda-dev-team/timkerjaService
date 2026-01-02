package controller

import (
	"net/http"
	"strconv"
	"strings"
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

// @Summary Clone susunan tim ke periode target
// @Description Clone susunan tim dari periode asal ke periode target
// @Tags Susunan Tim
// @Accept json
// @Produce json
// @Param request body web.CloneSusunanTimRequest true "Payload clone"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 409 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /susunantim/clone [post]
func (controller *SusunanTimControllerImpl) CloneSusunanTim(c echo.Context) error {
	var req web.CloneSusunanTimRequest
	if err := c.Bind(&req); err != nil {
		return badRequest(c, "payload tidak valid")
	}

	// --- Validasi dasar ---
	if req.KodeTim == "" {
		return badRequest(c, "kodeTim wajib diisi")
	}
	if req.Bulan < 1 || req.Bulan > 12 {
		return badRequest(c, "bulan tidak valid")
	}
	if req.BulanTarget < 1 || req.BulanTarget > 12 {
		return badRequest(c, "bulanTarget tidak valid")
	}
	if req.Tahun <= 0 || req.TahunTarget <= 0 {
		return badRequest(c, "tahun tidak valid")
	}

	// service
	err := controller.SusunanTimService.CloneByKodeTim(
		c.Request().Context(),
		req.Bulan,
		req.Tahun,
		req.KodeTim,
		req.BulanTarget,
		req.TahunTarget,
	)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "sudah ada"):
			return conflict(c, err.Error())
		case strings.Contains(err.Error(), "tidak ditemukan"):
			return conflict(c, err.Error())
		default:
			return internalError(c, err)
		}
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Clone Sukses",
	})
}

func badRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
		Data:   message,
	})
}

func conflict(c echo.Context, message string) error {
	return c.JSON(http.StatusConflict, web.WebResponse{
		Code:   http.StatusConflict,
		Status: "Conflict",
		Data:   message,
	})
}

func internalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
