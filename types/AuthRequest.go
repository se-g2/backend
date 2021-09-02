package types

import (
	"gorm.io/gorm"
	"time"
)

type AuthRequest struct {
	gorm.Model

	Number  string
	Session string
	Code    string
	Timeout time.Time
}
