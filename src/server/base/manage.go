package base

//不要在这里包含其他server目录的模块 除了msg/go
import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	protoMsg "server/msg/go"
	"strconv"
	"sync"
	"time"
)

//---[对象]

/*玩家Courier*/
type Player struct {
	UserID  uint64 		//ID
	Name    string 		//名字(角色)
	Age     int32 		//年龄
	Sex     int32		//性别
	Level   int32		//游戏级别(1000+ VIP级别)
	Account string      //账号(手机号码/邮箱/真名)
	Money   int64       //金币(与真实金币 扩大100倍)
	Agent   gate.Agent  //网络

	Sate    byte   		// 状态 0:旁观 1:坐下 2:同意  3:站起
	RoomNum uint32 		// 房间号 0:无效
	GameID  uint32 		// 所在游戏ID 0:无效
}

// 玩家管理
type PlayerManger struct {
	players sync.Map
	//players map[uint64]*Player
}

//游戏
type GameItem struct {
	KindID     uint32       //游戏标识 0表示无效
	Level      uint32       //游戏类别 0:普通 1:中级 2:高级 3:VIP 4:冠军
	PlayerList []uint64     //玩家列表
	Instance   IGameOperate //游戏实例
}

//资源
type Source struct {
	GameSet []*GameItem //子游戏集合
}

//房间信息
type RoomInfo struct {
	ID       uint32   //房间ID 0表示无效
	Key      string   //房间钥匙
	State    uint8    //房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
	Things   *Source  //资源 注:同一个房间内不出现两个一样kindID和level的房间
	UserList []uint64 //房间用户列表
}

type RoomManger struct {
	rooms sync.Map
	//rooms map[uint64]*RoomInfo
}

//

//----[接口]
//网络发包
//type IAsynNetwork interface {
//	SendPacket(*protoMsg.PacketData)    		// 异步发送的包
//}

//子游戏接口
type IGameOperate interface {
	Scene(args []interface{})        //场景
	Start(args []interface{})        //开始
	Playing(args []interface{})      //游戏
	Over(args []interface{})         //结算
	UpdateInfo(args []interface{})   //更新信息
	SuperControl(args []interface{}) //超级控制 可在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)

}

//用户行为
type IUserBehavior interface {
	Enter(args []interface{})     //入场
	Out(args []interface{})       //出场
	Offline(args []interface{})   //离线
	Reconnect(args []interface{}) //重入
	Ready(args []interface{})     //准备
	Host(args []interface{})      //抢庄/地主叫分
	SuperHost(args []interface{}) //超级抢庄
}

//房间管理接口
type IRoomMange interface {
	Create(roomID uint32) (*RoomInfo,bool)             //创建房间[房间ID和钥匙配对]
	Check(roomID uint32) (info *RoomInfo, isExit bool) //查找房间
	Delete(roomID uint32, key string) bool             //删除房间[房间ID和钥匙配对成功后,才能删除]
	Open(roomID uint32, key string) (*RoomInfo, bool)  //开启房间
	Close(roomID uint32, key string) bool              //关闭房间
	Clear(roomID uint32, key string) bool              //清理房间
}

//------实现-----------------
var manger *PlayerManger = nil
var once sync.Once
var lock *sync.Mutex = &sync.Mutex{}

//玩家管理对象(单例模式)
func GetPlayerManger() *PlayerManger {
	once.Do(func() {
		manger = &PlayerManger{}
		//manger.players = make(map[uint64]*Player)
	})
	return manger
}

//获取指定玩家[根据索引,即userID]
func (self *PlayerManger) Get(userID uint64) *Player {
	value, ok := self.players.Load(userID)
	if ok {
		return value.(*Player)
	}
	return nil
}

/*获取指定玩家[根据代理]*/
func (self *PlayerManger) Get_1(agent gate.Agent) *Player {
	var player *Player = nil
	self.players.Range(func(key, value interface{}) bool {
		if value.(*Player).Agent== agent {
			player = value.(*Player)
			return false
		}
		return true
	})
	return player
}

//玩家是否存在
func (self *PlayerManger) Exist(play *Player) bool {
	player := &Player{}
	isHas := false
	self.players.Range(func(key, value interface{}) bool {
		if value.(*Player).Agent.LocalAddr() == play.Agent.LocalAddr() || key.(uint64) == play.UserID {
			player = value.(*Player)
			isHas = true
			return false
		}
		return true
	})
	return isHas
}

