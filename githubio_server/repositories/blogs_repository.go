package repositories

import (
	"github.com/yuta519/githubio_server/infra"
)

func FetchBlogs(bucket_name string) {
	infra.FetchS3Objects(bucket_name)
}
