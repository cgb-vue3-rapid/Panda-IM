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
	TokenIsEmpty     = New(101, "token is empty")
	TokenFormatErr   = New(102, "token format error")
	TokenParseErr    = New(103, "token parse error")
	TokenExpired     = New(104, "token expired")
	TokenGenerateErr = New(105, "token generate error")
	TokenInvalid     = New(106, "token invalid")
)
