package xcode

var (
	OK                 = add(0, "OK")
	NoLogin            = add(101, "NOT_LOGIN")
	RequestErr         = add(400, "INVALID_ARGUMENT")
	Unauthorized       = add(401, "UNAUTHENTICATED")
	AccessDenied       = add(403, "PERMISSION_DENIED")
	NotFound           = add(404, "NOT_FOUND")
	MethodNotAllowed   = add(405, "METHOD_NOT_ALLOWED")
	Canceled           = add(498, "CANCELED")
	ServerErr          = add(500, "INTERNAL_ERROR")
	ServiceUnavailable = add(503, "UNAVAILABLE")
	Deadline           = add(504, "DEADLINE_EXCEEDED")
	LimitExceed        = add(509, "RESOURCE_EXHAUSTED")
)

//通用错误码

var (
	TokenIsEmpty     = New(101, "令牌为空")
	TokenFormatErr   = New(102, "令牌格式错误")
	TokenParseErr    = New(103, "令牌解析错误")
	TokenExpired     = New(104, "令牌已过期")
	TokenGenerateErr = New(105, "令牌生成错误")
	TokenInvalid     = New(106, "令牌无效")
)
