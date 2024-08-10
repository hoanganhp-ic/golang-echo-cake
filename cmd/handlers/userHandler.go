package handlers

import (
	"errors"
	"fitness-api/cmd/dto/request"
	"fitness-api/cmd/dto/response"
	"fitness-api/cmd/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) SignUp(ctx echo.Context) error {
	var u models.User
	req := &request.UserRegister{}
	if err := req.Bind(ctx, &u); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.userRepository.Create(&u); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, response.NewUserResponse(&u))
}

func (h *Handler) Login(ctx echo.Context) error {
	req := &request.UserLogin{}
	if err := req.Bind(ctx); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	u, err := h.userRepository.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusForbidden, "Invalid email or password")
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		return ctx.JSON(http.StatusForbidden, "Invalid email or password")
	}
	if !u.CheckPassword(req.Password) {
		return ctx.JSON(http.StatusForbidden, "Invalid email or password")
	}
	return ctx.JSON(http.StatusOK, response.NewUserResponse(u))

}

func (h *Handler) CurrentUser(ctx echo.Context) error {
	id := userIdFromToken(ctx)
	u, err := h.userRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, "User not found")
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}
	return ctx.JSON(http.StatusOK, u)
}

func userIdFromToken(ctx echo.Context) uint {
	id, ok := ctx.Get("userId").(uint)
	if !ok {
		return 0
	}
	return id
}
