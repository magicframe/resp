package resp

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RespInitMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(AppStartTime, time.Now().UnixMilli())
	}
}
