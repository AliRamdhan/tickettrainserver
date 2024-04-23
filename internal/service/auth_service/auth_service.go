package auth_service

import (
	"time"

	"github.com/AliRamdhan/trainticket/auth"
	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (au *AuthService) GetAllUser() ([]model.User, error) {
	var user []model.User
	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// func (au *AuthService) RegisterAuth(user *model.User) error {
// 	user.UserID = uuid.New()
// 	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
// 	return config.DB.Create(user).Error
// }

func (au *AuthService) RegisterUser(user *model.User) error {
	// Hash user password
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}

	// Create user record
	user.RoleUser = 2
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	record := config.DB.Create(user)
	if record.Error != nil {
		return record.Error
	}

	return nil
}

func (au *AuthService) LoginAuth(email, password string) (*model.User, error) {
	var user model.User
	// Check if email exists
	record := config.DB.Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}

	// Check if password is correct
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (au *AuthService) GenerateToken(email, username string, roleId int) (string, error) {
	return auth.GenerateJWT(email, username, roleId)
}

//	func (au *AuthService) Home() string {
//		return "Welcome to the home page!"
//	}

func (au *AuthService) Home(tokenString string) (string, error) {
	// Validate the token
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Here you can use claims.Username or claims.Email to retrieve user information from your database or wherever it's stored
	// For demonstration purposes, let's just return the username
	return claims.Username, nil
}

// func (au *AuthService) Home(userID uuid.UUID) (string, error) {
// 	// Retrieve user data based on userID
// 	var user model.User
// 	if err := config.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
// 		return "", err
// 	}

// 	// Construct welcome message with user data
// 	message := "Welcome " + user.Username + " to the home page!"
// 	return message, nil
// }

// func (au *AuthService) LoginAuth() ([]model.User, error) {
// 	var user []model.User
// 	if err := config.DB.Find(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (au *AuthService) LogoutAuth() ([]model.User, error) {
// 	var customer []model.User
// 	if err := config.DB.Find(&customer).Error; err != nil {
// 		return nil, err
// 	}
// 	return customer, nil
// }

// func (au *AuthService) HomeAuth() ([]model.User, error) {
// 	var customer []model.User
// 	if err := config.DB.Find(&customer).Error; err != nil {
// 		return nil, err
// 	}
// 	return customer, nil
// }
