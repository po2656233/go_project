package manger

import (
	"sync"
	"github.com/name5566/leaf/log"
	protoMsg "server/msg/go"
	. "server/base"
	"github.com/golang/protobuf/proto"
	"leaf/gate"
)

//玩家属性
type Player struct {
	UserID     uint64 // ID
	Name       string // 名字(角色)
	Age        int32  // 年龄
	Sex        int32  // 性别
	Level      int32  // 游戏级别(1000+ VIP级别)
	Account    string // 账号(手机号码/邮箱/真名)
	Gold       int64  // 金币(与真实金币 扩大100倍)
	Sate       byte   // 状态 0:旁观 1:坐下 2:同意  3:站起
	PlatformID  uint32 // 平台ID 0:无效
	RoomNum     uint32 // 房间号 0:无效
	GameID		uint32 // 所在游戏ID 0:无效
	Game  		*GameItem //所在游戏
}

//玩家行为
type IUserBehavior interface {
	Enter(args []interface{})     //入场
	Out(args []interface{})       //出场
	Offline(args []interface{})   //离线
	Reconnect(args []interface{}) //重入
	Ready(args []interface{})     //准备
	Host(args []interface{})      //抢庄/地主叫分
	SuperHost(args []interface{}) //超级抢庄
}

// 管理玩家
type PlayerManger struct {
	sync.Map
} // == players map[uint64]*Player

//------------------------管理接口--------------------------------------------------//
var manger *PlayerManger = nil
var once sync.Once

//玩家管理对象(单例模式)//manger.players = make(map[uint64]*Player)
func GetPlayerManger() *PlayerManger {
	once.Do(func() {
		manger = &PlayerManger{
			sync.Map{},
		}
	})
	return manger
}

//添加玩家
func (self *PlayerManger) Append(play *Player) bool {
	if _, ok := self.Load(play.UserID); !ok {
		log.Debug("新增一个玩家ID:%v", play.UserID)
		self.Store(play.UserID, play)
		return true
	} else {
		log.Debug("玩家ID:%v 已經存在", play.UserID)
		return false
	}
}

//获取指定玩家[根据索引,即userID]
func (self *PlayerManger) Get(userID uint64) *Player {
	value, ok := self.Load(userID)
	if ok {
		return value.(*Player)
	}
	return nil
}

//玩家是否存在
func (self *PlayerManger) Exist(userID uint64) bool {
	isHas := false
	self.Range(func(key, value interface{}) bool {
		if key.(uint64) == userID {
			isHas = true
			return false
		}
		return true
	})
	return isHas
}

//按索引删除玩家
func (self *PlayerManger) DeleteIndex(i uint64) {
	self.Delete(i)
}

//删除玩家
func (self *PlayerManger) DeletePlayer(play *Player) {
	//lock.Lock()
	//defer lock.Unlock()
	index := uint64(0)
	isOK := false
	self.Range(func(key, value interface{}) bool {
		if key.(uint64) == play.UserID {
			log.Debug("找到要删除的玩家:%v ", play.UserID)
			index = key.(uint64)
			isOK = true
			value = nil
			return false
		}
		return true
	})
	self.Delete(index)
	play = nil
}

///////////////////////////行为接口////////////////////////////////////////////
//进入场景(限制条件由外部转入)
func (self *Player) Enter(args []interface{}) { //入场
	gameInfo := args[0].(*protoMsg.GameBaseInfo)
	sender := args[1].(gate.Agent)

	if userData := sender.UserData(); userData != nil { //[0
		//玩家信息
		player := userData.(*Player)
		// 平台维度
		platformManger := GetPlatformManger().Get(player.PlatformID)
		// 房间维度
		if nil != platformManger.Roomer {
			if room, ok := platformManger.Roomer.Check(player.RoomNum); ok {
				// 游戏维度
				if game, isOk := room.GetGameHandle(gameInfo.KindID, gameInfo.Level); isOk {
					//-> 校验是否满足进入游戏的条件
					if gameInfo.EnterScore < uint32(self.Gold) && gameInfo.LessScore < uint32(self.Gold) { //[2-0
						var sceneArgs []interface{}
						sceneArgs = append(sceneArgs, gameInfo.Level, sender)
						game.Scene(sceneArgs) // 【进入-> 游戏场景】
					} else { //2]
						//MainLogin, SubEnterGameResult
						GetClientManger().SendData(sender,MainLogin,SubEnterGameResult,&protoMsg.ResResult{State: FAILD, Hints: string("error:not allow to enter!")})
					}
					log.Debug("[进入游戏]---游戏准入条件(LessScore:%v EnterScore:%v)--->user money:%v", gameInfo.LessScore, gameInfo.EnterScore, self.Gold)

				} else { //1]
					GetClientManger().SendData(sender,MainLogin,SubEnterGameResult,&protoMsg.ResResult{State:FAILD, Hints: string("error:not game handle!")})
				}
			} else { //0]]
				GetClientManger().SendData(sender,MainLogin,SubEnterGameResult,&protoMsg.ResResult{State:FAILD, Hints: string("error:not room info!")})
			}
		}
	}
}

func (self *Player) Out(args []interface{}) { //出场
	_=args[2]
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateOut), args[2], nil)
}
func (self *Player) Offline(args []interface{}) { //离线
	_=args[2]
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateOffline), args[2], nil)
}
func (self *Player) Reconnect(args []interface{}) { //重入
	_=args[2]
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateReconnect), args[2], nil)
}
func (self *Player) Ready(args []interface{}) { //准备
	_=args[2]
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateReady), args[2], nil)

}
func (self *Player) Host(args []interface{}) { //抢庄
	_=args[3]
	host := args[2].(*protoMsg.GameHost)
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateHost), args[3], host)

}

func (self *Player) SuperHost(args []interface{}) { //超级抢庄
	_=args[3]
	superHost := args[2].(*protoMsg.GameSuperHost)
	self.updateGameInfo(args[0].(uint32), args[1].(uint32), uint32(GameUpdateHost), args[3], superHost)
}

func (self *Player) updateGameInfo(kindID, level, flag uint32, agent interface{}, message proto.Message) bool {
	// 平台维度
	platformManger := GetPlatformManger().Get(self.PlatformID)
	// 房间维度
	if nil != platformManger.Roomer {
		if room, ok := platformManger.Roomer.Check(self.RoomNum); ok {
			// 游戏维度
			if handle, ok := room.GetGameHandle(kindID, level); ok { //[1-0
				var updateArgs []interface{}
				updateArgs = append(updateArgs, flag, agent, self.UserID, message)
				handle.UpdateInfo(updateArgs)
				return true
			}
		}
	}

	log.Debug("error:GameUpdate->:%v userID:%v", flag, self.UserID)
	return false
}
