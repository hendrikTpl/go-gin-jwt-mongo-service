package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/hendrikTpl/go-gin-jwt-mongo-service/controllers"
)

func AuthRoutes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controllers.SignUp)
        auth.POST("/login", controllers.Login)
    }
}
