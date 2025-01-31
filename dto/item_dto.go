package dto

type CreateItemDTO struct {
	Name       string  `json:"name" binding:"required" example:"Example Item"`
	Desc       string  `json:"desc" binding:"required" example:"Example Desc Item"`
	Price      float64 `json:"price" binding:"required" example:"1.00"`
	Photo      []byte  `json:"photo" binding:"required"` // Example: a byte array
	CategoryId uint    `json:"category_id" binding:"required" example:"1"`
}

type UpdateItemDTO struct {
	Name       string  `json:"name" binding:"required" example:"Example Update Item"`
	Desc       string  `json:"desc" binding:"required" example:"Example Update Desc Item"`
	Price      float64 `json:"price" binding:"required" example:"2.00"`
	Photo      []byte  `json:"photo" binding:"required"` // Example: a byte array
	CategoryId uint    `json:"category_id" binding:"required" example:"2"`
}
