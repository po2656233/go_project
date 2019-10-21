package manger

import (
	"sync"
	"github.com/name5566/leaf/log"
)

//----------平台----------------------
//平台信息
type PlatformInfo struct {
	ID       uint32    //平台ID 0表示无效
	Name      string   //平台名称
	Roomer 	*RoomManger //房间管理
}
//管理房间
type PlatformManger struct {
	sync.Map
	//rooms map[uint32]*RoomInfo
}

////////////////////平台管理////////////////////////////////////
var platformManger *PlatformManger = nil
var platformOnce sync.Once

//玩家管理对象(单例模式)//manger.players = make(map[uint32]*Player)
func GetPlatformManger() *PlatformManger {
	platformOnce.Do(func() {
		platformManger = &PlatformManger{
			sync.Map{},
		}
	})
	return platformManger
}
//添加玩家
func (self *PlatformManger) Append(plat *PlatformInfo) bool {
	if _, ok := self.Load(plat.ID); !ok {
		log.Debug("新增一个玩家ID:%v", plat.ID)
		if nil == plat.Roomer{
			plat.Roomer = CreateRoomManger()
		}
		self.Store(plat.ID, plat)
		return true
	} else {
		log.Debug("玩家ID:%v 已經存在", plat.ID)
		return false
	}
}

//获取指定平台[根据索引,即userID]
func (self *PlatformManger) Get(platformID uint32) *PlatformInfo {
	value, ok := self.Load(platformID)
	if ok {
		return value.(*PlatformInfo)
	}
	return nil
}

//平台是否存在
func (self *PlatformManger) Exist(platformID uint32) bool {
	isHas := false
	self.Range(func(key, value interface{}) bool {
		if key.(uint32) == platformID {
			isHas = true
			return false
		}
		return true
	})
	return isHas
}
