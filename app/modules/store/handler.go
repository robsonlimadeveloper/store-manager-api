package store

// Package store provides handlers for managing stores in the store manager API.

import (
	"database/sql"
	"net/http"
	"strconv"
	"store-manager-api/app/core"
	"github.com/labstack/echo/v4"
	dto "store-manager-api/app/modules/store/dto"
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

// @Summary Lista todas as lojas
// @Description Retorna a lista de todas as lojas cadastradas
// @Tags stores
// @Produce json
// @Success 200 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /api/stores [get]
func getAll(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		stores, err := service.GetAll()

		if err != nil {
			c.Logger().Error("Erro ao listar lojas: ", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrListStoresFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Lista de lojas", "data": stores})
	}
}

// @Summary Obtém loja por ID
// @Tags stores
// @Produce json
// @Param id path int true "ID da loja"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 404 {object} core.JsonResponse
// @Router /api/stores/{id} [get]
func getByID(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
		}
		store, err := service.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": ErrNotFound.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": store})
	}
}

// @Summary Cria uma nova loja
// @Tags stores
// @Accept json
// @Produce json
// @Param store body dto.CreateStoreDTO true "Dados da loja"
// @Success 201 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /api/stores [post]
func create(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input dto.CreateStoreDTO

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, core.JsonResponse{
				Message: ErrInvalidInput.Error(),
				Error:   err.Error(),
			})
		}

		store := Store{
			Number:          input.Number,
			Name:            input.Name,
			CorporateName:   input.CorporateName,
			Address:         input.Address,
			City:            input.City,
			State:           input.State,
			ZipCode:         input.ZipCode,
			StreetNumber:    input.StreetNumber,
			EstablishmentID: input.EstablishmentID,
		}

		if err := service.Create(&store); err != nil {
			return c.JSON(http.StatusInternalServerError, core.JsonResponse{
				Message: ErrInternalServer.Error(),
				Error:   err.Error(),
			})
		}

		response := dto.StoreResponseDTO{
			ID:              store.ID,
			Number:          store.Number,
			Name:            store.Name,
			CorporateName:   store.CorporateName,
			Address:         store.Address,
			City:            store.City,
			State:           store.State,
			ZipCode:         store.ZipCode,
			StreetNumber:    store.StreetNumber,
			EstablishmentID: store.EstablishmentID,
		}

		return c.JSON(http.StatusCreated, core.JsonResponse{
			Message: "Loja criada com sucesso",
			Data:    response,
		})
	}
}

// @Summary Atualiza loja existente
// @Tags stores
// @Accept json
// @Produce json
// @Param id path int true "ID da loja"
// @Param store body dto.UpdateStoreDTO true "Dados para atualização"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /api/stores/{id} [put]
func update(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidID.Error()})
		}

		var s Store
		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidInput.Error()})
		}
		s.ID = id

		if err := service.Update(&s); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrUpdateFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Loja atualizada com sucesso"})
	}
}

// @Summary Remove uma loja
// @Tags stores
// @Produce json
// @Param id path int true "ID da loja"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /api/stores/{id} [delete]
func delete(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidID.Error()})
		}
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrDeleteFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": "Loja removida com sucesso"})
	}
}
