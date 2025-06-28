package middleware

// Middleware for JWT authentication in Echo framework
// This middleware checks for a valid JWT token in the Authorization header.
// If the token is valid, it allows the request to proceed; otherwise, it returns an error response.

import (
	"net/http"
	"strings"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

// Middleware JWT Authentication
func CustomJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Authorization header is missing or invalid",
				})
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse JWT token
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// Confirm if the signing method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Algoritmo de assinatura inválido")
				}
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid token",
				})
			}

			// Se quiser acessar claims:
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("user", claims)
			}

			// Chama o próximo handler
			err = next(c)
			if err != nil {
				c.Logger().Error("Handler error:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Internal server error",
				})
			}

			return nil
		}
	}
}
