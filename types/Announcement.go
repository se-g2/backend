package types

import (
	"gorm.io/gorm"
	"time"
)

type Announcement struct {
	gorm.Model

	Title    string
	Category string
	Content  string
}

type AnnouncementResponse struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Publish  time.Time `json:"publish"`
	Update   time.Time `json:"update"`
	Category string    `json:"category"`
	Content  string    `json:"content"`
}
