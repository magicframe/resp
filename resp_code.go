package resp

var (
	// ErrSuccess 成功
	ErrSuccess = NewRespErr(0, "success")

	// ErrTokenFailed 签名认证错误 10000
	ErrTokenFailed       = NewRespErr(10000, "Token 校验失败")
	ErrTokenEmpty        = NewRespErr(10001, "Token 不能为空")
	ErrTokenBeforeFailed = NewRespErr(10002, "Token 必须在某个时间点后才能使用")
	ErrTokenExp          = NewRespErr(10003, "Token 已过期")
	// ErrTokenType Token刷新
	ErrTokenType = NewRespErr(10010, "RefreshToken 类型错误")
	// ErrAuthKeyFailed AuthKey验证
	ErrAuthKeyFailed       = NewRespErr(10020, "AuthKey 验证失败")
	ErrAuthKeyEmpty        = NewRespErr(10021, "AuthKey 不能为空")
	ErrAuthKeyBeforeFailed = NewRespErr(10022, "AuthKey 必须在某个时间点后才能使用")
	ErrAuthKeyExp          = NewRespErr(10023, "AuthKey 已过期")

	// ErrBadRequest 请求错误 10400
	ErrBadRequest           = NewRespErr(10400, "请求参数错误")
	ErrUnauthorized         = NewRespErr(10401, "访问受限，请登录后重试")
	ErrForbidden            = NewRespErr(10403, "您无权访问此资源")
	ErrNotFound             = NewRespErr(10404, "URL地址错误")
	ErrMethodNotAllowed     = NewRespErr(10405, "请求方法错误")                   // 例如post接口用get请求
	ErrNotAcceptable        = NewRespErr(10406, "服务器无法提供您请求的内容类型，请检查并修改请求") // 例如，服务器可能只支持 application/xml，但客户端要求 application/json
	ErrProxyAuthRequired    = NewRespErr(10407, "请求被拒绝，请通过代理服务器进行身份验证")
	ErrRequestTimeout       = NewRespErr(10408, "请求超时，请稍后重试")
	ErrConflict             = NewRespErr(10409, "不能重复请求")             // 在多用户环境中，两个客户端几乎同时尝试更新同一资源，导致冲突。
	ErrGone                 = NewRespErr(10410, "请求的内容已被删除")          // 服务器上的某个资源（如文件、页面、API 端点）已被删除，且服务器不再提供该资源
	ErrLengthRequired       = NewRespErr(10411, "请求信息不完整，请提供必要的长度信息") // 没有在请求中包含 Content-Length 头
	ErrPreconditionFailed   = NewRespErr(10412, "无法完成请求，因为资源状态已发生变化")
	ErrPayloadTooLarge      = NewRespErr(10413, "请求的内容过大")
	ErrURITooLong           = NewRespErr(10414, "请求的地址过长")
	ErrUnsupportedMediaType = NewRespErr(10415, "请求的文件类型不支持，请更换格式")
	ErrRangeNotSatisfiable  = NewRespErr(10416, "无法满足请求，因为请求的范围超出了可用内容")
	ErrExpectationFailed    = NewRespErr(10417, "服务器未能满足您的请求要求，请检查请求并重试")
	ErrTooManyRequest       = NewRespErr(10429, "请求过于频繁")

	// ErrServerError 服务器错误 10500
	ErrServerError        = NewRespErr(10500, "服务器错误，请稍后再试")
	ErrNotImplemented     = NewRespErr(10501, "服务器未能处理该请求，请稍后重试")
	ErrBadGateway         = NewRespErr(10502, "网关错误，请稍后重试")
	ErrServiceUnavailable = NewRespErr(10503, "功能维护中，请稍后再试")
	ErrGatewayTimeout     = NewRespErr(10504, "请求超时，请稍后再试")
)

var (
	ErrToast = NewRespErr(20000, "%s")

	// ErrUserNotFound 用户 20100
	ErrUserNotFound  = NewRespErr(20100, "该用户不存在")
	ErrInvalidUserID = NewRespErr(20101, "用户ID错误")
	ErrUserDel       = NewRespErr(20102, "用户 %s 已经被删除 %s 测试")
)
