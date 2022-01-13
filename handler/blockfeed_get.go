package handler

import (
	"go-block-web/platform/blockfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlockfeedGet(feed blockfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
