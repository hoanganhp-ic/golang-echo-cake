package request

import (
	"fitness-api/cmd/models"

	"github.com/labstack/echo/v4"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *UserLogin) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return err
	}
	return nil
}

func (r *UserRegister) Bind(ctx echo.Context, u *models.User) error {
	if err := ctx.Bind(r); err != nil {
		return err
	}
	u.Name = r.Name
	u.Email = r.Email
	hp, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = hp
	return nil
}
