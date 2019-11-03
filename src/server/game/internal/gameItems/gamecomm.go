package gameItems

import (
	"sort"
	"encoding/binary"
	"server/sql/mysql"
	"server/manger"
)

const (
	LOGIC_MASK_COLOR = 0xF0 //花色掩码
	LOGIC_MASK_VALUE = 0x0F //数值掩码
)
var GlobalSqlHandle = mysql.SqlHandle()
var GlobalSender *manger.ClientManger = manger.GetClientManger()
var GlobalPlatformManger *manger.PlatformManger = manger.GetPlatformManger()
var GlobalPlayerManger *manger.PlayerManger = manger.GetPlayerManger()

// 排序
func SortCards(cards []byte) []byte {
	var localCards PlayerCards = cards
	sort.Sort(localCards)
	return localCards
}
func SortCardX(cards []byte) []byte{
	tempCard := []byte{}
	curValue  := byte(0)
	for i := byte(0x03); i < byte(0x14); i++{
		if i == 0x0E || i == 0x0F {
			continue
		}else if i == 0x10 {
			curValue = 0x01
		}else if i == 0x11{
			curValue = 0x02
		}else if i == 0x12 {
			curValue = 0x0E
		}else if i == 0x13 {
			curValue = 0x0F
		}else{
			curValue = i
		}

		for _, v := range cards{
			if GetCardValue(v) == curValue{
				tempCard = append(tempCard,v)

				//花色排序
				colorCard := []byte{}
				for j := 0; j < 4; j++ {
					if 0 <= len(tempCard) - 1 - j && curValue == GetCardValue(tempCard[len(tempCard) - 1 - j]) {
						colorCard = append(colorCard, tempCard[len(tempCard) - 1 - j])
					}
				}

				if 1 < len(colorCard){
					var localCards PlayerCards = colorCard
					sort.Sort(localCards)
					tempCard = append(tempCard[:len(tempCard) - len(colorCard)], colorCard...)
				}

			}
		}
	}
	return tempCard
}

//获取数值
func GetCardValue(cbCardData byte) byte {
	return cbCardData & LOGIC_MASK_VALUE
}
//获取所有数值
func GetCardValues(cbCardData []byte) []byte {
	var cards   []byte
	for _,v :=range cbCardData {
		cards = append(cards,GetCardValue(v))
	}
	return cards
}

//获取花色
func GetCardColor(cbCardData byte) byte {
	return cbCardData & LOGIC_MASK_COLOR
}

//文字转换
func GetCardText(cbCardData byte) string {
	color := GetCardColor(cbCardData)
	value := GetCardValue(cbCardData)
	strTxt := string("")
	switch color {
	case 0x00:
		strTxt = "♦"
	case 0x10:
		strTxt = "♣"
	case 0x20:
		strTxt = "♥"
	case 0x30:
		strTxt = "♠"
	}

	switch value {
	case 0x00:
		return ""
	case 0x01:
		strTxt += "1"
	case 0x02:
		strTxt += "2"
	case 0x03:
		strTxt += "3"
	case 0x04:
		strTxt += "4"
	case 0x05:
		strTxt += "5"
	case 0x06:
		strTxt += "6"
	case 0x07:
		strTxt += "7"
	case 0x08:
		strTxt += "8"
	case 0x09:
		strTxt += "9"
	case 0x0A:
		strTxt += "10"
	case 0x0B:
		strTxt += "J"
	case 0x0C:
		strTxt += "Q"
	case 0x0D:
		strTxt += "K"
	case 0x0E:
		strTxt += "小王"
	case 0x0F:
		strTxt += "大王"
	}
	return strTxt
}
func GetCardsText(cbCardData []byte) string {
	strText := ""
	for _, v := range cbCardData {
		if card := GetCardText(v); card != "" {
			strText += GetCardText(v) + ","
		}

	}
	return strText
}

// 实现排序用
type PlayerCards []byte

//Len()
func (s PlayerCards) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s PlayerCards) Less(i, j int) bool {
	return s[i] < s[j]
}

//Swap()
func (s PlayerCards) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


//获取牌点
func GetCardPip(cbCardData byte) byte {
	//计算牌点
	cbCardValue := GetCardValue(cbCardData)
	var cbPipCount byte = 0
	if cbCardValue < 10 {
		cbPipCount = cbCardValue
	}
	return cbPipCount
}

//获取所有牌的最终点数
func GetCardListPip(cbCardData []byte) byte {
	//变量定义
	var cbPipCount byte = 0

	//获取牌点
	cbCardCount := len(cbCardData)
	for i := 0; i < cbCardCount; i++ {
		cbPipCount = (GetCardPip(cbCardData[i]) + cbPipCount) % 10
	}
	return cbPipCount
}

//整型转bytes数组
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}