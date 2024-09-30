package routes

import (
	"github.com/dfryer1193/thehardway/handlers"
	"github.com/dfryer1193/thehardway/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the Gin router and routes
func SetupRoutes(router *gin.Engine) {
	// Public Post routes
	router.GET("/posts", handlers.GetPublishedPosts)
	router.GET("/posts/:id", handlers.GetPost)

	// Comment routes
	router.POST("/posts/:id/comments", handlers.AddComment)

	// Public routes for login and 2FA
	router.GET("/challenge", handlers.LoginChallenge)
	router.POST("/login", handlers.Login)
	router.POST("/2fa", handlers.YubiKeyVerification)

	// Protected routes for the site owner (requires authentication)
	ownerRoutes := router.Group("/")
	ownerRoutes.Use(middleware.AuthMiddleware())
	{
		ownerRoutes.POST("/posts", handlers.CreatePost)
		ownerRoutes.PATCH("/posts/:id", handlers.UpdatePost)
		ownerRoutes.POST("/change-password", handlers.ChangePassword)
		ownerRoutes.DELETE("/comments/:id", handlers.DeleteComment)
		ownerRoutes.POST("/ban", handlers.BanUser)
		// Other protected routes (create post, delete comment, etc.)
	}
}
