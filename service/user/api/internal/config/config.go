package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ETCD struct {
		Endpoints string
		TTL       int64
	}
	Token struct {
		AccessSecret  string
		RefreshSecret string
	}
	//Auth struct {
	//	AccessSecret string
	//	AccessExpire int64
	//}
	DB struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`   // 最大连接数
		MaxIdleConns int `json:",default=100"`  //  最大空闲连接数
		MaxLifeTime  int `json:",default=3600"` // 连接最大存活时间
		Mode         string
	}
	BizRedis redis.RedisConf
	UserRPC  zrpc.RpcClientConf
}
