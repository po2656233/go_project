package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/name5566/leaf"
    lconf "github.com/name5566/leaf/conf" //该包的init函数会被调用
    "miniRobot/conf"
    "miniRobot/game"
    "miniRobot/gate"
    "miniRobot/login"
    "runtime"
)

//type Agent struct {
//	conn *network.TCPConn
//}
//
//func newAgent(conn *network.TCPConn) network.Agent {
//	a := new(Agent)
//	a.conn = conn
//	return a
//}
//
//func (a *Agent) Run() {}
//
//func (a *Agent) OnClose() {}
//

func main() {
    //确保并发执行
    runtime.GOMAXPROCS(runtime.NumCPU() * 64)

    // 返回当前处理器的数量
    fmt.Println(runtime.GOMAXPROCS(0))
    // 返回当前机器的逻辑处理器或者核心的数量
    fmt.Println(runtime.NumCPU())

    lconf.LogLevel = conf.Server.LogLevel
    lconf.LogPath = conf.Server.LogPath
    lconf.LogFlag = conf.LogFlag
    lconf.ConsolePort = conf.Server.ConsolePort
    lconf.ProfilePath = conf.Server.ProfilePath

    leaf.Run(
        game.Module,
        gate.Module,
        login.Module,
    )

    //client := new(network.TCPClient)
    //client.Addr = "127.0.0.1:9960"
    //client.ConnNum = 1
    //client.ConnectInterval = 3 * time.Second
    //client.PendingWriteNum = conf.PendingWriteNum
    //client.LenMsgLen = 4
    //client.MaxMsgLen = math.MaxUint32
    //client.NewAgent = newAgent
    //
    //client.Start()
    //
    //fmt.Println("goroutines:", runtime.NumGoroutine())
    //
    //
    //
    //c := make(chan os.Signal, 1)
    //signal.Notify(c, os.Interrupt, os.Kill)
    //sig := <-c
    //log.Release("Leaf closing down (signal: %v)", sig)
    //////////////
    ////性能测试
    ////cpu
    //cpuProfile, _ := os.Create("cpu_profile")
    //pprof.StartCPUProfile(cpuProfile)
    //defer pprof.StopCPUProfile()
    ////内存
    //memProfile, _ := os.Create("mem_profile")
    //pprof.WriteHeapProfile(memProfile)

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

    //---------平台（大厅）---------
    //task 模块，负责任务活动。领取与完成
    //notice 模块，负责游戏通告。平台通告、房间通告、邮箱通告。
    //spread 模块，负责游戏推广(分享)。
    //statistics 模块，统计面板。玩家输赢后的平台掉星或加星 亦或加减分

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
