package middleware

import (
	"github.com/labstack/echo/v4"
	"timkerjaService/internal"
)

func SessionIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Request().Header.Get("X-Session-Id")
		ctx := internal.WithSessionID(c.Request().Context(), sessionID)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
