package internal

import (
    "github.com/name5566/leaf/gate"
    "github.com/name5566/leaf/log"
    . "miniRobot/base"
    protoMsg "miniRobot/msg/go"
    "reflect"
    "sync/atomic"
)



func init() {

    // 向当前模块（login 模块）注册 Login 消息的消息处理函数 handleTest
    //handleMsg(&jsonMsg.UserLogin{}, handleLoginJson)
    handleMsg(&protoMsg.RegisterResp{}, handleRegister)       //反馈--->主页信息
    handleMsg(&protoMsg.LoginResp{}, handleLogin)             //反馈--->主页信息
    handleMsg(&protoMsg.ChooseClassResp{}, handleChooseClass) //反馈--->主页信息
    handleMsg(&protoMsg.ChooseGameResp{}, handleChooseGame)   //反馈--->主页信息
    handleMsg(&protoMsg.ResultResp{}, handleResultResp)       //反馈--->主页信息
    handleMsg(&protoMsg.ResultPopResp{}, handleResultPopResp) //反馈--->主页信息


    handleMsg(&protoMsg.PongResp{}, handlePongResp)             //反馈--->主页信息

}

//注册模块间的通信
func handleMsg(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

/////////////

func handlePongResp(args []interface{}) {
    _ = args[0].(*protoMsg.PongResp)
    //time.AfterFunc(time.Duration(20)*time.Second, func() {
    //   if a,ok := args[1].(gate.Agent);ok{
    //       a.WriteMsg(&protoMsg.PingReq{})
    //   }
    //})
}


//-----------------消息处理-----------------
func handleRegister(args []interface{}) {
    m := args[0].(*protoMsg.RegisterResp)
    log.Debug("注册成功:%v", m)
}


func handleLogin(args []interface{}) {
    m := args[0].(*protoMsg.LoginResp)
    a := args[1].(gate.Agent)
    //log.Debug("登录成功:%v", m)
    a.SetUserData(m.MainInfo.UserInfo)

    //time.AfterFunc(time.Duration(20)*time.Second, func() {
    //    if a,ok := args[1].(gate.Agent);ok{
    //        a.WriteMsg(&protoMsg.PingReq{})
    //    }
    //})

    //获取游戏分类列表
    msg := &protoMsg.ChooseClassReq{}
    for _, cls := range m.MainInfo.Classes.Classify {
        msg.ID = uint32(cls.ID)
        msg.TableKey = cls.Key
        a.WriteMsg(msg)
        return
    }

}

//获取房间列表
func handleChooseClass(args []interface{}) {
    m := args[0].(*protoMsg.ChooseClassResp)
    a := args[1].(gate.Agent)

    // person := a.UserData().(*protoMsg.UserInfo)
    //选择游戏
    index := atomic.AddInt32(&IndexGames, 1)
    if index < int32(len(m.Games.Items)) {
        msg := &protoMsg.ChooseGameReq{
            Info:    m.Games.Items[index].Info,
            PageNum: 0,
        }
        //  log.Debug("玩家'%v(ID:%v)' 请求游戏详情:ID:%v %v", person.Account, person.UserID, m.Games.Items[index].ID, msg.Info)
        a.WriteMsg(msg)
    } else {
        ok := atomic.CompareAndSwapInt32(&IndexGames, IndexGames, -1)
        log.Debug("---------------------%v---------------------", ok)
    }

}

//选择游戏
func handleChooseGame(args []interface{}) {
    m := args[0].(*protoMsg.ChooseGameResp)
    a := args[1].(gate.Agent)
    // person := a.UserData().(*protoMsg.UserInfo)
    for _, item := range m.Tables.Items {
        val, ok := MangerPerson.Load(item.GameID)
        if !ok {
            MangerPerson.Store(item.GameID, uint32(0))
            val, _ = MangerPerson.Load(item.GameID)
        }
        realCount := item.Info.MaxOnline + val.(uint32) + 1
        if realCount < item.Info.MaxChair || (0 == item.Info.MaxChair && val.(uint32)+1 < 30) {
            //不能坐满，留个座位给真实玩家
            msg := &protoMsg.EnterGameReq{
                GameID:   item.GameID,
                Password: item.Info.Password,
            }
            a.WriteMsg(msg)
            MangerPerson.Store(item.GameID, uint32(val.(uint32)+1))
            log.Debug("桌子名称:%v 游戏ID:%v 当前人数:%v 机器人:%v 最大可容纳:%v", item.Info.Name, item.GameID, realCount, val.(uint32)+1, item.Info.MaxChair)
            return
        }

    }

}

//
func handleResultResp(args []interface{}) {
    // m := args[0].(*protoMsg.ResultResp)
    //a := args[1].(gate.Agent)
    // log.Debug("提示:%v", m)
}

//选择游戏
func handleResultPopResp(args []interface{}) {
    m := args[0].(*protoMsg.ResultPopResp)
    //a := args[1].(gate.Agent)
    log.Debug("弹窗提示:%v", m)
}

/////////////////json-->测试用/////////////////////////////
// 消息处理
//func handleLoginJson(args []interface{}) {
//	// 收到的 Test 消息
//	m := args[0].(jsonMsg.UserLogin)
//	// 消息的发送者
//	a := args[1].(gate.Agent)
//	// 1 查询数据库--判断用户是不是合法
//	// 2 如果数据库返回查询正确--保存到缓存或者内存
//	// 输出收到的消息的内容
//	log.Debug("Test login %v", m.LoginName)
//	// 给发送者回应一个 Test 消息
//	a.WriteMsg(&jsonMsg.UserLogin{
//		LoginName: "client",
//	})
//}

//func handleRequestRoomInfoJson(args []interface{})  {
//	m := args[0].(jsonMsg.RequestRoomInfo)
//	// 消息的发送者
//	//a := args[1].(gate.Agent)
//	// 1 查询数据库--判断用户是不是合法
//	// 2 如果数据库返回查询正确--保存到缓存或者内存
//	// 输出收到的消息的内容
//	log.Debug("Test handleRequestRoomInfoJson %v", m)
//	// 给发送者回应一个 Test 消息
//	//a.WriteMsg(&jsonMsg.UserLogin{
//	//	LoginName: "client",
//	//})
//}

//////////////////数据库查询////////////////////////////

//[测试用]
//a := args[1].(gate.Agent)
//a.WriteMsg(&protoT.TestPro{
//	Name:*proto.String("kaile"),
//	Password:*proto.String("doo"),
//})
//Processor.Unmarshal(args[0].([]byte))
//
//buf := make([]byte, 32)
//// 接收消息
//n:=len(args)
//m := &proto.TestPro{}
//proto.Unmarshal(buf[4:n], m)
//
//// 消息的发送者
//a := args[1].(gate.Agent)
//defer a.Close()
//
//// 输出收到的消息的内容
//log.Debug("name:%v password:%v", m.GetName(), m.GetPassword())
//
//
//// 给发送者回应一个 Hello 消息
//a.WriteMsg(proto.UserLogin{})
