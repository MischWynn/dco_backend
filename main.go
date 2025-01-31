package main

import (
	"dco_mart/config"
	_ "dco_mart/docs" // Import generated docs
	"dco_mart/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title DCO Mart API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.SetupRoutes(e)

	// Start the Echo server
	port := ":8080" // Choose your desired port
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
