package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userService := user.NewService(userRepo, a.Config.JWTSecret)
	newUserHandler := user.NewHandler(userService)

	r.POST("/register", newUserHandler.RegisterUser)
	r.POST("/login", newUserHandler.LoginUser)

	api := r.Group("/api")
	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/protected", func(c *gin.Context) {
		userID, _ := middleware.GetUserID(c)
		role, _ := middleware.GetUserRole(c)
		c.JSON(200, gin.H{
			"message": "This is a protected route",
			"userID":  userID,
			"role":    role,
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RequireAdmin())

	admin.GET("/restricted", func(c *gin.Context) {
		role, _ := middleware.GetUserRole(c)
		if role != "admin" {
			c.JSON(403, gin.H{"error": "Forbidden: Admins only"})
			return
		}
		c.JSON(200, gin.H{
			"Role":    role,
			"message": "Welcome to the admin restricted route!",
		})
	})

	return r
}
