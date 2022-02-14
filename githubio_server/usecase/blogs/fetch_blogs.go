package blogs

import (
	"encoding/json"
	"net/http"

	"github.com/yuta519/githubio_server/repositories"
)

func FetchBlogs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repositories.FetchBlogs())
}
