package gcra

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func Test_Grca(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	burst := 5
	rate := 10
	period := time.Second * 10
	testDuration := time.Second * 15

	emissionRate := period.Seconds() / float64(rate)
	rl := NewRateLimiter(rdb, "rate-limit:gcra", time.Second)

	acceptedCount := 0
	reqCount := 0
	reqTime := time.Duration(0)

	start := time.Now()
	ticker := time.NewTicker(time.Millisecond * 10)
	for {
		<-ticker.C

		callStart := time.Now()
		can, err := rl.Take(context.Background(), "random-key", burst, emissionRate)

		reqTime += time.Since(callStart)
		reqCount++

		if err != nil {
			panic(err)
		}

		if can {
			acceptedCount++
			log.Println("#", acceptedCount, ": accepted > ", time.Since(start))
		} else {
			// log.Println("!can #", acceptedCount, ": accepted > ", time.Since(start))
		}

		elapsedTime := time.Since(start)

		if elapsedTime > testDuration {
			break
		}

		if elapsedTime > time.Second*7 && elapsedTime < time.Second*8 {
			// three seconds of rest, it should open up capacity for 3 more requests
			<-time.After(time.Second * 3)
		}
	}

	log.Println("Redis average response time ", time.Duration(reqTime.Nanoseconds()/int64(reqCount)))
}
