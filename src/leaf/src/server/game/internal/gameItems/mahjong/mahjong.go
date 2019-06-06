package mahjong

//麻将

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	. "server/base"
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
	"sync"
)

var timer *module.Skeleton = nil  //定时器
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表

var lock *sync.Mutex = &sync.Mutex{} //锁
//定时器
const (
	freeTime  = 5
	tableSite = 4
)

//继承于GameItem
type MahjongGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

	bankerID    uint64  //庄家ID
	bankerScore float32 //庄家积分
	readyCount  uint8   //准备的人数
}

//创建百家乐实例
func NewMahjong(level uint32, skeleton *module.Skeleton) *MahjongGame {
	log.Debug("------正在创建麻将实例------level:%v", level)
	p := &MahjongGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------麻将----------------------------------------------//
//初始化信息
func (self *MahjongGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  // 定时器的使用
	self.bStart = false               // 是否第一次启动
	self.Level = level                // 房间等级
	self.gameState = SubGameSenceFree // 场景状态
	self.bankerID = 0                 //庄家ID
}

func (self *MahjongGame) Scene(args []interface{}) {
	userID := args[0].(uint64)
	level := args[1].(uint32)

	player := manger.Get(userID)
	if player == nil {
		log.Debug("[Error][麻将] [未能查找到相关玩家] ID:%v", userID)
		return
	}

	// 获取玩家列表z
	self.AddPlayer(player.UserID) //加入玩家列表
	for _, uid := range self.PlayerList {
		if playerItem := manger.Get(uid); nil != playerItem {
			var playerInfo protoMsg.PlayerInfo
			playerInfo.UserID = playerItem.UserID
			playerInfo.Name = playerItem.Name
			playerInfo.Age = playerItem.Age
			playerInfo.Gold = sqlHandle.CheckMoney(playerItem.UserID) //玩家积分
			playerInfo.VipLevel = playerItem.Level
			playerInfo.Sex = playerItem.Sex
			playerList.AllInfos = CopyInsert(playerList.AllInfos, len(playerList.AllInfos), &playerInfo).([]*protoMsg.PlayerInfo)

		} else {
			manger.DeletePlayerIndex(uid)
		}
	}

	log.Debug("[麻将] [玩家列表新增] ID:%v", userID)
	senceInfo := &protoMsg.GameMahjongEnter{}
	senceInfo.UserID = player.UserID
	senceInfo.Players = &playerList //玩家列表[TODO]
	senceInfo.FreeTime = freeTime
	//senceInfo.AwardAreas // 录单
	//需优化[定时器中计算时长]
	senceInfo.TimeStamp = self.timeStamp //////已过时长 应当该为传时间戳
	switch level {
	case RoomGeneral:
	case RoomMiddle:
	case RoomHigh:
	default:
	}

	player.WillReceive(MainGameSence, self.gameState, senceInfo)
	log.Debug("[麻将场景]->玩家信息 ID:%v ", player.UserID)
}

//更新
func (self *MahjongGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]

	log.Debug("[麻将]更新信息:%v-> %v\n", args[0].(uint32), args[1])
	flag := args[0].(uint32)
	userID := args[1].(uint64)
	switch flag {
	case GameUpdateOut: //玩家离开 不再向该玩家广播消息[] 删除
		self.DeletePlayer(userID)
		manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdatePlayerList: //更新玩家列表
		self.AddPlayer(userID)
		manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdateHost: //更新玩家抢庄信息
	case GameUpdateSuperHost: //更新玩家超级抢庄信息
	case GameUpdateOffline: //更新玩家超级抢庄信息
	case GameUpdateReconnect: //更新玩家超级抢庄信息
	case GameUpdateReady: //统计准备的玩家
		self.readyCount++
		if tableSite == self.readyCount {
			self.Start(nil)
		}
	}
}

// 开始
func (self *MahjongGame) Start(args []interface{}) {
	//直接扣除金币
	log.Debug("游戏开始")
	//return
	//m := args[0].(*protoMsg.GameMahjongPlaying)

}

// 出牌
func (self *MahjongGame) Playing(args []interface{}) {
	//直接扣除金币
	log.Debug("出牌")
	//return
	//m := args[0].(*protoMsg.GameMahjongPlaying)

}

// 结算
func (self *MahjongGame) Over(args []interface{}) {
	//直接扣除金币
	log.Debug("结算")
	//return
	//m := args[0].(*protoMsg.GameMahjongPlaying)

}

// 操作
func (self *MahjongGame) SuperControl(args []interface{}) {
	//直接扣除金币
	log.Debug("操作")
	//return
	//m := args[0].(*protoMsg.GameMahjongPlaying)
}

// 游戏业务层
//重新初始化
func (self *MahjongGame) reset() {
	log.Release("[麻将]扫地僧出来干活了...")

	//hostList = []uint64{} 										//清理玩家

	//不要清理数据,因为数据清除之后,下一轮没法广播数据
	//playerCouriers = nil 【清内存】
	//for k,_:=range playerCouriers {//【数据清理】
	//	delete(playerCouriers,k)
	//}
}

//-----------------------逻辑层---------------------------
//发牌
func (self *MahjongGame) dispatchCard() {

}
