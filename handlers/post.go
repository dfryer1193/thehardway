package handlers

import (
	"github.com/dfryer1193/thehardway/models"
	"github.com/dfryer1193/thehardway/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB // Assumed initialized elsewhere

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreatePost(db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func GetPublishedPosts(c *gin.Context) {
	posts, err := services.GetPublishedPosts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func UpdatePost(c *gin.Context) {
	// Logic to update post state (draft/published)
}
