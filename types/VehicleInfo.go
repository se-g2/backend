package types

import "gorm.io/gorm"

type VehicleInfo struct {
	gorm.Model

	Owner    uint
	Brand    string
	Size     string
	Color    string
	Capacity int
	License  string
}

type VehicleInfoResponse struct {
	ID uint `json:"id"`

	Brand    string `json:"brand"`
	Size     string `json:"size"`
	Color    string `json:"color"`
	Capacity int    `json:"capacity"`
	License  string `json:"license"`
}
