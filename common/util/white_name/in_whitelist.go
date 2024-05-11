package white_name

import "github.com/zeromicro/go-zero/core/logx"

func InWhitelist(list []string, target string) bool {
	for _, v := range list {

		if v == target {
			return true
		}
	}
	// 如果没有找到匹配项，则抛出错误
	logx.Infof("当前请求不在白名单中:%s", target)
	return false
}
