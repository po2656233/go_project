package base

//MainID 200+ || SubID 100+

const (
	//注册
	MainRegister = 200 //-> 注册
	SubRegister  = 101 //-> 注册

	//登录
	MainLogin          = 201 //-> 登录
	SubLoginResult     = 101 //-> 登录结果
	SubMasterInfo      = 102 //-> 主页信息
	SubGameList        = 103 //-> 游戏列表
	SubEnterRoomResult = 104 //进入房间结果
	SubEnterGameResult = 105 //进入游戏结果

	//场景
	MainGameSence       = 300 //游戏场景
	SubGameSenceStart   = 101 //起始
	SubGameSencePlaying = 102 //游戏中
	SubGameSenceOver    = 103 //游戏结束
	SubGameSenceFree    = 104 //空闲

	//框架
	MainGameFrame         = 400 //游戏框架
	SubGameFrameStart     = 101 //开始
	SubGameFramePlaying   = 102 //游戏中(下注)
	SubGameFrameOver      = 103 //开奖
	SubGameFrameBetResult = 104 //下注结果
	SubGameFrameCheckout  = 105 //下注结果
	SubGameFrameSetHost   = 106 //定庄
	SubGameFrameHost      = 107 //抢庄
	SubGameFrameSuperHost = 108 //超级抢庄
	SubGameFrameResult    = 109 //结果信息
	SubGameFrameReady     = 110 //准备

	MainGameState       = 401 //游戏状态
	SubGameStateStart   = 101 //开始
	SubGameStatePlaying = 102 //游戏中(下注)
	SubGameStateOver    = 103 //开奖
    SubGameStateCall 	= 104	//叫分

//状态
	MainPlayerState = 500 //玩家状态
	PlayerSitDown   = 1   //坐下
	PlayerAgree     = 2   //同意
	PlayerAction    = 3   //游戏
	PlayerStandUp   = 4   //站起
	PlayerLookOn    = 5   //旁观

	//更新标识
	MainGameUpdate       = 501 //游戏状态
	GameUpdatePlayerList = 0   //玩家列表
	GameUpdateHost       = 1   //玩家抢庄
	GameUpdateSuperHost  = 2   //玩家超级抢庄
	GameUpdateOut        = 3   //玩家出场
	GameUpdateOffline    = 4   //玩家离线
	GameUpdateReconnect  = 5   //玩家重连
	GameUpdateReady      = 6   //玩家准备

	//源码中重要标识
	INVALID = 0 //无效(切记有效初始化,不要从零开始)
	FAILD   = 0 //失败
	SUCCESS = 1 //成功

	//房间级别
	RoomGeneral = 0 //普通
	RoomMiddle  = 1 //中级
	RoomHigh    = 2 //高级

	//游戏kindID
	Baccarat  = 2001
	FishLord  = 3003
	Landlords = 3001
	CowCow    = 1001
	Mahjong   = 3002
	ChinessChess = 8002
)
