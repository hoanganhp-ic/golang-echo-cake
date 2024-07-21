package handlers

import (
	"fitness-api/cmd/jwt"
	"fitness-api/cmd/utils"

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
	jwtMiddleware := jwt.JWT(utils.JWTSecret)
	// create api group
	apiGroup := e.Group("/api")
	auth := apiGroup.Group("/auth")
	auth.POST("/register", h.SignUp)
	auth.POST("/login", h.Login)

	// create user group
	user := apiGroup.Group("/users", jwtMiddleware)
	user.GET("/current", h.CurrentUser)
}
