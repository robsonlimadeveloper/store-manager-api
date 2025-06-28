package routes

import (
	"database/sql"
	"store-manager-api/app/middleware"
	"store-manager-api/app/modules/auth"
	"store-manager-api/app/modules/establishment"
	"store-manager-api/app/modules/store"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	RegisterLoginRoutes(e)
	RegisterSwaggerRoutes(e)
	RegisterHealthCheckRoute(e)

	// Protected routes
	api := e.Group("/v1/api")
	api.Use(middleware.CustomJWTMiddleware())
	
	RegisterModulesRoutes(api, db)
}

// RegisterPublicRoutes registers the public routes for the application.
func RegisterLoginRoutes(e *echo.Echo) {
	e.POST("/v1/api/login", auth.Login)
}

// RegisterSwaggerRoutes registers the Swagger documentation route for the application.
func RegisterSwaggerRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

// RegisterHealthCheckRoute registers the health check route for the application.
func RegisterHealthCheckRoute(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"status": "OK"})
	})
}

// RegisterModulesRoutes registers the routes for the modules in the application.
// This function is responsible for registering the routes for the store and establishment modules.
func RegisterModulesRoutes(api *echo.Group, db *sql.DB) {

	// TODO: Add all modules automatically
	store.RegisterRoutes(api, db)
	establishment.RegisterRoutes(api, db)
}