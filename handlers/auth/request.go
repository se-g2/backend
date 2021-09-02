package auth

import (
	"backend/global"
	"backend/handlers"
	"backend/types"
	"backend/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type authReq struct {
	Number string `json:"number"`
}

type authRes struct {
	Session string    `json:"session"`
	Timeout time.Time `json:"timeout"`
}

func RequestHandler(ctx *gin.Context) {

	var req authReq

	_ = ctx.BindJSON(&req)

	var reqInDB types.AuthRequest
	if tx := global.DB.First(&reqInDB, "number = ? AND timeout > ?", req.Number, time.Now());
		errors.Is(tx.Error, gorm.ErrRecordNotFound) {

			session, _ := uuid.NewV4()
			code := fmt.Sprintf("%06d", rand.Intn(999999))

			res := authRes{
				Session: session.String(),
				Timeout: time.Now().Add(time.Minute * 15),
			}
			reqInDB.Number = req.Number
			reqInDB.Session = session.String()
			reqInDB.Timeout = res.Timeout
			reqInDB.Code = code

			global.DB.Save(&reqInDB)

			utils.SendMessage(req.Number, fmt.Sprintf("您的验证码是 %s ，15 分钟内有效", code))

			handlers.HandleSuccess(ctx, res, "消息发送成功", http.StatusOK)
	} else {
		handlers.HandleError(ctx, "操作过快，请 15 分钟后再试", http.StatusTooManyRequests)
	}
}
