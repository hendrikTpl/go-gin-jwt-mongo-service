package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/hendrikTpl/go-gin-jwt-mongo-service/db"
	"github.com/hendrikTpl/go-gin-jwt-mongo-service/models"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))
var validate = validator.New()

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "details": err.Error()})
		return
	}

	// Validate user input based on struct tags
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}

	collection := db.GetCollection("users")

	// Check if email already exists
	count, err := collection.CountDocuments(context.Background(), bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = new(string)
	*user.Password = string(hashedPassword)

	// Add timestamps and ID
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert user into DB
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "details": err.Error()})
		return
	}

	var user models.User
	collection := db.GetCollection("users")
	err := collection.FindOne(context.Background(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": *user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
