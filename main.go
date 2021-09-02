package main

import (
	"backend/global"
	"backend/inits"
	"backend/middleware/cors"
	"backend/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// Load config

	log.Println("Loading config...")

	if err := inits.LoadConfig("config.yml"); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Config loaded successfully.")

	// Init database connect

	log.Println("Start connecting to database...")

	if err := inits.ConnectDB(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Database connected successfully")

	// Auto migrate database

	log.Println("Start auto migrating database...")

	if err := inits.MigrateDB(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Database migrated successfully")

	// Set gin work mode

	if global.Config.Debug {
		log.Println("Working on debug mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init routers

	log.Println("Initializing routers...")

	routers.Include(routers.AuthRouter, routers.UserRouters, routers.DriverRouter, routers.AdminRouter)

	r := routers.Init(cors.AllowAll())

	log.Println("Routers initialized successfully.")

	// Start gin server

	log.Println("Starting gin server...")

	if err := r.Run(); err != nil {
		panic(err)
	}

}
