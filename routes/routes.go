package routes

import (
	"github.com/dfryer1193/thehardway/handlers"
	"github.com/dfryer1193/thehardway/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the Gin router and routes
func SetupRoutes(router *gin.Engine) {
	// Post routes
	router.POST("/posts", handlers.CreatePost)
	router.GET("/posts", handlers.GetPublishedPosts)
	router.PATCH("/posts/:id", handlers.UpdatePost)

	// Comment routes
	router.POST("/posts/:id/comments", handlers.AddComment)
	router.DELETE("/comments/:id", handlers.DeleteComment)
	router.POST("/ban", handlers.BanUser)

	// Public routes for login and 2FA
	router.POST("/login", handlers.PasswordLogin)
	router.POST("/2fa", handlers.Complete2FA)

	// Protected routes for the site owner (requires authentication)
	ownerRoutes := router.Group("/")
	ownerRoutes.Use(middleware.AuthMiddleware())
	{
		ownerRoutes.POST("/change-password", handlers.ChangePassword)
		// Other protected routes (create post, delete comment, etc.)
	}
}
