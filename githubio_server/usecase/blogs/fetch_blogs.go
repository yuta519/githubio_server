package blogs

import (
	"strings"

	"github.com/yuta519/githubio_server/domain"
	"github.com/yuta519/githubio_server/repositories"
)

func FetchBlogs() []domain.Blog {
	var blogs []domain.Blog
	blog_files := repositories.FetchBlogFiles()
	for _, blog_file := range blog_files {
		blog_title := strings.Replace(blog_file, "_", " ", -1)
		blogs = append(
			blogs,
			domain.Blog{
				Title: repositories.GetFileNameWithoutExtension(blog_title),
				Url:   repositories.GgetUrlOfBlogFile(blog_file),
			},
		)
	}
	return blogs
}
