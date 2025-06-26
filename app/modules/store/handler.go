package store

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	g.GET("/stores", getAll(service))
	g.GET("/stores/:id", getByID(service))
	g.POST("/stores", create(service))
	g.PUT("/stores/:id", update(service))
	g.DELETE("/stores/:id", delete(service))
}

func getAll(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		stores, err := service.GetAll()

		if err != nil {
			c.Logger().Error("Erro ao listar lojas: ", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao listar lojas"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Lista de lojas", "data": stores})
	}
}

func getByID(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
		}
		store, err := service.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Loja não encontrada"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Loja encontrada", "data": store})
	}
}

func create(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var s Store
		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Dados inválidos"})
		}
		if err := service.Create(s); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao criar loja"})
		}
		return c.JSON(http.StatusCreated, echo.Map{"message": "Loja criada com sucesso"})
	}
}

func update(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
		}

		var s Store
		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Dados inválidos"})
		}
		s.ID = id

		if err := service.Update(s); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao atualizar loja"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Loja atualizada com sucesso"})
	}
}

func delete(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID inválido"})
		}
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao remover loja"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Loja removida com sucesso"})
	}
}
