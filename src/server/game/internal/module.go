package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"github.com/name5566/leaf/chanrpc"
)

var (
	skeleton  = base.NewSkeleton()
	ChanRPC   = skeleton.ChanRPCServer //模块之间可以通过这个进行交互
	AsyncChan = chanrpc.NewServer(10) // 测试异步chan用
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
	//PlayerManager.Close()
	//mgodb.Close()
	//log.Release("closed")
}
