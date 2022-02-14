package blogs

import (
	"fmt"
	"net/http"

	"github.com/yuta519/githubio_server/repositories"
)

func FetchBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := repositories.FetchBlogs()
	fmt.Println(blogs)
}
