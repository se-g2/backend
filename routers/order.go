package routers

import (
	"backend/handlers/order"
	"backend/middleware/auth"
	"github.com/gin-gonic/gin"
)

func OrderRouters(e *gin.Engine) {
	r := e.Group("/order")
	r.Use(auth.UserMiddleware())
	r.GET("/list/:identity", order.ListHandler)
	//r.GET("/search")
	r.GET("/:id/info", order.InfoHandler)
	r.POST("/new", order.NewHandler)
	r.POST("/:id/join", order.JoinHandler)
	r.DELETE("/:id/cancel", order.CancelHandler)
	r.POST("/:id/rating", order.RatingHandler)
	r.POST("/:id/feedback", order.FeedbackHandler)
}
