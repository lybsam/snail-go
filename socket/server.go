package socket

import (
	"encoding/hex"
	"log"
	"net"
)

type SocketTypes interface {
	ConnHandle(data string, conn net.Conn)
}

type Msf struct {
	SessionMaster *SessionM
	SocketType    SocketTypes
}

func NewMsf(socketType SocketTypes) *Msf {
	msf := &Msf{
		SocketType: socketType,
	}
	msf.SessionMaster = NewSessonM(msf)
	return msf
}

func (this *Msf) Listening(address string) {
	tcpListen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fd := uint32(0)
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		this.SessionMaster.SetSession(fd, conn)
		go connHandle(this, this.SessionMaster.GetSessionById(fd))
	}
}

func connHandle(msf *Msf, sess *Session) {
	var errs error
	data := make([]byte, 1024)
	for {
		n, err := sess.Con.Read(data)
		//更新接收时间
		sess.updateTime()
		if err != nil {
			return
		}
		if errs != nil {
			log.Println(errs)
			return
		}
		if len(data) == 0 {
			continue
		}
		str := hex.EncodeToString(data[:n])
		msf.SocketType.ConnHandle(str, sess.Con)
	}
}
