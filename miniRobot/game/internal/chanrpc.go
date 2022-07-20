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
	"sync"
	"sync/atomic"
	"time"
)

var once sync.Once
var isRegister = false

type nodeType struct {
	*sync.Mutex
	allNames []string
}

var count = int32(0)
var node = nodeType{allNames: make([]string, 0), Mutex: &sync.Mutex{}}

// 广播消息
// 这里是对所有玩家进行通知，通知单个游戏的所有玩家，请在单个游戏里实现
func init() {
	CreateRobot(ALLCount)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	AsyncChan.Register("Broadcast", rpcBroadcast)
	//AsyncChan.Register("Broadcast", func(args []interface{}) {
	//    fmt.Println("-------------->Broadcast------->Register")
	//    //a := args[0].(gate.Agent)
	//    //_ = a
	//    //a.WriteMsg(args[1])
	//    m := args[0].(*protoMsg.UserInfo)
	//    //a := args[1].(gate.Agent)
	//    node.Lock()
	//    defer node.Unlock()
	//    for k,v:=range  node.allNames{
	//        if v == m.Account{
	//            node.allNames = append(node.allNames[:k], node.allNames[k+1:]...)
	//            break
	//        }
	//    }
	//
	//}) // 广播消息 调用参考:game.ChanRPC.Go("Broadcast",agent,args)

}

var agentCount int = 0

func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent) //【模块间通信】跟路由之间的通信
	agentCount++
	fmt.Println("当前客户端数量", agentCount)
	login.InitData()
	//注册或登录
	if !isRegister {
		fmt.Println("-注册中-")
		register(agent)
	} else {
		//fmt.Println("-登录中-")
		Logins(agent)
	}
	once.Do(func() {
		go func(aa gate.Agent) {
			for {
				ticker := time.NewTicker(20 * time.Second)
				<-ticker.C
				aa.WriteMsg(&protoMsg.PingReq{})
			}
		}(agent)
	})

}

// 逐一注册
func register(a gate.Agent) {
	//node.Lock()
	//defer node.Unlock()
	size := len(node.allNames)
	curCount := atomic.LoadInt32(&count)
	log.Release("注册索引:%v 名字总数：%v", curCount, size)
	if size <= int(curCount)+1 {
		atomic.CompareAndSwapInt32(&count, curCount, 0)
		Logins(a)
		return
	}

	if 0 <= curCount && int(curCount) < size {

		msg := &protoMsg.RegisterReq{
			Name:       node.allNames[curCount],
			Gender:     0x0F,
			Password:   "rob",
			PlatformID: 1,
		}
		a.WriteMsg(msg)
		log.Debug("注册索引:%v", curCount)
		atomic.AddInt32(&count, 1)

		//time.AfterFunc(10*time.Millisecond, register)
	}

}

func Logins(a gate.Agent) {
	//node.Lock()
	//defer node.Unlock()

	size := len(node.allNames)
	curCount := atomic.LoadInt32(&count)
	if size <= int(curCount) {
		return
	}

	if 0 <= curCount && int(curCount) < size {
		msg := &protoMsg.LoginReq{
			Account:    node.allNames[curCount],
			Password:   "rob",
			PlatformID: 1,
		}
		a.WriteMsg(msg)
		log.Debug("第%v个正在进行登录", curCount)
		atomic.AddInt32(&count, 1)
	}

	//time.AfterFunc(100*time.Millisecond, Logins)
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	curCount := atomic.LoadInt32(&count)
	if 0 < curCount {
		curCount = atomic.AddInt32(&count, -1)
		if curCount < 0 {
			curCount = 0
		}
	}

	//  msg.GetClientManger().DeleteAll()
	log.Debug("当前count:%v", curCount)
	if curCount == 0 {
		log.Debug("重置了")
		MangerPerson.Range(func(k, v interface{}) bool {
			MangerPerson.Delete(k)
			return true
		})
		//  os.Exit(0)
	}
	a.Close()

	_ = a
}

func rpcBroadcast(args []interface{}) interface{} {
	//断线通知
	a := args[0].(gate.Agent)
	_ = a
	//a.WriteMsg(args[1])
	fmt.Println("-------------->Broadcast------->Register")
	//a := args[0].(gate.Agent)
	//_ = a
	//a.WriteMsg(args[1])
	m := args[0].(*protoMsg.UserInfo)
	//a := args[1].(gate.Agent)
	node.Lock()
	defer node.Unlock()
	for k, v := range node.allNames {
		if v == m.Account {
			node.allNames = append(node.allNames[:k], node.allNames[k+1:]...)
			break
		}
	}
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

// 创建
func CreateRobot(num int) []string {
	var f *os.File
	var err1 error
	node.Lock()
	defer node.Unlock()
	//node.allNames = make([]string, 0)
	name := ""
	filename := "./robotList.txt"
	if checkFileIsExist(filename) { //如果文件存在
		isRegister = true
		f, err1 = os.OpenFile(filename, os.O_RDONLY, 0644) //打开文件
		fmt.Println("文件存在")
		//存在的话，就不再生成
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			name = scanner.Text()
			node.allNames = append(node.allNames, name)
		}
		f.Close()
		return node.allNames
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)

	//获取名字列表
	somethings := ""
	for i := 0; i < num; i++ {
		name = GetFullName()
		node.allNames = append(node.allNames, name)
		somethings += name + "\n"
	}

	_, err1 = io.WriteString(f, somethings) //写入文件(字符串)
	check(err1)
	f.Close()
	return node.allNames
}
func GetCount() int32 {
	return atomic.LoadInt32(&count)
}

//////////////////////////////////////
