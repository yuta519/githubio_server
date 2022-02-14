package githubio_server

import (
	"net/http"

	"github.com/yuta519/githubio_server/infra"
)

func FetchBlogs(w http.ResponseWriter, r *http.Request) {
	infra.FetchS3Objects("md-host-bucket")
}
