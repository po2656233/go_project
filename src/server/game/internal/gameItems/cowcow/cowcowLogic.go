package cowcow

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
	"leaf/log"
)

const (
	//索引
	INDEX_Banker    	= (iota * PiceCount)
	INDEX_Tian
	INDEX_Xuan
	INDEX_Di
	INDEX_Huang

	//下注区域
	AREA_Banker 	= 1 // 不允许下注
	AREA_Tian 		= 2
	AREA_Xuan 		= 3
	AREA_Di 		= 4
	AREA_Huang 	= 5
	AREA_MAX   	= 6//最大区域

	//区域倍数multiple
	MULTIPLE_NULL	 = 0.0
	MULTIPLE_ONE   = 2.0
	MULTIPLE_TWO   = 2.0
	MULTIPLE_THREE = 2.0
	MULTIPLE_FOUR  = 2.0
	MULTIPLE_FIVE  = 2.0
	MULTIPLE_SIX 	= 6.0
	MULTIPLE_SEVEN 	= 7.0
	MULTIPLE_EIGHT 	= 8.0
	MULTIPLE_NINE 	= 9.0
	MULTIPLE_TEN 	= 10.0


	//牌数
	PiceCount			= 5 //每人的牌数
	CardCount 			= 52//总牌数(不带大小王)

	Lose = 0
	Win  = 1
)

//混乱扑克
func RandCardList(cbBufferCount int) []int {
	if cbBufferCount <= 0 {
		return []int{0}
	}

	//混乱准备
	var (
		cbCardBuffer            []int = make([]int, cbBufferCount)
		tempCardListData        [CardCount]int
		cbRandCount, cbPosition int = 0, 0
	)
	copy(tempCardListData[:CardCount], CardListData[:CardCount])

	//混乱扑克
	for {
		rand.Seed(time.Now().Unix())
		cbPosition = rand.Intn(CardCount - cbRandCount)
		cbCardBuffer[cbRandCount] = tempCardListData[cbPosition]
		cbRandCount++
		tempCardListData[cbPosition] = tempCardListData[CardCount-cbRandCount]
		if cbRandCount >= cbBufferCount {
			break
		}
	}
	return cbCardBuffer
}

//----------------------------------
type WinType struct {
	Name string
	Type int
	Odds float64
}

type Game struct {
	BetType       []*BetType
	SameAsOdds    float64
	WinType       []*WinType
	BetWinInfoMap *BetWinInfoMap
}

func (g *Game) Init() {
	g.SameAsOdds = 1.0

	//下注类型
	g.AddBetType("天", AREA_Tian, MULTIPLE_TEN)
	g.AddBetType("地", AREA_Di, MULTIPLE_TEN)
	g.AddBetType("玄", AREA_Xuan, MULTIPLE_TEN)
	g.AddBetType("黄", AREA_Huang, MULTIPLE_TEN)

	//赔率
	g.AddWinType("五小牛", 13, MULTIPLE_TEN)
	g.AddWinType("炸弹牛", 12, MULTIPLE_TEN)
	g.AddWinType("五花牛", 11, MULTIPLE_TEN)
	g.AddWinType("牛牛", 10, MULTIPLE_TEN)
	g.AddWinType("牛九", 9, MULTIPLE_NINE)
	g.AddWinType("牛八", 8, MULTIPLE_EIGHT)
	g.AddWinType("牛七", 7, MULTIPLE_SEVEN)
	g.AddWinType("牛六", 6, MULTIPLE_SIX)
	g.AddWinType("牛五", 5, MULTIPLE_FIVE)
	g.AddWinType("牛四", 4, MULTIPLE_FOUR)
	g.AddWinType("牛三", 3, MULTIPLE_THREE)
	g.AddWinType("牛二", 2, MULTIPLE_TWO)
	g.AddWinType("牛一", 1, MULTIPLE_ONE)
	g.AddWinType("无牛", 0, MULTIPLE_ONE)

}

// 获取最大的赔付金额
//func (g *Game) HasLockMoneyByOdds(betInfo *BetInfo) float64 {
//	var moneyCount float64 = 0
//	//var info = betInfo.GetAll()
//	//for k1 := range info {
//	//	for k2 := range info[k1] {
//	//		moneyCount += info[k1][k2].Odds * info[k1][k2].BetMoney
//	//	}
//	//}
//	return moneyCount
//}

