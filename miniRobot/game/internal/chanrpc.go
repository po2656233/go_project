package internal

import (
    "bufio"
    "fmt"
    "github.com/name5566/leaf/gate"
    "github.com/name5566/leaf/log"
    "io"
    . "miniRobot/base"
    "miniRobot/login"
    protoMsg "miniRobot/msg/go"
    "os"
    "time"
)

var haveNames bool
var allNames []string
var count = 0
var agentReg gate.Agent

//广播消息
//这里是对所有玩家进行通知，通知单个游戏的所有玩家，请在单个游戏里实现
func init() {
    haveNames = false
    CreateRobot(ALLCount)
    skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
    skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
    AsyncChan.Register("Broadcast", func(args []interface{}) {
        fmt.Println("-------------->Broadcast------->Register")
        //a := args[0].(gate.Agent)
        //_ = a
        //a.WriteMsg(args[1])
    }) // 广播消息 调用参考:game.ChanRPC.Go("Broadcast",agent,args)

}

//逐一注册
//并登录

func register() {
    size := len(allNames)
    if size < count {
        count = 0
        logins()
        return
    }
    if count == 0 {
        count = 1
    }

    if 0 < count && count < size {
        msg := &protoMsg.RegisterReq{
            Name:       allNames[count-1],
            Password:   "rob",
            PlatformID: 1,
        }
        agentReg.WriteMsg(msg)
        count++
        time.AfterFunc(10*time.Millisecond, register)
    }

}

func logins() {
    size := len(allNames)
    if size < count+1 {
        return
    }
    if 0 <= count && count < size {
        msg := &protoMsg.LoginReq{
            Account:    allNames[count],
            Password:   "rob",
            PlatformID: 1,
        }
        agentReg.WriteMsg(msg)
        count++
    }

    //time.AfterFunc(100*time.Millisecond, logins)
}

func rpcNewAgent(args []interface{}) {
    agent := args[0].(gate.Agent) //【模块间通信】跟路由之间的通信
    _ = agent
    //fmt.Println("-成功創建-")
    agentReg = agent
    //注册或登录
    if !haveNames {
        register()
    } else {
        logins()
    }

}

func rpcCloseAgent(args []interface{}) {
    a := args[0].(gate.Agent)
    if 0 < count {
        count--
        if count < 0 {
            count = 0
        }
    }
    if count == 0{
        log.Debug("重置了")
        login.Module.SetIndex(-1)
        MangerPerson.Range(func(k, v interface{}) bool {
            MangerPerson.Delete(k)
            return true
        })
    }
    a.Close()

    _ = a
}

func rpcBroadcast(args []interface{}) interface{} {
    //断线通知
    a := args[0].(gate.Agent)
    _ = a
    a.WriteMsg(args[1])
    return error(nil)
}

////////////////////////////////////

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

//创建
func CreateRobot(count int) []string {
    var f *os.File
    var err1 error
    allNames = make([]string, 0)
    name := ""
    filename := "./robotList.txt"
    if checkFileIsExist(filename) { //如果文件存在
        f, err1 = os.OpenFile(filename, os.O_RDONLY, 0644) //打开文件
        fmt.Println("文件存在")
        //存在的话，就不再生成
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            name = scanner.Text()
            allNames = append(allNames, name)
        }
        f.Close()
        haveNames = true
        return allNames
    } else {
        f, err1 = os.Create(filename) //创建文件
        fmt.Println("文件不存在")
    }
    check(err1)

    //获取名字列表
    somethings := ""
    for i := 0; i < count; i++ {
        name = GetFullName()
        allNames = append(allNames, name)
        somethings += name + "\n"
    }

    _, err1 = io.WriteString(f, somethings) //写入文件(字符串)
    check(err1)
    f.Close()
    haveNames = false
    return allNames
}

//////////////////////////////////////
