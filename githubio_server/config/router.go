package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuta519/githubio_server/controller/blogs"
	"github.com/yuta519/githubio_server/controller/ping"
)

func Router(port string) {
	http.HandleFunc("/", ping.Ping)
	http.HandleFunc("/blogs", blogs.Index)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
