package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cake := models.Cake{}
	c.Bind(&cake)
	cake, err := repositories.Create(cake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}
