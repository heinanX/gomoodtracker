package handlers

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChangeMood(c *gin.Context) {
	var mood models.Mood

	if err := c.BindJSON(&mood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	c.IndentedJSON(http.StatusCreated, mood)
}
