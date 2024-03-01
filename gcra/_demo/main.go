package main

import (
	"demo/gcra"
	"demo/rdb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RouterInit(r)
	rdb.RedisInit()
	gcra.DetectInit()

	srv := http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type response struct {
	Query  string `json:"query"`
	Status string `json:"status"`
}

func RouterInit(r *gin.Engine) {
	r.GET("", func(c *gin.Context) {
		resp := &response{}
		q := c.Query("q") // 查询关键字

		if q != "" {
			resp.Query = q
			resp.Status = gcra.Detect(q)
		}
		c.JSON(200, resp)
	})
}
