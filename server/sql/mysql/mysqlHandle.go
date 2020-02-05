package mysql

import (
	"database/sql"
	"github.com/name5566/leaf/log"
	. "server/base"
	protoMsg "server/msg/go"
	"strconv"
	"strings"
	"sync"
	"server/manger"
)

type SqlMan struct {
	user     string
	password string
	address  string
	port     string
	dbName   string
	db       *sql.DB
}

type ISqlOP interface {
	ConnectMySql(user, psw, addr, port, db string) //连接数据库
	CheckServerList() *protoMsg.GameList           //房间列表
	CheckLogin(user, password string) bool         //登陆
	CheckMoney(userID uint64) float32              //金额
	DeductMoney(userID uint64, money float32) bool //扣除金币
	CloseMysql()
}

var once sync.Once
var sqlobj *SqlMan

func SqlHandle() *SqlMan {
	once.Do(func() {
		sqlobj = &SqlMan{
			user:     string("root"),
			password: string("0000"),
			address:  string("127.0.0.1"),
			port:     string("3306"),
			dbName:   string("qipaiinfo"),
			db:       nil,
		}
		sqlobj.Init()
	})

	return sqlobj
}

/////////////////////////////////
//初始化
func (self *SqlMan) Init() {
	if self.db == nil {
		_, err := self.ConnectMySql("root", "0000", "127.0.0.1", "3306", "qipaiinfo")
		CheckError(err)
	}
	//获取平台信息
	self.initPlatformInfo()

	//获取所有房间信息

	//获取所有玩家信息
}

//获取平台信息
func (self *SqlMan) initPlatformInfo() bool {
	rows, err := self.db.Query("SELECT id,name FROM platform ")
	defer rows.Close()
	CheckError(err)
	var id uint32
	var name string
	for rows.Next() {

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err.Error())
			return false
		}
		manger.GetPlatformManger().Append(&manger.PlatformInfo{
			ID:   id,
			Name: name,
		})
	}
	return true
}

// 创建玩家信息

//关闭mysql
func (self *SqlMan) CloseMysql() {
	self.db.Close()
}

//游戏信息
func (self *SqlMan) CheckGameInfo(gameID uint32) *protoMsg.GameBaseInfo {
	//数据库查询
	rows, err := self.db.Query("SELECT Name,Type,Level,KindID,EnterScore,LessScore FROM game where ID=?;", gameID)
	defer rows.Close()
	CheckError(err)

	//log.Debug("读取数据库数据:->...游戏信息")
	var gameItem protoMsg.GameBaseInfo
	for rows.Next() {
		err := rows.Scan(&gameItem.Name, &gameItem.Type, &gameItem.Level, &gameItem.KindID, &gameItem.EnterScore, &gameItem.LessScore)
		CheckError(err)
		//log.Debug("GameName=%v GameType=%v KindID:%v\n", gameItem.Name, gameItem.Type, gameItem.KindID)
		return &gameItem
	}
	return nil
}

//服务列表查询
func (self *SqlMan) CheckGameList(roomID uint32) (name, key string, games *protoMsg.GameList) {
	//从房间中找到kindID和Level
	//数据库查询
	rows, err := self.db.Query("SELECT Name,RoomKey,Games FROM room where Num =?;", roomID)
	defer rows.Close()
	CheckError(err)
	var strGames string
	gameList := &protoMsg.GameList{}
	//log.Debug("服务列表查询")

	for rows.Next() {
		err := rows.Scan(&name, &key, &strGames)
		CheckError(err)
		//log.Debug("GameName=%v GameType=%v KindID:%v\n", gameItem.Name, gameItem.Type, gameItem.KindID)
		roomGames := strings.Split(strGames, ",")
		log.Debug("服务列表查询:%v", roomGames)
		for _, gameID := range roomGames {
			id, _ := strconv.Atoi(gameID)
			if gameInfo := self.CheckGameInfo(uint32(id)); gameInfo != nil {
				var gameItem protoMsg.GameItem
				gameItem.ID = uint32(id)
				gameItem.Info = gameInfo
				//log.Debug("GameName=%v GameType=%v Level:%v KindID:%v EnterScore:%v LessScore：%v\n",gameInfo.Name, gameInfo.Type,gameInfo.Level, gameInfo.KindID, gameInfo.EnterScore, gameInfo.LessScore)
				gameList.Items = append(gameList.Items, &gameItem)
			}
		}
	}

	return name, key, gameList
}

