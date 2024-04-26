package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	t.Run("No error", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return nil
		}
		handler := AuthMiddleware(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.SetBasicAuth(os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"))
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("No Basic Auth", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return nil
		}
		handler := AuthMiddleware(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("Wrong Basic Auth", func(t *testing.T) {
		mockHandler := func(c echo.Context) error {
			return nil
		}
		handler := AuthMiddleware(mockHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.SetBasicAuth("test", "test")
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		assert.NoError(t, handler(c))
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}
