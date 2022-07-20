package login

//对外暴露接口
import (
	"miniRobot/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func InitData() {
	//atomic.StoreInt32(&internal.IndexGames,0)
}
