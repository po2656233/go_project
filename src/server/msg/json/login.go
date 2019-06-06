package json

// 用户登陆协议
type UserLogin struct {
	LoginName 	string // 用户名
	LoginPW 	string   // 密码
}

// 注册协议
type UserRegister struct {
	LoginName string // 用户名
	LoginPW  string // 密码
	Mobiphone string // 手机号
	Email string // 邮箱
}

// 玩家的临时结构
// 玩家有角色的情况
type UserST struct {
	UID string // 账号ID
	ServerID string // 服务器ID
	RoleUID string // 角色UID
	RoleName string // 角色名字
	RoleLev string // 角色等级
	Coin  string // 金币
	// 其他的暂时不做定义
}
type RequestRoomInfo struct {
 	RoomID uint32//房间ID
 	KindID uint32//游戏标识
 	UserID uint64//用户标识
}