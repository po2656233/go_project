package baccarat

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	. "server/base"
	. "server/manger"
	. "server/game/internal/gameItems"
	protoMsg "server/msg/go"
	_ "server/sql/mysql" //仅仅希望导入 包内的init函数
	"time"
)

var timer *module.Skeleton = nil  //定时器
var playerList protoMsg.UserList  //玩家列表

//var lock *sync.Mutex = &sync.Mutex{} //锁
//定时器
const (
	freeTime = 3
	betTime  = 10
	openTime = 7
)

//继承于GameItem
type BaccaratGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

	cbPlayerCards []byte //闲家扑克
	cbBankerCards []byte //庄家扑克
	cbCardCount   []byte //首次拿牌数

	playerBetInfo map[uint64][]protoMsg.GameBet //下注信息
	openAreas     [][]byte                      //开奖纪录
	overResult    *protoMsg.GameBaccaratOver

	hostList    []uint64 //申请列表(默认11个)
	bankerID    uint64   //庄家ID
	superHostID uint64   //超级抢庄ID
	bankerScore int64    //庄家积分
	keepTwice   uint8    //连续坐庄次数
}

//创建百家乐实例
func NewBaccarat(level uint32, skeleton *module.Skeleton) *BaccaratGame {
	log.Debug("------正在创建百家乐实例------level:%v", level)
	p := &BaccaratGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------百家乐----------------------------------------------//
//初始化信息
func (self *BaccaratGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  //定时器的使用
	self.bStart = false               //是否第一次启动
	self.Level = level                //房间等级
	self.gameState = SubGameSenceFree //场景状态

	self.cbPlayerCards = make([]byte, CardCountPlayer)       // 闲家手牌
	self.cbBankerCards = make([]byte, CardCountBanker)       // 庄家手牌
	self.cbCardCount = []byte{2, 2}                          // 首次拿牌张数
	self.playerBetInfo = make(map[uint64][]protoMsg.GameBet) // 玩家下注信息
	self.openAreas = make([][]byte, 10, 10)                  // 开奖纪录
	self.overResult = &protoMsg.GameBaccaratOver{}           // 结算结果

	self.hostList = []uint64{} //申请列表(默认11个)
	self.bankerID = 0          //庄家ID
	self.superHostID = 0       //超级抢庄ID
	self.keepTwice = 0         //连续抢庄次数
}

//获取玩家列表
//func (self *BaccaratGame) getPlayersName() []string {
//	var name []string
//	for _, pid := range self.PlayerList {
//		log.Debug("----玩家：%v", pid)
//		name = append(name, sqlhandle.CheckName(pid))
//	}
//	return name
//}

////////////////////////////////////////////////////////////////
//场景信息
func (self *BaccaratGame) Scene(args []interface{}) {
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

	GlobalSender.SendData(agent,MainGameSence,self.gameState, senceInfo)


	// 更新游戏列表
	var updateArgs []interface{}
	updateArgs = append(updateArgs, uint32(GameUpdatePlayerList),player.UserID)
	playerList.AllInfos = append(playerList.AllInfos, &playerInfo)
	self.UpdateInfo(updateArgs)
	log.Debug("[百家乐场景]->玩家信息 ID:%v 当前玩家列表:%v ", player.UserID, self.PlayerList)
}

//开始下注前定庄
func (self *BaccaratGame) Start(args []interface{}) {
	if !self.bStart {
		self.onStart()
		self.bStart = true
	}
}

//下注 直接扣除金币
func (self *BaccaratGame) Playing(args []interface{}) {
	//【消息】
	m := args[0].(*protoMsg.GameBet)
	//【传输对象】
	sender := args[1].(gate.Agent)
	log.Debug("[百家乐]BaccaratPlaying:->%v", m)

	userData := sender.UserData()
	if nil == userData{
		log.Debug("[百家乐]BaccaratPlaying:->%v 无效玩家", m)
		return
	}
	player:=userData.(*Player)
	//反馈下注情况
	betResult := &protoMsg.GameBetResult{}
	betResult.UserID = player.UserID
	betResult.BetArea = m.BetArea
	betResult.BetScore = m.BetScore

	if self.gameState != SubGameSencePlaying {
		betResult.State = *proto.Int32(1)
		betResult.Hints = *proto.String("过了下注时间")
		GlobalSender.SendData(sender,MainGameFrame, SubGameFrameBetResult,betResult)
		return
	}

	//数据库中扣除玩家金币[下注成功]
	if money, ok := GlobalSqlHandle.DeductMoney(player.UserID, m.BetScore); !ok {
		betResult.State = *proto.Int32(1)
		betResult.Hints = *proto.String("数据库里的钱不够")
		log.Debug("下注失败 玩家ID:%v 现有金币:%v 下注金币:%v", player.UserID, money, m.BetScore)
		GlobalSender.SendData(sender,MainGameFrame, SubGameFrameBetResult,betResult)
		return
	}

	//下注成功
	betResult.State = *proto.Int32(0)
	betResult.Hints = *proto.String("[百家乐]下注成功")

	//下注数目累加
	if areaBetInfos, ok := self.playerBetInfo[player.UserID]; ok {
		ok = false
		for index, betItem := range areaBetInfos {
			if betItem.BetArea == m.BetArea {
				self.playerBetInfo[player.UserID][index].BetScore = betItem.BetScore + m.BetScore
				log.Debug("玩家:%v 累加:->%v", player.UserID, areaBetInfos[index].BetScore)
				ok = true
				break
			}
		}
		if !ok {
			self.playerBetInfo[player.UserID] = CopyInsert(areaBetInfos, len(areaBetInfos), *m).([]protoMsg.GameBet)
		}
	} else {
		log.Debug("[百家乐]第一次:%v", m)
		self.playerBetInfo[player.UserID] = CopyInsert(self.playerBetInfo[player.UserID], len(self.playerBetInfo[player.UserID]), *m).([]protoMsg.GameBet)
	}

	//通知其他玩家
	GlobalSender.SendData(sender,MainGameFrame, SubGameFrameBetResult, betResult)
	GlobalSender.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFramePlaying, m)

	//通知所有人
	log.Debug("[百家乐]反馈下注信息:->%v", m)
}

