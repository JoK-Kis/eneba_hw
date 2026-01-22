package handlers

import (
	"net/http"

	"github.com/JoK-Kis/eneba-backend/internal/database"
	"github.com/JoK-Kis/eneba-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	var products []models.Product

	if query == "" {
		if err := database.DB.Preload("Platform").Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
	} else {
		searchPattern := "%" + query + "%"
		if err := database.DB.Preload("Platform").
			Where("name ILIKE ? OR word_similarity(?, name) > 0.15", searchPattern, query).
			Order(gorm.Expr("word_similarity(?, name) DESC", query)).
			Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
			return
		}
	}

	c.JSON(http.StatusOK, products)
}