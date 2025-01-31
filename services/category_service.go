package services

import (
	"dco_mart/dto"
	"dco_mart/models"
	"errors"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return CategoryService{db: db}
}

func (s CategoryService) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Model(&models.Category{}).Preload("Items").Find(&categories).Error
	return categories, err
}

func (s CategoryService) GetByID(id string) (models.Category, error) {
	var category models.Category
	err := s.db.Model(&models.Category{}).Preload("Items").First(&category, id).Error
	return category, err
}

func (s CategoryService) GetImageByID(id string) (models.Category, error) {
	var category models.Category
	if err := s.db.Model(&models.Category{}).Select("photo").First(&category, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (s CategoryService) Create(input dto.CreateCategoryDTO) (models.Category, error) {
	// Check if a category with the same name already exists
	var existingCategory models.Category
	if err := s.db.Where("name = ?", input.Name).First(&existingCategory).Error; err == nil {
		return models.Category{}, errors.New("category already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Category{}, err
	}

	// Create a new Category model from the input data
	category := models.Category{Name: input.Name, Photo: input.Photo}
	if err := s.db.Create(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (s CategoryService) Update(id string, input dto.UpdateCategoryDTO) (models.Category, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return models.Category{}, err
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	if err := s.db.Save(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (s CategoryService) Delete(id string) error {
	return s.db.Delete(&models.Category{}, id).Error
}