//添加玩家
func (self *PlayerManger) AppandPlayer(play *Player) bool {
	if _, ok := self.players.Load(play.UserID); !ok {
		log.Debug("新增一个玩家ID:%v", play.UserID)
		self.players.Store(play.UserID, play)
		return true
	} else {
		log.Debug("玩家ID:%v 已經存在", play.UserID)
		return false
	}
}

//按索引删除玩家
func (self *PlayerManger) DeletePlayerIndex(i uint64) {
	self.players.Delete(i)
}

//删除玩家
func (self *PlayerManger) DeletePlayer(play *Player) {
	//lock.Lock()
	//defer lock.Unlock()
	index := uint64(0)
	isOK := false
	self.players.Range(func(key, value interface{}) bool {
		if value.(*Player).Agent == play.Agent || key.(uint64) == play.UserID {
			log.Debug("找到要删除的玩家:%v ", play.UserID)
			index = key.(uint64)
			isOK = true
			value = nil
			return false
		}
		return true
	})
	self.players.Delete(index)
}

//通知所有玩家
func (self *PlayerManger) NotifyAll(mainID, subID uint32, msg proto.Message) {
	self.players.Range(func(key, value interface{}) bool {
		player := value.(*Player)
		if nil == player.Agent {
			log.Debug("无效玩家:%v", key)
			return true
		}
		log.Debug("通知玩家：%v %v", key, player.Agent.LocalAddr())

		//指令+数据  包体内容
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}
		fmt.Println("发送数据(But):", len(data), data)
		player.Agent.NotifyMsg(packet)
		return true
	})
}

//cluster全网广播
func (self *PlayerManger) NotifyWorld(mainID, subID uint32, msg proto.Message) {
	self.players.Range(func(key, value interface{}) bool {
		player := value.(*Player)
		if nil == player.Agent {
			log.Debug("无效玩家:%v", key)
			return true
		}
		log.Debug("通知玩家：%v %v", key, player.Agent.LocalAddr())

		//指令+数据  包体内容
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}
		fmt.Println("发送数据(But):", len(data), data)
		player.Agent.NotifyMsg(packet)
		return true
	})
}




//通知除指定玩家外的玩家们
func (self *PlayerManger) NotifyButOthers(userIDs []uint64, mainID, subID uint32, msg proto.Message) {

	self.players.Range(func(key, value interface{}) bool {
		player := value.(*Player)
		if nil == player.Agent {
			log.Debug("无法通知:%v", key)
			return true
		}

		//不通知该部分玩家
		for _, uid := range userIDs {
			if uid == player.UserID {
				return true
			}
		}

		log.Debug("通知玩家(But)：%v %v", key, player.Agent.LocalAddr())
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}
		fmt.Println("发送数据(But):", len(data), data)
		player.Agent.NotifyMsg(packet)

		return true
	})
}

func (self *PlayerManger) NotifyOthers(userIDs []uint64, mainID, subID uint32, msg proto.Message) {
	self.players.Range(func(key, value interface{}) bool {
		player := value.(*Player)
		if nil == player.Agent {
			log.Debug("无法通知:%v", key)
			return true
		}

		//仅通知该部分玩家
		for _, uid := range userIDs {
			if uid == player.UserID {
				log.Debug("通知玩家(Others)：%v %v", key, player.Agent.LocalAddr())
				//指令+数据  包体内容
				data, _ := proto.Marshal(msg)
				packet := &protoMsg.PacketData{
					MainID:    mainID,
					SubID:     subID,
					TransData: data,
				}
				//fmt.Println("发送数据(Others):", len(data), data)
				player.Agent.NotifyMsg(packet)
				return true
			}
		}
		return true
	})
}

//////////////////////////////////////////////////

//数据发送
func (self *Player) WillReceive(mainID, subID uint32, message proto.Message) error {
	if nil == self || nil == self.Agent {
		fmt.Println("Not Aget Or Courier")
		return errors.New("Not Aget Or courier")
	}
	//指令+数据  包体内容
	data, error := proto.Marshal(message)
	packet := &protoMsg.PacketData{
		MainID:    mainID,
		SubID:     subID,
		TransData: data,
	}
	data, _ = proto.Marshal(packet)
	fmt.Println("发送数据:", len(data), data)
	self.Agent.WriteMsg(packet)
	return error
}


//结果反馈
func (self *Player) Feedback(mainID, subID, flag uint32, reason string) {
	enterResult := &protoMsg.ResResult{State: flag, Hints: reason}
	log.Error(reason)
	self.WillReceive(mainID, subID, enterResult)
}

