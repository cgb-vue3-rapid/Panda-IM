package etcdOp

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type RegisterInfomation struct {
	Cil *clientv3.Client
}

type EtcdOp interface {
	Put(serviceName, value string, ttl int64) error
	Get(key string) *clientv3.GetResponse
	Close() bool
	Watch(target []string, serviceName string)
}

// NewClient
//
//	@Description: 实例化RegisterInfomation
//	@param target
//	@return *RegisterInfomation
func NewClient(target []string) *RegisterInfomation {
	//获取endpoints用于实例化etcd对象
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   target,
			DialTimeout: 5 * time.Second,
		})
	if err != nil {
		logx.Errorf("etcd连接失败:%v", err)
		panic(err)
	}
	return &RegisterInfomation{Cil: cli}

}

// Put
//
//	@Description: 用于传入etcd,put的时候绑定租约
//	@receiver r *RegisterInfomation
//	@param serviceName
//	@param value
//	@param ttl
//	@return *RegisterInfomation
func (r *RegisterInfomation) Put(serviceName, value string, ttl int64) error {
	// 创建租约
	resp, err := r.Cil.Grant(context.Background(), ttl)
	if err != nil {
		logx.Errorf("租约创建失败: %v", err)
		return err
	}
	logx.Infof("租约ID: %v", resp.ID)

	// 自动续约
	ch, err := r.Cil.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		logx.Errorf("续约失败: %v", err)
		return err
	}

	// 使用租约进行操作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = r.Cil.Put(ctx, serviceName, value, clientv3.WithLease(resp.ID))
	if err != nil {
		logx.Errorf("put失败: %v", err)
		return err
	}
	logx.Info("put成功")

	// 监听租约状态
	go func() {
		for {
			select {
			case kaResp := <-ch:
				if kaResp == nil {
					logx.Errorf("续约失败: %v", err)
					return
				}
			}
		}
	}()

	return nil
}

// Get
//
//	@Description: 用于根据key获取value
//	@receiver r *RegisterInfomation
//	@param cli
//	@param key
//	@return *clientv3.GetResponse
func (r *RegisterInfomation) Get(key string) *clientv3.GetResponse {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := r.Cil.Get(ctx, key)
	if err != nil {
		logx.Errorf("etcd获取失败:%v", err)
		cancel()
		return nil
	}
	cancel()
	// 判断
	if len(resp.Kvs) == 0 {
		logx.Errorf("etcd获取失败")
		return nil
	}

	logx.Info("etcd获取成功")
	return resp
}

// Close
//
//	@Description: 用于关闭client连接
//	@receiver r *RegisterInfomation
//	@param cli
//	@return bool
func (r *RegisterInfomation) Close() bool {
	err := r.Cil.Close()
	if err != nil {
		logx.Errorf("关闭etcd连接失败:%v", err)
		return true
	}
	logx.Info("关闭etcd连接成功")
	return false
}

// Watch
//
//	@Description: 监听服务
//	@receiver *RegisterInfomation
//	@param cil
//	@param key
func (r *RegisterInfomation) Watch(serviceName string) {
	//获取endpoints用于实例化etcd对象
	cli, err := clientv3.New(
		clientv3.Config{
			Endpoints:   r.Cil.Endpoints(),
			DialTimeout: 5 * time.Second,
		})
	if err != nil {
		logx.Errorf("connect to etcd failed, err:%v", err)
		panic(err)
	}
	defer cli.Close()
	logx.Info("watch start!!")
	rec := cli.Watch(context.Background(), serviceName)
	for resp := range rec {
		for _, v := range resp.Events {
			logx.Infof("Type: %s serviceName:%s Value:%s\n", v.Type, v.Kv.Key, v.Kv.Value)
		}
	}
	logx.Info("watch end!!")
}
