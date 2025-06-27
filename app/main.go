package main

// main.go
// This is the entry point of the Store Manager API application.
// It initializes the Echo framework, sets up middleware, connects to the database,
// and registers routes for authentication, establishments, and stores.

import (
	"store-manager-api/app/database"
	"store-manager-api/app/modules/establishment"
	"store-manager-api/app/modules/store"
	"store-manager-api/app/modules/auth"
	"store-manager-api/app/middleware"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "store-manager-api/app/docs"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	
	// Inicializa o banco de dados
	db, err := database.Init()
	if err != nil {
		e.Logger.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Rota pública
	e.POST("/login", auth.Login)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Rota de verificação de saúde
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"status": "OK"})
	})
	
	// Rotas protegidas
	api := e.Group("/api")
	api.Use(middleware.CustomJWTMiddleware())

	// Roteamento dentro do grupo protegido
	establishment.RegisterRoutes(api, db)
	store.RegisterRoutes(api, db)

	e.Logger.Fatal(e.Start(":8080"))

}
