package main

import (
	"akita/panda-im/common/etcdOp"
	"akita/panda-im/common/xcode"
	"akita/panda-im/service/user/api/internal/config"
	"akita/panda-im/service/user/api/internal/handler"
	"akita/panda-im/service/user/api/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 上送etcd服务
	etcdOp.RegisterEtcd(c.ETCD.Endpoints, c.Name, fmt.Sprintf("%s:%d", c.Host, c.Port), c.ETCD.TTL)

	// 自定义错误方法
	httpx.SetErrorHandler(xcode.ErrHandler)
	logc.Infof(context.Background(), "UserAPI服务启动成功，运行在 %s:%d...\n", c.Host, c.Port)

	server.Start()
}
