package controllers

import (
	"dco_mart/dto"
	"dco_mart/models"
	"dco_mart/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// @Summary   Get All Order
// @Tags      Order
// @Router    /order [get]
// @Security BearerAuth
func (oc *OrderController) GetAll(ctx echo.Context) error {
	orders, err := oc.orderService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, orders)
}

// @Summary   Get Order By ID
// @Tags      Order
// @Param     id   path  string  true  "Order ID"
// @Router    /order/{id} [get]
// @Security BearerAuth
func (oc *OrderController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Id Order tidak boleh kosong"})
	}

	order, err := oc.orderService.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, order)
}

// @Summary   Create Order
// @Tags      Order
// @Param     order  body      dto.CreateOrderDTO true  "Create Order Data"
// @Router    /order [post]
// @Security BearerAuth
func (oc *OrderController) Create(ctx echo.Context) error {
	var input dto.CreateOrderDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := oc.orderService.Create(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, order)
}

// @Summary   Update Method Payment
// @Tags      Order
// @Param     id   path  string  true  "Order ID"
// @Param     order  body      dto.UpdateMethodDTO true  "Update Method Payment"
// @Router    /order/{id}/method [put]
// @Security BearerAuth
func (oc *OrderController) UpdateMethod(ctx echo.Context) error {
	id := ctx.Param("id")
	var input dto.UpdateMethodDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := oc.orderService.UpdateMethod(id, input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, order)
}

// @Summary   Update Status Order
// @Tags      Order
// @Param     id   path  string  true  "Order ID"
// @Param     order  body      dto.UpdateStatusDTO true  "Update Status Order"
// @Router    /order/{id}/status [put]
// @Security BearerAuth
func (oc *OrderController) UpdateStatus(ctx echo.Context) error {
	id := ctx.Param("id")
	var input dto.UpdateStatusDTO
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := oc.orderService.UpdateStatus(id, input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, order)
}

// @Summary   Checkout Order
// @Tags      Order
// @Param     order   body  dto.CreateOrderDTO  true  "Order payload"
// @Router    /order/checkout [post]
// @Security BearerAuth
func (oc *OrderController) CheckoutOrder(ctx echo.Context) error {
	request := dto.CreateOrderDTO{}
	user := ctx.Get("user").(models.User)
	request.UserId = user.ID

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	order, err := oc.orderService.CheckoutOrder(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, order)
}
