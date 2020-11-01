package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	protoMsg "miniRobot/msg/go"
)

//广播消息
//这里是对所有玩家进行通知，通知单个游戏的所有玩家，请在单个游戏里实现


func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	AsyncChan.Register("Broadcast", func(args []interface{}) {
		fmt.Println("-------------->Broadcast------->Register")
		//a := args[0].(gate.Agent)
		//_ = a
		//a.WriteMsg(args[1])
	}) // 广播消息 调用参考:game.ChanRPC.Go("Broadcast",agent,args)

}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent) //【模块间通信】跟路由之间的通信
	fmt.Println("-成功創建")
	//登录服务器
	msg:=&protoMsg.LoginReq{
		Account: "aaa",
		Password: "000",
	}
	a.WriteMsg(msg)
	_ = a

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a



}

func rpcBroadcast(args []interface{}) interface{} {
	//断线通知
	a := args[0].(gate.Agent)
	_ = a
	a.WriteMsg(args[1])
	return error(nil)
}
