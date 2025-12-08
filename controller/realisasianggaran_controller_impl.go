package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type RealisasiAnggaranControllerImpl struct {
	RealisasiAnggaranService service.RealisasiAnggaranService
}

func NewRealisasiAnggaranControllerImpl(realisasiAnggaranService service.RealisasiAnggaranService) *RealisasiAnggaranControllerImpl {
	return &RealisasiAnggaranControllerImpl{
		RealisasiAnggaranService: realisasiAnggaranService,
	}
}

// @Summary Delete Realisasi Anggaran
// @Description Delete Realisasi Anggaran
// @Tags Realisasi Anggaran
// @Accept json
// @Produce json
// @Param id path int true "Realisasi Anggaran ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /realisasianggaran/{id} [delete]
func (controller *RealisasiAnggaranControllerImpl) Delete(c echo.Context) error {
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

	err = controller.RealisasiAnggaranService.Delete(c.Request().Context(), idInt)
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

// @Summary Get Realisasi Anggaran by ID
// @Description Get Realisasi Anggaran detail by ID
// @Tags Realisasi Anggaran
// @Accept json
// @Produce json
// @Param id path int true "Realisasi Anggaran ID"
// @Success 200 {object} web.WebResponse{data=web.RealisasiAnggaranResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /realisasianggaran/detail/{id} [get]
func (controller *RealisasiAnggaranControllerImpl) FindById(c echo.Context) error {
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

	RealisasiAnggaranResponse, err := controller.RealisasiAnggaranService.FindById(c.Request().Context(), idInt)
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
		Data:   RealisasiAnggaranResponse,
	})
}

// @Summary Get All Realisasi Anggaran
// @Description Get All Realisasi Anggaran
// @Tags Realisasi Anggaran
// @Accept json
// @Param kode_tim path string true "Kode Tim"
// @Param id_rencana_kinerja path string true "ID Rencana Kinerja"
// @Param kode_subkegiatan path string true "Kode Subkegiatan"
// @Param bulan path string true "Bulan"
// @Param tahun path string true "Tahun"
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.RealisasiAnggaranResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /realisasianggaran/{kode_subkegiatan}/{kode_tim}/{id_rencana_kinerja}/{bulan}/{tahun} [get]
func (controller *RealisasiAnggaranControllerImpl) FindAll(c echo.Context) error {
	kodeSubkegiatan := c.Param("kode_subkegiatan")
	kodeTim := c.Param("kode_tim")
	idRencanaKinerja := c.Param("id_rencana_kinerja")
	bulan := c.Param("bulan")
	tahun := c.Param("tahun")
	if kodeSubkegiatan == "" || kodeTim == "" || idRencanaKinerja == "" || bulan == "" || tahun == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}
	RealisasiAnggaranResponses, err := controller.RealisasiAnggaranService.FindAll(c.Request().Context(), kodeSubkegiatan, kodeTim, idRencanaKinerja, bulan, tahun)
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
		Data:   RealisasiAnggaranResponses,
	})
}

// @Summary Upsert Realisasi Anggaran
// @Description Upsert Realisasi Anggaran
// @Tags Realisasi Anggaran
// @Accept json
// @Produce json
// @Param data body web.RealisasiAnggaranCreateRequest true "Realisasi Anggaran Create Request"
// @Success 200 {object} web.WebResponse{data=web.RealisasiAnggaranResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /realisasianggaran [post]
func (controller *RealisasiAnggaranControllerImpl) Upsert(c echo.Context) error {
	RealisasiAnggaranCreateRequest := web.RealisasiAnggaranCreateRequest{}
	err := c.Bind(&RealisasiAnggaranCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	RealisasiAnggaranResponse, err := controller.RealisasiAnggaranService.Upsert(c.Request().Context(), RealisasiAnggaranCreateRequest)
	if err != nil {
		if err.Error() == "bulan tidak valid" {
			return c.JSON(http.StatusBadRequest, web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   err.Error(),
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
		Data:   RealisasiAnggaranResponse,
	})
}
