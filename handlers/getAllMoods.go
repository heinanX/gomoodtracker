package handlers

import (
	"gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllMoods(c *gin.Context) {

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
	c.IndentedJSON(http.StatusOK, moods)
}
