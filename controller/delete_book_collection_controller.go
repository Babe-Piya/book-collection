package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *bookCollectionController) DeleteBookCollectionByID(c *gin.Context) {
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

	resp, err := ctrl.BookCollectionService.DeleteBookCollectionByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
