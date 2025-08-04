package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *bookCollectionController) UpdateBookCollectionByID(c *gin.Context) {
	ctx := c.Request.Context()

	var id int
	var err error
	idStr := c.Param("id")
	if idStr != "" {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			slog.Error("parse id fail", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
	}

	var req service.UpdateBookCollectionRequest
	err = c.BindJSON(&req)
	if err != nil {
		slog.Error("UpdateBookCollectionByID BindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}
	req.ID = id

	resp, err := ctrl.BookCollectionService.UpdateBookCollectionByID(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