//结算
func (self *BaccaratGame) Over(args []interface{}) {
	//【发牌】
	self.dispatchCard()
	//【计算结果并反馈】
	self.calculateScore()
}

//更新
func (self *BaccaratGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]

	log.Debug("[百家乐]更新信息:%v->*** %v\n", args[0].(uint32), args[1])
	flag := args[0].(uint32)
	userID := args[1].(uint64)
	switch flag {
	case GameUpdateOut: //玩家离开 不再向该玩家广播消息[] 删除
		self.DeletePlayer(userID)
		//
		for index, info := range playerList.AllInfos {
			if info.UserID == userID {
				log.Debug("正在从列表中剔除 玩家 %v", userID)
				playerList.AllInfos = append(playerList.AllInfos[:index], playerList.AllInfos[index+1:]...)
				break
			}
		}
		GlobalSender.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdatePlayerList: //更新玩家列表
		GlobalSender.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdateHost: //更新玩家抢庄信息
		self.host(args)
	case GameUpdateSuperHost: //更新玩家超级抢庄信息
		self.superHost(args)
	case GameUpdateOffline: //更新玩家超级抢庄信息
	case GameUpdateReconnect: //更新玩家超级抢庄信息

	}
}

//超级控制
func (self *BaccaratGame) SuperControl(args []interface{}) {

}

