package driver

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func FinishHandler(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	var order types.Order

	global.DB.First(&order, id)

	if order.DriverID != uint(ctx.GetInt("uid")) {
		handlers.HandleError(ctx, "不是你的订单", http.StatusForbidden)
		return
	}

	order.EndTime = time.Now()

	global.DB.Save(&order)

	for _, uid := range order.BookingPassengerIDs {
		global.DB.Create(&types.Transaction{
			UserID:   uid,
			OrderID:  order.ID,
			IsPaid:   false,
			IsRated:  false,
			Feedback: "",
			Reply:    "",
		})
	}

	handlers.HandleSuccess(ctx, nil, "订单确认完成", http.StatusOK)
}
