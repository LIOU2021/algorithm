package base64

import (
	rand2 "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	BlockSize = 16
)

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

func EncodeV2(src string) (string, error) {
	cipherText := make([]byte, BlockSize+len(src))
	iv := cipherText[:BlockSize]
	if _, err := io.ReadFull(rand2.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}
	copy(cipherText[BlockSize:], src)
	return EncodeToString(string(cipherText)), nil
}

func DecodeV2(s string) (string, error) {
	cipherText, err := DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	if len(cipherText) < BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	cipherText = cipherText[BlockSize:]

	return cipherText, nil
}
