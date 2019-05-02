package models

import (
	"github.com/bb3104/salmon/db"
)

type RssArticle struct {
	RssTypeName string `dynamo:"rss_type_name"`
	Title       string `dynamo:"title"`
	Description string `dynamo:"description"`
	Image       string `dynamo:"image"`
	Url         string `dynamo:"url"`
	PublishedAt int64  `dynamo:"published_at"`
	CreatedAt   int64  `dynamo:"created_at"`
}

func (h RssArticle) GetAll() (*[]RssArticle, error) {
	db := db.GetDB()
	table := db.Table("engineee")

	var results *[]RssArticle

	err := table.Scan().All(&results)
	return results, err
}