// 添加投注类型
func (g *Game) AddBetType(name string, betType int, odds float64) {
	g.BetType = append(g.BetType, &BetType{Name: name, Type: betType, Odds: odds})
}

// 添加赢钱类型
func (g *Game) AddWinType(name string, winType int, odds float64) {
	g.WinType = append(g.WinType, &WinType{Name: name, Type: winType, Odds: odds})
}

// 根据代号获取类型
func (g *Game) GetBetInfoByType(betType int) (*BetType, error) {
	for _, v := range g.BetType {
		if v.Type == betType {
			return v, nil
		}
	}
	return nil, fmt.Errorf("code is not exists")
}

func (g *Game) GetWinInfoByWinType(winType int) (*WinType, error) {
	for _, v := range g.WinType {
		if v.Type == winType {
			return v, nil
		}
	}
	return nil, fmt.Errorf("win is not exists:%v",winType)
}

// 获取赔率最大的类型
func (g *Game) GetMaxOddsType() *BetType {
	var max = 0.0
	var r *BetType
	for _, v := range g.BetType {
		if v.Odds > max {
			max = v.Odds
			r = v
		}
	}
	return r
}

func (g *Game) Compare(betType int, bankerType PokerType, playerType PokerType) *BetWinInfo {

	var betWinInfo = &BetWinInfo{}

	betTypeInfo, _ := g.GetBetInfoByType(betType)

	if IsBankerWin(bankerType, playerType) {
		// 如果庄家赢了 根据相应赔率算钱
		// 获取专家牌型的赔率
		// log.Printf("%+v",bankerType)
		isWin, err := g.GetWinInfoByWinType(bankerType.Type)
		if err!=nil{
			log.Debug(err.Error())
			return nil
		}
		betWinInfo = &BetWinInfo{
			LoseOdds: isWin.Odds,
			WinOdds:  0,
			BetType:  betTypeInfo,
			IsWin:    false,
		}
	} else {
		// log.Printf("%+v",playerType)
		isWin, err := g.GetWinInfoByWinType(playerType.Type)
		if err!=nil{
			log.Debug(err.Error())
			return nil
		}
		betWinInfo = &BetWinInfo{
			LoseOdds: 0,
			WinOdds:  isWin.Odds,
			BetType:  betTypeInfo,
			IsWin:    true,
		}
	}

	return betWinInfo
}

// CalcPoker CalcPoker
func CalcPoker(pokers Pokers) PokerType {

	var pokerType = PokerType{}

	var p1 = pokers.List[0]
	var p2 = pokers.List[1]
	var p3 = pokers.List[2]
	var p4 = pokers.List[3]
	var p5 = pokers.List[4]

	// 成员
	pokerType.Member = pokers.List

	// 最大牌
	pokerType.MaxPoint = p1.Number

	// 最大牌的花色
	pokerType.MaxColor = p1.Color

	// 最大牌的原始点数
	pokerType.MaxRaw = p1.Raw

	// 值
	pokerType.Value = p1.Number

	pokerType.ValueColor = p1.Color

	// 统计每个数字出现的个数
	var mapSlice = make(map[int]int)

	// 余数
	var leave = 0

	for i := 0; i < 5; i++ {
		leave += pokers.List[i].Mark
		if _, ok := mapSlice[pokers.List[i].Number]; ok {
			mapSlice[pokers.List[i].Number]++
		} else {
			mapSlice[pokers.List[i].Number] = 1
		}
	}

	leave = leave % 10

	// 五小牛
	if p1.Mark+p2.Mark+p3.Mark+p4.Mark+p5.Mark < 10 {
		if p1.Mark < 5 {
			pokerType.Type = 13
			return pokerType
		}
	}

	// 炸弹牛
	for k, v := range mapSlice {
		if v == 4 {
			pokerType.Value = k
			pokerType.Type = 12
			return pokerType
		}
	}

	// 五花牛
	if p1.Number > 10 && p2.Number > 10 && p3.Number > 10 && p4.Number > 10 && p5.Number > 10 {
		pokerType.Type = 11
		return pokerType
	}

	// 牛-?
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			// log.Println(pokers.List[i].Mark + pokers.List[j].Mark)
			if (pokers.List[i].Mark+pokers.List[j].Mark)%10 == leave {
				// 有牛
				if leave == 0 {
					// 牛牛
					pokerType.Type = 10
					return pokerType
				} else {
					// 牛X
					pokerType.Type = leave
					return pokerType
				}
			}
		}
	}

	pokerType.Type = 0
	return pokerType
}

