package dto

type CreateOrderDTO struct {
	Method      string                 `json:"method" binding:"required" example:"transfer"`
	Status      string                 `json:"status" binding:"required" example:"pending"`
	UserId      uint                   `json:"user_id" binding:"required" example:"1" swagger:"-"`
	OrderDetail []CreateOrderDetailDTO `json:"order_detail" binding:"required"`
}

type CreateOrderDetailDTO struct {
	OrderId  uint `json:"order_id" binding:"required" example:"1"`
	ItemId   uint `json:"item_id" binding:"required" example:"1"`
	Quantity int  `json:"quantity" binding:"required" example:"2"`
}

type UpdateMethodDTO struct {
	Method string `json:"method" binding:"required" example:"cod"`
}

type UpdateStatusDTO struct {
	Status string `json:"status" binding:"required" example:"delivered"`
}

type CategoryDTO struct {
	Name string `json:"name"`
}

type ItemDTO struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Desc     string      `json:"desc"`
	Price    float64     `json:"price"`
	Category CategoryDTO `json:"category"`
}

// type OrderDetailDTO struct {
// 	ID       uint    `json:"id"`
// 	ItemID   uint    `json:"item_id"`
// 	Quantity int     `json:"quantity"`
// 	Total    float64 `json:"total"`
// }

type OrderResponseDTO struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	// OrderDetails []OrderDetailDTO `json:"order_details"`
}
