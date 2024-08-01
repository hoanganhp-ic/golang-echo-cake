package handlers

import (
	repositoryImpl "fitness-api/cmd/repositories/repositoryImpl"
)

type Handler struct {
	userRepositoryImpl *repositoryImpl.UserRepositoryImpl
	cakeRepositoryImpl *repositoryImpl.CakeRepositoryImpl
}

func NewHandler(ur *repositoryImpl.UserRepositoryImpl, cr * repositoryImpl.CakeRepositoryImpl) *Handler {
	return &Handler{userRepositoryImpl: ur, cakeRepositoryImpl: cr}
}
