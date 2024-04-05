package main

import (
	"akita/panda-im/service/auth/internal/config"
	"akita/panda-im/service/auth/internal/handler"
	"akita/panda-im/service/auth/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//// 自定义错误方法
	//httpx.SetErrorHandler(xcode.ErrHandler)
	logc.Infof(context.Background(), "Auth服务启动成功，运行在 %s:%d...\n", c.Host, c.Port)
	server.Start()
}
