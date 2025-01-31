package controllers

import (
	"dco_mart/dto"
	"dco_mart/models"
	"dco_mart/services"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	itemService services.ItemService
}

func NewItemController(itemService services.ItemService) ItemController {
	return ItemController{itemService: itemService}
}

// @Summary   Get All Item
// @Tags      Item
// @Router    /item [get]
// @Security BearerAuth
func (c ItemController) GetAll(ctx echo.Context) error {
	items, err := c.itemService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, items)
}

// @Summary   Get Item By ID
// @Tags      Item
// @Param     id   path  string  true  "Item ID"
// @Router    /item/{id} [get]
// @Security BearerAuth
func (c ItemController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Id Item tidak boleh kosong"})
	}

	item, err := c.itemService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, item)
}

func (c ItemController) GetItemImage(ctx echo.Context) error {
	id := ctx.Param("id")
	item, err := c.itemService.GetImageByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.Blob(http.StatusOK, "image/jpeg", item.Photo)
}

// @Summary   Create Item
// @Tags      Item
// @Param     item  body    dto.CreateItemDTO true  "Create Item Data"
// @Router    /item [post]
// @Security BearerAuth
func (c ItemController) Create(ctx echo.Context) error {
	// Extract user claims from the context
	userClaims, ok := ctx.Get("user").(models.User)
	if !ok || userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	category_id, err := strconv.Atoi(ctx.FormValue("category_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid category ID"})
	}

	price, err := strconv.ParseFloat(ctx.FormValue("price"), 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid price"})
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid file"})
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

	// Bind request data to DTO
	input := dto.CreateItemDTO{
		Name:       name,
		Desc:       desc,
		CategoryId: uint(category_id),
		Price:      price,
		Photo:      fileBytes,
	}
	item, err := c.itemService.Create(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, item)
}

// @Summary   Update Item
// @Tags      Item
// @Param     id   path  string  true  "Item ID"
// @Param     item  body      dto.UpdateItemDTO true  "Update Item Data"
// @Router    /item/{id} [put]
// @Security BearerAuth
func (c ItemController) Update(ctx echo.Context) error {
	userClaims := ctx.Get("user").(models.User)
	if userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	id := ctx.Param("id")
	var input dto.UpdateItemDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	item, err := c.itemService.Update(id, input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, item)
}

// @Summary   Delete Item
// @Tags      Item
// @Param     id   path  string  true  "Item ID"
// @Router    /item/{id} [delete]
// @Security BearerAuth
func (c ItemController) Delete(ctx echo.Context) error {
	userClaims := ctx.Get("user").(models.User)
	if userClaims.Role != "admin" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"message": "Anda bukan Admin"})
	}

	id := ctx.Param("id")
	if err := c.itemService.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Item berhasil di hapus")
}
