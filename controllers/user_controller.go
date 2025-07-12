package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context) {
	userEmail := c.MustGet("email").(string)
	c.JSON(http.StatusOK, gin.H{
		"message": "Access granted to protected route",
		"email":   userEmail,
	})
}
