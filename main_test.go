package main

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	actual := ResponseData{}
	json.Unmarshal(w.Body.Bytes(), &actual)
	assert.Equal(t, ResponseData{
		Message: "pong",
	}, actual)
}
