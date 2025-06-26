package establishment

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	g := e.Group("/establishments")
	g.GET("", getAll(service))
	g.GET("/:id", getByID(service))
	g.POST("", create(service))
	g.PUT("/:id", update(service))
	g.DELETE("/:id", delete(service))
}

func getAll(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao listar estabelecimentos"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Lista de estabelecimentos", "data": data})
	}
}

func getByID(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		est, err := service.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Estabelecimento não encontrado"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Estabelecimento encontrado", "data": est})
	}
}

func create(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var est Establishment
		if err := c.Bind(&est); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Dados inválidos"})
		}
		if err := service.Create(est); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao criar estabelecimento"})
		}
		return c.JSON(http.StatusCreated, echo.Map{"message": "Estabelecimento criado com sucesso"})
	}
}

func update(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var est Establishment
		if err := c.Bind(&est); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Dados inválidos"})
		}
		est.ID = id
		if err := service.Update(est); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Erro ao atualizar estabelecimento"})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Estabelecimento atualizado com sucesso"})
	}
}

func delete(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Estabelecimento removido com sucesso"})
	}
}
