package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", HealthCheck)
	r.POST("/menu/new", CreateNewMenu)
	return r
}
