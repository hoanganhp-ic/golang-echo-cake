package handlers

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	id := userIdFromToken(c)
	cake := models.Cake{}
	c.Bind(&cake)
	cake.UserID = int(id)
	err := h.cakeRepositoryImpl.Create(cake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}

func (h *Handler) Get(c echo.Context) error {
	cake, err := repositories.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}

func (h *Handler) Search(c echo.Context) error {
	name := c.QueryParam("name")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("page_size")
	id := userIdFromToken(c)

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
	cakes, err := h.cakeRepositoryImpl.Search(dto.SearchCake{
		Name:     name,
		Page:     page,
		PageSize: pageSize,
		UserID:   int(id),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cakes)
}

func (h *Handler) GetByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cake, err := h.cakeRepositoryImpl.GetByID(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cake)
}

func (h *Handler) DeleteByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = repositories.DeleteByID(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Deleted")
}
