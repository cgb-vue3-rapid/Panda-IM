package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`   // 最大连接数
		MaxIdleConns int `json:",default=100"`  //  最大空闲连接数
		MaxLifeTime  int `json:",default=3600"` // 连接最大存活时间
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
