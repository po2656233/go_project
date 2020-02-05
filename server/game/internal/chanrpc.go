package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	."server/manger"
	. "server/base"
	"server/sql/mysql"
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
	GetClientManger().Append(a)
	_ = a

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	log.Debug("玩家断线:%v", a.RemoteAddr())

	//断线通知
	GetClientManger().Delete(a)
	userData := a.UserData()
	if nil != userData {
		player := userData.(*Player)
		if nil != player {
			// 平台维度
			platformManger := GetPlatformManger().Get(player.PlatformID)
			// 房间维度
			if nil != platformManger && nil != platformManger.Roomer {
				if room, ok := platformManger.Roomer.Check(player.RoomNum); ok {
					// 游戏维度 【重中之重】只有gameID的话，则需要从数据库中查出KindID和Level. 因为当kindID和Level一样时才是同一款游戏，游戏不是通过GameID区分的
					if gameInfo := mysql.SqlHandle().CheckGameInfo(player.GameID); gameInfo != nil {
						if game, isOk := room.GetGameHandle(gameInfo.KindID, gameInfo.Level); isOk {
							// 更新游戏内部信息
							var updateArgs []interface{}
							updateArgs = append(updateArgs, uint32(GameUpdateOut), player.UserID,a)
							game.UpdateInfo(updateArgs)
							log.Debug("游戏中清理玩家%v 数据", player.UserID)
						}
					}
				}
			}

			// 玩家维度
			log.Debug("管理列表中->清理玩家%v 数据", player.UserID)
			GetPlayerManger().DeletePlayer(player)
		} else {
			log.Debug(string("找不到对应的玩家数据:%v"), a.LocalAddr())
		}
	}


}

func rpcBroadcast(args []interface{}) interface{} {
	//断线通知
	a := args[0].(gate.Agent)
	_ = a
	a.WriteMsg(args[1])
	return error(nil)
}
