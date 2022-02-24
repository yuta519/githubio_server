package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yuta519/githubio_server/controller/blogs"
)

func TestBlogs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(blogs.Index))
	r := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	w := httptest.NewRecorder()
	ts.Config.Handler.ServeHTTP(w, r)

	// expectedResponse := "[{\"filename\":\"Why_I_use_Notion!.md\",\"title\":\"Why I use Notion!\",\"url\":\"https://md-host-bucket.s3.us-east-2.amazonaws.com/Why_I_use_Notion!.md\"}]"

	if w.Code != http.StatusOK {
		t.Errorf("Expected: %d, got %d", http.StatusOK, w.Code)
	}
	// if response := w.Body.String(); response != expectedResponse {
	// 	t.Errorf("Expected: %s, got %s", expectedResponse, response)
	// }
}
