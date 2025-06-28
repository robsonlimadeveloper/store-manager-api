package establishment

// Package establishment provides handlers for the establishment module.
// These handlers are responsible for handling requests and responses for the establishment module, such as creating, updating, and deleting establishments.

import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	dto "store-manager-api/app/modules/establishment/dto"
)

var _ = dto.CreateEstablishmentDTO{}

func RegisterRoutes(g *echo.Group, db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	g.GET("/establishments", getAll(service))
	g.GET("/establishments/:id", getByID(service))
	g.POST("/establishments", create(service))
	g.PUT("/establishments/:id", update(service))
	g.DELETE("/establishments/:id", delete(service))
}

// @Summary List all establishments
// @Description Returns a list of all establishments registered in the system.
// @Tags Establishments
// @Produce json
// @Security BearerAuth
// @Success 200 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/establishments [get]
func getAll(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := service.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrListEstablishmentsFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": data})
	}
}

// @Summary Get establishment by ID
// @Description Returns the establishment with the specified ID.
// @Tags Establishments
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID do estabelecimento"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 404 {object} core.JsonResponse
// @Router /v1/api/establishments/{id} [get]
func getByID(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		est, err := service.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": ErrNotFound.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": est})
	}
}

// @Summary Create a new establishment
// @Description Creates a new establishment with the provided data.
// @Tags Establishments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param establishment body dto.CreateEstablishmentDTO true "Dados do estabelecimento"
// @Success 201 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/establishments [post]
func create(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var est Establishment
		if err := c.Bind(&est); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
		}
		if err := service.Create(&est); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrCreateFailed.Error()})
		}
		return c.JSON(http.StatusCreated, echo.Map{"message": "Estabelecimento criado com sucesso"})
	}
}

// @Summary Update an existing establishment
// @Description Updates the establishment with the specified ID using the provided data.
// @Tags Establishments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID do estabelecimento"
// @Param establishment body dto.UpdateEstablishmentDTO true "Dados para atualização"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/establishments/{id} [put]
func update(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var est Establishment
		if err := c.Bind(&est); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
		}
		est.ID = id
		if err := service.Update(&est); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrUpdateFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Estabelecimento atualizado com sucesso"})
	}
}

// @Summary Remove an establishment
// @Description Removes the establishment with the specified ID.
// @Tags Establishments
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID do estabelecimento"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/establishments/{id} [delete]
func delete(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrDeleteFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Estabelecimento removido com sucesso"})
	}
}
