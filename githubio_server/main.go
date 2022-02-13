package main

import (
	controller "github.com/yuta519/githubio_server/controller"
	githubio_server "github.com/yuta519/githubio_server/infra"
)

func main() {
	controller.Router()
	githubio_server.FetchS3Objects("")
}
