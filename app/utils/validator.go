package utils

import (
	"strconv"
	"github.com/labstack/echo/v4"
	"errors"
)

// ParseIDParam extracts and validates an integer ID from the request context.
// It returns the ID if valid, or an error if the ID is not a valid integer or is less than or equal to zero.
func ParseIDParam(c echo.Context, paramName string) (int, error) {
	idStr := c.Param(paramName)
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID parameter: " + paramName + " must be a positive integer")
	}
	return id, nil
}