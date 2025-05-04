package handlers

import (
	"gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordMood(c *gin.Context) {
	var mood models.Mood

	if err := c.BindJSON(&mood); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	mood.TimeStamp = time.Now()

	c.IndentedJSON(http.StatusCreated, mood)
}
