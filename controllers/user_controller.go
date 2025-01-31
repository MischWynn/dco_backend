package controllers

import (
	"dco_mart/dto"
	"dco_mart/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary   Register User
// @Tags      User
// @Accept    json
// @Produce   json
// @Param     user  body      dto.RegisterDTO true  "User Registration Data"
// @Router    /user/register [post]
func (c *UserController) Register(ctx echo.Context) error {
	var input dto.RegisterDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := c.userService.Register(input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response := dto.RegisterResponse{
		Message: "User registered successfully",
		User:    dto.UserResponseDTO{ID: user.ID, Name: user.Name, Role: user.Role},
	}

	return ctx.JSON(http.StatusCreated, response)
}

// @Summary   Login User
// @Tags      User
// @Param     user  body      dto.LoginDTO true  "User Login"
// @Router    /user/login [post]
func (c *UserController) Login(ctx echo.Context) error {
	var input dto.LoginDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := c.userService.Login(input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)

}
