package internal

import (

	"reflect"




)

//初始化
func init() {
	handlerBehavior()
}

//注册传输消息
func handlerMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//玩家行为处理
func handlerBehavior() {

}
