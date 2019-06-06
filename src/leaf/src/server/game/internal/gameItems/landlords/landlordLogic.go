package landlords

import (
	"math/rand"
)

const (
	CardCount = 54
)

var CardListData = [CardCount]byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, //方块 A - K, 小王, 大王
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //梅花 A - K
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //红桃 A - K
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //黑桃 A - K ------
}

// 洗牌
func Shuffle(cards []byte) []byte {
	count := len(cards)
	var index int
	var temp byte
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
	data := make([]byte, count)
	var index int = 0

	if site == 0 {
		site = siteCount
		data[index] = cards[0]
		index++
	}
	for k, v := range cards {
		if site <= k && 0 == k%site {
			data[index] = v
			index++
		}
	}
	return data
}
