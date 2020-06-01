package tests

import (
	"github.com/dmitry-udod/codes_go/app/router"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	r.ServeHTTP(w, req)

	assert := is.New(t)
	assert.Equal(200, w.Code)
	assert.True(strings.Contains(w.Body.String(), "pong"))
}

func TestFindFopById(t *testing.T) {
	assert := is.New(t)
	checkConnectionToEsServer(t)
	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/view/" + TEST_FOP_ID, nil)
	r.ServeHTTP(w, req)
	assert.Equal(200, w.Code)
	assert.True(strings.Contains(w.Body.String(), TEST_FOP_NAME))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/fop/view/NOT_VALID_ID", nil)
	r.ServeHTTP(w, req)
	assert.Equal(404, w.Code)
	assert.True(! strings.Contains(w.Body.String(), TEST_FOP_NAME))
}

func TestFopLatest(t *testing.T) {
	checkConnectionToEsServer(t)
	r := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/latest", nil)
	r.ServeHTTP(w, req)

	assert := is.New(t)
	assert.Equal(200, w.Code)
	assert.True(strings.Contains(w.Body.String(), "full_name"))
	assert.True(strings.Contains(w.Body.String(), "metadata"))
	assert.True(strings.Contains(w.Body.String(), "total"))
}

func TestFopSearchQuery(t *testing.T) {
	checkConnectionToEsServer(t)
	r := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/fop/latest?q=Столяров" , nil)
	r.ServeHTTP(w, req)

	assert := is.New(t)
	assert.Equal(200, w.Code)
	assert.True(strings.Contains(w.Body.String(), "full_name"))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/fop/latest?q=" + TEST_FOP_NAME, nil)
	r.ServeHTTP(w, req)
	assert.Equal(200, w.Code)
	assert.True(strings.Contains(w.Body.String(), "full_name"))
}
