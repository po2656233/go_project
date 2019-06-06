package network

import (
	"net"
)

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	NotifyMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
}
