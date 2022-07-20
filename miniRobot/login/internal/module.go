package internal

import (
	"github.com/name5566/leaf/module"
	"miniRobot/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	//IndexGames int32
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	m.SetIndex(-1)
}

func (m *Module) OnDestroy() {
	//log.Debug("销毁")
}

func (m *Module) SetIndex(index int32) {
	//atomic.CompareAndSwapInt32(&IndexGames, IndexGames, index)
}
