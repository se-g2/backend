package types

import "gorm.io/gorm"

type UserBase struct {
	Avatar       string
	Bio          string
	SchoolID     string
	Name         string
	Tele         string
	Score        float32
	FinishOrders int
	Like         int
	Money        float32
	IsDriver     bool

	IdentityCode  string
	DriverLicense string
	VehicleIDs    []uint `gorm:"type:bigint[]"`
	AvailableTime string

	IsAdmin bool
}

type UserBaseResponse struct {
	ID            uint                  `json:"id"`
	Profile       UserProfile           `json:"profile"`
	SchoolID      string                `json:"schoolId"`
	Name          string                `json:"name"`
	Tele          string                `json:"tele"`
	Score         float32               `json:"score"`
	FinishOrders  int                   `json:"finishOrders"`
	Like          int                   `json:"like"`
	Money         float32               `json:"money"`
	IsDriver      bool                  `json:"isDriver"`
	DriverProfile DriverProfileResponse `json:"driverProfile"`
	IsAdmin       bool                  `json:"isAdmin"`
}

type UserProfile struct {
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

type User struct {
	gorm.Model

	UserBase

	PassSalt string
	PassKey  string
}
