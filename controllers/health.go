package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	fmt.Println("hogehoge")
	c.String(http.StatusOK, "Working!")
}
