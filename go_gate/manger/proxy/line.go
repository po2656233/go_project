package proxy

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/nothollyhigh/kiss/log"
	tnet "github.com/nothollyhigh/kiss/net"
	"github.com/nothollyhigh/kiss/util"
	 "go_gate/proto"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)
var BornTime   = time.Now() /* 进程启动时间 */
const (
	defaultTimeout  = time.Second * 5
	UnreachableTime = time.Duration(-1)
	Unpausecheck    = time.Second * 0
	COUNT_MINUTES   = 60
	MsgUserIP		= 20010
)

type FailedInMunite struct {
	Time      time.Time
	FailedNum int64
}
type Head struct {
	Length      int32
	MsgId       int32
	ErrorCode   int32
}

var (
	ErrorInvalidAddr = errors.New("Invalid Addr")
)

/* 线路 */
type Line struct {
	sync.RWMutex
	LineID string 	   	  `json:"lineID"` /* LineID==服务ID */
	Running  bool          `json:"running"` /* 线路检测是否在进行的标志 */
	Born     time.Time     `json:"_"`/* 线路出生时间 */
	Remote   string        `json:"remote"`/* 线路指向的服务器地址 */
	Delay    time.Duration `json:"delay"`/* 线路延迟 */
	Timeout  time.Duration `json:"timeout"`/* 进行线路检测时的超时时间 */
	Interval time.Duration `json:"interval"`/* 线路检测的时间周期 */
	Timer    *time.Timer   `json:"_"`/* 用于定时进行线路检测的定时器 */
	CurLoad  int64         `json:"curload"`/* 当前线路负载 */
	MaxLoad  int64         `json:"maxload"`/* 线路最大负载 */
	IsPaused bool 		   `json:"ispaused"`/* 线路暂停使用的标志 */
	Redirect bool 		   `json:"redirect"`/* 线路是否需要向服务器发送客户端真实IP的标志 */

	ChUpdateDelay chan time.Duration 	`json:"_"`/* 用于外部更新线路当前延迟的channel，外部进行更新后本线路的线路检测reset计时周期避免浪费 */

	FailedRecord     [COUNT_MINUTES]FailedInMunite `json:"_"`/* 环形队列，记录过去COUNT_MINUTES分钟内连接失败次数 */
	FailedRecordHead int                           `json:"_"`/* 环形队列头 */
}

/* 线路分数，小于0为线路不可用 */
func (line *Line) Score() int64 {
	if !line.IsPaused && line.Delay != UnreachableTime && line.CurLoad < line.MaxLoad {
		return (line.MaxLoad - line.CurLoad)
	}
	return -1
}

func (line *Line) CheckLine(now time.Time) {
	if line.IsPaused{
		log.Info("CheckLine (Addr: %s, lineID: %v is paused!)", line.Remote, line.LineID)
		return
	}
	if err := tnet.Ping(line.Remote, line.Timeout); err != nil {
		line.Delay = UnreachableTime //line.Timeout
		line.Running = false
		log.Error("CheckLine (Addr: %s) Failed, err: %v", line.Remote, err)
		return
	} else {
		line.Running = true
		line.Delay = time.Since(now)
	}
	log.Info("CheckLine (Addr: %s, Delay: %v lineID: %v)", line.Remote, line.Delay, line.LineID)
}

/* 启动线路检测 */
func (line *Line) Start(idx int) {
	line.Lock()
	defer line.Unlock()
	if line.Running {
		return
	}

	line.Running = true

	util.Go(func() {
		line.Born = time.Now()
		line.Timer = time.NewTimer(line.Interval)

		line.CheckLine(line.Born)
		for {
			select {
			case now, ok := <-line.Timer.C:
				if !ok {
					return
				}
				line.CheckLine(now)
			case delay, ok := <-line.ChUpdateDelay:
				if !ok {
					return
				}
				line.Delay = delay
			}
			line.Timer.Reset(line.Interval)
		}
	})
}

/* 更新线路延迟 */
func (line *Line) UpdateDelay(delay time.Duration) {
	line.RLock()
	defer line.RUnlock()
	if !line.Running {
		return
	}
	select {
	case line.ChUpdateDelay <- delay:
	default:
	}
}

/* 停止线路检测 */
func (line *Line) Stop() {
	line.Lock()
	defer line.Unlock()
	if !line.Running {
		return
	}

	line.Running = false

	line.Timer.Stop()
	close(line.ChUpdateDelay)
}

