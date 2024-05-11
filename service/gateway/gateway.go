package main

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/etcdOp"
	"akita/panda-im/common/util/white_name"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
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

// AuthResponse 定义结构体来表示认证响应
type AuthResponse struct {
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

// 网关处理函数
func gatewayHandler(res http.ResponseWriter, req *http.Request) {
	// 解析请求路径，获取服务地址
	serverAddr, err := parseRequestPath(req.URL.Path)
	if err != nil {
		// 替换为错误响应
		NewErrorResponse("Invalid request", res)
		return
	}
	logx.Infof("服务地址: %s", serverAddr)

	// 检查请求是否在白名单中，如果不在，则进行认证
	if !white_name.InWhitelist(c.Whitelist, req.URL.String()) {
		// 走认证服务
		err := authenticateRequest(req.RemoteAddr, req.Method, res, req)
		if err != nil {
			// 替换为错误响应
			logx.Errorf("认证失败: %v", err)
			NewErrorResponse("认证失败", res)
			return
		}
	}

	// 转发请求到服务
	err = forwardRequest(serverAddr, req, res)
	if err != nil {
		logx.Errorf("转发请求失败: %v", err)
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
	return addrList[1], nil
}

// 发起认证请求
func authenticateRequest(remoteAddr, method string, res http.ResponseWriter, req *http.Request) error {

	// 解析头部携带的token
	ok, tokenStr := authenticateAndAuthorizeToken(res, req)

	if !ok {
		logx.Errorf("认证失败")
		return fmt.Errorf("failed")
	}

	// 获取认证服务地址
	authAddr, err := getServiceAddress("auth")
	if err != nil {
		logx.Errorf("failed to get auth service address: %v", err)
		NewErrorResponse("认证失败", res)
		return err
	}
	logx.Infof("认证服务地址: %s", authAddr)
	authURL := fmt.Sprintf("http://%s/v1/api/auth/%s", authAddr, c.Auth.Authenticate)
	logx.Infof("认证服务请求地址: %s", authURL)
	authReq, err := http.NewRequest(method, authURL, nil)
	if err != nil {
		logx.Errorf("failed to create auth request: %v", err)
		return err
	}
	logx.Infof("指发起请求的客户端的地址: %s", remoteAddr)
	authReq.Header.Set("validPath", remoteAddr)
	authReq.Header.Set("token", tokenStr)

	// 发送认证请求
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Errorf("failed to send auth request: %v", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logx.Errorf("failed to close auth response body: %v", err)
		}
	}(authRes.Body)
	// 处理认证响应
	// 检查响应状态码
	if authRes.StatusCode != http.StatusOK {
		logx.Errorf("认证失败，响应状态码为: %d", authRes.StatusCode)
		return errors.New("failed")
	}

	// 读取响应的内容
	body, err := io.ReadAll(authRes.Body)
	if err != nil {
		logx.Errorf("failed to read auth response body: %v", err)
		return err
	}

	// 解析JSON数据到结构体
	var response AuthResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logx.Errorf("failed to decode auth response body: %v", err)
		return err
	}

	// 认证失败
	if response.Code != 0 {
		return errors.New("failed")
	}

	// 输出响应的内容
	logx.Infof("auth response: %s", string(body))
	return nil
}

// 转发请求到服务
func forwardRequest(serverAddr string, req *http.Request, res http.ResponseWriter) error {

	addr, err := getServiceAddress(serverAddr)
	if err != nil {
		logx.Errorf("failed to get service address: %v", err)
		return err
	}
	logx.Infof("转发请求到: %s", addr)
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
	resp := etcdClient.Get(serverAddr + constants.ApiPrefix)
	if string(resp.Kvs[0].Value) == "" {
		return "", fmt.Errorf("no such server: %s", serverAddr)
	}
	return string(resp.Kvs[0].Value), nil
}

func authenticateAndAuthorizeToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	// 获取请求中的 Authorization 头部
	token := r.Header.Get("token")

	// 如果 Authorization 头部为空，则返回 TokenIsEmpty 错误
	if token == "" {
		logx.Errorf("token is empty")
		return false, ""
	}

	//r.Header.Set("tokenStr", fmt.Sprintf("Bearer %s %s", parts[0], parts[1]))

	//// 解析 Token
	//parseToken, isExpire, err := token_manager.ParseToken(parts[0], parts[1], "panda@akita@AccessSecret", "panda@akita@RefreshSecret")
	//if err != nil {
	//	// 如果解析 Token 出错，则返回 TokenParseErr 错误
	//	httpx.ErrorCtx(r.Context(), w, xcode.TokenParseErr)
	//	return false
	//}

	// 刷新 Token
	//if isExpire {
	//	parts[0], parts[1] = token_manager.GenToken(parseToken.UserID, parseToken.Nickname, "panda@akita@AccessSecret", "panda@akita@RefreshSecret", parseToken.Role)
	//}

	//// 将用户 ID 和 Token 添加到请求的上下文中
	//ctx := context.WithValue(r.Context(), constants.UserId, parseToken.UserID)
	//*r = *r.WithContext(ctx)
	//
	//// 打印用户 ID
	//fmt.Println(r.Context().Value(constants.UserId))

	return true, token
}
