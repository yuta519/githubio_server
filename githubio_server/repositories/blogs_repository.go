package repositories

import (
	"github.com/yuta519/githubio_server/infra"
)

func FetchBlogs() []string {
	var blog_urls []string
	blog_files := infra.FetchS3Objects("md-host-bucket")
	for _, blog_file := range blog_files {
		blog_urls = append(blog_urls, infra.FetchUrlOfS3Object("md-host-bucket", blog_file))
	}
	return blog_urls
}
