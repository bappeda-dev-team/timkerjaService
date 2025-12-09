package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type PenilaianKinerjaControllerImpl struct {
	PenilaianKinerjaService service.PenilaianKinerjaService
}

func NewPenilaianKinerjaControllerImpl(penilaianKinerjaService service.PenilaianKinerjaService) *PenilaianKinerjaControllerImpl {
	return &PenilaianKinerjaControllerImpl{
		PenilaianKinerjaService: penilaianKinerjaService,
	}
}

// @Summary Penilaian kinerja by bulan tahun
// @Description Penilaian Kinerja berdasarkan bulan & tahun, dikelompokkan per kode tim dan individu
// @Tags Penilaian Kinerja
// @Accept json
// @Produce json
// @Param tahun query int true "Tahun penilaian (ex: 2025)"
// @Param bulan query int true "Bulan penilaian (ex: 12)"
// @Success 200 {object} web.WebResponse{data=[]web.LaporanPenilaianKinerjaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /penilaian_kinerja [get]
func (controller *PenilaianKinerjaControllerImpl) All(c echo.Context) error {
	// Ambil query param
	tahunStr := c.QueryParam("tahun")
	bulanStr := c.QueryParam("bulan")

	if tahunStr == "" || bulanStr == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "tahun dan bulan wajib diisi",
		})
	}

	// Convert ke int
	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "tahun harus berupa angka",
		})
	}

	bulan, err := strconv.Atoi(bulanStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "bulan harus berupa angka",
		})
	}

	// Panggil service
	result, err := controller.PenilaianKinerjaService.All(c.Request().Context(), tahun, bulan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	// Response sukses
	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result, // []LaporanPenilaianKinerjaResponse
	})
}

// @Summary Add New Penialaian to Individu
// @Description Add Penilaian to Individu by jenis
// @Tags Penilaian Kinerja
// @Accept json
// @Produce json
// @Param data body web.PenilaianKinerjaRequest true "Penilaian Tim Kerja Create Request"
// @Success 201 {object} web.WebResponse{data=web.PenilaianKinerjaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /penilaian_kinerja [post]
func (ctrl *PenilaianKinerjaControllerImpl) Create(c echo.Context) error {
	PenilaianCreateRequest := web.PenilaianKinerjaRequest{}
	err := c.Bind(&PenilaianCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	PenilaianKinerjaResponse, err := ctrl.PenilaianKinerjaService.Create(c.Request().Context(), PenilaianCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   PenilaianKinerjaResponse,
	})
}

// @Summary Update Penialaian Kinerja Individu
// @Description Update Penilaian Individu by id
// @Tags Penilaian Kinerja
// @Accept json
// @Produce json
// @Param data body web.PenilaianKinerjaRequest true "Penilaian Tim Kerja Create Request"
// @Success 201 {object} web.WebResponse{data=web.PenilaianKinerjaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /penilaian_kinerja/:id [put]
func (ctrl *PenilaianKinerjaControllerImpl) Update(c echo.Context) error {
	penilaianId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	PenilaianRequest := web.PenilaianKinerjaRequest{}
	err = c.Bind(&PenilaianRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := ctrl.PenilaianKinerjaService.Update(c.Request().Context(), PenilaianRequest, penilaianId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})
}
