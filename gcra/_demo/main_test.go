package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_API(t *testing.T) {

	ln := 20

	for i := 0; i < ln; i++ {

		if i == 15 {
			t.Log("sleep ...")
			time.Sleep(2 * time.Second)
		}
		resp, err := http.Get("http://127.0.0.1:8080?q=2330")
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Error("http status code not equal 200")
			return
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		data := &response{}
		json.Unmarshal(b, data)
		t.Logf("query: %s, status: %s", data.Query, data.Status)
	}
}
