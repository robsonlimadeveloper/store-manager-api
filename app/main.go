package main

import (
	"log"
	"store-manager-api/app/database"
	"store-manager-api/app/modules/establishment"
	"store-manager-api/app/modules/store"
	"store-manager-api/app/modules/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	

	db, err := database.Init()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Rota pública
	e.POST("/login", auth.Login)

	// Rotas protegidas
	api := e.Group("/api")
	api.Use(auth.JWTMiddleware())

	// Roteamento dentro do grupo protegido
	establishment.RegisterRoutes(api, db)
	store.RegisterRoutes(api, db)

	e.Logger.Fatal(e.Start(":8080"))

}
