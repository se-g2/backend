package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func JoinHandler(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	var order types.Order

	global.DB.First(&order, id)

	order.BookingPassengerIDs = append(order.BookingPassengerIDs, uint(ctx.GetInt("uid")))

	global.DB.Save(&order)

	handlers.HandleSuccess(ctx, nil, "加入成功", http.StatusOK)
}
