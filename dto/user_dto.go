package dto

type RegisterDTO struct {
	Name     string `json:"name" binding:"required" example:"Admin"`
	Email    string `json:"email" binding:"required,email" example:"admin@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"admin123"`
	Phone    string `json:"phone" binding:"required" example:"1234567890"`
	Address  string `json:"address" example:"123 Main St"`
	Role     string `json:"role" example:"admin"`
}

type RegisterResponse struct {
	Message string          `json:"message"`
	User    UserResponseDTO `json:"user"`
}

// UserDTO represents a simplified version of the user for response purposes
type UserResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email" example:"admin@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"admin123"`
}

type LoginResponse struct {
	Message string          `json:"message"`
	User    UserResponseDTO `json:"user"`
	Token   string          `json:"token"`
}
