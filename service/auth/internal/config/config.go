package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Token struct {
		AccessSecret  string
		RefreshSecret string
	}
	//Auth struct {
	//	AccessSecret string
	//	AccessExpire int64
	//}
	ETCD struct {
		Endpoints string
		TTL       int64
	}
	BizRedis redis.RedisConf
	UserRPC  zrpc.RpcClientConf
}
