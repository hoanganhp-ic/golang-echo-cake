package main

import (
	"fitness-api/cmd/handlers"
	repositoryimpl "fitness-api/cmd/repositories/repositoryImpl"
	"fitness-api/cmd/storage"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Test route
	handlers.InitTest(e)

	// connect to database
	storage.InitDB()
	db := storage.GetDB()
	us := repositoryimpl.NewUserRepositoryImpl(db)
	h := handlers.NewHandler(us)
	h.Register(e)

	// middleware
	e.Use(handlers.LogRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	// Init routes
	handlers.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
