package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf" //该包的init函数会被调用
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
)


func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	//db,err:=sql.ConnectMySql("root","000","127.0.0.1","3306","qipaidb")

	//gamedata.Example()
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

	//加解密数据
	//对应的模块实现
	//type Module interface {
	//	OnInit()
	//	OnDestroy()
	//	Run(closeSig chan bool)
	//}

	//操作步骤：
	//1、msg 构建传输协议（消息体）
	//2、gate 的router路由分发数据
	//3、game 的handle下进行消息处理

	//---------骨架---------
	//gate 模块，负责游戏客户端的接入。
	//login 模块，负责登录流程。
	//game 模块，负责游戏主逻辑。

	//---------平台---------
	//task 模块，负责任务活动。领取与完成
	//notice 模块，负责游戏通告。平台通告、房间通告、邮箱通告。
	//spread 模块，负责游戏推广(分享)。

	//---------玩家---------
	//stageBag 模块， 玩家道具背包。
	//chat 模块，负责用户聊天。大厅谈论，弹幕，私聊。
	//Recharge 模块，与在线充值，负责在线充值 |customer service 在线客服

	//--------本地信息---------
	//log 模块，负责日志，记录节点信息。
	//sql || database 模块，负责读写数据库。

	//一: 断线重连。 游戏时，保持长连接状态; 在游戏没结束，玩家都可以重新连接进来，并且处于托管状态。
	//二: 定时器。
	//三：事件派发。
	//配置信息是通过server.json配置的

}
