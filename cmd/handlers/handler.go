package handlers

import (
	repositoryimpl "fitness-api/cmd/repositories/repositoryImpl"
)

type Handler struct {
	userRepositoryImpl *repositoryimpl.UserRepositoryImpl
}

func NewHandler(ur *repositoryimpl.UserRepositoryImpl) *Handler {
	return &Handler{userRepositoryImpl: ur}
}
