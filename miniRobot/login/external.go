package login

//对外暴露接口
import (
	"miniRobot/login/internal"
	"sync/atomic"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func InitData(){
	atomic.StoreInt32(&internal.IndexGames,0)
}