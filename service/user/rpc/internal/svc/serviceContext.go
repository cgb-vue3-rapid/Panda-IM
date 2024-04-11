package svc

import (
	"akita/panda-im/common/orm"
	"akita/panda-im/service/user/rpc/internal/config"
	"akita/panda-im/service/user/rpc/models"
	"akita/panda-im/service/user/rpc/models/entity"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	Orm          *gorm.DB
	BizRedis     *redis.Redis
	UserModel    entity.UserModel
	UserModelDao *entity.UserModelDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	// todo 初始化其他模块
	conn := orm.ConnGorm(c.DB.DataSource, c.DB.Mode, c.DB.MaxIdleConns, c.DB.MaxOpenConns, c.DB.MaxLifeTime)
	err := models.Migrate(conn)
	if err != nil {
		conn.Logger.Error(context.Background(), "Auth服务数据库迁移失败", err)
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		Orm:          conn,
		UserModelDao: entity.NewUserModelDao(conn),
		BizRedis:     redis.MustNewRedis(c.BizRedis),
	}
}
