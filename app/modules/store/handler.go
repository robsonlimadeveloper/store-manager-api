package store

// Package store provides handlers for managing stores in the store manager API.

import (
	"database/sql"
	"net/http"
	"store-manager-api/app/core"
	"store-manager-api/app/utils"
	dto "store-manager-api/app/modules/store/dto"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"fmt"
)

var validate = validator.New()

type StoreRequest struct {
	Number         string `json:"number" validate:"required"`
	Name           string `json:"name" validate:"required,min=2,max=100"`
	CorporateName  string `json:"corporate_name" validate:"required"`
	Address        string `json:"address" validate:"required"`
	City           string `json:"city" validate:"required"`
	State          string `json:"state" validate:"required,len=2"` // Sigla de 2 letras
	ZipCode        string `json:"zip_code" validate:"required,len=8"` // Ex: 12345678
	StreetNumber   string `json:"street_number" validate:"required"`
	EstablishmentID int    `json:"establishment_id" validate:"required,gt=0"`
}

func RegisterRoutes(g *echo.Group, db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)

	g.GET("/stores", getAll(service))
	g.GET("/stores/:id", getByID(service))
	g.POST("/stores", create(service))
	g.PUT("/stores/:id", update(service))
	g.DELETE("/stores/:id", delete(service))
	g.GET("/establishments/:id/stores", getByEstablishmentID(service))
}

// @Summary List all stores
// @Description Returns a list of all stores registered in the system.
// @Tags Stores
// @Produce json
// @Security BearerAuth
// @Success 200 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/stores [get]
func getAll(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		stores, err := service.GetAll()

		if err != nil {
			c.Logger().Error("Error: ", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrListStoresFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": stores})
	}
}

// @Summary Get store by ID
// @Description Returns the store with the specified ID.
// @Tags Stores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Store ID"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 404 {object} core.JsonResponse
// @Router /v1/api/stores/{id} [get]
func getByID(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.ParseIDParam(c, "id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidID.Error()})
		}
		
		store, err := service.GetByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": store})
	}
}

// @Summary Create a new store
// @Description Creates a new store with the provided data.
// @Tags Stores
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param store body dto.CreateStoreDTO true "Store data"
// @Success 201 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/stores [post]
func create(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input dto.CreateStoreDTO

		// Bind JSON para struct
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, core.JsonResponse{
				Message: ErrInvalidInput.Error(),
				Error:   err.Error(),
			})
		}

		// Validação dos campos usando go-playground/validator
		if err := validate.Struct(input); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errors := make([]string, 0)
			for _, e := range validationErrors {
				errors = append(errors, fmt.Sprintf("Field '%s' failed on the '%s' validation", e.Field(), e.Tag()))
			}
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Validation failed",
				"errors":  errors,
			})
		}

		// Criação do store no banco
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

		return c.JSON(http.StatusCreated, echo.Map{
			"data": response,
		})
	}
}

// @Summary Update an existing store
// @Description Updates the store with the specified ID using the provided data.
// @Tags Stores
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Store ID"
// @Param store body dto.UpdateStoreDTO true "Update data for the store"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/stores/{id} [put]
func update(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.ParseIDParam(c, "id")
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
		return c.JSON(http.StatusOK, echo.Map{"data": dto.StoreResponseDTO{
			ID:              s.ID,
			Number:          s.Number,
			Name:            s.Name,
			CorporateName:   s.CorporateName,
			Address:         s.Address,
			City:            s.City,
			State:           s.State,
			ZipCode:         s.ZipCode,
			StreetNumber:    s.StreetNumber,
			EstablishmentID: s.EstablishmentID,
		}})
	}
}

// @Summary Remove a store
// @Description Removes the store with the specified ID.
// @Tags Stores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Store ID"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/stores/{id} [delete]
func delete(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.ParseIDParam(c, "id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidID.Error()})
		}
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrDeleteFailed.Error()})
		}
		return c.JSON(http.StatusOK, echo.Map{"data": dto.StoreResponseDTO{
			ID: id,
		}})
	}
}

// @Summary Get stores by establishment ID
// @Description Returns a list of stores associated with the specified establishment ID.
// @Tags Stores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Establishment ID"
// @Success 200 {object} core.JsonResponse
// @Failure 400 {object} core.JsonResponse
// @Failure 500 {object} core.JsonResponse
// @Router /v1/api/stores/establishment/{id} [get]
func getByEstablishmentID(service StoreService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := utils.ParseIDParam(c, "id")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrInvalidID.Error()})
		}

		stores, err := service.GetByEstablishmentID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not fetch stores"})
		}

		return c.JSON(http.StatusOK, echo.Map{"data": stores})
	}
}
