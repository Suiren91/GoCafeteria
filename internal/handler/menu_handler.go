package handler

import (
	"net/http"

	"github.com/Suiren91/go-cafeteria/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateNewMenuReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required,gte=0"`
	Stock       int    `json:"stock" binding:"required,gte=0"`
}

func CreateNewMenu(c *gin.Context) {
	var req CreateNewMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO: エラーハンドリング(テスト先行で)
	service.CreateNewMenu(req.Name, req.Description, req.Price, req.Stock)

	c.JSON(http.StatusOK, gin.H{"message": "New menu created"})
}
