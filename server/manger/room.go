package manger

import (
	"sync"
	"time"
	"strconv"
	"github.com/name5566/leaf/log"
	. "server/base"
)

//----------游戏----------------------
//游戏
type GameItem struct {
	KindID     uint32       //游戏标识 0表示无效
	Level      uint32       //游戏类别 0:普通 1:中级 2:高级 3:VIP 4:冠军
	PlayerList []uint64     //玩家列表
	Instance   IGameOperate //游戏实例
}
//子游戏接口
type IGameOperate interface {
	Scene(args []interface{})        //场景
	Start(args []interface{})        //开始
	Playing(args []interface{})      //游戏
	Over(args []interface{})         //结算
	UpdateInfo(args []interface{})   //更新信息
	SuperControl(args []interface{}) //超级控制 可在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)
}

//资源
type Source struct {
	GameSet []*GameItem //子游戏集合
}
//----------房间----------------------
//房间信息
type RoomInfo struct {
	ID       uint32   //房间ID 0表示无效
	Key      string   //房间钥匙
	State    uint8    //房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
	Things   *Source  //资源 注:同一个房间内不出现两个一样kindID和level的房间
	UserList []uint64 //房间用户列表
}


//房间接口
type IRoomMange interface {
	Create(roomID uint32) (*RoomInfo,bool)             //创建房间[房间ID和钥匙配对]
	Check(roomID uint32) (info *RoomInfo, isExit bool) //查找房间
	Delete(roomID uint32, key string) bool             //删除房间[房间ID和钥匙配对成功后,才能删除]
	Open(roomID uint32, key string) (*RoomInfo, bool)  //开启房间
	Close(roomID uint32, key string) bool              //关闭房间
	Clear(roomID uint32, key string) bool              //清理房间
}

//管理房间
type RoomManger struct {
	sync.Map
	//rooms map[uint64]*RoomInfo
}

///////////////房间和子游戏管理////////////////////////
var lock sync.Mutex

//由于新增平台，则这里的管理不应该再是单例模式
func CreateRoomManger() *RoomManger {
	return &RoomManger{
		sync.Map{},
	}
}

//测试
func (self *RoomManger) PrintfAll() { //打印当前所房间
	log.Debug("所有房间号:->")
	self.Range(func(index, value interface{}) bool {
		log.Debug("index:%v, 房间号码:%v", index, value)
		return true
	})
}

func (self *RoomManger) Create(roomID uint32) (*RoomInfo, bool) { //新增
	info := &RoomInfo{}
	bRet := true
	if v, ok := self.Load(roomID); !ok {
		log.Debug("创建房间:%v", roomID)
		timeUnix := time.Now().Unix()
		info.ID = roomID
		info.Key = strconv.FormatInt(timeUnix, 10) + "用户密钥" + strconv.Itoa(int(roomID))
		info.Things = nil

		self.Store(info.ID, info)
	} else {
		log.Debug("房间:%v 已经存在", roomID)
		info = v.(*RoomInfo)
		bRet = false
	}
	return info, bRet
}

func (self *RoomManger) DeleteRoom(roomID uint32, key string) bool { //删除
	bRet := false
	self.Range(func(index, value interface{}) bool {
		if roomID == index.(uint32) && key == value.(*RoomInfo).Key {
			bRet = true
			//for _,gameHandle :=range value.(*RoomInfo).Things.GameSet{
			//}

			self.Delete(index)
			return false
		}
		return true
	})
	return bRet
}

func (self *RoomManger) Check(roomID uint32) (*RoomInfo, bool) { //查找

	if value, ok := self.Load(roomID); ok {
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
	lock.Lock()
	defer lock.Unlock()


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
