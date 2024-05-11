package constants

/*
jwt相关
*/
// 604800
const JwtExpire = 24              // JWT的过期时间
const OrganizationName = " akita" // 签发人

/*
缓存相关
*/

const (
	// 缓存键
	PrefixUserLoginCache = "panda:user:login:id:%s" //  登录态缓存键前缀
	PrefixUserInfoCache  = "panda:user:info:id:%s"  //  用户信息缓存键前缀

	// 缓存键前缀
	PrefixActivation        = "biz#activation#%s"         //  激活码 Redis 键前缀
	PrefixVerificationCount = "biz#verification#count#%s" // 激活码计数的 Redis 键前缀

	// 发送验证码的频率控制
	VerificationLimitPerDay = 100 // 每日发送验证码的上限

	// 缓存过期时间
	PrefixUserInfoCacheExp = 60 * 60 * 24 * 7 //  用户信息缓存过期时间
	VerificationExpire     = 10 * 60          // 验证码过期时间
)

/*
网关相关
*/

const (
	ApiPrefix = ".api" // etcd key的前缀
	RpcPrefix = ".rpc" // etcd key的前缀
)

/*
请求相关
*/
const (
	UserId       = "user_id"
	AccessToken  = "accessToken"
	RefreshToken = "refreshToken"
)
