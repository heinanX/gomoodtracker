package main

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/moods/", handlers.GetAllMoods)
	r.GET("/api/moods/:id", handlers.GetMoodById)
	r.POST("/api/moods", handlers.RecordMood)
	r.PUT("/api/moods", handlers.ChangeMood)

	r.Run(":8080")
}
