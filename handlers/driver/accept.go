package driver

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type acceptReq struct {
	Vehicle uint `json:"vehicle"`
}

func AcceptHandler(ctx *gin.Context) {

	var req acceptReq

	_ = ctx.BindJSON(&req)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	var order types.Order

	global.DB.First(&order, id)

	if order.DriverID != 0 {
		handlers.HandleError(ctx, "订单无法接受", http.StatusForbidden)
		return
	}

	order.DriverID = uint(ctx.GetInt("uid"))
	order.VehicleInfoID = req.Vehicle

	global.DB.Save(&order)

}
