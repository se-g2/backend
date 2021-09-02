package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newRes struct {
	Id uint `json:"id"`
}

func NewHandler(ctx *gin.Context) {
	var order types.Order

	_ = ctx.BindJSON(&order)

	global.DB.Create(&order)

	handlers.HandleSuccess(ctx, order.ID, "订单创建成功", http.StatusOK)

}
