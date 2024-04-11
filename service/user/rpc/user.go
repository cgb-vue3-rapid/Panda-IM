package main

import (
	"akita/panda-im/common/interceptors"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logc"

	"akita/panda-im/service/user/rpc/internal/config"
	"akita/panda-im/service/user/rpc/internal/server"
	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	logc.Infof(context.Background(), "User服务启动成功，运行在 %s\n", c.ListenOn)
	s.Start()
}
