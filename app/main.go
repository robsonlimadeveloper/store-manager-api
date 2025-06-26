package main

import (
	"log"
	"store-manager-api/app/database"
	"store-manager-api/app/modules/establishment"
	"store-manager-api/app/modules/store"

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

	// Rotas
	e.POST("/login", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"message": "Login simulado com sucesso"})
	})

	establishment.RegisterRoutes(e, db)
	store.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))

}
