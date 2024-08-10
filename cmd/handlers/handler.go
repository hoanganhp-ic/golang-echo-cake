package handlers

import (
	"fitness-api/cmd/repositories"
)

type Handler struct {
	userRepository repositories.UserRepositories
	cakeRepository repositories.CakeRepositories
}

func NewHandler(ur repositories.UserRepositories, cr repositories.CakeRepositories) *Handler {
	return &Handler{userRepository: ur, cakeRepository: cr}
}
