package etcdOp

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestEtcd(t *testing.T) {
	client := NewClient([]string{"192.168.30.130:2379"})
	err := client.Put("wenzhuhao", "wenzhuhaobaba", 60)
	if err != nil {
		logx.Errorf("etcd put error: %v", err)
		return
	}
	resp := client.Get("wenzhuhao")
	for _, v := range resp.Kvs {
		fmt.Printf("%s:%s\n", v.Key, v.Value)
	}
}
