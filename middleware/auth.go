package middleware

import (
	"os"

	"github.com/anon-kae/assessment-tax/errortype"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, password, ok := c.Request().BasicAuth()
		if !ok {
			return errortype.AuthenticationError{Message: "Unauthorized."}
		}

		if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
			return errortype.ForbiddenError{Message: "Permission denied."}
		}

		return next(c)
	}
}
