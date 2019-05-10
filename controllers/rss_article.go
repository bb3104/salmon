package controllers

import (
	"github.com/bb3104/salmon/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RssArticleController struct{}

var rssArticleModel = new(models.RssArticle)

func (u RssArticleController) Index(c *gin.Context) {
	articles, err := rssArticleModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to Index RssArticle", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rss data founded!", "articles": articles})
	return
}
