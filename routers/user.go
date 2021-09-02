package routers

import (
	"backend/handlers/user"
	"backend/middleware/auth"
	"github.com/gin-gonic/gin"
)

func UserRouters(e *gin.Engine) {
	base := e.Group("")
	base.Use(auth.UserMiddleware())
	base.POST("/verify", user.VerifyHandler)
	base.GET("/announcements", user.AnnouncementsHandler)
	r := base.Group("/user")
	r.GET("/info", user.InfoHandler)
	r.PUT("/profile/modify", user.ModifyHandler)
	r.DELETE("/delete", user.DeleteHandler)
}
