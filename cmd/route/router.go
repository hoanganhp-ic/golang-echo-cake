package route

import (
	"fitness-api/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func InitTest(e *echo.Echo) {
	e.GET("/", handlers.Home)
}

func InitRoutes(e *echo.Echo) {
	// e.POST("/cakes", handlers.Create)
	// e.GET("/cakes", handlers.Get)
	// e.GET("/cakes/search", handlers.Search)

	// create api group
	apiGroup := e.Group("/api")

	// create cake group
	cake := apiGroup.Group("/cakes")

	// cake routes
	cake.POST("", handlers.Create)
	cake.GET("", handlers.Get)
	cake.GET("/search", handlers.Search)
	cake.GET("/:id", handlers.GetByID)
	cake.DELETE("/:id", handlers.DeleteByID)
}
