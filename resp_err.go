package resp

import "fmt"

type RespErr struct {
	Code    int64
	Message string
}

func (r *RespErr) Error() string {
	return r.Message
}

func NewRespErr(code int64, message string) *RespErr {
	return &RespErr{
		Code:    code,
		Message: message,
	}
}

func ShowRespErr(e *RespErr) *RespErr {
	return &RespErr{
		Code:    e.Code,
		Message: e.Message,
	}
}

// Msg 覆盖message消息
func (r *RespErr) Msg(message string) *RespErr {
	r.Message = message
	return r
}

// MsgArgs 将message的变量替换
func (r *RespErr) MsgArgs(args ...any) *RespErr {
	r.Message = fmt.Sprintf(r.Message, args...)
	return r
}
