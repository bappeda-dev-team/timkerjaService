package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type TimKerjaControllerImpl struct {
	TimKerjaService service.TimKerjaService
}

func NewTimKerjaControllerImpl(timKerjaService service.TimKerjaService) *TimKerjaControllerImpl {
	return &TimKerjaControllerImpl{
		TimKerjaService: timKerjaService,
	}
}

// @Summary Create Tim Kerja
// @Description Create new Tim Kerja
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Param data body web.TimKerjaCreateRequest true "Tim Kerja Create Request"
// @Success 201 {object} web.WebResponse{data=web.TimKerjaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /timkerja [post]
func (controller *TimKerjaControllerImpl) Create(c echo.Context) error {
	TimKerjaCreateRequest := web.TimKerjaCreateRequest{}
	err := c.Bind(&TimKerjaCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	TimKerjaResponse, err := controller.TimKerjaService.Create(c.Request().Context(), TimKerjaCreateRequest)
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
		Data:   TimKerjaResponse,
	})

}

// @Summary Update Tim Kerja
// @Description Update Tim Kerja
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Param data body web.TimKerjaUpdateRequest true "Tim Kerja Update Request"
// @Success 200 {object} web.WebResponse{data=web.TimKerjaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /timkerja/{id} [put]
func (controller *TimKerjaControllerImpl) Update(c echo.Context) error {
	TimKerjaUpdateRequest := web.TimKerjaUpdateRequest{}
	err := c.Bind(&TimKerjaUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	TimKerjaResponse, err := controller.TimKerjaService.Update(c.Request().Context(), TimKerjaUpdateRequest)
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
		Data:   TimKerjaResponse,
	})
}

// @Summary Delete Tim Kerja
// @Description Delete Tim Kerja
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Param id path int true "Tim Kerja ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /timkerja/{id} [delete]
func (controller *TimKerjaControllerImpl) Delete(c echo.Context) error {
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

	err = controller.TimKerjaService.Delete(c.Request().Context(), idInt)
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

// @Summary Get Tim Kerja by ID
// @Description Get Tim Kerja detail by ID
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Param id path int true "Tim Kerja ID"
// @Success 200 {object} web.WebResponse{data=web.TimKerjaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /timkerja/{id} [get]
func (controller *TimKerjaControllerImpl) FindById(c echo.Context) error {
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

	TimKerjaResponse, err := controller.TimKerjaService.FindById(c.Request().Context(), idInt)
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
		Data:   TimKerjaResponse,
	})
}

// @Summary Get All Tim Kerja
// @Description Get All Tim Kerja
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TimKerjaResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /only_timkerja [get]
func (controller *TimKerjaControllerImpl) FindAll(c echo.Context) error {
	TimKerjaResponses, err := controller.TimKerjaService.FindAll(c.Request().Context())
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
		Data:   TimKerjaResponses,
	})
}

// @Summary Get All Tim Kerja with Details
// @Description Get All Tim Kerja with their Susunan Tim details
// @Tags Tim Kerja
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.TimKerjaDetailResponse}
// @Failure 500 {object} web.WebResponse
// @Router /timkerja [get]
func (controller *TimKerjaControllerImpl) FindAllTm(c echo.Context) error {
	TimKerjaResponses, err := controller.TimKerjaService.FindAllTm(c.Request().Context())
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
		Data:   TimKerjaResponses,
	})
}
