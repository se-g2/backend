package user

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteHandler(ctx *gin.Context) {

	global.DB.Delete(&types.User{}, "id = ?", ctx.GetInt("uid"))

	handlers.HandleSuccess(ctx, nil, "帐号删除成功", http.StatusOK)

}
