package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	NotifyMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
