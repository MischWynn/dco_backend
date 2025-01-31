package services

import (
	"dco_mart/dto"
	"dco_mart/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Preload("User").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) GetByID(id string) (models.Order, error) {
	var order models.Order
	if err := s.db.Preload("User").First(&order, id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *OrderService) Create(input dto.CreateOrderDTO) (models.Order, error) {
	// var orderDetails []models.OrderDetail
	// var totalAmount float32
	// for _, detail := range input.OrderDetail {
	// 	// Check if the item exists
	// 	var item models.Item
	// 	if err := s.db.First(&item, detail.ItemId).Error; err != nil {
	// 		// If item not found, return error
	// 		return models.Order{}, fmt.Errorf("item with ID %d does not exist", detail.ItemId)
	// 	}

	// 	// Calculate total = price * quantity
	// 	total := float32(item.Price) * float32(detail.Quantity)
	// 	totalAmount += total

	// 	// Append the order detail with calculated total
	// 	orderDetails = append(orderDetails, models.OrderDetail{
	// 		OrderId:  detail.OrderId,
	// 		ItemId:   detail.ItemId,
	// 		Quantity: detail.Quantity,
	// 		Total:    float64(total),
	// 	})
	// }

	// var user models.User
	// if err := s.db.First(&user, input.UserId).Error; err != nil {
	// 	// If item not found, return error
	// 	return models.Order{}, fmt.Errorf("item with ID %d does not exist", input.UserId)
	// }

	// // Now create the order model
	// order := models.Order{
	// 	UserId:       input.UserId,
	// 	TotalAmount:  float64(totalAmount),
	// 	Status:       input.Status,
	// 	Method:       input.Method,
	// 	OrderDetails: orderDetails,
	// }

	// // Create the order in the database
	// if err := s.db.Create(&order).Error; err != nil {
	// 	return models.Order{}, err
	// }

	return models.Order{}, nil
}

func (s *OrderService) UpdateMethod(id string, input dto.UpdateMethodDTO) (models.Order, error) {
	var order models.Order
	if err := s.db.Preload("User").First(&order, id).Error; err != nil {
		return models.Order{}, err
	}

	if order.Method == input.Method {
		return models.Order{}, errors.New("Metode pembayaran pesanan dengan Id " + id + " sudah sesuai")
	}

	if order.Status != "pending" {
		return models.Order{}, fmt.Errorf("pesanan dengan Id %s sudah dibayar", id)
	}

	order.Method = input.Method
	if err := s.db.Save(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil

}

func (s *OrderService) UpdateStatus(id string, input dto.UpdateStatusDTO) (models.Order, error) {
	var order models.Order
	if err := s.db.Preload("User").First(&order, id).Error; err != nil {
		return models.Order{}, err
	}

	if order.Status == "delivered" || order.Status == "cancelled" {
		return models.Order{}, fmt.Errorf("status pesanan dengan Id %s tidak bisa diubah", id)
	}

	order.Status = input.Status
	if err := s.db.Save(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil

}

func (s *OrderService) Delete(id string) error {
	if err := s.db.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *OrderService) CheckoutOrder(request dto.CreateOrderDTO) (models.Order, error) {
	var order models.Order
	order.UserID = request.UserId
	var totalAmount float64
	for _, detail := range request.OrderDetail {
		var item models.Item
		if err := s.db.First(&item, detail.ItemId).Error; err != nil {
			return models.Order{}, fmt.Errorf("item with ID %d does not exist", detail.ItemId)
		}
		total := float64(item.Price) * float64(detail.Quantity)
		totalAmount += total
	}

	order.TotalAmount = totalAmount
	order.Status = request.Status
	order.Method = request.Method
	order.OrderDetails = make([]models.OrderDetail, len(request.OrderDetail))
	for i, detail := range request.OrderDetail {
		order.OrderDetails[i] = models.OrderDetail{
			OrderID:  detail.OrderId,
			ItemID:   detail.ItemId,
			Quantity: detail.Quantity,
		}
	}
	if err := s.db.Create(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}
