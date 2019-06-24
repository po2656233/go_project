package chineseChess

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	. "server/base"
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
)

var timer *module.Skeleton = nil  //定时器
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表(观看)

//继承于GameItem
type ChineseChessGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

}

//创建中国象棋实例
func NewChineseChess(level uint32, skeleton *module.Skeleton) *ChineseChessGame {
	log.Debug("------正在创建中国象棋实例------level:%v", level)
	p := &ChineseChessGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------中国象棋----------------------------------------------//
//初始化信息
func (self *ChineseChessGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  //定时器的使用
	self.bStart = false               //是否第一次启动
	self.Level = level                //房间等级
	self.gameState = SubGameSenceFree //场景状态
}

//游戏接口
//场景
func (self *ChineseChessGame) Scene(args []interface{}) {
	userID := args[0].(uint64)
	level := args[1].(uint32)

	player := manger.Get(userID)
	if player == nil {
		log.Debug("[Error][中国象棋场景] [未能查找到相关玩家] ID:%v", userID)
		return
	}

	log.Debug("当前玩家总数:%v %v ", len(playerList.AllInfos), self.PlayerList)
	// 获取玩家列表
	self.AddPlayer(player.UserID) //加入玩家列表
	senceInfo := &protoMsg.GameBaccaratEnter{}
	senceInfo.UserInfo = nil
	for _, uid := range self.PlayerList {
		if playerItem := manger.Get(uid); nil != playerItem {
			if uid == player.UserID {
				var playerInfo protoMsg.PlayerInfo
				playerInfo.UserID = playerItem.UserID
				playerInfo.Name = playerItem.Name
				playerInfo.Age = playerItem.Age
				playerInfo.Gold = int64(sqlHandle.CheckMoney(playerItem.UserID)) * 100 //玩家积分
				playerInfo.VipLevel = playerItem.Level
				playerInfo.Sex = playerItem.Sex
				senceInfo.UserInfo = &playerInfo
				isHave := false
				for _, info := range playerList.AllInfos {
					if info.UserID == uid {
						isHave = true
						break
					}
				}
				if !isHave {
					playerList.AllInfos = CopyInsert(playerList.AllInfos, len(playerList.AllInfos), &playerInfo).([]*protoMsg.PlayerInfo)
				}
			}
		} else {
			manger.DeletePlayerIndex(uid)
		}
	}
	if senceInfo.UserInfo == nil {
		log.Debug("[Error][中国象棋场景] [获取玩家ID:%v 信息失败]  ", userID)
		return
	}

	log.Debug("[中国象棋场景] [玩家列表新增] ID:%v 当前玩家总数:%v", userID, len(playerList.AllInfos))
	//senceInfo.AwardAreas // 录单
	//需优化[定时器中计算时长]
	senceInfo.TimeStamp = self.timeStamp //////已过时长 应当该为传时间戳
	switch level {
	case RoomGeneral:
		senceInfo.Chips = []int32{1, 5, 10, 100, 500, 1000} //筹码
	case RoomMiddle:
		senceInfo.Chips = []int32{10, 50, 100, 500, 1000, 5000} //筹码
	case RoomHigh:
		senceInfo.Chips = []int32{50, 100, 200, 500, 1000, 10000} //筹码
	default:
		senceInfo.Chips = []int32{1, 5, 10, 20, 50, 100} //筹码
	}

	//
	player.WillReceive(MainGameSence, self.gameState, senceInfo)
	manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)

	log.Debug("[中国象棋场景]->玩家信息 ID:%v ", player.UserID)
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