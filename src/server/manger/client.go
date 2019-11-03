package manger

import (
	"github.com/golang/protobuf/proto"
	protoMsg "server/msg/go"
	"sync"
	"github.com/name5566/leaf/log"
	"leaf/gate"
	"server/conf"
)

//
type ClientManger struct {
	sync.Map
}

var clientManger *ClientManger = nil
var clientOnce sync.Once
var wait sync.WaitGroup

//玩家管理对象(单例模式)//manger.players = make(map[uint64]*Player)
func GetClientManger() *ClientManger {
	clientOnce.Do(func() {
		clientManger = &ClientManger{
			sync.Map{},
		}
	})
	return clientManger
}


//添加玩家
func (self *ClientManger) Append(agent gate.Agent) bool {
	if _, ok := self.Load(agent); !ok {
		log.Debug("新增一个客户端IP:%v",agent.RemoteAddr() )
		self.Store(agent,struct{}{})
		return true
	} else {
		log.Debug("客户端IP:%v 已經存在", agent.RemoteAddr())
		return false
	}
}

//客户端连接是否存在
func (self *ClientManger) Get(userID uint64) ( gate.Agent,bool) {
	isHas := false
	var agent gate.Agent
	self.Range(func(key, value interface{}) bool {
		agent = key.(gate.Agent)
		if nil != agent && nil != agent.UserData() && agent.UserData().(*Player).UserID == userID {
			isHas = true
			return false
		}
		return true
	})
	return agent,isHas
}


//删除客户端
func (self *ClientManger) DeleteClient(userID uint64) {
	if agent,ok:=self.Get(userID);ok{
		self.Delete(agent)
	}
}

///



///------------------------广播---------------------------------------///
//全网广播
func (self *ClientManger) NotifyAll(mainID, subID uint32, msg proto.Message) {
	defer wait.Done()
	wait.Add(1)
	self.Range(func(key, value interface{}) bool {
		agent := key.(gate.Agent)
		if nil == agent {
			log.Debug("无效客户端:%v", key)
			return true
		}
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}


		//校验用户信息
		userData := agent.UserData()
		if nil != userData{
			player := userData.(*Player)
			log.Debug("通知玩家：%v %v", player.UserID, player.Name)
		}else {
			log.Debug("err:客户端IP：%v 无效玩家信息", agent.RemoteAddr())
		}

		//广播给客户端
		agent.NotifyMsg(packet)
		return true
	})
}

//
func (self *ClientManger) SendData(agent gate.Agent, mainID, subID uint32, msg proto.Message){
	if conf.Server.CarryMainSubID{
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}
		agent.WriteMsg(packet)
		return
	}

	agent.WriteMsg(msg)
}



//通知除指定玩家外的玩家们
func (self *ClientManger) NotifyButOthers(userIDs []uint64, mainID, subID uint32, msg proto.Message) {
	defer wait.Done()
	wait.Add(1)
	self.Range(func(key, value interface{}) bool {
		agent := key.(gate.Agent)
		if nil == agent {
			log.Debug("无效客户端:%v", key)
			return true
		}
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}


		//获取用户ID
		userData := agent.UserData()
		if nil != userData{
			userID := userData.(*Player).UserID
			//不通知该部分玩家
			for _, uid := range userIDs {
				if uid == userID {
					return true
				}
			}
			log.Debug("通知[部分]玩家：%v", userID)
		}else {
			log.Debug("err:[部分]客户端IP：%v 无效玩家信息", agent.RemoteAddr())
		}

		//广播给客户端
		agent.NotifyMsg(packet)
		return true
	})
}

func (self *ClientManger) NotifyOthers(userIDs []uint64, mainID, subID uint32, msg proto.Message) {

	defer wait.Done()
	wait.Add(1)
	self.Range(func(key, value interface{}) bool {
		agent := key.(gate.Agent)
		if nil == agent {
			log.Debug("无效客户端:%v", key)
			return true
		}
		//指令+数据  包体内容
		data, _ := proto.Marshal(msg)
		packet := &protoMsg.PacketData{
			MainID:    mainID,
			SubID:     subID,
			TransData: data,
		}


		//获取用户ID
		userData := agent.UserData()
		if nil != userData{
			userID := userData.(*Player).UserID
			//不通知该部分玩家
			for _, uid := range userIDs {
				if uid == userID {
					//广播给客户端
					agent.NotifyMsg(packet)
					log.Debug("通知[指定]玩家：%v", userID)
				}
			}

		}else {
			log.Debug("err:[指定]客户端IP：%v 无效玩家信息", agent.RemoteAddr())
		}


		return true
	})
}
