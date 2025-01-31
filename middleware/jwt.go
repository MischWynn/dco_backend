package middleware

import (
	"dco_mart/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// IsValidJWT is the middleware for validating JWT tokens
func IsValidJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please Login First")
		}

		// Expecting the token in the form of "Bearer <token>"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		// Remove "Bearer " prefix
		tokenString := authHeader[len("Bearer "):]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the token method (HMAC)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(os.Getenv("TOKEN_KEY")), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Error parsing token: %s", err.Error()))
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// Extract claims from the token
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		// Safely convert claims to user model
		user := models.User{
			ID:   uint((*claims)["id"].(float64)), // Ensure conversion to uint
			Name: (*claims)["name"].(string),
			Role: (*claims)["role"].(string),
		}
		c.Set("user", user)

		// Continue to the next handler if the token is valid
		return next(c)
	}
}
