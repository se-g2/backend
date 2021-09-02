package user

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verifyReq struct {
	UserNo string `json:"userno"`
	Name   string `json:"name"`
}

func VerifyHandler(ctx *gin.Context) {
	var req verifyReq
	_ = ctx.BindJSON(&req)

	var u types.User
	global.DB.First(&u, "id = ?", ctx.GetInt("uid"))
	u.Name = req.Name
	u.SchoolID = req.UserNo

	global.DB.Save(&u)

	handlers.HandleSuccess(ctx, nil, "认证成功", http.StatusOK)
}
