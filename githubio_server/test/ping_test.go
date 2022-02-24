package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yuta519/githubio_server/controller/ping"
)

func TestPing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ping.Ping))
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ts.Config.Handler.ServeHTTP(w, r)

	expectedResponse := "Ping, \"/\""

	if w.Code != http.StatusOK {
		t.Errorf("Expected: %d, got %d", http.StatusOK, w.Code)
	}
	if response := w.Body.String(); response != expectedResponse {
		t.Errorf("Expected: %s, got %s", expectedResponse, response)
	}
}
