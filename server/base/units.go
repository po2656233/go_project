package base

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"log"
)


//异常处理
func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())// 正式开服的时候，这里一定是错误类型 错误内容 等异常信息输出到日志上面的
	}
}

//插入某元素
func CopyInsert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

// 取整数绝对值
func IntAbs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

//打印函数名和行号
func PrintFuncInfo(layer int) {
	//参数 layer 函数所在的层数
	file, fileName, line, ok := runtime.Caller(layer + 1)
	if ok {
		funcName := runtime.FuncForPC(file).Name()
		log.Println( funcName, " -> ", line, " -> ", fileName)
		//fmt.Printf("%s:%d -> -> ->%s\n", funcName, line, fileName)
	}
}

//获取goroutine的id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	log.Println("id:", id)
	return id
}

//堆栈信息输出
func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// id 生成器
