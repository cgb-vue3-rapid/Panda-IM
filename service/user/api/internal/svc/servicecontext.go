package svc

import (
	"akita/panda-im/common/orm"
	"akita/panda-im/common/util/interceptors"
	"akita/panda-im/service/auth/models/dao"
	"akita/panda-im/service/user/api/internal/config"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	Orm      *gorm.DB
	BizRedis *redis.Redis
	//UserModel    entity.UserModel
	//UserModelDao *entity.UserModelDao
	UserRPC user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := orm.ConnGorm(c.DB.DataSource, c.DB.Mode, c.DB.MaxIdleConns, c.DB.MaxOpenConns, c.DB.MaxLifeTime)
	err := dao.Migrate(conn)
	if err != nil {
		conn.Logger.Error(context.Background(), "Auth服务数据库迁移失败", err)
		panic(err)
	}

	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config: c,
		Orm:    conn,
		//UserModelDao: entity.NewUserModelDao(conn),
		BizRedis: redis.MustNewRedis(c.BizRedis),
		UserRPC:  user.NewUser(userRPC),
	}
}
