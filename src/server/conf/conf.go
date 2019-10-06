package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags|log.Lshortfile ////Lshortfile:输出文件名和行号 LstdFlags：标准输出

	// gate conf
	PendingWriteNum        = 2000             //挂起数目
	MaxMsgLen       uint32 = 4096             //包体最大长度
	HTTPTimeout            = 10 * time.Second //HTTP网络延迟10秒
	LenMsgLen              = 2                //所占字节长度
	LittleEndian           = false            //小端模式

	// skeleton conf
	GoLen              = 10000 //go在骨架（skeleton）中可使用的数目
	TimerDispatcherLen = 10000 //定时器分配数
	AsynCallLen        = 10000 //异步调用的数目
	ChanRPCLen         = 10000 //信道RPC的数目
)
