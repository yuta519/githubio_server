package cmd

import (
	"log"

	"github.com/yuta519/githubio_server/config"
)

func Execute() {
	default_port := "8000"
	log.Printf("A server is starting on %s/TCP.", default_port)
	config.Router(default_port)
}
