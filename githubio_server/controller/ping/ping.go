package ping

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/yuta519/githubio_server/infra"
	"github.com/yuta519/githubio_server/utils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	utils.CorsHandler(w)
	fmt.Fprintf(w, "Ping, %q", html.EscapeString(r.URL.Path))
	log.Printf("Ping, %q", html.EscapeString(r.URL.Path))
	infra.ExportNotionPages()
}
