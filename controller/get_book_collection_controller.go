package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *bookCollectionController) GetBookCollectionByFilter(c *gin.Context) {
	ctx := c.Request.Context()

	var id, vol int
	var price float64
	var err error
	idStr := c.Query("id")
	if idStr != "" {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			slog.Error("parse id fail", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
	}

	volStr := c.Query("volume")
	if volStr != "" {
		vol, err = strconv.Atoi(volStr)
		if err != nil {
			slog.Error("parse vol fail", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
	}

	priceStr := c.Query("price")
	if priceStr != "" {
		price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			slog.Error("parse price fail", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		return
	}

	resp, err := ctrl.BookCollectionService.GetBookCollectionByFilter(ctx, service.GetBookCollection{
		ID:       id,
		BookName: c.Query("book_name"),
		Type:     c.Query("type"),
		Volume:   vol,
		Price:    price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
