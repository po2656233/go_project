package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
	protoMsg "server/msg/go"
	jsonMsg "server/msg/json"
	"sync"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var ProcessorJson = json.NewProcessor()
var ProcessorProto = protobuf.NewProcessor()
var wg sync.WaitGroup

func init() {
	//这里的注册顺序，必须，必须，必须与【客户端】一致 因为rpc根据ID解析字段的
	/////[protobuf]格式消息
	registerPacket()
}

//对外接口 【这里的注册函数并非线程安全】
func RegisterMessage(message proto.Message) {
	wg.Add(1)
	defer wg.Done()

	ProcessorProto.Register(message)

	//log.Debug("reg ID:%v",ProcessorProto.Register(message))
}

//-----------Protobuf-------------------
//【统一数据包】
func registerPacket() {
	RegisterMessage(&protoMsg.GameBaccaratEnter{})
	RegisterMessage(&protoMsg.GameBaccaratHost{})
	RegisterMessage(&protoMsg.GameBaccaratSuperHost{})
	RegisterMessage(&protoMsg.GameBaccaratBet{})
	RegisterMessage(&protoMsg.GameBaccaratBetResult{})
	RegisterMessage(&protoMsg.GameBaccaratOver{})
	RegisterMessage(&protoMsg.GameBaccaratCheckout{})
	RegisterMessage(&protoMsg.PacketData{})
	RegisterMessage(&protoMsg.GameCowcowEnter{})
	RegisterMessage(&protoMsg.GameCowcowHost{})
	RegisterMessage(&protoMsg.GameCowcowSuperHost{})
	RegisterMessage(&protoMsg.GameCowcowPlaying{})
	RegisterMessage(&protoMsg.GameCowcowBetResult{})
	RegisterMessage(&protoMsg.GameCowcowOver{})
	RegisterMessage(&protoMsg.GameCowcowCheckout{})
	RegisterMessage(&protoMsg.GameFishLordEnter{})
	RegisterMessage(&protoMsg.GameFishLordPlaying{})
	RegisterMessage(&protoMsg.GameFishLordBetResult{})
	RegisterMessage(&protoMsg.GameFishLordOver{})
	RegisterMessage(&protoMsg.PlayerInfo{})
	RegisterMessage(&protoMsg.UserList{})
	RegisterMessage(&protoMsg.PlayerRecord{})
	RegisterMessage(&protoMsg.GameReady{})
	RegisterMessage(&protoMsg.GameBet{})
	RegisterMessage(&protoMsg.GameBetResult{})
	RegisterMessage(&protoMsg.GameHost{})
	RegisterMessage(&protoMsg.GameSuperHost{})
	RegisterMessage(&protoMsg.GameRecord{})
	RegisterMessage(&protoMsg.GameRecordList{})
	RegisterMessage(&protoMsg.GameResult{})
	RegisterMessage(&protoMsg.GameLandLordsEnter{})
	RegisterMessage(&protoMsg.GameLandLordsPlayer{})
	RegisterMessage(&protoMsg.GameLandLordsBegins{})
	RegisterMessage(&protoMsg.GameLandLordsOutcard{})
	RegisterMessage(&protoMsg.GameLandLordsOperate{})
	RegisterMessage(&protoMsg.GameLandLordsAward{})
	RegisterMessage(&protoMsg.GameLandLordsCheckout{})
	RegisterMessage(&protoMsg.Register{})
	RegisterMessage(&protoMsg.RegisterResult{})
	RegisterMessage(&protoMsg.Login{})
	RegisterMessage(&protoMsg.ResResult{})
	RegisterMessage(&protoMsg.TaskItem{})
	RegisterMessage(&protoMsg.TaskList{})
	RegisterMessage(&protoMsg.GameList{})
	RegisterMessage(&protoMsg.UserInfo{})
	RegisterMessage(&protoMsg.RoomInfo{})
	RegisterMessage(&protoMsg.GameBaseInfo{})
	RegisterMessage(&protoMsg.GameItem{})
	RegisterMessage(&protoMsg.MasterInfo{})
	RegisterMessage(&protoMsg.ReqEnterRoom{})
	RegisterMessage(&protoMsg.ReqEnterGame{})
	RegisterMessage(&protoMsg.ReqExitGame{})
	RegisterMessage(&protoMsg.GameMahjongEnter{})
	RegisterMessage(&protoMsg.GameMahjongPlayer{})
	RegisterMessage(&protoMsg.GameMahjongBegins{})
	RegisterMessage(&protoMsg.GameMahjongOutcard{})
	RegisterMessage(&protoMsg.GameMahjongOperate{})
	RegisterMessage(&protoMsg.GameMahjongAward{})
	RegisterMessage(&protoMsg.GameMahjongCheckout{})

	//RegisterMessage(&protoMsg.PacketData{})
}