// IsBankerWin IsBankerWin
func IsBankerWin(bankerType PokerType, playerType PokerType) bool {

	// 直接对比类型 如果类型相等 比值的大小 值的大小相等 比花色 花色相等 庄家赢
	if bankerType.Type > playerType.Type {
		return true
	}

	if bankerType.Type < playerType.Type {
		return false
	}

	if bankerType.Type == playerType.Type {

		// 炸弹牛直接比值
		if bankerType.Type == 12 {
			if bankerType.Value > playerType.Value {
				return true
			}
			if bankerType.Value < playerType.Value {
				return false
			}
			if bankerType.Value == playerType.Value {
				return true
			}
		}

		// 不是炸弹牛 比最大的牌的点和花色
		if bankerType.MaxPoint > playerType.MaxPoint {
			return true
		}
		if bankerType.MaxPoint < playerType.MaxPoint {
			return false
		}
		if bankerType.MaxPoint == playerType.MaxPoint {
			return true
		}
	}

	return true
}

func removePoker(pokers []Poker, value int) []Poker {
	if len(pokers) == 0 {
		return pokers
	}
	for i, v := range pokers {
		if v.Points == value {
			pokers = append(pokers[:i], pokers[i+1:]...)
			return removePoker(pokers, value)
		}
	}
	return pokers
}



//------------Poker------------------------引入

type BetType struct {
	Name string  `json:"name" bson:"name"`
	Type int     `json:"type" bson:"type"`
	Odds float64 `json:"odds" bson:"odds"`
}

type BetWinInfoMap struct {
	WinInfoMap map[string]map[int]*BetWinInfo
	mux        sync.RWMutex
}

type BetWinInfo struct {
	LoseOdds float64
	WinOdds  float64
	IsWin    bool
	BetType  *BetType
}

func (m *BetWinInfoMap) Init() {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap = make(map[string]map[int]*BetWinInfo)
}

func (m *BetWinInfoMap) Set(sourceName string, betType int, betWinInfo *BetWinInfo) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap[sourceName][betType] = betWinInfo
}

func (m *BetWinInfoMap) Get(sourceName string, betType int) (*BetWinInfo, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if _, ok := m.WinInfoMap[sourceName][betType]; ok {
		return m.WinInfoMap[sourceName][betType], nil
	}
	return nil, fmt.Errorf("%d is not found", betType)
}

func (m *BetWinInfoMap) InitSourceMap(sourceName string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap[sourceName] = make(map[int]*BetWinInfo)
}

func (m *BetWinInfoMap) GetMap(sourceName string) (map[int]*BetWinInfo, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if _, ok := m.WinInfoMap[sourceName]; ok {
		return m.WinInfoMap[sourceName], nil
	}
	return nil, fmt.Errorf("%s is not found", sourceName)
}

// Poker Poker
type Poker struct {
	// 1-13 牌面
	Number int
	// 1-方块 2-梅花 3-红桃 4-黑桃
	Color int
	// 2-14 点数
	Points int
	// Mark JQK 为10点
	Mark int
	// Raw 原始
	Raw int
}

// Pokers Pokers
type Pokers struct {
	List []Poker
}

// PokerType PokerType
type PokerType struct {
	// 值
	Value      int
	ValueColor int
	// 第一位:9 豹子, 8 顺子, 7 对子, 6单张
	Type int
	// 第二位:最大点数
	MaxPoint int
	// 第三位:最大点数的花色
	MaxColor int
	// 最大牌值
	MaxRaw int
	// 第四位:是否是同花
	SameColor bool
	// 成员
	Member []Poker
}

// AddPokers AddPokers
func (p *Pokers) AddPokers(pokers ...Poker) *Pokers {
	p.List = append(p.List, pokers...)
	return p
}

