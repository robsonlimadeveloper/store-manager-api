package auth

// Package auth provides authentication functionalities for the application.
// It includes user login handling and JWT token generation.
// The package uses the Echo framework for HTTP handling and JWT for token management.

import (
	"net/http"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
	}

	// Simula autenticação
	if req.Username != "admin" || req.Password != "123" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": ErrInvalidCredentials.Error()})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrTokenExpired.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}
