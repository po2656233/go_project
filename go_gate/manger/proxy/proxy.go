package proxy

import (
	"bytes"
	"encoding/json"
	"github.com/nothollyhigh/kiss/log"
	"go_gate/config"
	"strings"
	"sync"
	"time"
)

var (
	PT_TCP       = "tcp"
	PT_WEBSOCKET = "websocket"
	HEAD_LEN     = 2

	/* 默认配置项 */
	DEFAULT_TCP_NODELAY            = true              /* tcp nodelay */
	DEFAULT_TCP_REDIRECT           = true              /* 是否向服务器发送客户端IP */
	DEFAULT_TCP_HEARTBEAT          = time.Second * 30  /* tcp代理设置的心跳时间 */
	DEFAULT_TCP_KEEPALIVE_INTERVAL = time.Second * 600 /* tcp代理设置的 keepalive 时间 */
	DEFAULT_TCP_READ_BUF_LEN       = 1024 * 8          /* tcp 接收缓冲区 */
	DEFAULT_TCP_WRITE_BUF_LEN      = 1024 * 8          /* tcp 发送缓冲区 */
	DEFAULT_TCP_READ_BLOCK_TIME    = time.Second * 35  /* tcp 读数据超时时间 */
	DEFAULT_TCP_WRITE_BLOCK_TIME   = time.Second * 5   /* tcp 写数据超时时间 */
	DEFAULT_TCP_CHECKLINE_INTERVAL = time.Second * 60  /* 线路检测周期 */
	DEFAULT_TCP_CHECKLINE_TIMEOUT  = time.Second * 10  /* 线路检测超时时间 */
)

type IProxy interface {
	GetLine(serverID, addr string) *Line       // 获取节点线路(确切)
	GetBestLine() *Line                        // 获取最优线路
	AssignLine(serverID string) *Line          // 指定线路
	ReserveLines(lines []*config.XMLLine) bool // 保留有效线路(不在配置内的线路被视为无效)
	LinesForJSON() []byte                      // 线路信息
}

/* 每个 ProxyBase 管理一组 Line ，Proxy is a ProxyBase */
type ProxyBase struct {
	sync.RWMutex

	name  string
	ptype string

	local string
	lines []*Line
}

type JsonLine struct {
	LineID   string        /* LineID==服务ID */
	Running  bool          /* 线路检测是否在进行的标志 */
	Born     int64         /* 线路出生时间 */
	Remote   string        /* 线路指向的服务器地址 */
	Delay    time.Duration /* 线路延迟 */
	Timeout  time.Duration /* 进行线路检测时的超时时间 */
	Interval time.Duration /* 线路检测的时间周期 */
	CurLoad  int64         /* 当前线路负载 */
	MaxLoad  int64         /* 线路最大负载 */
	IsPaused bool          /* 线路暂停使用的标志 */
	Redirect bool          /* 线路是否需要向服务器发送客户端真实IP的标志 */
}

/* 当前最适合的线路 */
func (mgr *ProxyBase) GetBestLine() *Line {
	mgr.RLock()
	defer mgr.RUnlock()

	return mgr.GetBestLineWithoutLock()
}

/* 根据serverID获取指定线路 */
func (mgr *ProxyBase) AssignLine(serverID string) *Line {
	mgr.RLock()
	defer mgr.RUnlock()

	if len(mgr.lines) > 0 {
		line := mgr.lines[0]
		for i := 0; i < len(mgr.lines); i++ {
			if mgr.lines[i].Score() > line.Score() && strings.EqualFold(mgr.lines[i].LineID, serverID) {
				line = mgr.lines[i]
				break
			}
		}
		if strings.EqualFold(line.LineID, serverID) && line.Score() >= 0 {
			return line
		}
	}

	return nil
}

/* 查找指定线路 */
func (mgr *ProxyBase) GetLine(serverID string, addr string) *Line {
	mgr.RLock()
	defer mgr.RUnlock()
	for i := 0; i < len(mgr.lines); i++ {
		if strings.EqualFold(mgr.lines[i].LineID, serverID) && mgr.lines[i].Remote == addr {
			return mgr.lines[i]
		}
	}
	return nil
}

/* 当前最适合的线路 */
func (mgr *ProxyBase) GetBestLineWithoutLock() *Line {
	if len(mgr.lines) > 0 {
		line := mgr.lines[0]
		for i := 1; i < len(mgr.lines); i++ {
			if mgr.lines[i].Score() > line.Score() {
				line = mgr.lines[i]
				break
			}
		}
		if line.Score() >= 0 {
			return line
		}
	}
	return nil
}

func (mgr *ProxyBase) ReserveLines(lines []*config.XMLLine) bool {

	have := false
	for i := len(mgr.lines) - 1; i > 0; i-- {
		for j := 0; j < len(lines); j++ {
			if strings.EqualFold(mgr.lines[i].LineID, lines[j].ServerID) && mgr.lines[i].Remote == lines[j].Addr {
				have = true
				break
			}
		}
		if !have {
			mgr.lines[i].Stop()
			mgr.RLock()
			mgr.lines = append(mgr.lines[:i], mgr.lines[i+1:]...)
			mgr.RUnlock()
		}
	}
	return true
}
func (mgr *ProxyBase) LinesForJSON() []byte {
	mgr.RLock()
	defer mgr.RUnlock()

	//格式
	var buffer bytes.Buffer
	linesLen := len(mgr.lines)
	buffer.Write([]byte("{\"" + mgr.name + "\":["))
	for i := 0; i < linesLen; i++ {
		b, err := json.Marshal(mgr.lines[i])
		if err != nil {
			log.Error("json err:%v data", err.Error())
			continue
		}
		buffer.Write(b)
		if i != (linesLen - 1) {
			buffer.Write([]byte(","))
		}
		log.Info("json info:%v", string(b))
	}
	buffer.Write([]byte("]}"))
	return buffer.Bytes()
}

/* 添加一个 Line */
func (mgr *ProxyBase) AddLine(serverid, addr string, timeout time.Duration, interval time.Duration, maxLoad int64, redirect bool) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.lines = append(mgr.lines, NewLine(serverid, addr, timeout, interval, maxLoad, redirect))
}

/* 开始检查所有 Line 状况 */
func (mgr *ProxyBase) StartCheckLines() {
	for i, line := range mgr.lines {
		line.Start(i)
	}
}

/* 停止检查所有 Line 状况 */
func (mgr *ProxyBase) StopCheckLines() {
	for _, line := range mgr.lines {
		line.Stop()
	}
}

// func (pbase *ProxyBase) GetBestLine() *Line {
// 	return pbase.GetBestLineWithoutLock()
// }
