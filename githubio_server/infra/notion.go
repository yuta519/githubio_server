package infra

import (
	"log"
	"os"

	"github.com/yuta519/notion_api"
)

func ExportNotionPages() {
	database_ids := notion_api.FetchDatabaseIds(os.Getenv("NOTION_SECRET"))
	log.Println(database_ids)
}
