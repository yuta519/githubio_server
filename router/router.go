package router

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/yuta519/githubio_server/controller/blogs"
	"github.com/yuta519/githubio_server/utils"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.CorsHandler(w)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/blogs", blogs.Index)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	port := "8000"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
