package user

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type modifyReq struct {
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}

func ModifyHandler(ctx *gin.Context) {

	var req modifyReq

	_ = ctx.BindJSON(&req)

	var u types.User

	global.DB.First(&u, "id = ?", ctx.GetInt("uid"))

	u.Avatar = req.Avatar
	u.Bio = req.Bio
	global.DB.Save(&u)
	handlers.HandleSuccess(ctx, nil, "用户信息更新成功", http.StatusOK)

}
