package auth

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uidO, exist := ctx.Get("uid")

		if !exist {
			handlers.HandleError(ctx, "中间件请求失败，请联系开发者获取更多帮助。", http.StatusInternalServerError)
			ctx.Abort()
			return
		}

		UID := uidO.(uint)

		var u types.User

		global.DB.First(&u, UID)

		if !u.IsAdmin {
			handlers.HandleError(ctx, "您不是管理员。", http.StatusForbidden)
			ctx.Abort()
		}

	}
}
