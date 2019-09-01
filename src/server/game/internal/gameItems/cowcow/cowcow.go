package cowcow

//牛牛
import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/gate"
	. "server/base"
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
	"sync"
	"time"
	"github.com/golang/protobuf/proto"
	. "server/game/internal/gameItems"
)

var timer *module.Skeleton = nil  //定时器
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表

var lock *sync.Mutex = &sync.Mutex{} //锁
//定时器
const (
	freeTime  = 3
	betTime = 10
	openTime = 7
	tableSite = 4
)

//继承于GameItem
type CowcowGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

	bankerID    uint64  //庄家ID
	bankerScore float32 //庄家积分
	readyCount  uint8   //准备的人数

	playerBetInfo map[uint64][]protoMsg.GameBet //下注信息

	overResult    *protoMsg.GameCowcowOver // 游戏结果
}

//创建百家乐实例
func NewCowcow(level uint32, skeleton *module.Skeleton) *CowcowGame {
	log.Debug("------正在创建牛牛实例------level:%v", level)
	p := &CowcowGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------牛牛----------------------------------------------//
//初始化信息
func (self *CowcowGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  // 定时器的使用
	self.bStart = false               // 是否第一次启动
	self.Level = level                // 房间等级
	self.gameState = SubGameSenceFree // 场景状态
	self.bankerID = 0                 //庄家ID
}

func (self *CowcowGame) Scene(args []interface{})   {
	userID := args[0].(uint64)
	level := args[1].(uint32)

	player := manger.Get(userID)
	if player == nil {
		log.Debug("[Error][牛牛场景] [未能查找到相关玩家] ID:%v", userID)
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
		log.Debug("[Error][牛牛场景] [获取玩家ID:%v 信息失败]  ", userID)
		return
	}

	log.Debug("[牛牛场景] [玩家列表新增] ID:%v 当前玩家总数:%v", userID, len(playerList.AllInfos))
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

	//
	player.WillReceive(MainGameSence, self.gameState, senceInfo)
	manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	log.Debug("[牛牛场景]->玩家信息 ID:%v  ", player.UserID)
}

//更新
func (self *CowcowGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]

	log.Debug("[牛牛]更新信息:%v-> %v\n", args[0].(uint32), args[1])
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
func (self *CowcowGame) Start(args []interface{}) {
	//直接扣除金币
	if !self.bStart {
		self.onStart()
		self.bStart = true
	}
	//return
	//m := args[0].(*protoMsg.GameCowcowPlaying)

}

// 出牌
func (self *CowcowGame) Playing(args []interface{}) {
	//【消息】
	m := args[0].(*protoMsg.GameBet)
	//【传输对象】
	agent := args[1].(gate.Agent)
	log.Debug("[牛牛]BaccaratPlaying:->%v", m)

	//var userID uint64 = 0
	player := manger.Get_1(agent)
	if nil == player {
		log.Debug("下注失败,")
		return
	}

	//反馈下注情况
	betResult := &protoMsg.GameBetResult{}
	betResult.UserID = player.UserID
	betResult.BetArea = m.BetArea
	betResult.BetScore = m.BetScore

	if self.gameState != SubGameSencePlaying {
		betResult.State = *proto.Int32(1)
		betResult.Hints = *proto.String("过了下注时间")
		player.WillReceive(MainGameFrame, SubGameFrameBetResult, betResult)
		return
	}

	//数据库中扣除玩家金币[下注成功]
	if money, ok := sqlHandle.DeductMoney(player.UserID, m.BetScore); !ok {
		betResult.State = *proto.Int32(1)
		betResult.Hints = *proto.String("数据库里的钱不够")
		log.Debug("[牛牛]下注失败 玩家ID:%v 现有金币:%v 下注金币:%v", player.UserID, money, m.BetScore)
		player.WillReceive(MainGameFrame, SubGameFrameBetResult, betResult)
		return
	}

	//下注成功
	betResult.State = *proto.Int32(0)
	betResult.Hints = *proto.String("[牛牛]下注成功")

	//

	//通知所有人
	log.Debug("[牛牛]反馈下注信息:->%v", m)

	//下注数目累加
	if areaBetInfos, ok := self.playerBetInfo[player.UserID]; ok {
		ok = false
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				self.playerBetInfo[player.UserID][index].BetScore = betItem.BetScore + m.BetScore
				log.Debug("[牛牛]玩家:%v 累加:->%v", player.UserID, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok {
			self.playerBetInfo[player.UserID] = CopyInsert(areaBetInfos, len(areaBetInfos), *m).([]protoMsg.GameBet)
		}
	} else {
		log.Debug("[牛牛]第一次:%v", m)
		self.playerBetInfo[player.UserID] = CopyInsert(self.playerBetInfo[player.UserID], len(self.playerBetInfo[player.UserID]), *m).([]protoMsg.GameBet)
	}
	player.WillReceive(MainGameFrame, SubGameFrameBetResult, betResult)

	//通知其他玩家
	//manger.NotifyButOthers(self.PlayerList, MainGameFrame, SubGameFramePlaying, m)
	manger.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFramePlaying, m)

}

// 结算
func (self *CowcowGame) Over(args []interface{}) {
	//直接扣除金币
	log.Debug("结算")
	//【发牌】
	self.dispatchCard()
	//【计算结果并反馈】
	self.calculateScore()

}

// 操作
func (self *CowcowGame) SuperControl(args []interface{}) {
	//直接扣除金币
	log.Debug("操作")
	//return
	//m := args[0].(*protoMsg.GameCowcowPlaying)
}

// 游戏业务层
//重新初始化
func (self *CowcowGame) reset() {
	log.Release("[牛牛]扫地僧出来干活了...")
	self.playerBetInfo = make(map[uint64][]protoMsg.GameBet) //玩家下注信息
	//hostList = []uint64{} 										//清理玩家

	//不要清理数据,因为数据清除之后,下一轮没法广播数据
	//playerCouriers = nil 【清内存】
	//for k,_:=range playerCouriers {//【数据清理】
	//	delete(playerCouriers,k)
	//}
}


func (self *CowcowGame)onStart(){
	log.Debug("游戏开始")
	self.reset() //重置
	log.Release("[牛牛]游戏开始")
	self.gameState = SubGameSenceStart
	self.timeStamp = time.Now().Unix()
	if timer == nil {
		time.AfterFunc(freeTime*time.Second, self.onPlay)
	} else {
		timer.AfterFunc(freeTime*time.Second, self.onPlay)
	}

	manger.NotifyOthers(self.PlayerList, MainGameState, SubGameStateStart, nil)
}

func (self *CowcowGame)onPlay(){
	log.Release("[百家乐]允许下注")
	self.gameState = SubGameSencePlaying
	self.timeStamp = time.Now().Unix()
	if timer == nil {
		time.AfterFunc(betTime*time.Second, self.onOver)
	} else {
		timer.AfterFunc(betTime*time.Second, self.onOver)
	}

	manger.NotifyOthers(self.PlayerList, MainGameState, SubGameStatePlaying, nil)
}

func (self *CowcowGame)onOver(){
	log.Release("[牛牛]开牌中...")
	self.gameState = SubGameSenceOver
	self.timeStamp = time.Now().Unix()
	if timer == nil {
		time.AfterFunc(openTime*time.Second, self.onStart)
	} else {
		timer.AfterFunc(openTime*time.Second, self.onStart)
	}
	//当有玩家结算信息时,该
	manger.NotifyOthers(self.PlayerList, MainGameState, SubGameStateOver, nil)

	log.Release("[牛牛]结算中...")
	self.Over(nil)
	log.Release("[牛牛]完成一轮游戏")
}

//-----------------------逻辑层---------------------------
//发牌
func (self *CowcowGame) dispatchCard() {
	//25 张牌
	tableCards := RandCardList( PiceCount * AREA_MAX )
	log.Debug("[牛牛-->开始发牌啦<---]  牌堆中取牌:%v ", GetCardsText(tableCards))
	copy(self.overResult.BankerCard,tableCards[INDEX_Banker:INDEX_Banker+PiceCount])
	copy(self.overResult.BankerCard,tableCards[INDEX_Tian: INDEX_Tian+PiceCount])
	copy(self.overResult.BankerCard,tableCards[INDEX_Xuan: INDEX_Xuan+PiceCount])
	copy(self.overResult.BankerCard,tableCards[INDEX_Di: INDEX_Di+PiceCount])
	copy(self.overResult.BankerCard,tableCards[INDEX_Huang: INDEX_Huang+PiceCount])
}
//结算
func (self *CowcowGame) calculateScore() {

}
func (self *CowcowGame) deduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	//庄家的牌
	//判断牌型
	//judgeCardType(self.overResult.BankerCard)

	//庄家区域
	return pWinArea
}

//区域赔额
func (self *CowcowGame) bonusArea(CardType int, betScore int64) int64 {
	multiple := int64(0)
	switch CardType {
	case MULTIPLE_Normal:
		multiple = MULTIPLE_Normal
	case MULTIPLE_Middle:
		multiple = MULTIPLE_Middle
	case MULTIPLE_High:
		multiple = MULTIPLE_High
	case MULTIPLE_WuHuaNiu:
		multiple = MULTIPLE_WuHuaNiu
	case MULTIPLE_QuanHuaNiu:
		multiple = MULTIPLE_QuanHuaNiu
	case MULTIPLE_ZhaDan:
		multiple = MULTIPLE_ZhaDan
	default:
		multiple = int64(0)
	}
	return betScore * multiple
}

