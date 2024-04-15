package main

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/etcdOp"
	"akita/panda-im/common/util/white_name"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// 配置文件路径
var (
	configFile = flag.String("f", "gateway.yaml", "the config file")
	c          config
)

// ErrorResponse 是自定义的错误响应结构体
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse 用于初始化并返回一个新的 ErrorResponse 结构体实例
func NewErrorResponse(message string, res http.ResponseWriter) {
	// 初始化 ErrorResponse 实例
	errResponse := ErrorResponse{
		Code:    3001,
		Message: message,
	}

	// 序列化 ErrorResponse 实例
	errJSON, err := json.Marshal(errResponse)
	if err != nil {
		// 如果序列化失败，则写入默认错误信息
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to serialize error response"))
		return
	}

	// 设置响应头为 JSON 格式
	res.Header().Set("Content-Type", "application/json")

	// 写入序列化后的 JSON 错误信息
	res.WriteHeader(http.StatusBadRequest)
	res.Write(errJSON)
}

// 配置结构体
type config struct {
	Addr string // 网关服务地址
	Etcd struct {
		Endpoints []string // etcd 服务器地址列表
	}
	Auth struct {
		Authenticate string // 认证服务路径
	}
	Whitelist []string // 白名单列表
	Log       logx.LogConf
}

// 网关服务入口函数
func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &c)
	http.HandleFunc("/", gatewayHandler) // 注册网关处理函数
	logx.SetUp(c.Log)
	logx.Infof("starting gateway server at %s", c.Addr)
	err := http.ListenAndServe(c.Addr, nil)
	if err != nil {
		logx.Errorf("failed to start gateway server: %v", err)
		return
	}
}

// 网关请求处理函数
func gatewayHandler(res http.ResponseWriter, req *http.Request) {
	// 解析请求路径，获取服务地址
	serverAddr, err := parseRequestPath(req.URL.Path)
	if err != nil {
		// 替换为错误响应
		NewErrorResponse("Invalid request", res)
		return
	}
	logx.Infof("serverAddr: %s", serverAddr)

	// 检查请求是否在白名单中，如果不在，则进行认证
	if !white_name.InWhitelist(c.Whitelist, req.URL.String()) {
		err := authenticateRequest(serverAddr, req.RemoteAddr, req.Method)
		if err != nil {
			// 替换为错误响应
			NewErrorResponse("Authentication failed", res)
			return
		}
	}

	// 转发请求到服务
	err = forwardRequest(serverAddr, req, res)
	if err != nil {
		// 替换为错误响应
		NewErrorResponse("Failed to forward request", res)
		return
	}
}

// 解析请求路径，获取服务地址
func parseRequestPath(path string) (string, error) {
	regex, _ := regexp.Compile(`/v1/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(path)
	if len(addrList) != 2 {
		return "", fmt.Errorf("invalid request path: %s", path)
	}
	logx.Infof("serverAddr: %s", addrList[1])
	return addrList[1], nil
}

// 发起认证请求
func authenticateRequest(serverAddr, remoteAddr, method string) error {
	authAddr, err := getServiceAddress(serverAddr)
	if err != nil {
		logx.Errorf("failed to get auth service address: %v", err)
		return err
	}
	logx.Infof("authAddr: %s", authAddr)
	authURL := fmt.Sprintf("http://%s/v1/api/auth/%s", authAddr, c.Auth.Authenticate)
	logx.Infof("authURL: %s", authURL)
	authReq, err := http.NewRequest(method, authURL, nil)
	if err != nil {
		logx.Errorf("failed to create auth request: %v", err)
		return err
	}
	logx.Infof("remoteAddr: %s", remoteAddr)
	authReq.Header.Set("validPath", remoteAddr)
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Errorf("failed to send auth request: %v", err)
		return err
	}
	defer authRes.Body.Close()
	// 处理认证响应
	// ...

	return nil
}

// 转发请求到服务
func forwardRequest(serverAddr string, req *http.Request, res http.ResponseWriter) error {
	addr, err := getServiceAddress(serverAddr)
	if err != nil {
		logx.Errorf("failed to get service address: %v", err)
		return err
	}
	logx.Infof("forwarding request to: %s", addr)
	url := "http://" + addr + req.URL.String()
	all, err := io.ReadAll(req.Body)
	if err != nil {
		logx.Errorf("failed to read request body: %v", err)
		return err
	}
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(all))
	if err != nil {
		logx.Errorf("failed to create proxy request: %v", err)
		return err
	}
	proxyReq.Header = req.Header
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	proxyRes, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		logx.Errorf("failed to send proxy request: %v", err)
		return err
	}
	defer proxyRes.Body.Close()
	_, err = io.Copy(res, proxyRes.Body)
	if err != nil {
		logx.Errorf("failed to copy response body: %v", err)
		return err
	}
	return nil
}

// 从etcd中获取服务地址
func getServiceAddress(serverAddr string) (string, error) {
	etcdClient := etcdOp.NewClient(c.Etcd.Endpoints)
	resp := etcdClient.Get(serverAddr + constants.Prefix)
	if string(resp.Kvs[0].Value) == "" {
		return "", fmt.Errorf("no such server: %s", serverAddr)
	}
	return string(resp.Kvs[0].Value), nil
}
