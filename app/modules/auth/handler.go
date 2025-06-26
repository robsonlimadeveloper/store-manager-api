package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("secret123") // 🔐 troque em produção!

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Dados inválidos"})
	}

	// Simula autenticação
	if req.Username != "admin" || req.Password != "123" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Credenciais inválidas"})
	}

	// Cria token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 dias
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao gerar token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}
