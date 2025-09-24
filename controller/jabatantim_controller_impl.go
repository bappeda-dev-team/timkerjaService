package controller

import (
	"net/http"
	"strconv"
	"timkerjaService/model/web"
	"timkerjaService/service"

	"github.com/labstack/echo/v4"
)

type JabatanTimControllerImpl struct {
	JabatanTimService service.JabatanTimService
}

func NewJabatanTimControllerImpl(jabatanTimService service.JabatanTimService) *JabatanTimControllerImpl {
	return &JabatanTimControllerImpl{
		JabatanTimService: jabatanTimService,
	}
}

// @Summary Create Jabatan Tim
// @Description Create new Jabatan Tim
// @Tags Jabatan Tim
// @Accept json
// @Produce json
// @Param data body web.JabatanTimCreateRequest true "Jabatan Tim Create Request"
// @Success 201 {object} web.WebResponse{data=web.JabatanTimResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/timkerja/jabatantim [post]
func (controller *JabatanTimControllerImpl) Create(c echo.Context) error {

	JabatanTimCreateRequest := web.JabatanTimCreateRequest{}
	err := c.Bind(&JabatanTimCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	JabatanTimResponse, err := controller.JabatanTimService.Create(c.Request().Context(), JabatanTimCreateRequest)
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
		Data:   JabatanTimResponse,
	})
}

// @Summary Update Jabatan Tim
// @Description Update Jabatan Tim
// @Tags Jabatan Tim
// @Accept json
// @Produce json
// @Param id path int true "Jabatan Tim ID"
// @Param data body web.JabatanTimUpdateRequest true "Jabatan Tim Update Request"
// @Success 200 {object} web.WebResponse{data=web.JabatanTimResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/timkerja/jabatantim/{id} [put]
func (controller *JabatanTimControllerImpl) Update(c echo.Context) error {
	JabatanTimUpdateRequest := web.JabatanTimUpdateRequest{}
	err := c.Bind(&JabatanTimUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	JabatanTimUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	JabatanTimResponse, err := controller.JabatanTimService.Update(c.Request().Context(), JabatanTimUpdateRequest)
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
		Data:   JabatanTimResponse,
	})
}

// @Summary Delete Jabatan Tim
// @Description Delete Jabatan Tim
// @Tags Jabatan Tim
// @Accept json
// @Produce json
// @Param id path int true "Jabatan Tim ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/timkerja/jabatantim/{id} [delete]
func (controller *JabatanTimControllerImpl) Delete(c echo.Context) error {
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

	err = controller.JabatanTimService.Delete(c.Request().Context(), idInt)
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

// @Summary Get Jabatan Tim by ID
// @Description Get Jabatan Tim detail by ID
// @Tags Jabatan Tim
// @Accept json
// @Produce json
// @Param id path int true "Jabatan Tim ID"
// @Success 200 {object} web.WebResponse{data=web.JabatanTimResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/timkerja/jabatantim/{id} [get]
func (controller *JabatanTimControllerImpl) FindById(c echo.Context) error {
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

	JabatanTimResponse, err := controller.JabatanTimService.FindById(c.Request().Context(), idInt)
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
		Data:   JabatanTimResponse,
	})
}

// @Summary Get All Jabatan Tim
// @Description Get All Jabatan Tim
// @Tags Jabatan Tim
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.JabatanTimResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /api/v1/timkerja/jabatantim [get]
func (controller *JabatanTimControllerImpl) FindAll(c echo.Context) error {
	JabatanTimResponses, err := controller.JabatanTimService.FindAll(c.Request().Context())
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
		Data:   JabatanTimResponses,
	})
}
