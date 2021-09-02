package types

import (
	"gorm.io/gorm"
)

type ActiveSession struct {
	gorm.Model

	UserID  uint
	Code    string
}
