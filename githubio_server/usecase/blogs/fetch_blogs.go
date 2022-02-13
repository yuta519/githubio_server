package blogs

import (
	"net/http"

	"github.com/yuta519/githubio_server/repositories"
)

func FetchBlogs(w http.ResponseWriter, r *http.Request) {
	repositories.FetchBlogs("md-host-bucket")
}
