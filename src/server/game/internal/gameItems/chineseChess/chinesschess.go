package chineseChess

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	. "server/base"
	. "server/manger"
	protoMsg "server/msg/go"
	"server/sql/mysql"
	_ "server/sql/mysql"
)

var timer *module.Skeleton = nil  //定时器
var sqlHandle = mysql.SqlHandle() //数据库
var manger = GetPlayerManger()    //玩家管理类
var playerList protoMsg.UserList  //玩家列表(观看)

//继承于GameItem
type ChineseChessGame struct {
	GameItem
	bStart    bool   //第一次启动
	gameState uint32 //游戏状态
	timeStamp int64  //时间戳

}

//创建中国象棋实例
func NewChineseChess(level uint32, skeleton *module.Skeleton) *ChineseChessGame {
	log.Debug("------正在创建中国象棋实例------level:%v", level)
	p := &ChineseChessGame{}
	p.Init(level, skeleton)
	return p
}

//---------------------------中国象棋----------------------------------------------//
//初始化信息
func (self *ChineseChessGame) Init(level uint32, skeleton *module.Skeleton) {
	timer = skeleton                  //定时器的使用
	self.bStart = false               //是否第一次启动
	self.Level = level                //房间等级
	self.gameState = SubGameSenceFree //场景状态
}

//游戏接口
//场景
func (self *ChineseChessGame) Scene(args []interface{}) {



}
//开始
func (self *ChineseChessGame) Start(args []interface{}) {

}

//游戏
func (self *ChineseChessGame) Playing(args []interface{}) {

}

//结算
func (self *ChineseChessGame) Over(args []interface{}) {

}

//更新信息
func (self *ChineseChessGame) UpdateInfo(args []interface{}){

}

//超级控制 可在检测到没真实玩家时,且处于空闲状态时,自动关闭(未实现)
func (self *ChineseChessGame) SuperControl(args []interface{}) {

}