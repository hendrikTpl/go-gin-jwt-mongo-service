package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/hendrikTpl/go-gin-jwt-mongo-service/db"
	"github.com/hendrikTpl/go-gin-jwt-mongo-service/routes"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	db.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
