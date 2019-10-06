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
	"strconv"
)

var timer *module.Skeleton = nil  //定时器
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表

var lock *sync.Mutex = &sync.Mutex{} //锁
//定时器
const (
	freeTime  = 3
	betTime = 3
	openTime = 6
	tableSite = 4
)

//继承于GameItem
type CowcowGame struct {
	GameItem
	bStart    bool   	//第一次启动
	gameState uint32 	//游戏状态
	timeStamp int64  	//时间戳

	logic *Game		 	//游戏逻辑
	roundNumber string 	//游戏局号

	bankerID    uint64  //庄家ID
	bankerScore float32 //庄家积分
	readyCount  uint8   //准备的人数

	playerBetInfo map[uint64][]protoMsg.GameBet //下注信息

	cbBankerCards []byte      	// 庄家手牌
	cbTianCards []byte      	// 天家手牌
	cbXuanCards []byte      	// 玄家手牌
	cbDiCards []byte        	// 地家手牌
	cbHuangCards []byte      	// 黄家手牌
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

	self.logic = &Game{}
	self.logic.Init()
	self.roundNumber = ""

	self.cbBankerCards = make([]byte, PiceCount)       // 庄家手牌
	self.cbTianCards = make([]byte, PiceCount)       	// 天家手牌
	self.cbXuanCards = make([]byte, PiceCount)       	// 玄家手牌
	self.cbDiCards = make([]byte, PiceCount)       	// 地家手牌
	self.cbHuangCards = make([]byte, PiceCount)       	// 黄家手牌
	self.overResult = &protoMsg.GameCowcowOver{} 		// 结算结果
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
				playerInfo.Gold = int64(sqlHandle.CheckMoney(playerItem.UserID)* 100)  //玩家积分
				player.Money = playerInfo.Gold
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
	log.Debug("[牛牛场景]->玩家信息 ID:%v gold:%v ", player.UserID, player.Money)
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
	self.cbBankerCards = make([]byte, PiceCount)       // 庄家手牌
	self.cbTianCards = make([]byte, PiceCount)       	// 天家手牌
	self.cbXuanCards = make([]byte, PiceCount)       	// 玄家手牌
	self.cbDiCards = make([]byte, PiceCount)       	// 地家手牌
	self.cbHuangCards = make([]byte, PiceCount)       	// 黄家手牌
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

	self.gameState = SubGameSenceStart
	self.timeStamp = time.Now().Unix()

	self.roundNumber = strconv.FormatInt(self.timeStamp,10)
	if timer == nil {
		time.AfterFunc(freeTime*time.Second, self.onPlay)
	} else {
		timer.AfterFunc(freeTime*time.Second, self.onPlay)
	}
	log.Release("[牛牛:%v]游戏开始",self.roundNumber)
	manger.NotifyOthers(self.PlayerList, MainGameState, SubGameStateStart, nil)
}

