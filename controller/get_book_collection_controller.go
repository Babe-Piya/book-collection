package controller

import (
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	volStr := c.Query("volume")
	if volStr != "" {
		vol, err = strconv.Atoi(volStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	priceStr := c.Query("price")
	if priceStr != "" {
		price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
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
	}

	c.JSON(http.StatusOK, resp)
}
