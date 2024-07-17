package route

import (
	"fitness-api/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func InitTest(e *echo.Echo) {
	e.GET("/", handlers.Home)
}

func InitRoutes(e *echo.Echo) {
	e.POST("/cake", handlers.Create)
	e.GET("/cake", handlers.Get)
	e.GET("/cake/search", handlers.Search)
}
