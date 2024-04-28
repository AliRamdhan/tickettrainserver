package service

import (
	"time"

	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetAllUser() ([]model.User, error) {
	var user []model.User
	if err := config.DB.Preload("Role").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUserById(userId uint) (*model.User, error) {
	var user model.User
	if err := config.DB.Preload("Role").First(&user, "user_id = ?", userId).Error; err != nil {
		return nil, err // Ticket not found
	}
	return &user, nil
}

func (us *UserService) UpdateUser(userId uint, roleId uint, updateUser *model.User) error {
	var existingUser model.User
	if err := config.DB.First(&existingUser, "user_id = ?", userId).Error; err != nil {
		return err // Product not found
	}
	existingUser.Username = updateUser.Username
	existingUser.Email = updateUser.Email
	existingUser.RoleUser = roleId
	existingUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	// Save the updated product
	return config.DB.Save(&existingUser).Error
}

func (us *UserService) DeleteUser(userId uint) error {
	var user model.User
	if err := config.DB.First(&user, "user_id = ?", userId).Error; err != nil {
		return err // Product not found
	}
	return config.DB.Delete(user).Error
}
