// @title Store Manager API
// @version 1.0
// @description API REST to manage stores and establishments.

// @contact.name Robson Soares
// @contact.url https://github.com/robsonlimadeveloper/
// @contact.email robsonlimadeveloper@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

// main.go
// This is the entry point of the Store Manager API application.
// It initializes the Echo framework, sets up middleware, connects to the database,
// and registers routes for authentication, establishments, and stores.

import (
	routes "store-manager-api/app/core/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "store-manager-api/app/docs"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"store-manager-api/app/database"
)

func main() {
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS()) // Allows all origins, methods, and headers by default
	
	// Start database connection
	db, err := database.Init()
	if err != nil {
		e.Logger.Fatal("Error to connect to database", err)
	}

	routes.RegisterRoutes(e, db)

	e.Logger.Info("Store Manager API is running on port 8080")
	
	e.Logger.Fatal(e.Start(":8080"))

}