//定时器控制
//开始|下注|结算
func (self *BaccaratGame) onStart() {
	self.reset()
	self.gameState = SubGameSenceStart
	self.timeStamp = time.Now().Unix()
	self.playerBetInfo = make(map[uint64][]protoMsg.GameBet) // 清空玩家下注信息
	if timer == nil {
		time.AfterFunc(freeTime*time.Second, self.onPlay)
	} else {
		timer.AfterFunc(freeTime*time.Second, self.onPlay)
	}

	// 开始状态
	m := self.permitHost() //反馈定庄信息
	GlobalSender.NotifyOthers(self.PlayerList, MainGameState, SubGameStateStart, m)
	log.Release("[百家乐]游戏开始")
}

//[下注减法]
func (self *BaccaratGame) onPlay() {

	self.gameState = SubGameSencePlaying
	self.timeStamp = time.Now().Unix()
	if timer == nil {
		time.AfterFunc(betTime*time.Second, self.onOver)
	} else {
		timer.AfterFunc(betTime*time.Second, self.onOver)
	}

	// 下注状态
	GlobalSender.NotifyOthers(self.PlayerList, MainGameState, SubGameStatePlaying, nil)
}

//[结算加法]
func (self *BaccaratGame) onOver() {
	log.Release("[百家乐]开奖中...")
	self.gameState = SubGameSenceOver
	self.timeStamp = time.Now().Unix()
	if timer == nil {
		time.AfterFunc(openTime*time.Second, self.onStart)
	} else {
		timer.AfterFunc(openTime*time.Second, self.onStart)
	}

	// 玩家结算(框架消息)
	self.Over(nil)

	// 开奖状态
	GlobalSender.NotifyOthers(self.PlayerList, MainGameState, SubGameStateOver, nil)

}

//抢庄
func (self *BaccaratGame) host(args []interface{}) {
	//【消息】
	sender := args[1].(gate.Agent)
	host := args[2].(*protoMsg.GameHost)
	msg := &protoMsg.GameResult{
		Flag:   1,
		Reason: []byte("无效玩家"),
	}

	userData := sender.UserData()
	if nil == userData{
		sender.WriteMsg(msg)
		return
	}

	userID := userData.(*Player).UserID
	size := len(self.hostList)
	if self.bankerID == userID {
		msg.Flag = 1
		msg.Reason = []byte("已经是庄家了")
	} else if host.IsWant { //申请(已经是庄家)
		//列表上是否已经申请了
		hasID := false
		for _, pid := range self.hostList {
			if pid == userID {
				hasID = true
				break
			}
		}
		if hasID {
			msg.Flag = 1
			msg.Reason = []byte("[百家乐]已在申请列表")
		} else if size < 12 {
			self.hostList = CopyInsert(self.hostList, size, userID).([]uint64)
		} else {
			msg.Flag = 1
			msg.Reason = []byte("[百家乐]申请列表的人数已满")
		}
	} else { //取消
		for index, pid := range self.hostList {
			if userID == pid {
				self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
				break
			}
		}
		if size == len(self.hostList) {
			msg.Flag = 1
			msg.Reason = []byte("[百家乐]该玩家已不在列表中")
		}
	}

	log.Debug("[百家乐]有人来抢庄啦:%d 列表人数%d", userID, len(self.hostList))

	sender.WriteMsg(msg)

	if msg.Flag == 0 { //广播申请上庄成功的玩家
		msgAll := &protoMsg.GameBaccaratHost{
			UserID: userID,
			IsWant: host.IsWant,
		}
		GlobalSender.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFrameHost, msgAll)
	}
}

//超级抢庄
func (self *BaccaratGame) superHost(args []interface{}) {
	//【消息】
	sender := args[1].(gate.Agent)
	host := args[2].(*protoMsg.GameHost)
	msg := &protoMsg.GameResult{
		Flag:   0,
		Reason: []byte("success"),
	}

	userData := sender.UserData()
	if nil == userData{
		sender.WriteMsg(msg)
		return
	}
	userID := userData.(*Player).UserID
	log.Debug("[百家乐]有人要超级抢庄--->:%d", userID)
	if host.IsWant {
		if self.superHostID == 0 {
			self.superHostID = userID
			//广播给所有玩家
			msgAll := &protoMsg.GameBaccaratSuperHost{
				UserID: self.bankerID,
				IsWant: true,
			}
			//超级抢庄放申请列表首位
			self.hostList = CopyInsert(self.hostList, 0, userID).([]uint64)
			GlobalSender.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFrameSuperHost, msgAll)
		} else {
			msg.Flag = 1
			msg.Reason = []byte("faild")
		}
	}

	//反馈
	sender.WriteMsg(msg)
}

