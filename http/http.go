package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {
	Code int         `json:"Code"`
	Data interface{} `json:"Data"`
	Msg  string      `json:"Msg"`
}

func HtpReq(c *gin.Context, code int, msg string, data interface{}) {
	httpCode := http.StatusOK
	if code > 2000 {
		httpCode = http.StatusBadGateway
	}
	c.JSON(httpCode, Req{
		Code: code,
		Data: data,
		Msg:  msg,
	})

}
