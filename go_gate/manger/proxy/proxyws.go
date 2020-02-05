package proxy

import (
	"crypto/tls"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"github.com/nothollyhigh/kiss/log"
	"github.com/nothollyhigh/kiss/util"
	"github.com/tomasen/realip"
	knet "github.com/nothollyhigh/kiss/net"
	"go_gate/config"

	"io"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)
var (
	DefaultSocketOpt = &knet.SocketOpt{
		NoDelay:           true,
		Keepalive:         false,
		ReadBufLen:        1024 * 8,
		WriteBufLen:       1024 * 8,
		ReadTimeout:       time.Second * 35,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 5,
		MaxHeaderBytes:    4096,
	}

	DefaultUpgrader = &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

/* websocket 代理 */
type ProxyWebsocket struct {
	*ProxyBase
	Running       bool
	EnableTls     bool
	Listener      net.Listener
	Heartbeat     time.Duration
	AliveTime     time.Duration
	RecvBlockTime time.Duration
	RecvBufLen    int
	SendBlockTime time.Duration
	SendBufLen    int
	linelay       bool
	ConnCount     uint64
	RealIpMode	  string
	Certs         []config.XMLCert
	Routes        map[string]func(w http.ResponseWriter, r *http.Request)
}

func (pws *ProxyWebsocket) InitConn(conn *net.TCPConn) bool {
	if err := conn.SetKeepAlivePeriod(pws.AliveTime); err != nil {
		log.Info("ProxyWebsocket(TLS: %v) InitConn SetKeepAlivePeriod Err: %v", pws.EnableTls, err)
		return false
	}

	if err := conn.SetReadBuffer(pws.RecvBufLen); err != nil {
		log.Info("ProxyWebsocket(TLS: %v) InitConn SetReadBuffer Err: %v", pws.EnableTls, err)
		return false
	}
	if err := conn.SetWriteBuffer(pws.SendBufLen); err != nil {
		log.Info("ProxyWebsocket(TLS: %v) InitConn SetWriteBuffer Err: %v", pws.EnableTls, err)
		return false
	}
	if err := conn.SetNoDelay(pws.linelay); err != nil {
		log.Info("ProxyWebsocket(TLS: %v) InitConn Setlinelay Err: %v", pws.EnableTls, err)
		return false
	}
	return true
}

func (pws *ProxyWebsocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := pws.Routes[r.URL.Path]; ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}

func (pws *ProxyWebsocket) OnNew(w http.ResponseWriter, r *http.Request) {
	defer util.HandlePanic()

	atomic.AddUint64(&(pws.ConnCount), 1)

	var (
		serverConn *net.TCPConn
		tcpAddr    *net.TCPAddr
		wsaddr     = r.RemoteAddr
	)

	//获取客户端的serverID  从http.head的Sec-Websocket-Protocol字段中获取,是与客户端商定的
	whead:=w.Header()
	serverID := r.Header.Get("Sec-Websocket-Protocol")
	if "" == serverID{//默认走大厅
		serverID = "HALL"
	}else {
		whead.Add("Sec-Websocket-Protocol",serverID)
	}


	// 向后端透传真实IP的方式
	if "http" == pws.RealIpMode{
		wsaddr = realip.FromRequest(r)
	}

	//根据serverID获取有效线路
	wsline := pws.AssignLine(serverID)
	if wsline == nil{
		log.Info("Session(%s -> null, TLS: %v serverID:%v)  Failed", wsaddr, pws.EnableTls, serverID)
		http.NotFound(w, r)
		return
	}

	// http升级至websocket
	wsConn, err := DefaultUpgrader.Upgrade(w, r, whead)
	line := wsline
	ConnMgr.UpdateInNum(1)
	defer ConnMgr.UpdateInNum(-1)


	//服务端根据域名获取IP
	addrs,_ := net.LookupHost(line.Remote)
	if 0 < len(addrs){
		line.Remote = addrs[0]
	}

	//检测IP是否可用
	if tcpAddr, err = net.ResolveTCPAddr("tcp", line.Remote); err != nil {
		log.Info("Session(%s -> %s, TLS: %v) ResolveTCPAddr Err: %s", wsaddr, line.Remote, pws.EnableTls, err.Error())
		wsConn.Close()
		line.UpdateDelay(UnreachableTime)
		line.UpdateFailedNum(1)
		ConnMgr.UpdateFailedNum(1)
		return
	}

	log.Info("ServerID: %v  name: %v Remote: %v", line.LineID, pws.name, line.Remote)
	var (
		clientRecv int64 = 0
		clientSend int64 = 0
		serverRecv int64 = 0
		serverSend int64 = 0
	)

	// 服务端 --> 客户端
	s2c := func() {
		defer util.HandlePanic()
		defer func() {
			wsConn.Close()
			if serverConn != nil {
				serverConn.Close()
			}
		}()

		var headlen = HEAD_LEN
		var nread int
		var bodylen int
		var err error
		var buf = make([]byte, pws.RecvBufLen)
		for {
			serverConn.SetReadDeadline(time.Now().Add(pws.RecvBlockTime))
			nread, err = io.ReadFull(serverConn, buf[:headlen])
			if err != nil {
				wsConn.Close()
				log.Info("Session(%s -> %s, TLS: %v) Closed, Server Read Err: %s",
					wsaddr, line.Remote, pws.EnableTls, err.Error())
				break
			}

			serverRecv += int64(nread)
			ConnMgr.UpdateServerInSize(int64(nread))

			//pbLen = len - msgID - errCode
			bodylen = int(binary.BigEndian.Uint32(buf[:4])) - MSGID_LEN - ERRCODE_LEN
			if bodylen > 0 {
				if cap(buf) < headlen+bodylen {
					if (headlen+bodylen)%1024 != 0 {
						newBuf := make([]byte, 1024*((headlen+bodylen)/1024+1))
						copy(newBuf, buf)
						buf = newBuf
					} else {
						newBuf := make([]byte, headlen+bodylen)
						copy(newBuf, buf)
						buf = newBuf
					}

				}
				serverConn.SetReadDeadline(time.Now().Add(pws.RecvBlockTime))
				nread, err = io.ReadFull(serverConn, buf[headlen:headlen+bodylen])
				if err != nil {//没法响应客户端的线路进行连接失败统计,并在客户端重连时进行切换
					wsConn.Close()
					log.Info("Session(%s -> %s, TLS: %v) Closed, Server Read Err: %s",
						wsaddr, line.Remote, pws.EnableTls, err.Error())

					//线路延迟
					line.UpdateDelay(UnreachableTime)
					//统计连接失败数
					line.UpdateFailedNum(1)
					ConnMgr.UpdateFailedNum(1)
					return
				}

				serverRecv += int64(nread)
				ConnMgr.UpdateServerInSize(int64(nread))

				nread += headlen
			}
			wsConn.SetWriteDeadline(time.Now().Add(pws.SendBlockTime))
			err = wsConn.WriteMessage(websocket.BinaryMessage, buf[:nread])
			if err != nil {
				log.Info("Session(%s -> %s, TLS: %v) Closed, Server WriteMessage Err: %s",
					wsaddr, line.Remote, pws.EnableTls, err.Error())
				break
			}

			serverSend += int64(nread)
			ConnMgr.UpdateServerOutSize(int64(nread))
			log.Info("server:[%v] send-->>> <%v> MsgLen::%v", line.Remote, wsaddr, nread)
		}
	}

	// 客户端 --> 服务端
	c2s := func() {
		defer func() {
			wsConn.Close()
			if serverConn != nil {
				serverConn.Close()
			}
		}()

		var nwrite int
		var err error
		var message []byte
		for {
			err = wsConn.SetReadDeadline(time.Now().Add(pws.RecvBlockTime))
			if err != nil {
				log.Info("Session(%s -> %s, TLS: %v) Closed, Client ReadMessage Err: %s",
					wsaddr, line.Remote, pws.EnableTls, err.Error())
				break
			}
			_, message, err = wsConn.ReadMessage()
			if err != nil {
				log.Info("Session(%s -> %s, TLS: %v) Closed, Client ReadMessage Err: %s",
					wsaddr, line.Remote, pws.EnableTls, err.Error())
				break
			}


			// 建立连接
			if serverConn == nil {
				// 校验第一个数据包是否有效
				nReadLen := len(message)
				if nReadLen < HEAD_LEN || nReadLen - HEAD_LEN != int(binary.BigEndian.Uint32(message[:4])) - MSGID_LEN - ERRCODE_LEN{
					log.Info("Session(%s -> %s ReadLen:%v length:%v) protocol Failed",wsaddr,  line.Remote, nReadLen,int(binary.BigEndian.Uint32(message[:4])))
					wsConn.Close()

					//线路延迟
					line.UpdateDelay(UnreachableTime)
					//统计连接失败数
					line.UpdateFailedNum(1)
					ConnMgr.UpdateFailedNum(1)
					return
				}

				t1 := time.Now()
				serverConn, err = net.DialTCP("tcp", nil, tcpAddr)
				if err != nil {
					log.Info("Session(%s -> %s, TLS: %v) DialTCP Err: %s",
						wsaddr, line.Remote, pws.EnableTls, err.Error())
					wsConn.Close()

					//线路延迟
					line.UpdateDelay(UnreachableTime)
					//统计连接失败数
					line.UpdateFailedNum(1)
					ConnMgr.UpdateFailedNum(1)
					return
				}
				line.UpdateDelay(time.Since(t1))
				pws.InitConn(serverConn)

				//统计负载量
				line.UpdateLoad(1)
				defer line.UpdateLoad(-1)

				//统计当前连接数
				ConnMgr.UpdateOutNum(1)
				defer ConnMgr.UpdateOutNum(-1)

				//统计连接成功数
				ConnMgr.UpdateSuccessNum(1)

				log.Info("Session(%s -> %s, TLS: %v) Established", wsaddr, line.Remote, pws.EnableTls)

				//传真实IP
				if err = line.HandleRedirect(serverConn, wsaddr); err != nil {
					log.Info("Session(%s -> %s) HandleRedirect Failed: %s", wsaddr, line.Remote, err.Error())
					return
				}

				util.Go(s2c)
			}

			clientRecv += int64(len(message))
			ConnMgr.UpdateClientInSize(int64(len(message)))

			serverConn.SetWriteDeadline(time.Now().Add(pws.SendBlockTime))
			nwrite, err = serverConn.Write(message)
			if nwrite != len(message) || err != nil {
				log.Info("Session(%s -> %s, TLS: %v) Closed, Client Write len:%v Err: %s ",
					wsaddr, line.Remote, pws.EnableTls, nwrite, err.Error())
				break
			}

			clientSend += int64(nwrite)
			ConnMgr.UpdateClientOutSize(int64(nwrite))

			//仅打印长度
			log.Info("client:[%v] send-->> <%v> MsgLen:%v",wsaddr, line.Remote, nwrite)
		}
	}

	// 客户端-->服务端
	c2s()

	log.Info("Session(%s -> %s, TLS: %v SID: %s) Over, DataInfo(CR: %d, CW: %d, SR: %d, SW: %d)",
		wsaddr, line.Remote, pws.EnableTls,line.LineID, clientRecv, clientSend, serverRecv, serverSend)
}


func (pws *ProxyWebsocket) Start() {
	if len(pws.lines) == 0 {
		log.Info("ProxyWebsocket(%v TLS: %v) Start Err: No Line !", pws.name, pws.EnableTls)
		return
	}

	// 监听数据
	if pws.Listener == nil{
		var err error
		pws.Listener, err = knet.NewListener(pws.local, DefaultSocketOpt)
		if err != nil {
			log.Fatal("ProxyWebsocket(%v TLS: %v) NewListener Failed: %v", pws.name, pws.EnableTls, err)
		}
	}


	util.Go(func() {
		pws.Lock()
		defer pws.Unlock()
		if !pws.Running {
			pws.Running = true
			util.Go(func() {

				//由于部分线路共用busline的端口,故牵至goroutine外
				//l, err := knet.NewListener(pws.local, DefaultSocketOpt)
				//if err != nil {
				//	log.Fatal("ProxyWebsocket(%v TLS: %v) NewListener Failed: %v", pws.name, pws.EnableTls, err)
				//}else{
				//	log.Info(" Listen:%v local:%v",pws.name, pws.local)
				//}

				s := &http.Server{
					Addr:              pws.local,
					Handler:           pws,
					ReadTimeout:       DefaultSocketOpt.ReadTimeout,
					ReadHeaderTimeout: DefaultSocketOpt.ReadHeaderTimeout,
					WriteTimeout:      DefaultSocketOpt.WriteTimeout,
					MaxHeaderBytes:    DefaultSocketOpt.MaxHeaderBytes,
				}

				if pws.EnableTls {
					if len(pws.Routes) == 0 {
						pws.Routes["/gate/wss"] = pws.OnNew
					}

					log.Info("ProxyWebsocket(%v TLS: %v) Running On: %s, Routes: %+v, Certs: %+v", pws.name, pws.EnableTls, pws.local, pws.Routes, pws.Certs)

					pws.StartCheckLines()
					defer pws.StopCheckLines()

					if len(pws.Certs) == 0 {
						log.Fatal("ProxyWebsocket(%v TLS: %v) ListenAndServeTLS Error: No Cert And Key Files", pws.name, pws.EnableTls)
					}

					s.TLSConfig = &tls.Config{}
					for _, v := range pws.Certs {
						cert, err := tls.LoadX509KeyPair(v.Certfile, v.Keyfile)
						if err != nil {
							log.Fatal("ProxyWebsocket(%v TLS: %v) tls.LoadX509KeyPair(%v, %v) Failed: %v", pws.name, pws.EnableTls, v.Certfile, v.Keyfile, err)
						}
						s.TLSConfig.Certificates = append(s.TLSConfig.Certificates, cert)
					}

					tlsListener := tls.NewListener(pws.Listener, s.TLSConfig)

					if err := s.Serve(tlsListener); err != nil {
						log.Fatal("ProxyWebsocket(%v TLS: %v) Serve Error: %v", pws.name, pws.EnableTls, err)
					}
				} else {
					if len(pws.Routes) == 0 {
						pws.Routes["/gate/ws"] = pws.OnNew
					}

					log.Info("ProxyWebsocket(%v TLS: %v, Routes: %+v) Running On: %s", pws.name, pws.EnableTls, pws.Routes, pws.local)

					//线路检测
					pws.StartCheckLines()
					defer pws.StopCheckLines()

					if err := s.Serve(pws.Listener); err != nil {
						log.Fatal("ProxyWebsocket(TLS: %v) Serve Error: %v", pws.EnableTls, err)
					}
				}
			})
		}
	})
}


func (pws *ProxyWebsocket) Stop() {
	pws.Lock()
	defer pws.Unlock()
	if pws.Running {
		pws.Running = false
	}
}




func NewWebsocketProxy(name string, local string, realIpModel string, paths []string, tls bool, certs []config.XMLCert) *ProxyWebsocket {
	pws := &ProxyWebsocket{
		Running:       false,
		EnableTls:     tls,
		Listener:      nil,
		Heartbeat:     DEFAULT_TCP_HEARTBEAT,
		AliveTime:     DEFAULT_TCP_KEEPALIVE_INTERVAL,
		RecvBlockTime: DEFAULT_TCP_READ_BLOCK_TIME,
		RecvBufLen:    DEFAULT_TCP_READ_BUF_LEN,
		SendBlockTime: DEFAULT_TCP_WRITE_BLOCK_TIME,
		SendBufLen:    DEFAULT_TCP_WRITE_BUF_LEN,
		linelay:       DEFAULT_TCP_NODELAY,
		Certs:         certs,
		Routes:        map[string]func(w http.ResponseWriter, r *http.Request){},
		RealIpMode:   realIpModel,
		ProxyBase: &ProxyBase{
			name:  name,
			ptype: PT_WEBSOCKET,
			local: local,
			lines: []*Line{},
		},
	}

	for _, path := range paths {
		pws.Routes[path] = pws.OnNew
	}
	return pws
}
