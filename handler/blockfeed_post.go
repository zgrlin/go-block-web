package handler

import (
	"go-block-web/platform/blockfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type blockfeedPostRequest struct {
	Height  string `json:"number"`
	Minedby string `json:"minedby"`
}

func BlockfeedPost(feed blockfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := blockfeedPostRequest{}
		c.Bind(&requestBody)

		block := blockfeed.Block{
			Height:  requestBody.Height,
			Minedby: requestBody.Minedby,
		}
		feed.Add(block)
		c.Status(http.StatusNoContent)
	}
}