func (self *CowcowGame)onPlay(){
	log.Release("[牛牛]允许下注")
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
//发牌(并未实际发送牌值)
func (self *CowcowGame) dispatchCard() {
	//25 张牌
	tableCards := RandCardList( PiceCount * AREA_MAX )

	cards := make([]byte,PiceCount * AREA_MAX )
	for k,v:=range tableCards{
		cards[k] = byte(v)
	}
	log.Debug("[牛牛-->开始发牌啦<---]  牌堆中取牌:%v ", GetCardsText(cards))
	copy(self.cbBankerCards,cards[INDEX_Banker:INDEX_Banker+PiceCount])
	copy(self.cbTianCards,cards[INDEX_Tian: INDEX_Tian+PiceCount])
	copy(self.cbXuanCards,cards[INDEX_Xuan: INDEX_Xuan+PiceCount])
	copy(self.cbDiCards,cards[INDEX_Di: INDEX_Di+PiceCount])
	copy(self.cbHuangCards,cards[INDEX_Huang: INDEX_Huang+PiceCount])


	log.Debug("[牛牛各家牌值]\n庄家:%v\n天:%v\t\t玄:%v\n地:%v\t\t黄:%v",
		GetCardsText(self.cbBankerCards),
		GetCardsText(self.cbTianCards), GetCardsText(self.cbXuanCards),
		GetCardsText(self.cbDiCards), GetCardsText(self.cbHuangCards))


	var banker = Pokers{}
	var tian = Pokers{}
	var xuan = Pokers{}
	var di = Pokers{}
	var huang = Pokers{}

	pokerList := CreatePoker(tableCards)

	banker.AddPokers(pokerList[0], pokerList[1], pokerList[2], pokerList[3], pokerList[4]).ArrangeByNumber()
	tian.AddPokers(pokerList[5], pokerList[6], pokerList[7], pokerList[8], pokerList[9]).ArrangeByNumber()
	xuan.AddPokers(pokerList[10], pokerList[11], pokerList[12], pokerList[13], pokerList[14]).ArrangeByNumber()
	di.AddPokers(pokerList[15], pokerList[16], pokerList[17], pokerList[18], pokerList[19]).ArrangeByNumber()
	huang.AddPokers(pokerList[20], pokerList[21], pokerList[22], pokerList[23], pokerList[24]).ArrangeByNumber()

	bankerType := CalcPoker(banker)
	tianType := CalcPoker(tian)
	xuanType := CalcPoker(xuan)
	diType := CalcPoker(di)
	huangType := CalcPoker(huang)


	if self.logic.BetWinInfoMap == nil {
		self.logic.BetWinInfoMap = &BetWinInfoMap{}
	}

	if self.logic.BetWinInfoMap.WinInfoMap == nil {
		self.logic.BetWinInfoMap.Init()
	}

	// 清空当前彩源的开奖信息
	self.logic.BetWinInfoMap.InitSourceMap(self.roundNumber)

	//各个区域的开奖结果
	self.logic.BetWinInfoMap.Set(self.roundNumber, AREA_Tian, self.logic.Compare(AREA_Tian, bankerType, tianType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, AREA_Xuan, self.logic.Compare(AREA_Xuan, bankerType, xuanType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, AREA_Di, self.logic.Compare(AREA_Di, bankerType, diType))
	self.logic.BetWinInfoMap.Set(self.roundNumber, AREA_Huang, self.logic.Compare(AREA_Huang, bankerType, huangType))

	//开奖结果
	self.overResult.CardValue = make([]byte,AREA_MAX)
	self.overResult.CardValue[AREA_Banker] = byte(bankerType.Type)
	self.overResult.CardValue[AREA_Tian] = byte(tianType.Type)
	self.overResult.CardValue[AREA_Xuan] = byte(xuanType.Type)
	self.overResult.CardValue[AREA_Di] = byte(diType.Type)
	self.overResult.CardValue[AREA_Huang] = byte(huangType.Type)

	log.Debug("牌值:Banker:%v ",self.overResult.CardValue)

}
//结算
func (self *CowcowGame) calculateScore() {

	self.overResult.BankerCard = self.cbBankerCards
	self.overResult.TianCard = self.cbTianCards
	self.overResult.XuanCard = self.cbXuanCards
	self.overResult.DiCard = self.cbDiCards
	self.overResult.HuangCard = self.cbHuangCards

	odds := make([]float64, AREA_MAX)
	self.overResult.AwardArea,odds = self.deduceWin()

	others := make([]uint64, 10)
	copy(others, self.PlayerList)

	playerAwardScroe := int64(0)
	for userID, betInfos := range self.playerBetInfo {
		//每一次的下注信息
		playerAwardScroe = 0
		for _, betInfo := range betInfos {
			log.Debug("玩家:%v,下注区域:%v 下注金额:%v", userID, betInfo.BetArea, betInfo.BetScore)
			//玩家奖金
			if Win == self.overResult.AwardArea[betInfo.BetArea] {
				playerAwardScroe += int64(odds[int(betInfo.BetArea)])*betInfo.BetScore
			}
		}
		//发送给指定玩家
		checkout := &protoMsg.GameCowcowCheckout{}
		checkout.Acquire = playerAwardScroe

		//写入数据库
		if 0 != playerAwardScroe {
			if money, ok := sqlHandle.DeductMoney(userID, -playerAwardScroe); ok {
				log.Debug("结算成功:%v 当前金币:%v!!", playerAwardScroe, money)
			} else {
				log.Debug("结算失败:%v 当前金币:%v!!", playerAwardScroe, money)
			}
		}

		manger.Get(userID).WillReceive(MainGameFrame, SubGameFrameCheckout, checkout)
		//userIDs = CopyInsert(userIDs, len(userIDs), userID).([]uint64)
		//for k, v := range others { //获取没下注玩家
		//	if v == userID {
		//		//lock.Lock()
		//		//defer lock.Unlock()
		//		others = append(others[:k], others[k+1:]...)
		//
		//	}
		//}
	}

	// 发给没下注玩家
	manger.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFrameOver, self.overResult)
}
func (self *CowcowGame) deduceWin() ([]byte,[]float64){
	pWinArea := make([]byte, AREA_MAX)
	pOdds := make([]float64, AREA_MAX)

	betTianInfo,_ := self.logic.BetWinInfoMap.Get(self.roundNumber, AREA_Tian)
	betXuanInfo,_ :=self.logic.BetWinInfoMap.Get(self.roundNumber, AREA_Xuan)
	betDiInfo,_ :=self.logic.BetWinInfoMap.Get(self.roundNumber, AREA_Di)
	betHuangInfo,_ :=self.logic.BetWinInfoMap.Get(self.roundNumber, AREA_Huang)



	if betTianInfo.IsWin{
		pWinArea[AREA_Tian] = Win
		pOdds[AREA_Tian] = betTianInfo.WinOdds
		log.Debug("天赢:%v ->%v",betTianInfo.WinOdds, betTianInfo.LoseOdds)
	}else{
		pOdds[AREA_Tian] = 0.0
	}

	if betXuanInfo.IsWin{
		pWinArea[AREA_Xuan] = Win
		pOdds[AREA_Xuan] = betXuanInfo.WinOdds
		log.Debug("玄赢:%v ->%v",betXuanInfo.WinOdds,betXuanInfo.LoseOdds)
	}else{
		pOdds[AREA_Xuan] = 0.0
	}


	if betDiInfo.IsWin{
		pWinArea[AREA_Di] = Win
		pOdds[AREA_Di] = betDiInfo.WinOdds
		log.Debug("地赢:%v ->%v",betDiInfo.WinOdds,betDiInfo.LoseOdds)
	}else{
		pOdds[AREA_Di] = 0.0
	}


	if betHuangInfo.IsWin{
		pWinArea[AREA_Huang] = Win
		pOdds[AREA_Huang] = betHuangInfo.WinOdds
		log.Debug("黄赢:%v ->%v",betHuangInfo.WinOdds,betHuangInfo.LoseOdds)
	}else{
		pOdds[AREA_Huang] = 0.0
	}


	//庄家区域
	return pWinArea,pOdds
}
