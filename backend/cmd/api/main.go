package main

import (
	"github.com/JoK-Kis/eneba-backend/config"
	"github.com/JoK-Kis/eneba-backend/internal/database"
	"github.com/JoK-Kis/eneba-backend/internal/handlers"
	"github.com/JoK-Kis/eneba-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.Migrate()
	database.Seed()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.CORS())
	router.GET("/api/search", handlers.SearchProducts)
	router.Run(":" + config.GetPort())
}