// ArrangeByPoints ArrangeByPoints
func (p *Pokers) ArrangeByPoints() []Poker {
	for i := 0; i < len(p.List); i++ {
		for j := i + 1; j < len(p.List); j++ {
			if p.List[i].Points < p.List[j].Points {
				p.List[j], p.List[i] = p.List[i], p.List[j]
			}
			if p.List[i].Points == p.List[j].Points {
				if p.List[i].Color < p.List[j].Color {
					p.List[j], p.List[i] = p.List[i], p.List[j]
				}
			}
		}
	}
	return p.List
}

// ArrangeByNumber ArrangeByNumber
func (p *Pokers) ArrangeByNumber() []Poker {
	for i := 0; i < len(p.List); i++ {
		for j := i + 1; j < len(p.List); j++ {
			if p.List[i].Number < p.List[j].Number {
				p.List[j], p.List[i] = p.List[i], p.List[j]
			}
			if p.List[i].Number == p.List[j].Number {
				if p.List[i].Color < p.List[j].Color {
					p.List[j], p.List[i] = p.List[i], p.List[j]
				}
			}
		}
	}
	return p.List
}

// CreatePoker
func CreatePoker(list []int) []Poker {

	var pokerList []Poker

	for _, v := range list {

		// 牌面
		var number = int(v % 16)
		if number == 0 {
			number = 13
		}

		// 点数
		var point = int(v % 16)
		if point == 0 {
			point = 13
		}
		//if point == 1 {//A作为10时可解开
		//	point = 14
		//}

		var mark = point
		if mark >= 11 && mark <= 16 {
			mark = 10
		}
		if mark == 14 {
			mark = 1
		}

		var color = int((v-1)/16 + 1)

		var poker = Poker{Number: number, Color: color, Points: point, Mark: mark, Raw: int(v)}

		pokerList = append(pokerList, poker)
	}

	return pokerList
}




//大数定义统一置文本后
var CardListData = [CardCount]int{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //方块 A - K
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //梅花 A - K
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //红桃 A - K
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //黑桃 A - K ------
}

//炸弹——————————6 倍与压注筹码
//全花牌牛牛—————— 5 倍与压注筹码
//四张花牌牛牛—————4 倍与压注筹码
//牛牛—————————3 倍与压注筹码
//牛7、牛8、牛9－－－－2 倍与压注筹码
//无牛—————————1 倍与压注筹码


//数字比较： k>q>j>10>9>8>7>6>5>4>3>2>a。
//花色比较：黑桃>红桃>梅花>方块。
//牌型比较：无牛<有牛<牛牛<银牛<金牛<炸弹<五小牛。
//无牛牌型比较：取其中最大的一张牌比较大小，牌大的赢，大小相同比花色。
//有牛牌型比较：比牛数；牛数相同庄吃闲。
//牛牛牌型比较：取其中最大的一张牌比较大小，牌大的赢，大小相同比花色。
//银牛牌型比较：取其中最大的一张牌比较大小，牌大的赢，大小相同比花色。
//金牛牌型比较：取其中最大的一张牌比较大小，牌大的赢，大小相同比花色。
//炸弹之间大小比较：取炸弹牌比较大小。
//五小牛牌型比较：庄吃闲。


//名称
//赔率（闲家下注）
//说 明
//有牛
//1~2倍
//五张牌中有三张的点数之和为10点的整数倍，并且另外两张牌之和与10进行取余，所得之数即为牛几。例如: 2、8、j、6、3，即为牛9。牛一到牛6为1倍，牛七到牛九位2倍！
//无牛
//五张牌中没有任意三张牌点数之和为10的整数倍。例如: a、8、4、k、7。
//牛牛
//3倍
//五张牌中第一组三张牌和第一组二张牌之和分别为10的整数倍。 3、7、k、10、j。
//银牛
//4倍
//五张牌全由10～k组成且只有一张10，例如10、j、j、q、k。
//金牛
//5倍
//五张牌全由j～k组成，例如j、j、q、q、k。
//炸弹
//6倍
//五张牌中有4张牌点数相同的牌型，例如：2、2、2、2、k。
//五小牛
//10倍
//五张牌的点数加起来小于10，且每张牌点数都小于5，例如a、3、2、a、2。