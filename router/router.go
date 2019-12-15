package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/pkg/jwt"
	"snail/pkg/upload"
)

type HttpType int32

type SRouter struct {
	Mode     string
	GPath    string
	Path     string
	Listener ISRListener
}

type ISRListener interface {
	With(sr *SRouter)
}

var sr *gin.Engine
var sapi *gin.RouterGroup

func (this *SRouter) SnailRouters() *gin.Engine {
	sr = gin.New()
	gin.SetMode(this.Mode)
	sr.Use(gin.Logger())
	sr.Use(gin.Recovery())
	sapi = sr.Group(this.Path + this.GPath)
	sapi.Use(jwt.JWT())
	sr.StaticFS("upload/images", http.Dir(upload.GetImageFullPath()))
	sapi.POST("/file", UploadImage)
	sapi.DELETE("/file", DelImage)
	this.Listener.With(this)
	return sr
}

func (this *SRouter) GetEx(k string, v gin.HandlerFunc) {
	sr.GET(this.Path+k, v)
}

func (this *SRouter) PostEx(k string, v gin.HandlerFunc) {
	sr.POST(this.Path+k, v)
}

func (this *SRouter) PutEx(k string, v gin.HandlerFunc) {
	sr.PUT(this.Path+k, v)
}

func (this *SRouter) DELEx(k string, v gin.HandlerFunc) {
	sr.DELETE(this.Path+k, v)
}

func (this *SRouter) GetExt(k string, v gin.HandlerFunc) {
	sapi.GET(k, v)

}

func (this *SRouter) PostExt(k string, v gin.HandlerFunc) {
	sapi.POST(k, v)
}

func (this *SRouter) PutExt(k string, v gin.HandlerFunc) {
	sapi.PUT(k, v)
}

func (this *SRouter) DELExt(k string, v gin.HandlerFunc) {
	sapi.DELETE(k, v)
}
