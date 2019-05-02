package server

import (
	"github.com/bb3104/salmon/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	health := new(controllers.HealthController)

	r.GET("/health", health.Status)

	u := r.Group("/v1")
	{
		rss_article := new(controllers.RssArticleController)
		u.GET("/rss_articles", rss_article.Index)
	}

	return r
}