//获取房间列表
func (self *SqlMan) CheckRoomList(userID uint64) (roomsInfo []*protoMsg.RoomInfo) {
	//数据库查询
	rows, err := self.db.Query("SELECT RoomNums FROM user where ID=?;", userID)
	defer rows.Close()
	CheckError(err)

	log.Debug("读取数据库数据:->...房间列表")
	var strNums string
	for rows.Next() {
		err := rows.Scan(&strNums)
		CheckError(err)
		log.Debug("[数据库]:rooms=%v\n", strNums)

	}

	//提取数据
	allRooms := strings.Split(strNums, ",")
	for _, room := range allRooms {
		if num, error := strconv.Atoi(room); error == nil {
			var info protoMsg.RoomInfo
			info.RoomNum = uint32(num)
			info.RoomName, info.RoomKey, info.Games = self.CheckGameList(uint32(num))
			roomsInfo = CopyInsert(roomsInfo, len(roomsInfo), &info).([]*protoMsg.RoomInfo)
		}
	}
	return roomsInfo
}

func (self *SqlMan) CheckTaskList(userID uint64) (tasks []uint32) {
	return tasks
}

//登录查询，改为redis
func (self *SqlMan) CheckLogin(user, password string) (uid uint64, isSuccessful bool) {
	isSuccessful = false
	rows, err := self.db.Query("SELECT ID,Name,Password FROM user where Name = ? and Password=?;", user, password)
	defer rows.Close()
	CheckError(err)

	log.Debug("读取数据库数据:->...登录查询")
	for rows.Next() {
		var name string
		var psw string
		err := rows.Scan(&uid, &name, &psw)
		CheckError(err)
		//log.Debug("[数据库]Name=%v Password=%v\n", name, psw)
		if name == user && psw == password {
			isSuccessful = true
			break
		}
	}

	if isSuccessful {
		log.Debug("登陆成功")
	} else {
		log.Debug("登陆失败,请检视<账号|密码>")
	}
	return uid, isSuccessful
}

//获取玩家账户金币
func (self *SqlMan) CheckMoney(userID uint64) float64 {

	//rows,err := db.Query("SELECT 'ID','money-count' FROM userinfo WHERE ID=?",userID)
	rows, err := self.db.Query("SELECT Money FROM user WHERE ID in(?)", userID)
	defer rows.Close()

	CheckError(err)

	userMoney := float64(0.0)
	for rows.Next() {
		if err := rows.Scan(&userMoney); err != nil {
			log.Fatal(err.Error())
		}
		break
	}
	return userMoney
}

func (self *SqlMan) CheckName(userID uint64) string {
	var name string = ""
	rows, err := self.db.Query("SELECT Name FROM user WHERE ID in(?)", userID)
	defer rows.Close()
	CheckError(err)

	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err.Error())
		}
		break
	}

	return name
}

//获取玩家信息
func (self *SqlMan) CheckUserInfo(userID uint64) (name string, age, sex, vipLevel uint32, money int64) {
	rows, err := self.db.Query("SELECT Name,Age,Sex,Money,VipLevel FROM user WHERE ID in(?)", userID)
	defer rows.Close()
	CheckError(err)

	fMoney := 0.0
	for rows.Next() {
		if err := rows.Scan(&name, &age, &sex, &fMoney, &vipLevel); err != nil {
			log.Fatal(err.Error())
		}
		break
	}
	money = int64(fMoney * 100)
	return name, age, sex, vipLevel, money
}

