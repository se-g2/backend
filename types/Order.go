package types

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model

	Creator             uint
	Departure           Position
	Arrival             Position
	StartTime           time.Time
	EndTime             time.Time
	Distance            float32
	Price               float32
	BookingPassengerIDs []uint `gorm:"type:bigint[]"`
	TargetPassenger     uint
	VehicleInfoID       uint
	DriverID            uint
	Status              string
	IsComplete          bool
}

type OrderResponse struct {
	ID               int                   `json:"id"`
	CreateTime       time.Time             `json:"createTime"`
	Creator          int                   `json:"creator"`
	Departure        Position              `json:"departure"`
	StartTime        time.Time             `json:"startTime"`
	EndTime          time.Time             `json:"endTime"`
	Distance         float32               `json:"distance"`
	Price            float32               `json:"price"`
	BookingPassenger []UserBaseResponse    `json:"bookingPassenger"`
	TargetPassenger  uint                  `json:"targetPassenger"`
	Vehicle          VehicleInfo           `json:"vehicle"`
	Driver           DriverProfileResponse `json:"driver"`
	Status           string                `json:"status"`
	Complete         bool                  `json:"complete"`
}
