package handlers

import (
	"fitness-api/cmd/dto"
	"fitness-api/cmd/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
)

func (h *Handler) Create(c echo.Context) error {
	id := userIdFromToken(c)
	cake := models.Cake{}
	c.Bind(&cake)
	cake.UserID = int(id)
	cake.Name = c.FormValue("name")
	cake.Description = c.FormValue("description")
	cake.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)

	file, err := c.FormFile("image")
	if err != nil {
		log.Errorf("Invalid data!")
	}
	fileName, err := saveFile(file)
	if err != nil {
		log.Errorf("An internal server error occurred when saving the image!")
	} else {
		cake.ImageUrl = fileName
	}

	// err = h.cakeRepositoryImpl.Create(cake)
	err = h.cakeRepository.Create(cake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}

// func (h *Handler) Get(c echo.Context) error {
// 	cake, err := h.cakeRepository.GetAll()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	return c.JSON(http.StatusOK, cake)
// }

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
	cakes, err := h.cakeRepository.Search(dto.SearchCake{
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
	cake, err := h.cakeRepository.GetByID(idInt)
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
	err = h.cakeRepository.DeleteByID(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (h *Handler) UpdateByID(c echo.Context) error {

	idUser := userIdFromToken(c)
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cake, _ := h.cakeRepository.GetByID(idInt)
	nameC := c.FormValue("name")
	descriptionC := c.FormValue("description")
	cake.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)
	cake.UserID = int(idUser)
	if nameC != "" {
		cake.Name = nameC
	}

	if descriptionC != "" {
		cake.Description = descriptionC
	}

	file, err := c.FormFile("image")
	if file != nil {
		if err != nil {
			log.Errorf("Invalid data!")
		} else {
			fileName, err := saveFile(file)
			if err != nil {
				log.Errorf("An internal server error occurred when saving the image!")
			} else {
				cake.ImageUrl = fileName
			}
		}
	}

	err = h.cakeRepository.UpdateByID(idInt, cake)
	if err != nil {
		log.Errorf("An internal server error occurred when updating the cake!")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cake)
}

func saveFile(file *multipart.FileHeader) (string, error) {

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	var fileName string = xid.New().String() + filepath.Ext(file.Filename)
	dst, err := os.Create(os.Getenv("PATH_TO_UPLOAD") + fileName)
	if err != nil {
		return "", err
	}

	// Copy the uploaded content to the destination file.
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}
