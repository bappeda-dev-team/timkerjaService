package helper

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"timkerjaService/model/web"
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
