package controller

import (
	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

type BookCollectionController interface {
	CreateBookCollection(c *gin.Context)
	GetBookCollectionByFilter(c *gin.Context)
}

type bookCollectionController struct {
	BookCollectionService service.BookCollectionService
}

func NewBookCollectionController(service service.BookCollectionService) BookCollectionController {
	return &bookCollectionController{
		BookCollectionService: service,
	}
}
