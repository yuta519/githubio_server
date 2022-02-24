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
}
