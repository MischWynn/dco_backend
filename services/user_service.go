// services/user_service.go
package services

import (
	"dco_mart/config"
	"dco_mart/dto"
	"dco_mart/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService instance
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// Register creates a new user in the database
func (s *UserService) Register(input dto.RegisterDTO) (models.User, error) {
	// Check if the email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return models.User{}, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}

	// Create the new user
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Phone:    input.Phone,
		Address:  input.Address,
		Role:     input.Role,
	}

	// Save the user to the database
	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Login checks if the provided email and password match a user in the database
func (s *UserService) Login(input dto.LoginDTO) (dto.LoginResponse, error) {
	var user models.User
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return dto.LoginResponse{}, errors.New("user not Found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid email or password")
	}

	token, err := config.GenerateToken(user)
	if err != nil {
		return dto.LoginResponse{}, errors.New("could not generate token")
	}

	return dto.LoginResponse{
		Message: "Login successful",
		User:    dto.UserResponseDTO{ID: user.ID, Name: user.Name, Role: user.Role},
		Token:   token,
	}, nil
}
