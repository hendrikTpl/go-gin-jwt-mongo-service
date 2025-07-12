package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/hendrikTpl/go-gin-jwt-mongo-service/controllers"
    "github.com/hendrikTpl/go-gin-jwt-mongo-service/middleware"
)

func UserRoutes(router *gin.Engine) {
    user := router.Group("/user")
    user.Use(middleware.AuthMiddleware())
    {
        user.GET("/profile", controllers.UserProfile)
    }
}
