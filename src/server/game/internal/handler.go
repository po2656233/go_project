package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	. "server/base"
	"server/game/internal/gameItems/baccarat"
	"server/game/internal/gameItems/cowcow"
	landlord "server/game/internal/gameItems/landlords"
	"server/game/internal/gameItems/mahjong"
	protoMsg "server/msg/go"
	"server/game/internal/gameItems/chineseChess"
)

//初始化
func init() {
	handlerBehavior()
}

//注册传输消息
func handlerMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//玩家行为处理
func handlerBehavior() {
	handlerMsg(&protoMsg.ReqEnterGame{}, enter)
	handlerMsg(&protoMsg.ReqExitGame{}, exit)
	handlerMsg(&protoMsg.GameBet{}, playing)
	handlerMsg(&protoMsg.GameHost{}, host)
	handlerMsg(&protoMsg.GameSuperHost{}, superHost)
	handlerMsg(&protoMsg.GameReady{}, ready)

}

//创建游戏对象
func productGame(kindID, level uint32) interface{} {
	switch kindID {
	case Baccarat:
		return baccarat.NewBaccarat(level, skeleton)
	case FishLord:
	case Landlords:
		return landlord.NewLandlord(level, skeleton)
	case CowCow:
		return cowcow.NewCowcow(level, skeleton)
	case Mahjong:
		return mahjong.NewMahjong(level, skeleton)
	case ChinessChess:
		return chineseChess.NewChineseChess(level, skeleton)
	}

	return nil
}

//进入
func enter(args []interface{}) {
	//查找玩家
	m := args[0].(*protoMsg.ReqEnterGame)
	agent := args[1].(gate.Agent)
	if player := playerManger.Get_1(agent); player != nil { //[0
		//补充用户信息
		player.Sate = PlayerLookOn
		player.GameID = m.GetGameID()
		var enterArgs []interface{}
		if gameLimit := sqlHandle.CheckGameInfo(player.GameID); gameLimit != nil { //[1-0
			//补充游戏实例
			if room, _ := roomManger.Create(player.RoomNum); room != nil { //[1
				if _, isOk := room.GetGameHandle(gameLimit.KindID, gameLimit.Level); !isOk { //[2
					//创建游戏实例
					if instance := productGame(gameLimit.KindID, gameLimit.Level); instance != nil { //[3
						item := &GameItem{KindID: gameLimit.KindID, Level: gameLimit.Level}
						//游戏实例
						item.Instance = instance.(IGameOperate)
						//房间添加游戏资源
						room.AddSource(item)
					}
				}
			}
			//游戏种类 游戏级别
			enterArgs = append(enterArgs, gameLimit)
		}
		log.Debug("场景参数:%v GameID:%v", enterArgs, player.GameID)
		//玩家进入(游戏场景)
		if 0 < len(enterArgs) {
			player.Enter(enterArgs)
		}

	}
}

func exit(args []interface{}) {
	m := args[0].(*protoMsg.ReqExitGame)
	//查找玩家
	if player := playerManger.Get_1(args[1].(gate.Agent)); player != nil { //[0
		if gameInfo := sqlHandle.CheckGameInfo(m.GameID); gameInfo != nil {
			var hostArgs []interface{}
			hostArgs = append(hostArgs, gameInfo, args[0])
			player.Out(hostArgs)
		}

	}
}

//游戏
func playing(args []interface{}) {
	agent := args[1].(gate.Agent)
	if player := playerManger.Get_1(agent); player != nil { //[0
		if gameInfo := sqlHandle.CheckGameInfo(player.GameID); gameInfo != nil {
			if room, ok := roomManger.Check(player.RoomNum); ok {
				if handle, ok := room.GetGameHandle(gameInfo.KindID, gameInfo.Level); ok {
					//下注
					handle.Playing(args)
				} else {
					log.Debug("game bet:INVALID!")
				}
			} else {
				log.Debug("room:INVALID!")
			}
		}
	} else {
		log.Debug("no player!")
		tempPlayer := &Player{Agent: agent}
		tempPlayer.Feedback(MainGameFrame, SubGameFrameBetResult, FAILD, string("no player!"))
	}

}

//抢庄
func host(args []interface{}) {
	//查找玩家
	if player := playerManger.Get_1(args[1].(gate.Agent)); player != nil { //[0
		if gameInfo := sqlHandle.CheckGameInfo(player.GameID); gameInfo != nil {
			var hostArgs []interface{}
			hostArgs = append(hostArgs, gameInfo, args[0])
			player.Host(hostArgs)
		}

	}
}

//超级抢庄
func superHost(args []interface{}) {
	//查找玩家
	if player := playerManger.Get_1(args[1].(gate.Agent)); player != nil { //[0
		if gameInfo := sqlHandle.CheckGameInfo(player.GameID); gameInfo != nil {
			var hostArgs []interface{}
			hostArgs = append(hostArgs, gameInfo, args[0])
			player.SuperHost(hostArgs)
		}

	}
}

// 玩家准备
func ready(args []interface{}) {
	//查找玩家
	if player := playerManger.Get_1(args[1].(gate.Agent)); player != nil { //[0
		if gameInfo := sqlHandle.CheckGameInfo(player.GameID); gameInfo != nil {
			var readyArgs []interface{}
			readyArgs = append(readyArgs, gameInfo, args[0])
			player.Ready(readyArgs)
		}

	}
}

//测试用
//func handleHello(args []interface{}) {
////////////////////////
//// 收到的 Hello 消息
//m := args[0].(*msg.Hello)
//// 消息的发送者
//a := args[1].(gate.Agent)
//
//// 输出收到的消息的内容
//log.Debug("hello %v", m.Name)
//
//// 给发送者回应一个 Hello 消息
//a.WriteMsg(&msg.Hello{
//	Name: "client",
//})
//}
