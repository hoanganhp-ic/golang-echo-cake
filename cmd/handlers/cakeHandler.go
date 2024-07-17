package handlers

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"fmt"
	"net/http"
	"strconv"

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

func Get(c echo.Context) error {
	cake, err := repositories.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}

func Search(c echo.Context) error {
	name := c.QueryParam("name")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("page_size")

	var page, pageSize int
	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if page < 1 {
			page = 1
		}
	} else {
		page = 1
	}

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if pageSize < 1 {
			pageSize = 3
		}
	} else {
		pageSize = 3
	}
	fmt.Println(page, pageSize)
	cakes, err := repositories.Search(dto.SearchCake{
		Name:     name,
		Page:     page,
		PageSize: pageSize,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cakes)
}
