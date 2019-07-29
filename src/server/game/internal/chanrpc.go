package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	. "server/base"
	"server/sql/mysql"
)

//广播消息
//这里是对所有玩家进行通知，通知单个游戏的所有玩家，请在单个游戏里实现
var agents = make(map[gate.Agent]struct{})
var sqlHandle = mysql.SqlHandle()
var playerManger *PlayerManger = GetPlayerManger()
var roomManger *RoomManger = GetRoomManger()
//

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
	_ = a
	agents[a] = struct{}{}

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	delete(agents, a)

	log.Debug("玩家断线:%v", a.LocalAddr())

	//断线通知
	if player := playerManger.Get_1(a); nil != player {
		//游戏内通知有玩家断开
		if room, ok := roomManger.Check(player.RoomNum); ok {
			if gameInfo := sqlHandle.CheckGameInfo(player.GameID); gameInfo != nil {
				if game, isOk := room.GetGameHandle(gameInfo.KindID, gameInfo.Level); isOk {
					var updateArgs []interface{}
					updateArgs = append(updateArgs, uint32(GameUpdateOut), player.UserID)
					game.UpdateInfo(updateArgs)
					log.Debug("游戏中清理玩家%v 数据", player.UserID)
				}
			}
		}
		log.Debug("管理列表中->清理玩家%v 数据", player.UserID)
		playerManger.DeletePlayer(player)
	} else {
		log.Debug(string("找不到对应的玩家数据:%v"), a.LocalAddr())
	}
}

func rpcBroadcast(args []interface{}) interface{} {
	//断线通知
	a := args[0].(gate.Agent)
	_ = a
	a.WriteMsg(args[1])
	return error(nil)
}
