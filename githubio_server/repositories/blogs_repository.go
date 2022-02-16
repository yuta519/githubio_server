package repositories

import (
	"path/filepath"

	"github.com/yuta519/githubio_server/infra"
)

func FetchBlogFiles() []string {
	return infra.FetchS3Objects("md-host-bucket")
}

func GgetUrlOfBlogFile(filename string) string {
	return infra.FetchUrlOfS3Object("md-host-bucket", filename)
}

func GetFileNameWithoutExtension(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
