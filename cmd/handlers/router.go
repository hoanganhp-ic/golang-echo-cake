package handlers

import (
	"fitness-api/cmd/jwt"
	"fitness-api/cmd/utils"
	"os"

	"github.com/labstack/echo/v4"
)

func InitTest(e *echo.Echo) {
	e.GET("/", Home)
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

	// create cake group
	cake := apiGroup.Group("/cakes", jwtMiddleware)
	// cake routes
	cake.POST("", h.Create)
	// cake.GET("", h.Get)
	cake.GET("/search", h.Search)
	cake.GET("/:id", h.GetByID)
	cake.DELETE("/:id", h.DeleteByID)
	cake.PUT("/:id", h.UpdateByID)

	// create category group
	apiCate := apiGroup.Group("/categories")
	apiCate.GET("", h.GetAllCategories)
	apiCate.GET("/:id", h.GeCateByID)
	apiCate.POST("", h.CreateCategory)
	apiCate.PUT("/:id", h.UpdateCategory)

	// Serve static files (profile pictures) from the 'picture' directory.
	e.Static("/picture", os.Getenv("PATH_TO_UPLOAD"))
}
