package main

import (
	"go-block-web/httpd/handler"
	"go-block-web/platform/blockfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := blockfeed.New()

	r := gin.Default()

	r.GET("/web", handler.BlockTest())
	r.GET("/blockfeed", handler.BlockfeedGet(feed))
	r.POST("/blockfeed", handler.BlockfeedPost(feed))

	r.Run()

	// Console new block feed print method
	// feed := blockfeed.New()
	// fmt.Println(feed)

	// Example add block
	// feed.Add(blockfeed.Block{"1", "Fake block added"})
	// fmt.Println(feed)
}
