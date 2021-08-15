package manger

import (
	"fmt"
	"github.com/nothollyhigh/kiss/log"
	"go_gate/config"
	. "go_gate/manger/proxy"
	"strings"
	"time"
)

var (
	ProxyMgr = &ProxyManger{Proxys: make(map[string]IProxy)}
)

//
type ProxyManger struct {
	Proxys map[string]IProxy
}

func (mgr *ProxyManger) addProxy(name string, proxy IProxy) {
	if _, ok := mgr.Proxys[name]; ok {
		log.Fatal("Duplicate Proxy Name: %v", name)
	}
	mgr.Proxys[name] = proxy
}

func (mgr *ProxyManger) AddProxy(busline *config.XMLBusLine) {
	//校验有效地址
	var addrs []string
	isLineValid := func(port string, lineaddr string) bool {
		for _, localaddr := range addrs {
			if localaddr == lineaddr {
				return false
			}
		}
		addrs = append(addrs, lineaddr)
		return true
	}

	/* 创建并启动一个TcpLine */
	newOneTcpLine := func(proxy *ProxyTcp, serverID, addr string, redictip bool, nodes []config.XMLNode) {
		port := strings.Split(addr, ":")[1]
		for _, node := range nodes {
			node.Addr = fmt.Sprintf("%s:%s", node.Ip, node.Port)
			//地址是否有效
			if !isLineValid(port, node.Addr) {
				log.Fatal("Proxy(%s, %s) AddLine Error: Recursive, Shouldn't Use The Proxy Self's Addr(%s) As Target Addr", serverID, node.Addr, node.Addr)
			}
			log.Info("WebSocket:current ->addr:%v serverID:%v ", node.Addr, serverID) //node.Addr真实IP
			proxy.AddLine(serverID, node.Addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, node.Maxload, redictip)
		}
	}

	/* 创建一个WebsocketLine */
	newOneWSLine := func(proxy *ProxyWebsocket, serverID, addr string, redictip bool, nodes []config.XMLNode) {
		port := strings.Split(addr, ":")[1]
		for _, node := range nodes {
			node.Addr = fmt.Sprintf("%s:%s", node.Ip, node.Port)
			if !isLineValid(port, node.Addr) {
				log.Fatal("Proxy(%s, %s) AddLine Error: Recursive, Shouldn't Use The Proxy Self's Addr(%s) As Target Addr", serverID, node.Addr, node.Addr)
			}
			log.Info("WebSocket:current ->addr:%v serverID:%v  ", node.Addr, serverID)
			proxy.AddLine(serverID, node.Addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, node.Maxload, redictip)
		}
	}

	//代理
	var proxy IProxy
	name := busline.Name
	isRedirect := DEFAULT_TCP_REDIRECT
	switch busline.Type {
	case PT_TCP:
		proxy = NewTcpProxy(name, busline.Addr)
	case PT_WEBSOCKET:
		paths := make([]string, 0)
		for _, route := range busline.Routes {
			paths = append(paths, route.Path)
		}
		proxy = NewWebsocketProxy(name, busline.Addr, busline.RealIpMode, paths, busline.TLS, busline.Certs)
	}

	//集成线路
	for _, line := range busline.Lines {
		if "" == name {
			name = line.ServerID
		}
		if "" == line.Addr {
			line.Addr = busline.Addr
		}
		if "" == line.RealIpMode {
			line.RealIpMode = busline.RealIpMode
		}
		if "" == line.Type {
			line.Type = busline.Type
		}
		if "" != line.Redirect {
			isRedirect = line.Redirect == "true"
		}

		switch line.Type {
		case PT_TCP:
			newOneTcpLine(proxy.(*ProxyTcp), line.ServerID, line.Addr, isRedirect, line.Nodes)
		case PT_WEBSOCKET:
			newOneWSLine(proxy.(*ProxyWebsocket), line.ServerID, line.Addr, isRedirect, line.Nodes)
		}
	}

	//启动代理
	switch busline.Type {
	case PT_TCP:
		proxy.(*ProxyTcp).Start()
	case PT_WEBSOCKET:
		proxy.(*ProxyWebsocket).Start()
	}
	mgr.addProxy(name, proxy)
}

func (mgr *ProxyManger) GetProxy(name string) (IProxy, bool) {
	if _, ok := mgr.Proxys[name]; ok {
		return mgr.Proxys[name], true
	}
	return nil, false
}

func (mgr *ProxyManger) InitProxy() {
	options := config.GlobalXmlConfig.Options
	proxy := config.GlobalXmlConfig.Proxy
	DEFAULT_TCP_REDIRECT = options.Redirect
	DEFAULT_TCP_CHECKLINE_INTERVAL = time.Second * time.Duration(options.Heartbeat.Interval)
	DEFAULT_TCP_CHECKLINE_TIMEOUT = time.Second * time.Duration(options.Heartbeat.Timeout)

	for _, busLine := range proxy.BusLines {
		mgr.AddProxy(busLine)
	}
}
