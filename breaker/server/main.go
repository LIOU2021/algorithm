package main

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var flag atomic.Uint32

func main() {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		// v := flag.Add(1)
		// if v%2 == 0 {
		// 	ctx.JSON(200, "success")
		// } else {
		// 	ctx.JSON(http.StatusTooManyRequests, "too many request")
		// }
		ctx.JSON(http.StatusTooManyRequests, "too many request")
	})

	r.Run(":8081")
}
