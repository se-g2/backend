package routers

import (
	"backend/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(e *gin.Engine) {
	r := e.Group("/auth")
	r.POST("/tele/request", auth.RequestHandler)
	r.POST("/tele/code", auth.CodeHandler)
}
