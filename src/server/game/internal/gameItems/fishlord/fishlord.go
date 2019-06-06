package fishlord

//牛牛

import (
	"github.com/name5566/leaf/log"
	. "server/base"
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
	"time"
)

var sqlhandle = mysql.SqlHandle()

//游戏状态 默认空闲
var gameState uint32 = SubGameSenceFree

//定时器
const (
	freeTime = 5
	betTime  = 15
	openTime = 10
)

func init() {
	onStart()
}

//
func Enter(args []interface{}) {

}

//下注
func Playing(args []interface{}) {
	//直接扣除金币
	log.Debug("玩家下注")
	//return
	m := args[0].(*protoMsg.GameFishLordPlaying)
	//【传输对象】
	//courier := &Courier{Agent: args[1].(gate.Agent)}
	log.Debug("BaccaratPlaying:->%v", m)

	//广播给所有玩家
	//下注成功
}

//定时器控制
func cowsTimer() {
	time.AfterFunc(5*time.Second, func() { log.Debug("func") })
	onStart()
	//for i := 0; i < 10; i++ {
	//	log.Debug("->[%d]", i)
	//	time.Sleep(time.Second)
	//}
}

//开始|下注|结算
func onStart() {
	gameState = SubGameSenceStart
	time.AfterFunc(freeTime*time.Second, onPlayerBet)
}
func onPlayerBet() {
	gameState = SubGameSencePlaying
	time.AfterFunc(betTime*time.Second, onOver)
}
func onOver() {
	gameState = SubGameSenceOver
	time.AfterFunc(openTime*time.Second, onStart)
}

//逻辑层