//进入场景(限制条件由外部转入)
func (self *Player) Enter(args []interface{}) { //入场
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	if room, _ := GetRoomManger().Check(self.RoomNum); room != nil && self.RoomNum != 0 { //[0-0
		if handle, ok := room.GetGameHandle(gameInfo.KindID, gameInfo.Level); ok { //[1-0
			//-> 校验是否满足进入游戏的条件
			if gameInfo.EnterScore < uint32(self.Money) && gameInfo.LessScore < uint32(self.Money) { //[2-0
				var sceneArgs []interface{}
				sceneArgs = append(sceneArgs, self.UserID, gameInfo.Level)
				handle.Scene(sceneArgs) // 【进入-> 游戏场景】
			} else { //2]
				self.Feedback(MainLogin, SubEnterGameResult, FAILD, string("error:not allow to enter!"))
			}
			log.Debug("[进入游戏]---游戏准入条件(LessScore:%v EnterScore:%v)--->user money:%v", gameInfo.LessScore, gameInfo.EnterScore, self.Money)
		} else { //1]
			self.Feedback(MainLogin, SubEnterGameResult, FAILD, string("error:not game handle!"))
		}
	} else { //0]
		self.Feedback(MainLogin, SubEnterGameResult, FAILD, string("error:not room info!"))
	}
}

func (self *Player) Out(args []interface{}) { //出场
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateOut), self.UserID, nil)
}
func (self *Player) Offline(args []interface{}) { //离线
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateOffline), self.UserID, nil)
}
func (self *Player) Reconnect(args []interface{}) { //重入
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateReconnect), self.UserID, nil)
}
func (self *Player) Ready(args []interface{}) { //准备
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateReady), self.UserID, nil)

}
func (self *Player) Host(args []interface{}) { //抢庄
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	host := args[1].(*protoMsg.GameHost)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateHost), self.UserID, host)
}

func (self *Player) SuperHost(args []interface{}) { //超级抢庄
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	superHost := args[1].(*protoMsg.GameSuperHost)
	self.updateGameInfo(gameInfo.KindID, gameInfo.Level, uint32(GameUpdateSuperHost), self.UserID, superHost)
}

func (self *Player) updateGameInfo(kindID, level, flag uint32, userID uint64, message proto.Message) bool {
	if room, _ := GetRoomManger().Check(self.RoomNum); room != nil && self.RoomNum != 0 { //[0-0
		if handle, ok := room.GetGameHandle(kindID, level); ok { //[1-0
			var updateArgs []interface{}
			updateArgs = append(updateArgs, flag, userID, message)
			handle.UpdateInfo(updateArgs)
			return true
		}
	}
	log.Debug("error:GameUpdate->:%v userID:%v", flag, userID)
	return false
}

///////////////房间和子游戏管理////////////////////////
var roomManger *RoomManger = nil
var onceRoom sync.Once

func GetRoomManger() *RoomManger {
	onceRoom.Do(func() {
		roomManger = &RoomManger{
			rooms: sync.Map{},
		}
		//manger.players = make(map[uint64]*Player)
	})
	return roomManger
}

//测试
func (self *RoomManger) PrintfAll() { //打印当前所房间
	log.Debug("所有房间号:->")
	self.rooms.Range(func(index, value interface{}) bool {
		log.Debug("index:%v, 房间号码:%v", index, value)
		return true
	})
}

func (self *RoomManger) Create(roomID uint32) (*RoomInfo, bool) { //新增
	info := &RoomInfo{}
	bRet := true
	if v, ok := self.rooms.Load(roomID); !ok {
		log.Debug("创建房间:%v", roomID)
		timeUnix := time.Now().Unix()
		info.ID = roomID
		info.Key = strconv.FormatInt(timeUnix, 10) + "用户密钥" + strconv.Itoa(int(roomID))
		info.Things = nil

		self.rooms.Store(info.ID, info)
	} else {
		log.Debug("房间:%v 已经存在", roomID)
		info = v.(*RoomInfo)
		bRet = false
	}
	return info, bRet
}

func (self *RoomManger) Delete(roomID uint32, key string) bool { //删除
	bRet := false
	self.rooms.Range(func(index, value interface{}) bool {
		if roomID == index.(uint32) && key == value.(*RoomInfo).Key {
			bRet = true
			//for _,gameHandle :=range value.(*RoomInfo).Things.GameSet{
			//}

			self.rooms.Delete(index)
			return false
		}
		return true
	})
	return bRet
}

