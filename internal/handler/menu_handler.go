package handler

import (
	"context"
	"errors"
	"log/slog"
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
	Price       *int   `json:"price" binding:"required"`
	Stock       *int   `json:"stock" binding:"required"`
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
	if req.Price == nil || req.Stock == nil {
		// PriceとStockのnilはJSONデコード時に弾かれるので通常このコードは通らない
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}
	id, err := h.svc.CreateNewMenu(c.Request.Context(), req.Name, req.Description, *req.Price, *req.Stock)
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
		// TODO: 将来，ミドルウェアでログ出力する
		slog.ErrorContext(c.Request.Context(), "create menu failed",
			"err", err, "path", c.FullPath())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
