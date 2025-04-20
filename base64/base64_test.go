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
