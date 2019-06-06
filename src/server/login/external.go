package login

//对外暴露接口
import (
	"server/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