/* 更新负载 */
func (line *Line) UpdateLoad(delta int64) {
	atomic.AddInt64(&(line.CurLoad), delta)
}

/* 暂停在此线路选路和进行代理连接 */
func (line *Line) Pause() {
	line.Lock()
	defer line.Unlock()
	line.IsPaused = true
}

/* 恢复在此线路选路和进行代理连接 */
func (line *Line) UnPause() {
	line.Lock()
	defer line.Unlock()
	line.IsPaused = false

	line.CheckLine(time.Now())
}

/* 根据线路配置决定是否向服务器发送重定向包，应在刚建立与服务器的连接时首先进行然后再发送其他包 */
func (line *Line) HandleRedirect(conn *net.TCPConn, addr string) error {
	if line.Redirect {
		// 消息头
		head := &Head{MsgId:MsgUserIP,}
		// 消息体
		addrstr := strings.Split(addr, ":")[0]
		msg := &gate.MsgSetUserAddrReq{Addr:addrstr}

		log.Info("RemoteAddr: %v realIP: %v", conn.RemoteAddr(),  addrstr)
		if buffer,err:= makeBuf(head, msg); err == nil{
			if err = conn.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
				return err
			}
			// 向conn传输buf
			_, err = conn.Write(buffer)
			return err
		}else{
			return err
		}
	}

	return nil
}

/* 更新为客户端与服务器建立连接的失败总次数，以及近期每分钟内的失败次数记录 */
func (line *Line) UpdateFailedNum(delta int64) {
	line.Lock()
	defer line.Unlock()

	currHead := int(time.Since(BornTime).Minutes()) % COUNT_MINUTES
	if currHead != line.FailedRecordHead || time.Since(line.FailedRecord[line.FailedRecordHead].Time).Minutes() >= 1.0 {
		line.FailedRecordHead = currHead
		line.FailedRecord[currHead] = FailedInMunite{
			Time:      time.Now(),
			FailedNum: 1,
		}
	} else {
		line.FailedRecord[currHead].FailedNum++
	}
}

/* 获取近期n分钟内为客户端与服务器建立连接的失败次数 */
func (line *Line) GetFailedInLastNMinutes(n int) int64 {
	line.Lock()
	defer line.Unlock()

	if n > 0 && n <= COUNT_MINUTES {
		var total int64 = 0
		for i := 0; i < n; i++ {
			if time.Since(line.FailedRecord[(line.FailedRecordHead+COUNT_MINUTES-i)%COUNT_MINUTES].Time).Minutes() >= float64(n) {
				break
			}
			total += line.FailedRecord[(line.FailedRecordHead+COUNT_MINUTES-i)%COUNT_MINUTES].FailedNum
		}
		return total
	}
	return -1
}

/* 获取线路负载量 */

/* 获取线路负载量 */


/*字节转换*/
func makeBuf(head *Head, msg proto.Message) ([]byte, error) {
	var body []byte
	var err error
	if msg != nil {
		if body, err = proto.Marshal(msg); err != nil {
			return nil, fmt.Errorf("GateClient.makeBuf proto.Marshal err %v", err)
		}
	}

	btHead := make([]byte,HEAD_LEN)
	head.Length = int32(MSGID_LEN) + int32(ERRCODE_LEN) + int32(len(body))
	binary.BigEndian.PutUint32(btHead[:4], uint32(head.Length))
	binary.BigEndian.PutUint32(btHead[4:8], uint32(head.MsgId))
	binary.BigEndian.PutUint32(btHead[8:12], uint32(head.ErrorCode))

	var buffer bytes.Buffer
	buffer.Write(btHead)
	buffer.Write(body)
	return buffer.Bytes(),nil
}

/* 创建新线路 */
func NewLine(id, remote string, timeout time.Duration, interval time.Duration, maxLoad int64, redirect bool) *Line {
	line := &Line{
		LineID:	   	   id,
		Remote:        remote,
		Delay:         UnreachableTime,
		Timeout:       timeout,
		Interval:      interval,
		Timer:         nil,
		CurLoad:       0,
		MaxLoad:       maxLoad,
		IsPaused:      false,
		Redirect:      redirect,
		ChUpdateDelay: make(chan time.Duration, 1024),
	}
	failed := FailedInMunite{
		Time: time.Now(),
	}
	for i := 0; i < COUNT_MINUTES; i++ {
		line.FailedRecord[i] = failed
	}
	return line
}




