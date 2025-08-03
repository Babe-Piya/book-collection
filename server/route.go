package server

import (
	"github.com/gin-gonic/gin"
	"github/Babe-piya/book-collection/appconfig"
	"gorm.io/gorm"
	"net/http"
)

func Routes(r *gin.Engine, db *gorm.DB, config *appconfig.AppConfig) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "service available",
		})
	})
}
