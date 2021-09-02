package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleError(ctx *gin.Context, msg string, status int) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    status,
		"ok":      false,
		"message": msg,
	})
	log.Println(msg)

}

func HandleSuccess(ctx *gin.Context, data interface{}, msg string, status int) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    status,
		"ok":      true,
		"message": msg,
		"data":    data,
	})

}
