package handlers

import (
	"fitness-api/cmd/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCategories(ctx echo.Context) error {
	categories, err := h.categoryRepositories.GetAll()
	if err != nil {
		log.Errorf("An internal server error occurred when getting all categories!")
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, categories)
}

func (h *Handler) GeCateByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("Invalid data!")
		return ctx.JSON(http.StatusBadRequest, err)
	}
	category, err := h.categoryRepositories.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Errorf("Category not found by ID: %d", id)
			return ctx.JSON(http.StatusNotFound, fmt.Sprintf("Category not found by ID: %d", id))

		}
		log.Errorf("An internal server error occurred when getting category by ID!")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, category)
}

func (h *Handler) CreateCategory(ctx echo.Context) error {
	category := models.Category{}
	ctx.Bind(&category)
	exist, err := h.categoryRepositories.ExistsByName(category.Name)
	if err != nil {
		log.Errorf("internal server error: %v", err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if exist {
		log.Errorf("Category already exists!")
		return ctx.JSON(http.StatusBadRequest, "Category already exists!")
	}
	_, err = h.categoryRepositories.Create(&category)
	if err != nil {
		log.Errorf("An internal server error occurred when creating category!")
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, category)
}

func (h *Handler) UpdateCategory(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("Invalid data!")
		return ctx.JSON(http.StatusBadRequest, err)
	}
	category := models.Category{}
	ctx.Bind(&category)
	_, err = h.categoryRepositories.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Errorf("Category not found by ID: %d", id)
			return ctx.JSON(http.StatusNotFound, fmt.Sprintf("Category not found by ID: %d", id))
		}
		log.Errorf("An internal server error occurred when getting category by ID!")
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	err = h.categoryRepositories.Update(id, category)
	if err != nil {
		log.Errorf("An internal server error occurred when updating category!")
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, category)
}
