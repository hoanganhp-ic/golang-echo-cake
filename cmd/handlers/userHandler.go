package handlers

import (
	"fitness-api/cmd/dto/request"
	"fitness-api/cmd/dto/response"
	"fitness-api/cmd/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func CreateUser(c echo.Context) error {
// 	user := models.User{}
// 	c.Bind(&user)
// 	newUser, err := repositories.CreateUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusCreated, newUser)
// }

func (h *Handler) SignUp(ctx echo.Context) error {
	var u models.User
	req := &request.UserRegister{}
	if err := req.Bind(ctx, &u); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.userRepositoryImpl.Create(&u); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, response.NewUserResponse(&u))
}

func (h *Handler) Login(ctx echo.Context) error {
	req := &request.UserLogin{}
	if err := req.Bind(ctx); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	u, err := h.userRepositoryImpl.GetByEmail(req.Email)
	if err != nil {
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

func (h *Handler) CurrentUser(cxt echo.Context) error {
	id := userIdFromToken(cxt)
	u, err := h.userRepositoryImpl.GetByID(id)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, err.Error())
	}
	if u == nil {
		return cxt.JSON(http.StatusNotFound, "User not found")
	}
	return cxt.JSON(http.StatusOK, u)
}

func userIdFromToken(ctx echo.Context) uint {
	id, ok := ctx.Get("userId").(uint)
	if !ok {
		return 0
	}
	return id
}
