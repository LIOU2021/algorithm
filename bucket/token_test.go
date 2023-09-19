package bucket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func Test_Bucket_Token(t *testing.T) {
	// 3秒產1個token，bucket最多存放3個token
	rateLimit := NewIPRateLimiter(rate.Every(time.Second*3), 3)
	limiter := rateLimit.GetLimiter("127.0.0.1")

	for i := 0; i < 3; i++ {
		ok := limiter.Allow()
		if i < 3 {
			assert.Equal(t, true, ok)
		} else {
			assert.Equal(t, false, ok)
		}
	}

}
