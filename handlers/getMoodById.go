package handlers

import (
	"gin/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMoodById(c *gin.Context) {
	id := c.Param("id")
	numId, _ := strconv.Atoi(id)

	moods := []models.Mood{
		{
			Id:        1,
			MoodScore: 7,
			TimeStamp: time.Date(2025, 5, 4, 10, 0, 0, 0, time.UTC),
		},
		{
			Id:        2,
			MoodScore: 4,
			TimeStamp: time.Date(2025, 5, 3, 14, 30, 0, 0, time.UTC),
		},
		{
			Id:        3,
			MoodScore: 9,
			TimeStamp: time.Date(2025, 5, 2, 18, 15, 0, 0, time.UTC),
		},
	}
	for _, m := range moods {
		if m.Id == numId {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "mood not found"})
}
