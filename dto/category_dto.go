package dto

type CreateCategoryDTO struct {
	Name string `json:"name" binding:"required" example:"Example Category"`
}

type UpdateCategoryDTO struct {
	Name string `json:"name" binding:"required" example:"Example Update Category"`
}
