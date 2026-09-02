package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Suiren91/go-cafeteria/internal/domain"
	"github.com/gin-gonic/gin"
)

type MenuService interface {
	CreateNewMenu(ctx context.Context, name, description string, price, stock int) (int, error)
	ListMenus(ctx context.Context) ([]*domain.Menu, error)
}

type MenuHandler struct {
	svc MenuService
}

type CreateNewMenuReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required,gte=0"`
	Stock       int    `json:"stock" binding:"required,gte=0"`
}

func NewMenuHandler(svc MenuService) *MenuHandler {
	return &MenuHandler{
		svc: svc,
	}
}

func (h *MenuHandler) CreateNewMenu(c *gin.Context) {
	var req CreateNewMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.svc.CreateNewMenu(c.Request.Context(), req.Name, req.Description, req.Price, req.Stock)
	// TODO: ログ書き込み処理
	if err != nil {
		if errors.Is(err, domain.ErrInvalidName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidName.Error()})
			return
		}
		if errors.Is(err, domain.ErrNegativePrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrNegativePrice.Error()})
			return
		}
		if errors.Is(err, domain.ErrNegativeStock) {
			c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrNegativeStock.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": " Unexpected error. Please try later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "New menu created",
		"id":      id,
	})
}
