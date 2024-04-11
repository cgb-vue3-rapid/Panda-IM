package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// GormConnPoolConf 可以在配置引用
type GormConnPoolConf struct {
	MaxIdleConns int `json:",default=20"`  // 最大空闲连接数
	MaxOpenConns int `json:",default=20"`  // 最大打开连接数
	MaxLifeTime  int `json:",default=300"` // 连接最大生命周期（秒）
}

// GormLogger 自定义日志
type GormLogger struct {
	SlowThreshold time.Duration // 慢查询阈值
	Mode          string        // 模式
}

func NewGormLogger(mode string) *GormLogger {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond, // 一般超过200毫秒就算慢查所以不使用配置进行更改
		Mode:          mode,
	}
}

var _ logger.Interface = (*GormLogger)(nil)

func (l *GormLogger) LogMode(lev logger.LogLevel) logger.Interface {
	return &GormLogger{}
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *GormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 语句和返回条数
	sql, rows := fc()
	// 通用字段
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("time", microsecondsStr(elapsed)),
		logx.Field("rows", rows),
	}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(ctx).Infow("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, logx.Field("catch error", err))
			logx.WithContext(ctx).Errorw("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logx.WithContext(ctx).Sloww("Database Slow Log", logFields...)
	}

	// 非生产模式下，记录所有 SQL 请求
	if l.Mode != service.ProMode {
		logx.WithContext(ctx).Infow("Database Query", logFields...)
	}
}

func microsecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// ConnGorm 连接数据库
func ConnGorm(dataSource, mode string, maxIdleConns, maxOpenConns, maxLifeTime int) *gorm.DB {
	// 使用自定义日志
	conn, err := gorm.Open(
		mysql.Open(dataSource),
		&gorm.Config{Logger: NewGormLogger(mode)},
	)
	if err != nil {
		panic("连接数据库失败:" + err.Error())
		return nil
	}
	connPoolInit(maxIdleConns, maxOpenConns, maxLifeTime, conn)
	logx.Info("连接数据库成功")
	return conn
}

// 连接池初始化
func connPoolInit(maxIdleConns, maxOpenConns, maxLifeTime int, conn *gorm.DB) {
	if sqlDB, err := conn.DB(); err == nil {
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)
	}
}
