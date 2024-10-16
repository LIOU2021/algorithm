package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("打到error-1")
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("打到error-2")
			return nil, err
		}

		if resp.StatusCode != 200 { // 我自己加的，方便测试错误
			return nil, fmt.Errorf("status code not equal 200 : %d", resp.StatusCode)
		}
		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {
	for i := 0; i < 50; i++ {
		body, err := Get("http://127.0.0.1:8081/ping")
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}

		fmt.Println(string(body), cb.State())
	}

}
