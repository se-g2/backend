package order

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type feedbackReq struct {
	Comment string `json:"comment"`
}

func FeedbackHandler(ctx *gin.Context) {
	var req feedbackReq

	_ = ctx.BindJSON(&req)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	var tx types.Transaction

	global.DB.First(&tx, "user_id = ? AND order_id = ?", ctx.GetInt("uid"), id)

	if tx.ID == 0 || !tx.IsPaid {
		handlers.HandleError(ctx, "订单无效", http.StatusForbidden)
		return
	}

	tx.Feedback = req.Comment

	global.DB.Save(&tx)

	handlers.HandleSuccess(ctx, nil, "订单评价成功", http.StatusOK)

}
