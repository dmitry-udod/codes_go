package tests

import (
	router "github.com/dmitry-udod/codes_go/app/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestFindFopById(t *testing.T) {
	checkConnectionToEsServer(t)

	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/view/" + TEST_FOP_ID, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), TEST_FOP_NAME)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/fop/view/NOT_VALID_ID", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	assert.NotContains(t, w.Body.String(), TEST_FOP_NAME)
}

func TestFopLatest(t *testing.T) {
	checkConnectionToEsServer(t)
	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/latest", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), TEST_FOP_NAME)
	assert.Contains(t, w.Body.String(), TEST_FOP_ID)
	assert.Contains(t, w.Body.String(), "metadata")
	assert.Contains(t, w.Body.String(), `"total":`)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/fop/latest?page=2", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.NotContains(t, w.Body.String(), TEST_FOP_NAME)
}
