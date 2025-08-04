package server

import (
	"gorm.io/gorm"
	"net/http"

	"github/Babe-piya/book-collection/controller"
	"github/Babe-piya/book-collection/repositories"
	"github/Babe-piya/book-collection/service"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	bookCollectionRepo := repositories.NewBookCollection(db)
	bookCollectionService := service.NewBookCollectionService(bookCollectionRepo)
	bookCollectionController := controller.NewBookCollectionController(bookCollectionService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "service available",
		})
	})

	bookCollectApi := r.Group("/api/v1/book-collection")
	bookCollectApi.POST("/create", bookCollectionController.CreateBookCollection)
	bookCollectApi.GET("", bookCollectionController.GetBookCollectionByFilter)
	bookCollectApi.PUT("update/:id", bookCollectionController.UpdateBookCollectionByID)
}
