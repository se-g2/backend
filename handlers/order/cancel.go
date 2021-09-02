package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CancelHandler(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	var order types.Order

	global.DB.First(&order, id)

	for index, uid := range order.BookingPassengerIDs {
		if uint(ctx.GetInt("uid")) == uid {
			order.BookingPassengerIDs = append(
				order.BookingPassengerIDs[:index],
				order.BookingPassengerIDs[index+1:]...,
			)
			break
		}
	}

	global.DB.Save(&order)

	handlers.HandleSuccess(ctx, nil, "取消成功", http.StatusOK)
}