//重新初始化
func (self *BaccaratGame) reset() {
	log.Release("[百家乐]扫地僧出来干活了...")
	self.cbPlayerCards = make([]byte, CardCountPlayer)
	self.cbBankerCards = make([]byte, CardCountBanker)
	self.cbCardCount = []byte{2, 2}
	self.playerBetInfo = make(map[uint64][]protoMsg.GameBet) //玩家下注信息
	//hostList = []uint64{} 										//清理玩家

	//不要清理数据,因为数据清除之后,下一轮没法广播数据
	//playerCouriers = nil 【清内存】
	//for k,_:=range playerCouriers {//【数据清理】
	//	delete(playerCouriers,k)
	//}
}

//定庄(说明: 随时可申请上庄,申请列表一共11位。如果有超级抢庄,则插入列表首位。)
func (self *BaccaratGame) permitHost() *protoMsg.GameBaccaratHost {
	//校验是否满足庄家条件 [5000 < 金额] 不可连续坐庄15次
	tempList := self.hostList
	log.Debug("[百家乐]定庄.... 列表数据:%v", self.hostList)

	//befBankerId := self.bankerID 避免重复
	for index, pid := range tempList {
		player := GlobalPlayerManger.Get(pid)
		if player == nil {
			continue
		}
		if 2 == self.keepTwice && self.bankerID == player.UserID {
			log.Debug("[百家乐]不再连续坐庄：%v", player.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
			self.keepTwice = 0
		} else if player.Gold < 5000 {
			log.Debug("[百家乐]玩家%d 金币%lf 少于5000不能申请坐庄", player.UserID, player.Gold)
			self.hostList = append(self.hostList[:index], self.hostList[index+1:]...)
		}

	}

	//取第一个作为庄家
	if 0 < len(self.hostList) {
		if self.bankerID == self.hostList[0] {
			log.Debug("[百家乐]连续坐庄次数:%d", self.keepTwice)
			self.keepTwice++
		} else {
			self.keepTwice = 0
		}
		self.bankerID = self.hostList[0]

		if banker := GlobalPlayerManger.Get(self.bankerID); banker != nil {
			self.bankerScore = banker.Gold
		}

		log.Debug("[百家乐]确定庄家:%d", self.bankerID)
	} else {
		self.bankerID = 0 //系统坐庄
		self.bankerScore = 1000000
		log.Debug("[百家乐]系统坐庄")
	}
	//完成定庄后,初始化超级抢庄ID
	self.superHostID = 0
	msg := &protoMsg.GameBaccaratHost{
		UserID: self.bankerID,
		IsWant: true,
	}
	log.Debug("[百家乐]广播上庄")
	return msg
}

//-----------------------逻辑层---------------------------
//发牌
func (self *BaccaratGame) dispatchCard() {
	//6张牌
	tableCards := RandCardList(CardCountPlayer + CardCountBanker)

	//首次发牌
	log.Debug("[百家乐-->开始发牌啦<---]  牌堆中取牌:%v ", GetCardsText(tableCards))
	copy(self.cbPlayerCards[:CardCountPlayer], tableCards[0:self.cbCardCount[INDEX_PLAYER]])
	copy(self.cbBankerCards[:CardCountBanker], tableCards[CardCountBanker:CardCountBanker+self.cbCardCount[INDEX_BANKER]])
	log.Debug("[百家乐 第一次发牌]\t\t\t闲家:%v \t庄家:%v", GetCardsText(self.cbPlayerCards), GetCardsText(self.cbBankerCards))
	//计算点数
	cbBankerCount := GetCardListPip(self.cbBankerCards)
	cbPlayerTwoCardCount := GetCardListPip(self.cbPlayerCards)

	//闲家补牌
	var cbPlayerThirdCardValue byte = 0 //第三张牌点数
	if cbPlayerTwoCardCount <= 5 && cbBankerCount < 8 {
		//计算点数
		self.cbCardCount[INDEX_PLAYER]++
		cbPlayerThirdCardValue = GetCardPip(tableCards[CardCountPlayer-1])
		self.cbPlayerCards[CardCountPlayer-1] = tableCards[CardCountPlayer-1]
	}

	//庄家补牌
	if cbPlayerTwoCardCount < 8 && cbBankerCount < 8 {
		switch cbBankerCount {
		case 0:
			self.cbCardCount[INDEX_BANKER]++
		case 1:
			self.cbCardCount[INDEX_BANKER]++
		case 2:
			self.cbCardCount[INDEX_BANKER]++
		case 3:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 8) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 4:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 1 && cbPlayerThirdCardValue != 8 && cbPlayerThirdCardValue != 9 && cbPlayerThirdCardValue != 0) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 5:
			if (self.cbCardCount[INDEX_PLAYER] == 3 && cbPlayerThirdCardValue != 1 && cbPlayerThirdCardValue != 2 && cbPlayerThirdCardValue != 3 && cbPlayerThirdCardValue != 8 && cbPlayerThirdCardValue != 9 && cbPlayerThirdCardValue != 0) || self.cbCardCount[INDEX_PLAYER] == 2 {
				self.cbCardCount[INDEX_BANKER]++
			}
		case 6:
			if self.cbCardCount[INDEX_PLAYER] == 3 && (cbPlayerThirdCardValue == 6 || cbPlayerThirdCardValue == 7) {
				self.cbCardCount[INDEX_BANKER]++
			}
		}
	}
	if self.cbCardCount[INDEX_BANKER] == 3 {
		self.cbBankerCards[CardCountBanker-1] = tableCards[CardCountBanker+CardCountPlayer-1]
	}

	log.Debug("[百家乐 补牌结果]\t\t\t\t闲家:%v  庄家:%v", GetCardsText(self.cbPlayerCards), GetCardsText(self.cbBankerCards))
}

