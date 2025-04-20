package base64

import (
	"encoding/base64"
	"math/rand"
	"time"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var encodeStd2 = ""

func shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rand2 := rand.New(source)

	// 將 encodeStd 轉為 rune 切片以便操作
	runes := []rune(encodeStd)

	// 使用 Fisher-Yates 洗牌算法打亂順序
	for i := len(runes) - 1; i > 0; i-- {
		j := rand2.Intn(i + 1)                  // 隨機選一個 0 到 i 的索引
		runes[i], runes[j] = runes[j], runes[i] // 交換
	}

	// 將打亂後的 rune 切片轉回字符串
	encodeStd2 = string(runes)
}

func init() {
	shuffle()
}

func GetEncodeStd2() string {
	return encodeStd2
}

func EncodeToString(src string) string {
	return base64.NewEncoding(encodeStd2).EncodeToString([]byte(src))
}

func DecodeString(s string) (string, error) {
	d, err := base64.NewEncoding(encodeStd2).DecodeString(s)
	return string(d), err
}
