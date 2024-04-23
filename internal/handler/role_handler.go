package handler

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(ps *service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: ps}
}

func (th *RoleHandler) CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.roleService.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "role created successfully", "product": role})
}

func (th *RoleHandler) GetAllRole(c *gin.Context) {
	roles, err := th.roleService.GetAllRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All roles", "roles": roles})
}

func (th *RoleHandler) UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleIdStr := c.Param("roleId")
	// roleID, err := uuid.Parse(roleIdStr)
	var roleId uint
	_, err := fmt.Sscanf(roleIdStr, "%d", &roleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := th.roleService.UpdateRole(roleId, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "role updated successfully", "product": role})
}

func (th *RoleHandler) DeleteRole(c *gin.Context) {
	roleIdStr := c.Param("roleId")
	// productID, err := uuid.Parse(roleIdStr)
	var roleId uint
	_, err := fmt.Sscanf(roleIdStr, "%d", &roleId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := th.roleService.DeleteRole(roleId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
