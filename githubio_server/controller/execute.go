package githubio_server

import (
	"log"
	"net/http"

	"github.com/yuta519/githubio_server/usecase/blogs"
)

func Execute() {
	http.HandleFunc("/", blogs.FetchBlogs)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	log.Fatal(http.ListenAndServe(":8000", nil))
}
