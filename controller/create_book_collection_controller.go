package controller

import (
	"net/http"

	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *bookCollectionController) CreateBookCollection(c *gin.Context) {
	ctx := c.Request.Context()

	var req service.BookCollectionRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	resp, err := ctrl.BookCollectionService.CreateBookCollection(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, resp)
}
