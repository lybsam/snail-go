package socket

import (
	"net"
	"sync"
	"time"
)

//---------------------------------------------一个session代表一个连接--------------------------------------------------
type Session struct {
	Id    uint32
	Con   net.Conn
	times int64
	lock  sync.Mutex
}

func NewSession(id uint32, con net.Conn) *Session {
	return &Session{
		Id :    id,
		Con:   con,
		times: time.Now().Unix(),
	}
}

func (this *Session) write(msg string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_ ,errs := this.Con.Write([]byte(msg))
	return errs
}

func (this *Session)close(){
	this.Con.Close()
}

func (this *Session)updateTime(){
	this.times = time.Now().Unix()
}
//---------------------------------------------------SESSION管理类------------------------------------------------------

type SessionM struct {
	isWebSocket bool
	ser     *Msf
	sessions sync.Map
}

func NewSessonM(msf *Msf) *SessionM {
	if msf == nil {
		return nil
	}

	return &SessionM{
		ser : msf,
	}
}

func (this *SessionM) GetSessionById(id uint32) *Session {
	tem ,exit := this.sessions.Load(id)
	if exit {
		if sess, ok := tem.(*Session) ; ok {
			return sess
		}
	}
	return nil
}

func (this *SessionM) SetSession(fd uint32, conn net.Conn) {
	sess := NewSession(fd, conn)
	this.sessions.Store(fd,sess)
}

//关闭连接并删除
func (this *SessionM) DelSessionById(id uint32) {
	tem ,exit := this.sessions.Load(id)
	if exit {
		if sess, ok := tem.(*Session) ; ok {
			sess.close()
		}
	}
	this.sessions.Delete(id)
}
