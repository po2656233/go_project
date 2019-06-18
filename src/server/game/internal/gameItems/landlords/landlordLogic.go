package landlords

import (
	"math/rand"
	"time"
	. "server/game/internal/gameItems" // 注意这里不能这样导入 "../../gameItems" 因为本地导入是根据gopath路径设定的
)

const (
	CardCount = 54
	SiteCount = 3
)

// 牌型
const (
	SINGLE_CARD          int32 = 1 * iota //单牌
	DOUBLE_CARD                           //对子
	THREE_CARD                            //3不带
	BOMB_CARD                             //炸弹
	MISSILE_CARD                          //火箭
	THREE_ONE_CARD                        //3带1
	THREE_TWO_CARD                        //3带2
	BOMB_TWO_CARD                         //四个带2张单牌
	BOMB_TWOOO_CARD                       //四个带2对
	CONNECT_CARD                          //连牌
	COMPANY_CARD                          //连对
	AIRCRAFT_CARD                         //飞机不带
	AIRCRAFT_SINGLE_CARD                  //飞机带单牌
	AIRCRAFT_DOUBLE_CARD                  //飞机带对子
	ERROR_CARD                            //错误的牌型
)

// 每种牌的信息
type CardInfo struct {
	Cards []byte
	Type  int32;
	Value int32 // 0+ 单张 20+对子 40+三张 60+单顺子 80+双顺子 100+炸弹 120+火箭
}

//记录牌值
type CardIndex struct {
	single_index []byte //记录单张的牌
	double_index []byte //记录对子的牌
	three_index  []byte //记录3张
	four_index   []byte //记录4张
}

var CardListData = [CardCount]byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, //方块 A - K, 小王, 大王
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D,             //梅花 A - K
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D,             //红桃 A - K
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D,             //黑桃 A - K ------
}

// 洗牌
func Shuffle(cards []byte) []byte {
	count := len(cards)
	var index int
	var temp byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		index = rand.Int() % count
		temp = cards[i]
		cards[i] = cards[index]
		cards[index] = temp
	}
	return cards
}

// 发牌 参数:牌\座号\座位总数   座位号==索引号
func Deal(cards []byte, site int, siteCount int) []byte {
	if siteCount <= 0 {
		return nil
	}

	count := int(len(cards) / siteCount)
	//if count*siteCount != len(cards){
	//	count++
	//}

	data := make([]byte, count)
	var index int = 0

	if site == 0 {
		site = siteCount
		data[index] = cards[0]
		index++
	}

	for k, v := range cards {
		if site <= k && 0 == k%site {
			if count == index {
				break
			}
			data[index] = v
			index++
		}
	}
	return data
}

