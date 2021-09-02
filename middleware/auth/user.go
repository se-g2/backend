package auth

import (
	"backend/global"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		code := ctx.GetHeader("X-Token")

		var session types.ActiveSession
		global.DB.First(&session, "code = ?", code)

		if session.UserID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			ctx.Set("uid", session.UserID)
		}
	}
}
