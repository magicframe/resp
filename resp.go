package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Duration string `json:"duration"`
	Trade    string `json:"trade"`
}

// JsonOK 成功响应
func JsonOK(c *gin.Context, data any) {
	JsonResp(c, ErrSuccess.Code, ErrSuccess.Message, data)
}

// JsonErr 状态码返回
func JsonErr(c *gin.Context, err error) {
	if respErr, ok := err.(*RespErr); ok {
		// 如果是自定义错误类型，返回对应的错误信息
		JsonResp(c, respErr.Code, respErr.Message, nil)
		return
	}

	// 处理其他未知错误，返回通用的服务器错误响应
	JsonResp(c, ErrServerError.Code, ErrServerError.Message, nil)
}

// JsonToast 自定义错误消息
//func JsonToast(c *gin.Context, message string) {
//	JsonResp(c, ErrToast.Code, fmt.Sprintf(ErrToast.Message, message), nil)
//}

func JsonResp(c *gin.Context, code int64, message string, data any) {
	// 记录日志
	// ...
	// 返回结果
	c.AbortWithStatusJSON(http.StatusOK, &Response{
		Code:     code,
		Message:  message,
		Data:     data,
		Duration: GetDuration(c),
		Trade:    GetTraceID(c),
	})
}