//计算得分[只做加法]
func (self *BaccaratGame) calculateScore() {

	//反馈下注情况
	self.overResult.PlayerCard = self.cbPlayerCards
	self.overResult.BankerCard = self.cbBankerCards
	self.overResult.AwardArea = self.deduceWin()

	var  playerAwardScroe int64
	for userID, betInfos := range self.playerBetInfo {
		//每一次的下注信息
		playerAwardScroe = 0
		for _, betInfo := range betInfos {
			log.Debug("玩家:%v,下注区域:%v 下注金额:%v", userID, betInfo.BetArea, betInfo.BetScore)
			//玩家奖金
			if Win == self.overResult.AwardArea[betInfo.BetArea] {
				playerAwardScroe += self.bonusArea(betInfo.BetArea, betInfo.BetScore)
			}
		}
		//发送给指定玩家
		checkout := &protoMsg.GameBaccaratCheckout{}
		checkout.Acquire = playerAwardScroe

		//写入数据库(todo 写入注单表)
		if 0 != playerAwardScroe {
			if money, ok := GlobalSqlHandle.DeductMoney(userID, -playerAwardScroe); ok {
				GlobalSender.SendTo(userID, MainGameFrame, SubGameFrameCheckout, checkout)
				log.Debug("结算成功:%v 当前金币:%v!!", playerAwardScroe, money)
			} else {
				log.Debug("结算失败:%v 当前金币:%v!!", playerAwardScroe, money)
			}
		}
	}

	// 开奖状态
	GlobalSender.NotifyOthers(self.PlayerList, MainGameFrame, SubGameFrameOver, self.overResult)
	log.Release("[百家乐]结算中...")
}

