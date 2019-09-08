package snail

import (
	"fmt"
	"log"
	"net/http"
	"snail/pkg/logging"
	"snail/pkg/model"
	"snail/pkg/setting"
	"snail/pkg/util"
	"snail/router"
	"snail/socket"
)

type Snail struct {
	setting.Config
	router.SRouter
	TcpListener socket.SocketTypes
}

func InitSnail(snail Snail) *Snail {
	setting.Setup(&snail.Config)
	model.Setup()
	util.Setup()
	logging.Setup()
	return &snail
}

func (this *Snail) StartRESTServer() {
	routersInit := this.SRouter.SnailRouters()
	readTimeout := setting.Conf.ReadTimeout
	writeTimeout := setting.Conf.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.Conf.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

func (this *Snail) StartTCPServer() {
	msf := socket.NewMsf(this.TcpListener)
	msf.Listening(this.Address)
}
