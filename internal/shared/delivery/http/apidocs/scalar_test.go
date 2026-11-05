package apidocs

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	RegisterRoutes(r)

	t.Run("UI", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "@scalar/api-reference")
	})

	t.Run("spec", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/swagger/doc.json", nil))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.True(t, strings.Contains(w.Body.String(), `"paths"`))
	})

	t.Run("redirect", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/swagger", nil))
		assert.Equal(t, http.StatusMovedPermanently, w.Code)
	})
}