//登录
func registerLoginProtoMsg() {
	// 这个是protobuf的消息体
	//登录注册
	RegisterMessage(&protoMsg.Login{})
	RegisterMessage(&protoMsg.MasterInfo{})
	RegisterMessage(&protoMsg.ResResult{})
	RegisterMessage(&protoMsg.Register{})
	RegisterMessage(&protoMsg.ReqEnterRoom{})
	RegisterMessage(&protoMsg.ReqEnterGame{})
	RegisterMessage(&protoMsg.ReqExitGame{})

	//游戏房间列表
	RegisterMessage(&protoMsg.GameList{})
	RegisterMessage(&protoMsg.GameBet{})
	RegisterMessage(&protoMsg.GameBetResult{})
	RegisterMessage(&protoMsg.GameHost{})
	RegisterMessage(&protoMsg.GameSuperHost{})
	RegisterMessage(&protoMsg.GameReady{})
}

//游戏
func registerGameProtoMsg() {
	// 这个是protobuf的消息体
	//百家乐
	baccaratRegister()
	// 麻将
	mahjongRegister()
	// 斗地主
	landLordsRegister()

	//百人类[统一]
	cowcowRegister()

}

//子游戏
//[百家乐]
func baccaratRegister() {
	RegisterMessage(&protoMsg.GameBaccaratEnter{})     //入场
	RegisterMessage(&protoMsg.GameBaccaratHost{})      //抢庄
	RegisterMessage(&protoMsg.GameBaccaratSuperHost{}) //超级抢庄
	RegisterMessage(&protoMsg.GameBaccaratBet{})       //下注
	RegisterMessage(&protoMsg.GameBaccaratBetResult{}) //下注结果
	RegisterMessage(&protoMsg.GameBaccaratOver{})      //开奖
	//RegisterMessage(&protoMsg.GameBaccaratCheckout{})	//结算
}

func mahjongRegister() {
	RegisterMessage(&protoMsg.GameMahjongEnter{})    //入场
	RegisterMessage(&protoMsg.GameMahjongPlayer{})   //玩家信息
	RegisterMessage(&protoMsg.GameMahjongBegins{})   //开始信息
	RegisterMessage(&protoMsg.GameMahjongOutcard{})  //出牌
	RegisterMessage(&protoMsg.GameMahjongOperate{})  //操作
	RegisterMessage(&protoMsg.GameMahjongAward{})    //个人得分
	RegisterMessage(&protoMsg.GameMahjongCheckout{}) //所有得分
}
func landLordsRegister() {
	RegisterMessage(&protoMsg.GameLandLordsEnter{})    //入场
	RegisterMessage(&protoMsg.GameLandLordsPlayer{})   //玩家信息
	RegisterMessage(&protoMsg.GameLandLordsBegins{})   //开始信息
	RegisterMessage(&protoMsg.GameLandLordsOutcard{})  //出牌
	RegisterMessage(&protoMsg.GameLandLordsOperate{})  //操作
	RegisterMessage(&protoMsg.GameLandLordsAward{})    //个人得分
	RegisterMessage(&protoMsg.GameLandLordsCheckout{}) //所有得分
}

//[百人类游戏]
func cowcowRegister() {
	RegisterMessage(&protoMsg.GameCowcowEnter{})     //入场
	RegisterMessage(&protoMsg.GameCowcowPlaying{})   //下注
	RegisterMessage(&protoMsg.GameCowcowBetResult{}) //下注结果
	RegisterMessage(&protoMsg.GameCowcowOver{})      //结算
}

//-----------JSON-------------------
func registerLoginJsonMsg() {
	//json
	ProcessorJson.Register(&jsonMsg.UserLogin{})
}
