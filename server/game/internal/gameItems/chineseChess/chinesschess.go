package chineseChess

import (
	"github.com/name5566/leaf/log"
	. "server/base"
	. "server/manger"
	. "server/game/internal/gameItems"
	protoMsg "server/msg/go"
	_ "server/sql/mysql"
	"github.com/name5566/leaf/gate"
)


var playerList protoMsg.UserList  //玩家列表(观看)


//定时器
const (
	freeTime = 3
	betTime  = 10
	openTime = 7
)

//继承于GameItem
type ChineseChessGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

}

//创建中国象棋实例
func NewChineseChess(level uint32) *ChineseChessGame {
	log.Debug("------正在创建中国象棋实例------level:%v", level)
	p := &ChineseChessGame{}
	p.Init(level)
	return p
}

//---------------------------中国象棋----------------------------------------------//
//初始化信息
func (self *ChineseChessGame) Init(level uint32) {
	self.bStart = false               //是否第一次启动
	self.Level = level                //房间等级
	self.gameState = SubGameSenceFree //场景状态
}

//游戏接口
//场景
func (self *ChineseChessGame) Scene(args []interface{}) {

	level := args[0].(uint32)
	agent := args[1].(gate.Agent)
	userData := agent.UserData()
	if userData == nil {
		log.Debug("[Error][百家乐场景] [未能查找到相关玩家] ")
		return
	}
	//加入玩家列表
	player := userData.(*Player)
	self.AddPlayer(player.UserID)

	//场景信息
	senceInfo := &protoMsg.GameBaccaratEnter{}

	//玩家信息
	var playerInfo protoMsg.PlayerInfo
	playerInfo.UserID = player.UserID
	playerInfo.Name = player.Name
	playerInfo.Age = player.Age
	playerInfo.Gold = int64(GlobalSqlHandle.CheckMoney(player.UserID)* 100)  //玩家积分
	playerInfo.VipLevel = player.Level
	playerInfo.Sex = player.Sex
	senceInfo.UserInfo = &playerInfo

	if senceInfo.UserInfo == nil {
		log.Debug("[Error][百家乐场景] [获取玩家ID:%v 信息失败]  ", player.UserID)
		return
	}

	senceInfo.FreeTime = freeTime
	senceInfo.BetTime = betTime
	senceInfo.OpenTime = openTime
	//senceInfo.AwardAreas // 录单
	//需优化[定时器中计算时长]
	senceInfo.TimeStamp = self.timeStamp //////已过时长 应当该为传时间戳
	switch level {
	case RoomGeneral:
		senceInfo.Chips = []int32{1, 5, 25, 50, 100} //筹码
	case RoomMiddle:
		senceInfo.Chips = []int32{10, 50, 100, 500, 1000, 5000} //筹码
	case RoomHigh:
		senceInfo.Chips = []int32{50, 100, 200, 500, 1000, 10000} //筹码
	default:
		senceInfo.Chips = []int32{1, 5, 10, 20, 50, 100} //筹码
	}

	//数据回包

	GlobalSender.SendData(agent,MainGameSence, self.gameState, senceInfo)


	// 更新游戏列表
	var updateArgs []interface{}
	updateArgs = append(updateArgs, uint32(GameUpdatePlayerList),player.UserID)
	playerList.AllInfos = append(playerList.AllInfos, &playerInfo)
	self.UpdateInfo(updateArgs)
	log.Debug("[百家乐场景]->玩家信息 ID:%v 当前玩家列表:%v ", player.UserID, self.PlayerList)

}
//开始
func (self *ChineseChessGame) Start(args []interface{}) {

}

//游戏
func (self *ChineseChessGame) Playing(args []interface{}) {

}

//结算
func (self *ChineseChessGame) Over(args []interface{}) {

}

//更新信息
func (self *ChineseChessGame) UpdateInfo(args []interface{}){

}

//超级控制 可在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)
func (self *ChineseChessGame) SuperControl(args []interface{}) {

}