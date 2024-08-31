package handlers

import (
	"fitness-api/cmd/repositories"
)

type Handler struct {
	userRepository repositories.UserRepositories
	cakeRepository repositories.CakeRepositories
	categoryRepositories repositories.CategoryRepositories
}

func NewHandler(ur repositories.UserRepositories, cr repositories.CakeRepositories, cgr repositories.CategoryRepositories) *Handler {
	return &Handler{userRepository: ur, cakeRepository: cr, categoryRepositories: cgr}
}
