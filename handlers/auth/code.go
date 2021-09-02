package auth

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"time"
)

type codeReq struct {
	Number  string `json:"number"`
	Session string `json:"session"`
	Code    string `json:"code"`
}

type codeRes struct {
	Code string `json:"code"`
}

func CodeHandler(ctx *gin.Context) {

	var req codeReq

	_ = ctx.BindJSON(&req)

	var reqInDB types.AuthRequest
	global.DB.First(&reqInDB, "number = ?", req.Number)

	if req.Session == reqInDB.Session &&
		req.Code == reqInDB.Code {
		if reqInDB.Timeout.Before(time.Now()) {
			handlers.HandleError(ctx, "验证码已过期", http.StatusForbidden)
		} else {
			var user types.User
			// 不存在则开户
			global.DB.First(&user, "tele = ?", req.Number)
			if user.ID == 0 {
				user.IsAdmin = false
				user.IsDriver = false
				global.DB.Save(&user)
			}

			// 记录 ID
			var res codeRes
			code, _ := uuid.NewV4()
			sessionInDB := types.ActiveSession{
				UserID: user.ID,
				Code: code.String(),
			}
			global.DB.Save(&sessionInDB)
			res.Code = code.String()
			handlers.HandleSuccess(ctx, res, "登录成功", http.StatusOK)
		}
	} else {
		handlers.HandleError(ctx, "短信验证码错误", http.StatusUnauthorized)
	}

}
