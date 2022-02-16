package blogs

import (
	"encoding/json"
	"net/http"

	"github.com/yuta519/githubio_server/usecase/blogs"
	"github.com/yuta519/githubio_server/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.CorsHandler(w)
	json.NewEncoder(w).Encode(blogs.FetchBlogs())
}
