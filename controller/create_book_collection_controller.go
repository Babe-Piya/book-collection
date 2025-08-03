package controller

import (
	"log/slog"
	"net/http"

	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *bookCollectionController) CreateBookCollection(c *gin.Context) {
	ctx := c.Request.Context()

	var req service.BookCollectionRequest
	err := c.BindJSON(&req)
	if err != nil {
		slog.Error("BookCollectionController BindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	resp, err := ctrl.BookCollectionService.CreateBookCollection(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
