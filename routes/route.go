// routes/route.go
package routes

import (
	"dco_mart/config"
	"dco_mart/controllers"
	"dco_mart/middleware"
	"dco_mart/services"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Initialize services and controllers
	userService := services.NewUserService(config.DB)
	userController := controllers.NewUserController(userService)
	user := e.Group("/user")
	user.POST("/register", userController.Register)
	user.POST("/login", userController.Login)

	categoryService := services.NewCategoryService(config.DB)
	categoryController := controllers.NewCategoryController(categoryService)
	category := e.Group("/category", middleware.IsValidJWT)
	category.GET("", categoryController.GetAll)
	category.GET("/:id", categoryController.GetByID)
	category.POST("", categoryController.Create)
	category.PUT("/:id", categoryController.Update)
	category.DELETE("/:id", categoryController.Delete)

	itemService := services.NewItemService(config.DB)
	itemController := controllers.NewItemController(itemService)
	item := e.Group("/item", middleware.IsValidJWT)
	item.GET("", itemController.GetAll)
	item.GET("/:id", itemController.GetByID)
	item.POST("", itemController.Create)
	item.PUT("/:id", itemController.Update)
	item.DELETE("/:id", itemController.Delete)
	item.GET("/image/:id", itemController.GetItemImage)

	orderService := services.NewOrderService(config.DB)
	orderController := controllers.NewOrderController(orderService)
	order := e.Group("/order", middleware.IsValidJWT)
	order.GET("", orderController.GetAll)
	order.GET("/:id", orderController.GetByID)
	order.POST("", orderController.Create)
	order.PUT("/:id/method", orderController.UpdateMethod)
	order.PUT("/:id/status", orderController.UpdateStatus)
	order.POST("/checkout", orderController.CheckoutOrder)
}
