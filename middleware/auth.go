package middleware

import (
	"net/http"
	"os"

	"github.com/anon-kae/assessment-tax/helper"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, password, ok := c.Request().BasicAuth()
		if !ok {
			return helper.FailedHandler(c, map[string]interface{}{
				"type":    "AuthenticationError",
				"message": "Unauthorized",
			}, http.StatusUnauthorized)
		}

		if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
			return helper.FailedHandler(c, map[string]interface{}{
				"type":    "ForbiddenError",
				"message": "Permission denied.",
			}, http.StatusForbidden)
		}

		return next(c)
	}
}
