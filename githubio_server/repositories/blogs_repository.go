package repositories

import (
	"path/filepath"

	"github.com/yuta519/githubio_server/domain"
	"github.com/yuta519/githubio_server/infra"
)

func FetchBlogs() []domain.Blog {
	var blogs []domain.Blog
	blog_files := infra.FetchS3Objects("md-host-bucket")
	for _, blog_file := range blog_files {
		blogs = append(
			blogs,
			domain.Blog{
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
