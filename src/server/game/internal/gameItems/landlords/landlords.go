package landlords

//斗地主

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	. "server/base"
	. "server/game/internal/gameItems" // 注意这里不能这样导入 "../../gameItems" 因为本地导入是根据gopath路径设定的
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
)

var timer *module.Skeleton = nil  //定时器
//var lock *sync.Mutex = &sync.Mutex{} //锁
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表

//定时器
const (
	freeTime  = 5
	tableSite = 3
)

//继承于GameItem
type LandlordGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳 20秒准备时间,到了20秒后被提出房间

	siteInfo map[uint8]uint64 // 座位上的玩家
	bankerID    uint64  //庄家ID
	bankerScore float32 //庄家积分
	readyCount  uint8   //准备的人数
}

//创建百家乐实例
func NewLandlord(level uint32, skeleton *module.Skeleton) *LandlordGame {
	log.Debug("------正在创建斗地主实例------level:%v", level)
	p := &LandlordGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------斗地主----------------------------------------------//
//初始化信息
func (self *LandlordGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  // 定时器的使用
	self.bStart = false               // 是否第一次启动
	self.Level = level                // 房间等级
	self.gameState = SubGameSenceFree // 场景状态
	self.bankerID = 0                 //庄家ID
	self.siteInfo = make(map[uint8]uint64)
	self.readyCount = 0

}

func (self *LandlordGame) Scene(args []interface{}) {
	userID := args[0].(uint64)
	level := args[1].(uint32)

	player := manger.Get(userID)
	if player == nil {
		log.Debug("[Error][斗地主] [未能查找到相关玩家] ID:%v", userID)
		return
	}


	// 获取玩家列表
	self.AddPlayer(player.UserID) //加入玩家列表
	for _, uid := range self.PlayerList {
		if playerItem := manger.Get(uid); nil != playerItem {
			var playerInfo protoMsg.PlayerInfo
			playerInfo.UserID = playerItem.UserID
			playerInfo.Name = playerItem.Name
			playerInfo.Age = playerItem.Age
			playerInfo.Gold = int64(sqlHandle.CheckMoney(playerItem.UserID))*100 //玩家积分
			playerInfo.VipLevel = playerItem.Level
			playerInfo.Sex = playerItem.Sex
			playerList.AllInfos = CopyInsert(playerList.AllInfos, len(playerList.AllInfos), &playerInfo).([]*protoMsg.PlayerInfo)

		} else {
			manger.DeletePlayerIndex(uid)
		}
	}


	log.Debug("[斗地主] [玩家列表新增] ID:%v", userID)
	senceInfo := &protoMsg.GameLandLordsEnter{}
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
	log.Debug("[斗地主场景]->玩家信息 ID:%v ", player.UserID)
}

//更新
func (self *LandlordGame) UpdateInfo(args []interface{}) { //更新玩家列表[目前]

	log.Debug("[斗地主]更新信息:%v-> %v\n", args[0].(uint32), args[1])
	flag := args[0].(uint32)
	userID := args[1].(uint64)
	player := manger.Get(userID)

	switch flag {
	case GameUpdateOut: //玩家离开 不再向该玩家广播消息[] 删除
		// 判断状态,如果玩家准备了,再进行退出
		if self.gameState == SubGameSenceFree && player.Sate == PlayerAgree && 0 < self.readyCount {
			self.readyCount--
			delete(self.siteInfo,self.readyCount)
		}

		self.DeletePlayer(userID)
		manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdatePlayerList: //更新玩家列表
		self.AddPlayer(userID)
		manger.NotifyOthers(self.PlayerList, MainGameUpdate, GameUpdatePlayerList, &playerList)
	case GameUpdateHost: //更新玩家叫分

	case GameUpdateSuperHost: //更新玩家明牌叫分
	case GameUpdateOffline: //更新玩家超级抢庄信息
	case GameUpdateReconnect: //更新重入
	case GameUpdateReady: //统计准备的玩家
		self.siteInfo[self.readyCount] = userID 		//保存座位信息
		self.readyCount++ //统计准备人数
		player.Sate = PlayerAgree

		manger.NotifyOthers(self.PlayerList,MainGameFrame,SubGameFrameReady, &protoMsg.GameReady{IsReady:true,UserID:userID})
		if tableSite == self.readyCount { //桌面上三个玩家都准备好了,才进行发牌
			self.Start(nil)
		}

		self.Start(nil)//测试用
		log.Debug("[斗地主]玩家准备就绪...")
	}
}

// 开始
func (self *LandlordGame) Start(args []interface{}) {
	//直接扣除金币
	log.Debug("游戏开始")

	tempCards := make([]byte, len(CardListData))
	copy(tempCards[0:], CardListData[0:])
	cards := Shuffle(tempCards)
	log.Debug("洗牌之后:%v", cards)

	// 取54-3张牌
	for site,userID := range self.siteInfo{
		player := manger.Get(userID)
		if player == nil {
			log.Debug("[Error][斗地主] [未能查找到相关玩家] ID:%v", userID)
			continue
		}
		if SiteCount <= site{
			continue
		}

		playerCard := Deal(cards[0:len(CardListData)-3], int(site), SiteCount)
		log.Debug("玩家的牌:%v", playerCard)

		// 排序其实可以交给客户端,以减少服务端运算压力
		sortCards := SortCardX(playerCard)
		log.Debug("排序之后:%v\n %v", sortCards, GetCardsText(sortCards))
		log.Debug("底牌:%v", GetCardsText(cards[0:3]))
		msg := &protoMsg.GameLandLordsBegins{}
		msg.CardsBottom = cards[0:3]
		msg.CardsHand = sortCards

		player.WillReceive(MainGameState,SubGameStateStart,msg)
	}



	//manger.NotifyOthers(self.PlayerList, MainGameState, SubGameStateStart, msg)
	//return
	//m := args[0].(*protoMsg.GameLandlordPlaying)

}

// 出牌
func (self *LandlordGame) Playing(args []interface{}) {
	//直接扣除金币
	log.Debug("出牌")
	//return
	//m := args[0].(*protoMsg.GameLandLordsOutcard)

}

// 结算
func (self *LandlordGame) Over(args []interface{}) {
	//直接扣除金币
	log.Debug("结算")
	self.readyCount = 0
	self.siteInfo = make(map[uint8]uint64)
	//return
	//m := args[0].(*protoMsg.GameLandlordPlaying)

}

// 操作
func (self *LandlordGame) SuperControl(args []interface{}) {
	//直接扣除金币
	log.Debug("操作")
	//return
	//m := args[0].(*protoMsg.GameLandlordPlaying)
}

// 游戏业务层
//重新初始化
func (self *LandlordGame) reset() {
	log.Release("[斗地主]扫地僧出来干活了...")

	//hostList = []uint64{}
	//
	//	//不要清理数据,因为数据清除之后,下一轮没法广播数据
	//	//playerCouriers = nil 【清内存】
	//	//for k,_:=range playerCouriers {//【数据清理】			//清理玩家
	//	delete(playerCouriers,k)
	//}
}

//-----------------------逻辑层---------------------------
//发牌
func (self *LandlordGame) dispatchCard() {

}