//扣除金币 注:返回类型最好使用error
func (self *SqlMan) DeductMoney(userID uint64, money int64) (nowMoney int64, isOK bool) {
	rows, err := self.db.Query("SELECT Money FROM user WHERE ID=?", userID)
	defer rows.Close()

	CheckError(err)
	preMoney := 0.0
	isOK = false
	for rows.Next() {
		if err := rows.Scan(&preMoney); err != nil {
			log.Fatal(err.Error())
			return int64(preMoney * 100), isOK
		}
		break
	}

	log.Debug("ID:%v 金额:%v ", userID, int64(preMoney*100))
	nowMoney = int64(preMoney*100) - money
	if nowMoney <= 0 {
		log.Debug("金额成负数了:%v  %v %v", nowMoney/100, money/100, userID)
		return nowMoney, isOK
	}

	isOK = true
	log.Debug("当前:%.3f  扣除%.2f", float64(nowMoney/100), float64(money/100))
	_, err = self.db.Exec("UPdate user set money=? where ID in(?)  ", nowMoney/100, userID)
	CheckError(err)
	return nowMoney, isOK
}

//完成注册 【新增一个玩家】TODO
func (self *SqlMan) AddUser(user, password, adress, port, dbName string) error {
	statement := "INSERT INTO table_name ( field1, field2,...fieldN )VALUES( value1, value2,...valueN );"
	_, err := self.db.Exec(statement)
	CheckError(err)
	return err
}

//获取平台信息
func (self *SqlMan) CheckPlatformInfo(uid uint64) (platformID uint32) {
	rows, err := self.db.Query("SELECT PlatformID FROM user WHERE ID in(?)", uid)
	defer rows.Close()
	CheckError(err)

	for rows.Next() {
		if err := rows.Scan(&platformID); err != nil {
			log.Fatal(err.Error())
		}
		break
	}
	return platformID
}



//数据库连接
func (self *SqlMan) ConnectMySql(user, password, adress, port, dbName string) (*sql.DB, error) {
	dataSourceName := user + ":" + password + "@tcp(" + adress + ":" + port + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open("mysql", dataSourceName)
	self.db = db
	return db, err
}

///////////////////////////////////////////

//////////////////////////////////////////

