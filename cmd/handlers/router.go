package handlers

import (
	"github.com/labstack/echo/v4"
)

func InitTest(e *echo.Echo) {
	e.GET("/", Home)
}

func InitRoutes(e *echo.Echo) {

	// create api group
	apiGroup := e.Group("/api")

	// create cake group
	cake := apiGroup.Group("/cakes")

	// cake routes
	cake.POST("", Create)
	cake.GET("", Get)
	cake.GET("/search", Search)
	cake.GET("/:id", GetByID)
	cake.DELETE("/:id", DeleteByID)
}

func (h *Handler) Register(e *echo.Echo) {
	// create api group
	apiGroup := e.Group("/api")	
	guestUsers := apiGroup.Group("/users")
	guestUsers.POST("/register", h.SignUp)
	guestUsers.POST("/login", h.Login)
}
