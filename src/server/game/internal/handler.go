package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	. "server/manger"
	. "server/base"
	"server/game/internal/gameItems/baccarat"
	"server/game/internal/gameItems/cowcow"
	landlord "server/game/internal/gameItems/landlords"
	"server/game/internal/gameItems/mahjong"
	protoMsg "server/msg/go"
	"server/game/internal/gameItems/chineseChess"

	"server/sql/mysql"
	"server/game/internal/gameItems"
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
	//斗地主出牌
	handlerMsg(&protoMsg.GameLandLordsOutcard{}, playing)

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
	_ = args[1]
	m := args[0].(*protoMsg.ReqEnterGame)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		//玩家坐下
		player := userData.(*Player)
		player.Sate = PlayerSitDown
		player.GameID = m.GetGameID()

		if platform := GetPlatformManger().Get(player.PlatformID); nil != platform {
			if gameLimit := mysql.SqlHandle().CheckGameInfo(player.GameID); gameLimit != nil { //[1-0
				//补充游戏实例
				if room, _ := platform.Roomer.Create(player.RoomNum); room != nil { //[1
					if gameItem, isOk := room.CheckGame(gameLimit.KindID, gameLimit.Level); !isOk { //[2
						//创建游戏实例
						if instance := productGame(gameLimit.KindID, gameLimit.Level); instance != nil { //[3
							item := &GameItem{KindID: gameLimit.KindID, Level: gameLimit.Level}
							//游戏实例
							item.Instance = instance.(IGameOperate)
							//房间增加游戏
							room.AddSource(item)
							//玩家所在游戏
							player.Game = item
						}
					} else {
						//玩家所在游戏
						player.Game = gameItem
					}
					//玩家进入(游戏场景)
					var enterArgs []interface{}
					enterArgs = append(enterArgs, gameLimit, agent)
					player.Enter(enterArgs)
					log.Debug("场景参数:%v GameID:%v", enterArgs, player.GameID)
				} else {
					gameItems.GlobalSender.SendData(agent, MainLogin, SubEnterGameResult, &protoMsg.ResResult{State: FAILD, Hints: string("error:room Permission denied!")})
					log.Debug(" 房间:%v 不存在!", player.GameID,player.RoomNum)
				}
			} else {
				gameItems.GlobalSender.SendData(agent, MainLogin, SubEnterGameResult, &protoMsg.ResResult{State: FAILD, Hints: string("error:not Game Handle!")})
				log.Debug("游戏ID:%v 没有对应的游戏实例", player.GameID)
			}
		} else {
			gameItems.GlobalSender.SendData(agent, MainLogin, SubEnterGameResult, &protoMsg.ResResult{State: FAILD, Hints: string("error:not platform info!")})
			log.Debug("不存在 %v平台 ", player.PlatformID)
		}
	}
}

//游戏
func playing(args []interface{}) {
	//查找玩家
	_ = args[1]
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		player := userData.(*Player)
		//玩家状态:游戏中
		player.Sate = PlayerPlaying
		if game := player.Game; nil != game {
			if gameHandel := game.Instance; nil != gameHandel {
				gameHandel.Playing(args)
			}
		}
	}

}

// 玩家准备
func ready(args []interface{}) {
	_ = args[1]
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		player := userData.(*Player)
		//玩家状态:游戏中
		player.Sate = PlayerAgree
		if game := player.Game; nil != game && nil != game.Instance {
			var readyArgs []interface{}
			readyArgs = append(readyArgs, game.KindID, game.Level, player.UserID, agent)
			player.Ready(readyArgs)
		}
	}
}
func exit(args []interface{}) {
	_ = args[1]
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		player := userData.(*Player)
		//玩家状态:游戏中
		player.Sate = PlayerStandUp
		if game := player.Game; nil != game && nil != game.Instance {
			var outArgs []interface{}
			outArgs = append(outArgs, game.KindID, game.Level, player.UserID, agent)
			player.Out(outArgs)
			player.Game = nil
		}
	}
}

//抢庄
func host(args []interface{}) {
	_ = args[1]
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		player := userData.(*Player)
		if game := player.Game; nil != game && nil != game.Instance {
			var hostArgs []interface{}
			hostArgs = append(hostArgs, game.KindID, game.Level, args[0], args[1])
			player.Host(hostArgs)
		}
	}
}

//超级抢庄
func superHost(args []interface{}) {
	_ = args[1]
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		player := userData.(*Player)
		if game := player.Game; nil != game && nil != game.Instance {
			var hostArgs []interface{}
			hostArgs = append(hostArgs, game.KindID, game.Level, args[0], args[1])
			player.SuperHost(hostArgs)
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