func (self *RoomManger) Check(roomID uint32) (*RoomInfo, bool) { //查找

	if value, ok := self.rooms.Load(roomID); ok {
		return value.(*RoomInfo), ok
	} else {
		return nil, ok
	}

}
func (self *RoomManger) Open(roomID uint32, key string) (*RoomInfo, bool) { //开启
	lock.Lock()
	defer lock.Unlock()
	info := &RoomInfo{}
	return info, true
}
func (self *RoomManger) Close(roomID uint32, key string) bool { //关闭
	lock.Lock()
	defer lock.Unlock()

	return true
}
func (self *RoomManger) Clear(roomID uint32, key string) bool { //清理房间

	return true
}

//--------------------------游戏实例----------------------------
func (self *RoomInfo) AddSource(game *GameItem) (IGameOperate, bool) {
	if item, ok := self.CheckGame(game.KindID, game.Level); !ok {
		if nil == self.Things {
			self.Things = &Source{}
		}
		self.Things.GameSet = CopyInsert(self.Things.GameSet, len(self.Things.GameSet), game).([]*GameItem)
		game.Instance.Start(nil)
		return game.Instance, true
	} else {
		return item.Instance, false
	}
}
func (self *RoomInfo) GetGameHandle(kindID, level uint32) (IGameOperate, bool) {
	//self.Things
	if item, ok := self.CheckGame(kindID, level); ok {
		return item.Instance, true
	} else {
		return nil, false
	}
}

func (self *RoomInfo) CheckGame(kindID, level uint32) (*GameItem, bool) {
	if nil == self.Things {
		return nil, false
	}
	for _, v := range self.Things.GameSet {
		if v.KindID == kindID && v.Level == level {
			return v, true
		}
	}
	return nil, false
}

//更新玩家列表[目前]
func (self *RoomInfo) UpdateInfo(args []interface{}) bool {
	return true
}

//加入玩家列表
func (self *GameItem) AddPlayer(userID uint64) bool {
	lock.Lock()
	defer lock.Unlock()
	for _, pid := range self.PlayerList {
		if pid == userID {
			log.Debug("无法添加玩家到列表%v", pid)
			return false
		}
	}
	self.PlayerList = CopyInsert(self.PlayerList, len(self.PlayerList), userID).([]uint64)
	return true
}

//踢出玩家列表
func (self *GameItem) DeletePlayer(userID uint64) {
	lock.Lock()
	defer lock.Unlock()
	for index, pid := range self.PlayerList {
		if userID == pid {
			log.Debug("剔除玩家:%v", userID)
			self.PlayerList = append(self.PlayerList[:index], self.PlayerList[index+1:]...)
			break
		}
	}
}

///////////////定时器管理////////////////////////
//定时器
var onceTimer sync.Once
var timeManager *TimerManager

type Timer struct {
	key      string
	cb       func()
	interval int
	loop     bool
}

type TimerManager struct {
	timerMap map[interface{}][]*Timer
	skeleton module.Skeleton
}

func GetTimerManager() *TimerManager {
	onceTimer.Do(func() {
		timeManager = &TimerManager{
			timerMap: make(map[interface{}][]*Timer),
		}
	})
	return timeManager
}

func (m *TimerManager) addTimer(obj interface{}, key string, cb func(), interval int, loop bool) {
	m.skeleton.AfterFunc(time.Millisecond*time.Duration(interval), func() {
		if !m.timerVaild(obj, key) {
			return
		}
		if !loop {
			m.RmvTimer(obj, key)
		}

		cb()
		if loop {
			m.addTimer(obj, key, cb, interval, loop)
		}
	})
}

func (m *TimerManager) getTimerIndex(l []*Timer, key string) int {
	for i, timer := range l {
		if timer.key == key {
			return i
		}
	}
	return -1
}

func (m *TimerManager) timerVaild(obj interface{}, key string) bool {
	l, ok := m.timerMap[obj]
	if !ok {
		return false
	}

	return m.getTimerIndex(l, key) >= 0
}

func (m *TimerManager) AddTimer(obj interface{}, key string, interval int, loop bool, cb func()) {
	l, _ := m.timerMap[obj]
	if m.getTimerIndex(l, key) >= 0 {
		log.Error("add repeated timer:", key)
		return
	}

	l = append(l, &Timer{key, cb, interval, loop})
	m.timerMap[obj] = l
	m.addTimer(obj, key, cb, interval, loop)
}

func (m *TimerManager) RmvTimer(obj interface{}, key string) {
	l, ok := m.timerMap[obj]
	if !ok {
		return
	}

	index := m.getTimerIndex(l, key)
	if index >= 0 {
		l = append(l[:index], l[index+1:]...)
	}
	m.timerMap[obj] = l
}

func (m *TimerManager) RmvAllTimer(obj interface{}) {
	delete(m.timerMap, obj)
}
