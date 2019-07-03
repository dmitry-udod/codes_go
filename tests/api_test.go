package tests

import (
	router2 "github.com/dmitry-udod/codes_go/app/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := router2.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestFindFopById(t *testing.T) {
	checkConnectionToEsServer(t)

	router := router2.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/" + TEST_FOP_ID, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), TEST_FOP_NAME)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/fop/NOT_VALID_ID", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	assert.NotContains(t, w.Body.String(), TEST_FOP_NAME)
}
