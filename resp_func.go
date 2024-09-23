package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

// GetTraceID 获取或生成 Trace ID
func GetTraceID(c *gin.Context) string {
	trace := c.Request.Header.Get(AppTradeName)
	if trace == "" {
		trace = uuid.New().String()
		c.Request.Header.Set(AppTradeName, trace)
	}
	return trace
}

// GetDuration API响应时间
func GetDuration(c *gin.Context) string {
	val, ok := c.Get(AppStartTime)
	if !ok {
		return ""
	}
	start := val.(int64)
	return fmt.Sprintf("%.4f", float64(time.Now().UnixMilli()-start)/1000.0)
}
