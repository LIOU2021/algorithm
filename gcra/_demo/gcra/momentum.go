package gcra

import (
	"context"
	"demo/rdb"
	"log"
	"sync"
	"time"
)

var rl *RateLimit
var emissionRate float64 // 速率。产一个token需要几秒
var DetectInitOnce = sync.Once{}

func DetectInit() {
	DetectInitOnce.Do(func() {
		// n秒内产x个token
		rate := 10                 // x
		period := time.Second * 10 // n

		emissionRate = period.Seconds() / float64(rate)
		rl = NewRateLimiter(rdb.Client, "rate-limit:gcra", time.Second)
	})
}

func Detect(query string) string {
	level1 := 5  // 上升
	level2 := 10 // 急升

	can1, err1 := rl.Take(context.Background(), query, level2, emissionRate)
	if err1 != nil {
		log.Fatal(err1)
	}

	if !can1 {
		return "急升"
	}

	can2, err2 := rl.Take(context.Background(), query, level1, emissionRate)
	if err2 != nil {
		log.Fatal(err2)
	}

	if !can2 {
		return "上升"
	}

	return ""
}
