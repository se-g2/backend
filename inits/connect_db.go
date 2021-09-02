package inits

import (
	"backend/global"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() error {
	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s`,
		global.Config.DB.Host,
		global.Config.DB.User,
		global.Config.DB.Password,
		global.Config.DB.DBName,
		global.Config.DB.Port,
		global.Config.DB.SSLMode,
		global.Config.DB.TimeZone,
	)

	var err error

	global.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err

}
