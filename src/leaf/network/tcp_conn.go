package network

import (
	"github.com/name5566/leaf/log"
	"net"
	"sync"
)

type ConnSet map[net.Conn]struct{}

type TCPConn struct {
	sync.Mutex
	conn      net.Conn
	writeChan chan []byte
	closeFlag bool
	msgParser *MsgParser
	notifyChan chan []byte // 新增异步通知 2019/6/6
	writeFin bool // 写完成
	notifyFin bool //通知完成
}

func newTCPConn(conn net.Conn, pendingWriteNum int, msgParser *MsgParser) *TCPConn {
	tcpConn := new(TCPConn)
	tcpConn.conn = conn
	tcpConn.writeChan = make(chan []byte, pendingWriteNum)
	tcpConn.notifyChan = make(chan []byte, 20) // 绑定数据不需要过大
	tcpConn.msgParser = msgParser
	tcpConn.writeFin = false
	tcpConn.notifyFin = false
	tcpConn.closeFlag = false

	go func() {
		tcpConn.writeFin = false
		for b := range tcpConn.writeChan {
			if b == nil {
				break
			}

			_, err := conn.Write(b)
			if err != nil {
				break
			}
		}
		tcpConn.writeFin = true

		if tcpConn.writeFin && tcpConn.notifyFin && !tcpConn.closeFlag{
			conn.Close()
			tcpConn.Lock()
			tcpConn.closeFlag = true
			tcpConn.Unlock()
		}
	}()

	//异步通知
	go func() {
		tcpConn.notifyFin = false
		for nb := range tcpConn.notifyChan {
			if nb == nil {
				break
			}

			_, err := conn.Write(nb)
			if err != nil {
				break
			}
		}
		tcpConn.notifyFin = true

		if tcpConn.writeFin && tcpConn.notifyFin && !tcpConn.closeFlag{
			conn.Close()
			tcpConn.Lock()
			tcpConn.closeFlag = true
			tcpConn.Unlock()
		}
	}()

	return tcpConn
}

func (tcpConn *TCPConn) doDestroy() {
	tcpConn.conn.(*net.TCPConn).SetLinger(0)
	tcpConn.conn.Close()

	if !tcpConn.closeFlag {
		close(tcpConn.writeChan)
		close(tcpConn.notifyChan) // 新增
		tcpConn.closeFlag = true
	}
}

func (tcpConn *TCPConn) Destroy() {
	tcpConn.Lock()
	defer tcpConn.Unlock()

	tcpConn.doDestroy()
}

func (tcpConn *TCPConn) Close() {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	if tcpConn.closeFlag {
		return
	}

	tcpConn.doWrite(nil)
	tcpConn.doNotify(nil) // 新增
	tcpConn.closeFlag = true
}

func (tcpConn *TCPConn) doWrite(b []byte) {
	if len(tcpConn.writeChan) == cap(tcpConn.writeChan) {
		log.Debug("close conn: channel full")
		tcpConn.doDestroy()
		return
	}

	tcpConn.writeChan <- b
}

func (tcpConn *TCPConn) doNotify(b []byte) {
	if len(tcpConn.notifyChan) == cap(tcpConn.notifyChan) {
		log.Debug("close conn: channel full -")
		tcpConn.doDestroy()
		return
	}

	tcpConn.notifyChan <- b
}

// b must not be modified by the others goroutines
func (tcpConn *TCPConn) Write(b []byte) {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	if tcpConn.closeFlag || b == nil {
		return
	}

	tcpConn.doWrite(b)
}

func (tcpConn *TCPConn) Notify(b []byte) {
	tcpConn.Lock()
	defer tcpConn.Unlock()
	if tcpConn.closeFlag || b == nil {
		return
	}

	tcpConn.doNotify(b)
}



func (tcpConn *TCPConn) Read(b []byte) (int, error) {
	return tcpConn.conn.Read(b)
}

func (tcpConn *TCPConn) LocalAddr() net.Addr {
	return tcpConn.conn.LocalAddr()
}

func (tcpConn *TCPConn) RemoteAddr() net.Addr {
	return tcpConn.conn.RemoteAddr()
}

func (tcpConn *TCPConn) ReadMsg() ([]byte, error) {
	return tcpConn.msgParser.Read(tcpConn)
}

func (tcpConn *TCPConn) WriteMsg(args ...[]byte) error {
	return tcpConn.msgParser.Write(tcpConn, args...)
}
//// 新增一个异步通知
func (tcpConn *TCPConn) NotifyMsg(args ...[]byte) error {
	return tcpConn.msgParser.Notify(tcpConn, args...)
}