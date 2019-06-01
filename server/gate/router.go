package gate

import (
	"server/game"
	"server/login"
	"server/msg"
	protoMsg "server/msg/go"
)

//路由模块分发消息【模块间使用 ChanRPC 通讯，消息路由也不例外】
//注:需要解析的结构体才进行路由分派，即用客户端主动发起的
func init() {

	//大厅登陆
	loginSection()

	//入口节点
	enterSection()

	//游戏节点
	gameSection()

	//【任务】
	taskSection()
}

func loginSection() {
	//登录注册信息[protobuf]
	msg.ProcessorProto.SetRouter(&protoMsg.Login{}, login.ChanRPC)
	msg.ProcessorProto.SetRouter(&protoMsg.Register{}, login.ChanRPC)

	//个人中心信息//游戏房间信息//大厅通告信息
}

func enterSection() {
	//进入房间
	msg.ProcessorProto.SetRouter(&protoMsg.ReqEnterRoom{}, login.ChanRPC) //[proto]
	//进入游戏
	msg.ProcessorProto.SetRouter(&protoMsg.ReqEnterGame{}, game.ChanRPC) //[proto]
	//退出游戏
	msg.ProcessorProto.SetRouter(&protoMsg.ReqExitGame{}, game.ChanRPC) //[proto]
}

func gameSection() {
	//子游戏消息[protobuf]
	msg.ProcessorProto.SetRouter(&protoMsg.GameBet{}, game.ChanRPC)       //[proto]
	msg.ProcessorProto.SetRouter(&protoMsg.GameHost{}, game.ChanRPC)      //[proto]
	msg.ProcessorProto.SetRouter(&protoMsg.GameSuperHost{}, game.ChanRPC) //[proto]
	msg.ProcessorProto.SetRouter(&protoMsg.GameReady{}, game.ChanRPC)     //[proto]

	baccaratSetRouter()

	mahjongSetRouter()

	landLordsSetRouter()

	cowcowSetRouter()

	//这里指定消息 Hello 路由到 game 模块[json格式]
	//msg.ProcessorJson.SetRouter(&jsonMsg.UserST{}, game.ChanRPC) //json
}

func taskSection() {

}

//百家乐
func baccaratSetRouter() {
	msg.ProcessorProto.SetRouter(&protoMsg.GameBaccaratHost{}, game.ChanRPC)      //抢庄
	msg.ProcessorProto.SetRouter(&protoMsg.GameBaccaratSuperHost{}, game.ChanRPC) //超级抢庄
	msg.ProcessorProto.SetRouter(&protoMsg.GameBaccaratBet{}, game.ChanRPC)       //下注
}

// 麻将
func mahjongSetRouter() {
	msg.ProcessorProto.SetRouter(&protoMsg.GameMahjongOutcard{}, game.ChanRPC) // 出牌
	msg.ProcessorProto.SetRouter(&protoMsg.GameMahjongOperate{}, game.ChanRPC) // 操作
}

// 斗地主
func landLordsSetRouter() {
	msg.ProcessorProto.SetRouter(&protoMsg.GameLandLordsOutcard{}, game.ChanRPC) // 出牌
	msg.ProcessorProto.SetRouter(&protoMsg.GameLandLordsOperate{}, game.ChanRPC) // 操作
}

//牛牛GameBet
func cowcowSetRouter() {
	msg.ProcessorProto.SetRouter(&protoMsg.GameCowcowPlaying{}, game.ChanRPC) //下注
}
