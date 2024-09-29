package handlers

import (
	"github.com/dfryer1193/thehardway/models"
	"github.com/dfryer1193/thehardway/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.AddComment(db, &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add comment"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	// Logic to delete comment
}

func BanUser(c *gin.Context) {
	// Logic to ban user by email
}
