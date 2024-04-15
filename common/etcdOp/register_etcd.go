package etcdOp

import (
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

// RegisterEtcd 注册到etcd
func RegisterEtcd(etcdAddr, key, ip string, ttl int64) {
	// 切割ip地址,并判断长度是否为2
	list := strings.Split(ip, ":")
	if len(list) != 2 {
		logx.Error("ip地址错误")
		return
	}
	// 判断ip是否为0.0.0.0
	if list[0] == "0.0.0.0" {
		//newIP := util.GetSecondToLastIp()
		ip = strings.ReplaceAll(ip, "0.0.0.0", "192.168.30.1")
	}
	// 连接etcd
	etcd := NewClient([]string{etcdAddr})
	// put
	err := etcd.Put(key, ip, ttl)
	if err != nil {
		return
	}
	logx.Infof("put key %s value %s success", key, ip)
}
