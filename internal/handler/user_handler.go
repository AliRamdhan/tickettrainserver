package handler

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (uh *UserHandler) GetAllUser(c *gin.Context) {
	users, err := uh.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All users", "users": users})
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	userIdStr := c.Param("userId")
	var userId uint
	_, err := fmt.Sscanf(userIdStr, "%d", &userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id format"})
		return
	}
	users, err := uh.userService.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Details users", "users": users})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIdStr := c.Param("userId")
	var userId uint
	_, err := fmt.Sscanf(userIdStr, "%d", &userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id format"})
		return
	}
	// Extract train ticket ID from the request body
	roleId := user.RoleUser

	// Check if train ticket ID is missing or invalid
	if roleId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing user ID"})
		return
	}

	if err := uh.userService.UpdateUser(userId, roleId, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "user": user})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("userId")
	// productID, err := uuid.Parse(userIdStr)
	var userId uint
	_, err := fmt.Sscanf(userIdStr, "%d", &userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := uh.userService.DeleteUser(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
