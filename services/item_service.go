package services

import (
	"dco_mart/dto"
	"dco_mart/models"

	"gorm.io/gorm"
)

type ItemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) ItemService {
	return ItemService{db: db}
}

func (s ItemService) GetAll() ([]models.Item, error) {
	var items []models.Item
	if err := s.db.Find(&items).Preload("Category").Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s ItemService) GetByID(id string) (models.Item, error) {
	var item models.Item
	if err := s.db.First(&item, id).Preload("Category").Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (s ItemService) GetImageByID(id string) (models.Item, error) {
	var item models.Item
	if err := s.db.Model(&models.Item{}).Select("photo").First(&item, "id = ?", id).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (s ItemService) Create(input dto.CreateItemDTO) (models.Item, error) {
	var category models.Category
	if err := s.db.First(&category, "id = ?", input.CategoryId).Error; err != nil {
		return models.Item{}, err
	}
	item := models.Item{
		Name:     input.Name,
		Desc:     input.Desc,
		Price:    input.Price,
		Photo:    input.Photo,
		Category: category,
	}
	if err := s.db.Create(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (s ItemService) Update(id string, input dto.UpdateItemDTO) (models.Item, error) {
	var item models.Item
	if err := s.db.First(&item, id).Error; err != nil {
		return models.Item{}, err
	}
	item.Name = input.Name
	item.Desc = input.Desc
	item.Price = input.Price
	// item.CategoryId = input.CategoryId
	if err := s.db.Save(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (s ItemService) Delete(id string) error {
	return s.db.Delete(&models.Item{}, id).Error
}
