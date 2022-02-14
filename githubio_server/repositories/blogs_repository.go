package repositories

import (
	"path/filepath"

	"github.com/yuta519/githubio_server/infra"
)

type Blog struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func FetchBlogs() []Blog {
	var blogs []Blog
	blog_files := infra.FetchS3Objects("md-host-bucket")
	for _, blog_file := range blog_files {
		blogs = append(
			blogs,
			Blog{
				Title: getFileNameWithoutExtension(blog_file),
				Url:   infra.FetchUrlOfS3Object("md-host-bucket", blog_file),
			},
		)
	}
	return blogs
}

func getFileNameWithoutExtension(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
