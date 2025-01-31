package dto

type CreateCategoryDTO struct {
	Name  string `json:"name" binding:"required" example:"Example Category"`
	Photo []byte `json:"photo" binding:"required"` // Example: a byte array
}

type UpdateCategoryDTO struct {
	Name  string `json:"name" binding:"required" example:"Example Update Category"`
	Photo []byte `json:"photo" binding:"required"` // Example: a byte array
}
