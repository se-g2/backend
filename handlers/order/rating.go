package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ratingReq struct {
	Rating int `json:"rating"`
}

func RatingHandler(ctx *gin.Context) {

	var req ratingReq

	_ = ctx.BindJSON(&req)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	var tx types.Transaction

	global.DB.First(&tx, "user_id = ? AND order_id = ?", ctx.GetInt("uid"), id)

	if tx.ID == 0 || !tx.IsPaid || tx.IsRated {
		handlers.HandleError(ctx, "订单无效或已经评价", http.StatusForbidden)
		return
	}

	if req.Rating > 0 {
		var order types.Order

		global.DB.First(&order, tx.OrderID)

		var driver types.User

		global.DB.First(&driver, order.DriverID)

		driver.Like++

		global.DB.Save(&driver)
	}

	tx.IsRated = true

	global.DB.Save(&tx)

	handlers.HandleSuccess(ctx, nil, "订单评价成功", http.StatusOK)

}
