package order

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InfoHandler(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res := getOrder(uint(id))

	handlers.HandleSuccess(ctx, res, "订单详细信息获取成功", http.StatusOK)

}
