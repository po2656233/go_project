package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"reflect"
	. "server/base"
	. "server/manger"
	protoMsg "server/msg/go"
	"server/sql/mysql"
)

var sqlHandle = mysql.SqlHandle()    //数据库管理
var playerManger = GetPlayerManger() //玩家管理

func init() {
	// 向当前模块（login 模块）注册 Login 消息的消息处理函数 handleTest
	handleMsg(&protoMsg.Register{}, handleRegister)      //反馈--->用户信息
	handleMsg(&protoMsg.Login{}, handleLogin)            //反馈--->主页信息
	handleMsg(&protoMsg.ReqEnterRoom{}, handleEnterRoom) //反馈--->游戏列表
	//handleMsg(&jsonMsg.UserLogin{}, handleLoginJson)
}

//注册模块间的通信
func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//-----------------消息处理-----------------

//注册
func handleRegister(args []interface{}) {
	m := args[0].(*protoMsg.Register)
	//a := args[1].(gate.Agent)
	log.Debug("msg: %v psw:%v", m.GetName(), m.GetPassword())
	//sqlHandle.AddUser(m.GetName(), m.GetPassword())

}

//登录(房间列表)
func handleLogin(args []interface{}) {
	m := args[0].(*protoMsg.Login)
	log.Debug("[receive]LoginInfo:->%v", m)

	//
	agent := args[1].(gate.Agent)
	player := &Player{}
	//数据库校验玩家数据
	if uid, ok := sqlHandle.CheckLogin(m.GetAccount(), m.GetPassword()); ok {
		log.Debug("Login Success!")

		//获取游戏列表
		////据说这种遍历可以防止编码问题
		//for index, item := range gameList.Items {
		//	log.Debug("%d->[游戏名称:%s\t类型:%d\t标识:%v\t房间:%v  入场:%v\t坐下:%v]\n", index+1, item.Name, item.Type, item.KindID, item.RoomID, item.EnterScore, item.LessScore)
		//}
		name, age, sex, vipLevel, gold := sqlHandle.CheckUserInfo(uid)
		//房间列表
		msg := &protoMsg.MasterInfo{}
		userInfo := &protoMsg.UserInfo{}
		userInfo.Name = name
		userInfo.Accounts = name
		userInfo.Age = age
		userInfo.Gender = sex
		userInfo.Level = vipLevel
		userInfo.Gold = gold
		msg.UserInfo = userInfo
		msg.RoomsInfo = sqlHandle.CheckRoomList(uid)

		//log.Debug("房间号列表%v", msg.RoomsInfo)

		//添加到用户管理
		player.UserID = uid
		player.Name = name
		player.Account = name
		player.Gold = gold
		player.Sate = INVALID
		player.PlatformID = sqlHandle.CheckPlatformInfo(uid)
		player.RoomNum = INVALID
		player.GameID = INVALID

		// 保存玩家数据
		playerManger.Append(player)

		// 往agent里添加数据
		agent.SetUserData(player)

		//发送【房间列表】
		GetClientManger().SendData(agent, MainLogin, SubMasterInfo, msg)
	} else {
		//失败信息
		loginResult := &protoMsg.ResResult{}
		loginResult.State = *proto.Uint32(FAILD)
		loginResult.Hints = *proto.String("Failed")

		//【返回结果】[MainID|SubID]
		GetClientManger().SendData(agent, MainLogin, SubLoginResult, loginResult)

		//日志打印
		log.Error("Login Failed!")
	}
}

//进入房间(游戏列表)
func handleEnterRoom(args []interface{}) {
	m := args[0].(*protoMsg.ReqEnterRoom)
	agent := args[1].(gate.Agent)

	log.Debug("[进入房间]:%v", m)
	//查找代理信息
	if userData := agent.UserData(); userData != nil { //[0
		//用户信息
		player := userData.(*Player)
		//获取房间号码
		player.RoomNum = m.GetRoomNum()

		//游戏列表信息
		_, _, msg := sqlHandle.CheckGameList(player.RoomNum)

		//发送数据
		GetClientManger().SendData(agent, MainLogin, SubGameList, msg)

	} else {
		//失败信息
		enterResult := &protoMsg.ResResult{}
		enterResult.State = *proto.Uint32(FAILD)
		enterResult.Hints = *proto.String("Failed")

		//【返回结果】[MainID|SubID]
		GetClientManger().SendData(agent, MainLogin, SubEnterRoomResult, enterResult)

		//日志打印
		log.Error("Login Failed!")
	}

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
