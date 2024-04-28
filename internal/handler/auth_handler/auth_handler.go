package auth_handler

import (
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service/auth_service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth_service.AuthService
}

func NewAuthHandler(au *auth_service.AuthService) *AuthHandler {
	return &AuthHandler{authService: au}
}

func (au *AuthHandler) GetAllUser(c *gin.Context) {
	users, err := au.authService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All users", "User": users})
}

func (au *AuthHandler) RegisterAuth(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register user
	if err := au.authService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": user.UserID, "email": user.Email, "username": user.Username})
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (au *AuthHandler) Login(c *gin.Context) {
	var request TokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	user, err := au.authService.LoginAuth(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate token
	tokenString, err := au.authService.GenerateToken(user.Email, user.Username, int(user.RoleUser))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user})
}

// func (au *AuthHandler) Home(context *gin.Context) {
// 	message := au.authService.Home()
// 	context.JSON(http.StatusOK, gin.H{"message": message})
// }

func (au *AuthHandler) Home(context *gin.Context) {
	// Retrieve the token from the request header
	tokenString := context.GetHeader("Authorization")
	if tokenString == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Call the Home function in the auth service to retrieve user information
	username, err := au.authService.Home(tokenString)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the username in the response
	context.JSON(http.StatusOK, gin.H{"message": "Welcome, " + username + "!"})
}
