package routers

import (
	"backend/middleware/auth"
	"github.com/gin-gonic/gin"
)

func AdminRouter(e *gin.Engine) {
	r := e.Group("/admin")
	r.Use(auth.UserMiddleware(), auth.AdminMiddleware())

}
