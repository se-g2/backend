package routers

import (
	"backend/handlers/driver"
	"backend/middleware/auth"
	"github.com/gin-gonic/gin"
)

func DriverRouter(e *gin.Engine) {
	r := e.Group("/order")
	r.Use(auth.UserMiddleware(), auth.DriverMiddleware())
	r.POST("/:id/accept", driver.AcceptHandler)
	r.POST("/:id/finish", driver.FinishHandler)
}
