package controllers

import (
	"dco_mart/dto"
	"dco_mart/models"
	"dco_mart/services"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return CategoryController{categoryService: categoryService}
}

// @Summary   Get All Categories
// @Tags      Category
// @Router    /category [get]
// @Security BearerAuth
func (c CategoryController) GetAll(ctx echo.Context) error {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, categories)
}

// @Summary   Get Category By ID
// @Tags      Category
// @Param     id   path  string  true  "Category ID"
// @Router    /category/{id} [get]
// @Security  BearerAuth
func (c CategoryController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Id kategori tidak boleh kosong"})
	}

	category, err := c.categoryService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	// Return the retrieved category
	return ctx.JSON(http.StatusOK, category)
}

func (c CategoryController) GetImageByID(ctx echo.Context) error {
	id := ctx.Param("id")
	category, err := c.categoryService.GetImageByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.Blob(http.StatusOK, "image/jpeg", category.Photo)
}

// @Summary   Create Category
// @Tags      Category
// @Param     category  body      dto.CreateCategoryDTO true  "Create Category Data"
// @Router    /category [post]
// @Security BearerAuth
func (c CategoryController) Create(ctx echo.Context) error {
	userClaims := ctx.Get("user").(models.User)
	if userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	name := ctx.FormValue("name")
	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid file : " + err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Read file content
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	input := dto.CreateCategoryDTO{
		Name:  name,
		Photo: fileBytes,
	}
	category, err := c.categoryService.Create(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, category)
}

// @Summary   Update Category
// @Tags      Category
// @Param     id   path  string  true  "Category ID"
// @Param     category  body      dto.CreateCategoryDTO true  "Create Category Data"
// @Router    /category/{id} [put]
// @Security BearerAuth
func (c CategoryController) Update(ctx echo.Context) error {
	userClaims := ctx.Get("user").(models.User)
	if userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	id := ctx.Param("id")
	var input dto.UpdateCategoryDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	category, err := c.categoryService.Update(id, input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, category)
}

// @Summary   Delete Category
// @Tags      Category
// @Param     id   path  string  true  "Category ID"
// @Router    /category/{id} [delete]
// @Security BearerAuth
func (c CategoryController) Delete(ctx echo.Context) error {
	userClaims := ctx.Get("user").(models.User)
	if userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	id := ctx.Param("id")
	if err := c.categoryService.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, "Kategori berhasil dihapus")
}
