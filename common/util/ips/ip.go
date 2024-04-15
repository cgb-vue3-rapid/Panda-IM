package ips

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net"
)

func GetSecondToLastIp() (addr string) {
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		logx.Errorf("Failed to get network interfaces: %v", err)
		return
	}

	for i := len(interfaces) - 1; i >= 0; i-- {
		addrs, err := interfaces[i].Addrs()
		if err != nil {
			logx.Errorf("Failed to get addresses for interface %s: %v", interfaces[i].Name, err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	if len(ips) >= 2 {
		return ips[len(ips)-2]
	}

	return ips[0]
}
