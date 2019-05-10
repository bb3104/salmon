package server

import (
	"github.com/bb3104/salmon/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := router()
	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("/v1")
	{
		rss_article := new(controllers.RssArticleController)
		v1.GET("/rss_articles", rss_article.Index)
	}

	return router
}