//开奖区域
func (self *BaccaratGame) deduceWin() []byte {
	pWinArea := make([]byte, AREA_MAX)
	//计算牌点
	cbPlayerCount := GetCardListPip(self.cbPlayerCards)
	cbBankerCount := GetCardListPip(self.cbBankerCards)

	//胜利区域--------------------------
	//平
	if cbPlayerCount == cbBankerCount {
		pWinArea[AREA_PING] = Win
		log.Debug("平")
		// 同平点
		if self.cbCardCount[INDEX_PLAYER] == self.cbCardCount[INDEX_BANKER] {
			var wCardIndex byte = 0
			for wCardIndex = 0; wCardIndex < self.cbCardCount[INDEX_PLAYER]; wCardIndex++ {
				cbBankerValue := GetCardValue(self.cbBankerCards[wCardIndex])
				cbPlayerValue := GetCardValue(self.cbPlayerCards[wCardIndex])
				if cbBankerValue != cbPlayerValue {
					break
				}
			}

			if wCardIndex == self.cbCardCount[INDEX_PLAYER] {
				pWinArea[AREA_TONG_DUI] = Win
				log.Debug("同点平")
			}
		}
	} else if cbPlayerCount < cbBankerCount { //庄
		pWinArea[AREA_ZHUANG] = Win
		log.Debug("庄")
		//天王判断
		if cbBankerCount == 8 || cbBankerCount == 9 {
			pWinArea[AREA_ZHUANG_TIAN] = Win
			log.Debug("庄天王")
		}
	} else { //闲
		pWinArea[AREA_XIAN] = Win
		log.Debug("闲")
		//天王判断
		if cbPlayerCount == 8 || cbPlayerCount == 9 {
			pWinArea[AREA_XIAN_TIAN] = Win
			log.Debug("闲天王")
		}
	}

	//对子判断(前两张牌比较)
	if GetCardValue(self.cbPlayerCards[0]) == GetCardValue(self.cbPlayerCards[1]) ||
		(0 != self.cbPlayerCards[2] && (GetCardValue(self.cbPlayerCards[0]) == GetCardValue(self.cbPlayerCards[2]) ||
			GetCardValue(self.cbPlayerCards[1]) == GetCardValue(self.cbPlayerCards[2]))) {
		pWinArea[AREA_XIAN_DUI] = Win
		log.Debug("闲对子")
	}

	if GetCardValue(self.cbBankerCards[0]) == GetCardValue(self.cbBankerCards[1]) ||
		(0 != self.cbBankerCards[2] && (GetCardValue(self.cbBankerCards[0]) == GetCardValue(self.cbBankerCards[2]) ||
			GetCardValue(self.cbBankerCards[1]) == GetCardValue(self.cbBankerCards[2]))) {
		pWinArea[AREA_ZHUANG_DUI] = Win
		log.Debug("庄对子")
	}
	return pWinArea
}

//区域赔额
func (self *BaccaratGame) bonusArea(area int32, betScore int64) int64 {
	multiple := int64(0)
	switch area {
	case AREA_XIAN:
		multiple = MULTIPLE_XIAN
	case AREA_PING:
		multiple = MULTIPLE_PING
	case AREA_ZHUANG:
		multiple = MULTIPLE_ZHUANG
	case AREA_XIAN_TIAN:
		multiple = MULTIPLE_XIAN_TIAN
	case AREA_ZHUANG_TIAN:
		multiple = MULTIPLE_ZHUANG_TIAN
	case AREA_TONG_DUI:
		multiple = MULTIPLE_TONG_DIAN
	case AREA_XIAN_DUI:
		multiple = MULTIPLE_XIAN_PING
	case AREA_ZHUANG_DUI:
		multiple = MULTIPLE_ZHUANG_PING
	default:
		multiple = int64(0)
	}
	return betScore * multiple
}