// 判断牌型
func JudgeCarType(cards []byte) *CardInfo {
	info := &CardInfo{}
	length := len(cards)
	// ->获取牌值
	info.Cards = SortCards(cards)
	//小于5张牌
	//单牌，对子，3不带,炸弹通用算法
	if 0 < length && length < 5 {
		// 单牌/对子/三不带/炸弹
		if cards[0] == cards[length-1] {
			// ->获取类型
			info.Type = int32(length)
			if length == 4 {
				info.Value = int32(100 + cards[0])
			} else {
				info.Value = int32((length-1)*20) + int32(cards[0])
			}
			return info
		}

		// 火箭
		if cards[0] == 0x0E && cards[1] == 0x0F && length == 2 {
			info.Type = MISSILE_CARD
			info.Value = 120
			return info
		}

		// 三带一
		if (cards[0] == cards[length-2] && length == 4) || (cards[1] == cards[length-1] && length == 4) {
			info.Type = THREE_ONE_CARD
			if cards[0] == cards[length-2] && length == 4 {
				info.Value = 40 + int32(cards[0])
			} else {
				info.Value = 40 + int32(cards[1])
			}
			return info
		}
	} else if length >= 5 {
		// 连牌
		if IsContinuous(cards) && IsLessTwo(cards) {
			info.Type = CONNECT_CARD
			info.Value = 60 + int32(cards[0])
			return info
		}

		// 判断是否大于六张，且为双数
		if length >= 6 && length%2 == 0 {
			// 判断连对 是否都是对子
			is_all_double := true
			for i := 0; i < length; i += 2 {
				if cards[i] != cards[i+1] {
					is_all_double = false
					break
				}
			}
			if is_all_double {
				vec_single := []byte{}
				for i := 0; i < length; i += 2 {
					vec_single = append(vec_single, cards[i])
				}
				if IsContinuous(vec_single) {
					info.Type = COMPANY_CARD
					info.Value = 80 + int32(cards[0])
					return info
				}
			}
		}

		//判断三张以上
		card_index := CardIndex{}

		for i := 0; i < length; {
			if i+1 < length && cards[i] == cards[i+1] {
				if i+2 < length && cards[i+1] == cards[i+2] {
					if i+3 < length && cards[i+2] == cards[i+3] {
						card_index.four_index = append(card_index.four_index, cards[i])
						i += 4
					} else {
						card_index.three_index = append(card_index.three_index, cards[i])
						i += 3
					}
				} else {
					card_index.double_index = append(card_index.double_index, cards[i])
					i += 2
				}
			} else {
				card_index.single_index = append(card_index.single_index, cards[i])
				i++
			}
		}

		// 3带对
		if len(card_index.three_index) == 1 && len(card_index.double_index) == 1 && len(card_index.four_index) == 0 && len(card_index.single_index) == 0 {
			info.Type = THREE_TWO_CARD
			info.Value = 40 + int32(card_index.three_index[0])
			return info
		}

		// 飞机
		// 前提：两个连续三张的
		if len(card_index.four_index) == 0 && len(card_index.three_index) == 2 && card_index.three_index[0]+1 == card_index.three_index[1] {
			// 333444
			if len(card_index.single_index) == 0 && len(card_index.double_index) == 0 {
				info.Type = AIRCRAFT_CARD
				info.Value = 80 + int32(cards[0])
				return info
			}

			// 33344456
			if len(card_index.double_index) == 0 && len(card_index.single_index) == 2 {
				info.Type = AIRCRAFT_SINGLE_CARD
				info.Value = 80 + int32(cards[0])
				return info
			}

			// 33344455
			if len(card_index.single_index) == 0 && len(card_index.double_index) == 1 {
				info.Type = AIRCRAFT_SINGLE_CARD
				info.Value = 80 + int32(cards[0])
				return info
			}

			// 3334445566
			if len(card_index.single_index) == 0 && len(card_index.double_index) == 2 {
				info.Type = AIRCRAFT_DOUBLE_CARD
				info.Value = 80 + int32(cards[0])
				return info
			}
		}

		// 4带2
		if len(card_index.four_index) == 1 && length%2 == 0 && len(card_index.three_index) == 0 {
			// 444423
			if len(card_index.single_index) == 2 && len(card_index.double_index) == 0 {
				info.Type = BOMB_TWO_CARD
				info.Value = 100 + int32(cards[0])
				return info
			}

			// 444422
			if len(card_index.double_index) == 1 && len(card_index.single_index) == 0 {
				info.Type = BOMB_TWO_CARD
				info.Value = 80 + int32(cards[0])
				return info
			}

			// 44442233
			if len(card_index.double_index) == 2 && len(card_index.single_index) == 0 {
				info.Type = BOMB_TWOOO_CARD;
				info.Value = 80 + int32(cards[0])
				return info
			}

		}

	}

	info.Type = ERROR_CARD
	info.Value = 0
	return info
}

//是否连续
func IsContinuous(cards []byte) bool {
	sortCards := SortCards(cards)
	for i := 0; i < len(sortCards); i++ {
		if sortCards[i+1]-sortCards[i] != 1 {
			return false
		}
	}
	return true
}

//是否所有牌值 都小于2
func IsLessTwo(cards []byte) bool {
	for i := 0; i < len(cards); i++ {
		if 0x02 == GetCardValue(cards[i]) || 0x0E == GetCardValue(cards[i]) || 0x0F == GetCardValue(cards[i]){
			return false
		}
	}
	return true
}
