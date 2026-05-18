package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"db-go.com/api/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	db.InitDB()
	os.Exit(m.Run())
}

func TestEvents(t *testing.T) {
	router := gin.New()
	RegisterRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/events", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
