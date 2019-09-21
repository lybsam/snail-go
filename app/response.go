package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/pkg/e"
)

type Gin struct {
	C *gin.Context
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (it *Gin) Resp(ok bool, data interface{}) {
	if !ok {
		it.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	it.Response(http.StatusOK, e.SUCCESS, data)
}

func (g *Gin) resp(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (it *Gin) Respext(errCode int, msg string, data interface{}) {
	it.resp(http.StatusOK, errCode, msg, data)
}