//////////////////////////////////////////
//func update() {
//	//方式1 update
//	start := time.Now()
//	for i := 1001; i <= 1100; i++ {
//		db.Exec("UPdate userinfo set age=? where uid=? ", i, i)
//	}
//	end := time.Now()
//	fmt.Println("方式1 update total time:", end.Sub(start).Seconds())
//
//	//方式2 update
//	start = time.Now()
//	for i := 1101; i <= 1200; i++ {
//		stm, _ := db.Prepare("UPdate userinfo set age=? where uid=? ")
//		stm.Exec(i, i)
//		stm.Close()
//	}
//	end = time.Now()
//	fmt.Println("方式2 update total time:", end.Sub(start).Seconds())
//
//	//方式3 update
//	start = time.Now()
//	stm, _ := db.Prepare("UPdate userinfo set age=? where uid=?")
//	for i := 1201; i <= 1300; i++ {
//		stm.Exec(i, i)
//	}
//	stm.Close()
//	end = time.Now()
//	fmt.Println("方式3 update total time:", end.Sub(start).Seconds())
//
//	//方式4 update
//	start = time.Now()
//	tx, _ := db.Begin()
//	for i := 1301; i <= 1400; i++ {
//		tx.Exec("UPdate userinfo set age=? where uid=?", i, i)
//	}
//	tx.Commit()
//
//	end = time.Now()
//	fmt.Println("方式4 update total time:", end.Sub(start).Seconds())
//
//	//方式5 update
//	start = time.Now()
//	for i := 1401; i <= 1500; i++ {
//		tx, _ := db.Begin()
//		tx.Exec("UPdate userinfo set age=? where uid=?", i, i)
//		tx.Commit()
//	}
//	end = time.Now()
//	fmt.Println("方式5 update total time:", end.Sub(start).Seconds())
//
//}
//
//func delete() {
//	//方式1 delete
//	start := time.Now()
//	for i := 1001; i <= 1100; i++ {
//		db.Exec("DELETE FROM userinfo WHERE uid=?", i)
//	}
//	end := time.Now()
//	fmt.Println("方式1 delete total time:", end.Sub(start).Seconds())
//
//	//方式2 delete
//	start = time.Now()
//	for i := 1101; i <= 1200; i++ {
//		stm, _ := db.Prepare("DELETE FROM userinfo WHERE uid=?")
//		stm.Exec(i)
//		stm.Close()
//	}
//	end = time.Now()
//	fmt.Println("方式2 delete total time:", end.Sub(start).Seconds())
//
//	//方式3 delete
//	start = time.Now()
//	stm, _ := db.Prepare("DELETE FROM userinfo WHERE uid=?")
//	for i := 1201; i <= 1300; i++ {
//		stm.Exec(i)
//	}
//	stm.Close()
//	end = time.Now()
//	fmt.Println("方式3 delete total time:", end.Sub(start).Seconds())
//
//	//方式4 delete
//	start = time.Now()
//	tx, _ := db.Begin()
//	for i := 1301; i <= 1400; i++ {
//		tx.Exec("DELETE FROM userinfo WHERE uid=?", i)
//	}
//	tx.Commit()
//
//	end = time.Now()
//	fmt.Println("方式4 delete total time:", end.Sub(start).Seconds())
//
//	//方式5 delete
//	start = time.Now()
//	for i := 1401; i <= 1500; i++ {
//		tx, _ := db.Begin()
//		tx.Exec("DELETE FROM userinfo WHERE uid=?", i)
//		tx.Commit()
//	}
//	end = time.Now()
//	fmt.Println("方式5 delete total time:", end.Sub(start).Seconds())
//
//}
//
//func query() {
//
//	//方式1 query
//	start := time.Now()
//	rows, _ := db.Query("SELECT uid,userinfoname FROM userinfo")
//	defer rows.Close()
//	for rows.Next() {
//		var name string
//		var id int
//		if err := rows.Scan(&id, &name); err != nil {
//			//log.Fatal(err)
//		}
//		//fmt.Printf("name:%s ,id:is %d\n", name, id)
//	}
//	end := time.Now()
//	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())
//
//	//方式2 query
//	start = time.Now()
//	stm, _ := db.Prepare("SELECT uid,username FROM userinfo")
//	defer stm.Close()
//	rows, _ = stm.Query()
//	defer rows.Close()
//	for rows.Next() {
//		var name string
//		var id int
//		if err := rows.Scan(&id, &name); err != nil {
//			//log.Fatal(err)
//		}
//
//		// fmt.Printf("name:%s ,id:is %d\n", name, id)
//	}
//	end = time.Now()
//	fmt.Println("方式2 query total time:", end.Sub(start).Seconds())
//
//	//方式3 query
//	start = time.Now()
//	tx, _ := db.Begin()
//	defer tx.Commit()
//	rows, _ = tx.Query("SELECT uid,username FROM userinfo")
//	defer rows.Close()
//	for rows.Next() {
//		var name string
//		var id int
//		if err := rows.Scan(&id, &name); err != nil {
//			//log.Fatal(err)
//		}
//		//fmt.Printf("name:%s ,id:is %d\n", name, id)
//	}
//	end = time.Now()
//	fmt.Println("方式3 query total time:", end.Sub(start).Seconds())
//}
//
//func insert() {
//
//	//方式1 insert
//	//strconv,int转string:strconv.Itoa(i)
//	start := time.Now()
//	for i := 1001; i <= 1100; i++ {
//		//每次循环内部都会去连接池获取一个新的连接，效率低下
//		db.Exec("INSERT INTO userinfo(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
//	}
//	end := time.Now()
//	fmt.Println("方式1 insert total time:", end.Sub(start).Seconds())
//
//	//方式2 insert
//	start = time.Now()
//	for i := 1101; i <= 1200; i++ {
//		//Prepare函数每次循环内部都会去连接池获取一个新的连接，效率低下
//		stm, _ := db.Prepare("INSERT INTO userinfo(uid,username,age) values(?,?,?)")
//		stm.Exec(i, "userinfo"+strconv.Itoa(i), i-1000)
//		stm.Close()
//	}
//	end = time.Now()
//	fmt.Println("方式2 insert total time:", end.Sub(start).Seconds())
//
//	//方式3 insert
//	start = time.Now()
//	stm, _ := db.Prepare("INSERT INTO userinfo(uid,username,age) values(?,?,?)")
//	for i := 1201; i <= 1300; i++ {
//		//Exec内部并没有去获取连接，为什么效率还是低呢？
//		stm.Exec(i, "userinfo"+strconv.Itoa(i), i-1000)
//	}
//	stm.Close()
//	end = time.Now()
//	fmt.Println("方式3 insert total time:", end.Sub(start).Seconds())
//
//	//方式4 insert
//	start = time.Now()
//	//Begin函数内部会去获取连接
//	tx, _ := db.Begin()
//	for i := 1301; i <= 1400; i++ {
//		//每次循环用的都是tx内部的连接，没有新建连接，效率高
//		tx.Exec("INSERT INTO userinfo(uid,username,age) values(?,?,?)", i, "userinfo"+strconv.Itoa(i), i-1000)
//	}
//	//最后释放tx内部的连接
//	tx.Commit()
//
//	end = time.Now()
//	fmt.Println("方式4 insert total time:", end.Sub(start).Seconds())
//
//	//方式5 insert
//	start = time.Now()
//	for i := 1401; i <= 1500; i++ {
//		//Begin函数每次循环内部都会去连接池获取一个新的连接，效率低下
//		tx, _ := db.Begin()
//		tx.Exec("INSERT INTO userinfo(uid,username,age) values(?,?,?)", i, "userinfo"+strconv.Itoa(i), i-1000)
//		//Commit执行后连接也释放了
//		tx.Commit()
//	}
//	end = time.Now()
//	fmt.Println("方式5 insert total time:", end.Sub(start).Seconds())
//}
//
//////////////////////////////////////////////////////////////////////////
/////封装测试用
//////////////////////////////////////////////////////////////////////////
//
////数据库语句执行
//func exec(statement string) ([]string, error) {
//	//筹备mysql语句
//	stmt, err := db.Prepare(statement)
//	if (err != nil) {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	//通过Statement执行查询
//	rows, err := stmt.Query()
//	if err != nil {
//		return nil, err
//	}
//	//建立一个列数组
//	cols, err := rows.Columns()
//	var colsdata = make([]interface{}, len(cols))
//	for i := 0; i < len(cols); i++ {
//		colsdata[i] = new(interface{})
//		//打印字段
//		//fmt.Print(cols[i])
//		//fmt.Print(" ")
//	}
//	printCol(cols)
//
//	//遍历每一行
//	for rows.Next() {
//		rows.Scan(colsdata...) //将查到的数据写入到这行中
//		printRow(colsdata)     //打印此行
//	}
//	defer rows.Close()
//	return nil, err
//}
//
//func printCol(clos []string) {
//	for _, value := range clos {
//		fmt.Printf("%v ", value)
//	}
//	fmt.Println()
//}
//
//func printRow(colsdata []interface{}) {
//	for _, val := range colsdata {
//		//反射
//		switch v := (*(val.(*interface{}))).(type) {
//		case nil:
//			fmt.Print("NULL")
//		case bool:
//			if v {
//				fmt.Print("True")
//			} else {
//				fmt.Print("False")
//			}
//		case []byte:
//			fmt.Print(string(v))
//		case time.Time:
//			fmt.Print(v.Format("2016-01-02 15:05:05"))
//		default:
//			fmt.Print(v)
//		}
//		fmt.Print("\t")
//	}
//	fmt.Println()
//}
