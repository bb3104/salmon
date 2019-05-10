package models

import (
	"testing"
)

func TestGetAllSuccess(t *testing.T) {
	rssArticleModel := new(RssArticle)
	_, err := rssArticleModel.GetAll()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
