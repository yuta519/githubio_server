package blogs

import (
	"github.com/yuta519/githubio_server/domain"
	"github.com/yuta519/githubio_server/repositories"
)

func FetchBlogs() []domain.Blog {
	var blogs []domain.Blog
	blog_files := repositories.FetchBlogFiles()
	for _, blog_file := range blog_files {
		blogs = append(
			blogs,
			domain.Blog{
				Title: repositories.GetFileNameWithoutExtension(blog_file),
				Url:   repositories.GgetUrlOfBlogFile(blog_file),
			},
		)
	}
	return blogs
}
