package types

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	UserID   uint
	OrderID  uint
	IsPaid   bool
	IsRated  bool
	Feedback string
	Reply    string
}
