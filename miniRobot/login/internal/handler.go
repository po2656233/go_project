package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	. "miniRobot/base"
	"miniRobot/conf"
	protoMsg "miniRobot/msg/go"
	"reflect"
)

var allgames = make([]*protoMsg.GameItem, 0)
var count int = 0

func init() {

	// 向当前模块（login 模块）注册 Login 消息的消息处理函数 handleTest
	//handleMsg(&jsonMsg.UserLogin{}, handleLoginJson)
	handleMsg(&protoMsg.RegisterResp{}, handleRegister) //反馈--->主页信息
	handleMsg(&protoMsg.LoginResp{}, handleLogin)       //反馈--->主页信息

	handleMsg(&protoMsg.ChooseClassResp{}, handleChooseClass) //反馈--->主页信息
	handleMsg(&protoMsg.ChooseGameResp{}, handleChooseGame)   //反馈--->主页信息
	handleMsg(&protoMsg.ResultResp{}, handleResultResp)       //反馈--->主页信息
	handleMsg(&protoMsg.ResultPopResp{}, handleResultPopResp) //反馈--->主页信息
	//充值
	handleMsg(&protoMsg.PongResp{}, handlePongResp) //反馈--->主页信息

}

// 注册模块间的通信
func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/////////////

func handlePongResp(args []interface{}) {
	_ = args[0].(*protoMsg.PongResp)
	//log.Debug("收到心跳")
}

// -----------------消息处理-----------------
func handleRegister(args []interface{}) {
	m := args[0].(*protoMsg.RegisterResp)
	log.Debug("注册成功:%v", m)
}

func handleLogin(args []interface{}) {
	m := args[0].(*protoMsg.LoginResp)
	a := args[1].(gate.Agent)
	count++
	a.SetUserData(m.MainInfo.UserInfo)
	log.Debug("机器人:%v-%v  登录成功! 总人数:%v", m.MainInfo.UserInfo.UserID, m.MainInfo.UserInfo.Account, count)

	//获取游戏分类列表
	for _, cls := range m.MainInfo.Classes.Classify {
		log.Debug("%v-%v 进入房间%v", m.MainInfo.UserInfo.UserID, m.MainInfo.UserInfo.Account, cls)
		msg := &protoMsg.ChooseClassReq{
			ID:       uint32(cls.ID),
			TableKey: cls.Key,
		}
		a.WriteMsg(msg)
		//break

	}

}

// 获取房间列表
func handleChooseClass(args []interface{}) {
	m := args[0].(*protoMsg.ChooseClassResp)
	a := args[1].(gate.Agent)

	//person := a.UserData().(*protoMsg.UserInfo)
	for _, item := range m.Games.Items {
		have := false
		for _, game := range allgames {
			if game.ID == item.ID {
				have = true
				break
			}
		}
		if !have {
			allgames = append(allgames, item)
		}
	}

	for _, game := range allgames {
		msg := &protoMsg.ChooseGameReq{
			Info:    game.Info,
			PageNum: 1,
		}
		// log.Debug("共有游戏%v个",len(allgames))
		// log.Debug("玩家'%v(ID:%v)' GameSize:%v 请求游戏详情:ID:%v  %v", person.Account, person.UserID,len(allgames),allgames[atomic.LoadInt32(&IndexGames)].ID, msg.Info)
		a.WriteMsg(msg)
	}
}

// 选择游戏
func handleChooseGame(args []interface{}) {
	m := args[0].(*protoMsg.ChooseGameResp)
	a := args[1].(gate.Agent)
	person := a.UserData().(*protoMsg.UserInfo)

	for _, item := range m.Tables.Items {
		if person.Age == 300 {
			return
		}
		val, ok := MangerPerson.Load(item.GameID)
		if !ok {
			MangerPerson.Store(item.GameID, uint32(0))
			val, _ = MangerPerson.Load(item.GameID)
		}

		if item.GameID == 0 {
			continue
		}
		//  log.Debug("[进前]桌子名称:%v 游戏ID:%v 当前人数:%v ", item.Info.Name, item.GameID, item.Info.MaxOnline)
		chair := val.(uint32) + 1
		if int64(item.Info.EnterScore) < person.Money && item.Info.HostID == 0 &&
			(chair < item.Info.MaxChair || (0 == item.Info.MaxChair && chair < uint32(conf.Server.TablePeopleMax))) {
			//不能坐满，留个座位给真实玩家
			msg := &protoMsg.EnterGameReq{
				GameID:   item.GameID,
				Password: item.Info.Password,
			}
			//将游戏ID保存至玩家信息
			person.Age = 300
			person.AgentID = msg.GameID
			person.ReferralCode = msg.Password
			a.SetUserData(person)

			MangerPerson.Store(item.GameID, chair)
			a.WriteMsg(msg)
			log.Debug("[坐下]桌子名称:%v 游戏ID:%v 机器人ID:%v 当前人数:%v 机器人数:%v 最大可容纳:%v", item.Info.Name, item.GameID, person.GetUserID(), item.Info.MaxOnline+chair, chair, item.Info.MaxChair)
			return
		}
	}

}

func handleResultResp(args []interface{}) {
	m := args[0].(*protoMsg.ResultResp)
	//a := args[1].(gate.Agent)
	if m.Hints == "下注失败: 无效金额!" {
		//msg := &protoMsg.RechargeReq{
		//	Info:    allgames[atomic.LoadInt32(&IndexGames)].Info,
		//	PageNum: 0,
		//}
		//a.WriteMsg(msg)
		log.Debug("提示:%v", m)
	}
	// log.Debug("提示:%v", m)
}

// 弹窗提示
func handleResultPopResp(args []interface{}) {
	m := args[0].(*protoMsg.ResultPopResp)
	log.Debug("弹窗提示:%v", m)
	//	a := args[1].(gate.Agent)
	//	person := a.UserData().(*protoMsg.UserInfo)
	//if m.Hints == "您的账号已经在异地登录了!" {
	//	game.ChanRPC.Go("Broadcast", a, person)
	//}
}

// ///////////////////////////////////////////////////////////////////////////
