package models

import (
	"time"
)

type Mood struct {
	Id        int
	MoodScore int       `json:"MoodScore"`
	TimeStamp time.Time `json:"TimeStamp"`
}
