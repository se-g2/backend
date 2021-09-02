package inits

import (
	"backend/global"
	"backend/types"
)

func MigrateDB() error {
	return global.DB.AutoMigrate(
		&types.ActiveSession{},
		&types.Announcement{},
		&types.AuthRequest{},
		&types.Order{},
		&types.Transaction{},
		&types.User{},
		&types.VehicleInfo{},
	)
}
