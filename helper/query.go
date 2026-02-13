package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"timkerjaService/model/web"

	"github.com/labstack/echo/v4"
)

func GetQueryIntWithDefault(
	c echo.Context,
	param string,
	defaultVal int,
) (int, error) {

	valStr := c.QueryParam(param)
	if valStr == "" {
		return defaultVal, nil
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   param + " harus berupa angka",
		})
	}

	return val, nil
}

func GetQueryToInt(
	c echo.Context,
	param string,
) (int, error) {

	valStr := c.QueryParam(param)
	if valStr == "" {
		return 0, fmt.Errorf("%s wajib diisi", param)
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("%s harus berupa angka", param)
	}

	return val, nil
}
