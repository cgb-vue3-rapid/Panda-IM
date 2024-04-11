package svc

import (
	"akita/panda-im/common/interceptors"
	"akita/panda-im/service/auth/internal/config"
	"akita/panda-im/service/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//Orm      *gorm.DB
	BizRedis *redis.Redis
	UserRPC  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	// todo 初始化其他模块
	//conn := orm.ConnGorm(c.DB.DataSource, c.DB.Mode, c.DB.MaxIdleConns, c.DB.MaxOpenConns, c.DB.MaxLifeTime)
	//err := dao.Migrate(conn)
	//if err != nil {
	//	conn.Logger.Error(context.Background(), "Auth服务数据库迁移失败", err)
	//	panic(err)
	//}

	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config: c,
		//Orm:      conn,
		BizRedis: redis.MustNewRedis(c.BizRedis),
		UserRPC:  user.NewUser(userRPC),
	}
}
