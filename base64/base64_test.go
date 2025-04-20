package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Base64(t *testing.T) {
	assert.NotEmpty(t, encodeStd2)
	assert.Equal(t, encodeStd2, GetEncodeStd2())

	plainText := "hello我，你好，==是jld，握寿司体育运动，,."
	encodeText := EncodeToString(plainText)
	t.Log(encodeText)
	result, err := DecodeString(encodeText)

	assert.NoError(t, err)
	assert.Equal(t, plainText, string(result))
}

func Test_Base64V2(t *testing.T) {
	plainText := "hello我，你好，==是jld，握寿司体育运动，,."
	duplicateMap := make(map[string]bool)
	for i := 0; i < 10; i++ {
		encodeText, err := EncodeV2(plainText)

		value, exists := duplicateMap[encodeText]
		assert.False(t, exists)
		assert.False(t, value)

		assert.NoError(t, err)
		t.Log(encodeText)
		result, err := DecodeV2(encodeText)
		assert.NoError(t, err)
		assert.Equal(t, plainText, string(result))
	}
}
