package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/route"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Test route
	route.InitTest(e)

	// connect to database
	storage.InitDB()

	// middleware
	e.Use(handlers.LogRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))	  

	// Init routes
	route.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
