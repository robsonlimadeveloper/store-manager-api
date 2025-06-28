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

// LoginRequest represents the structure of the login request.
//
// @Description Structure containing the username and password for user authentication.
type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123"`
}

// LoginResponse represents the structure of the login response.
//
// @Description Structure containing the JWT token returned upon successful login.
// @Description The token is used for subsequent authenticated requests.
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// ErrorResponse represents the structure of an error response.
//
// @Description Structure containing an error message.
type ErrorResponse struct {
	Message string `json:"message" example:"Invalid credentials"`
}

// @Summary User Authentication
// @Description Handles user login and returns a JWT token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login request with username and password"
// @Success 200 {object} LoginResponse "Token generated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input data"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /v1/api/login [post]
func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
	}

	// Authenticate simulation
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
