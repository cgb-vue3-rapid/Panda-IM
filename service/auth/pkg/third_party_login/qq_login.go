package third_party_login

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//type QQInfo struct {
//	NickName     string `json:"nickname"`
//	Gender       string `json:"gender"`
//	Avatar       string `json:"avatar"`
//	OpenID       string `json:"third_party_id"`
//}
//}
//
//type QQLoginInfo struct {
//	appID        string
//	appKey       string
//	redirect     string
//	code         string
//	accessToken  string
//	OpenID        string
//}
//
//type QQConfig struct {
//	AppID       string
//	AppKey      string
//	RedirectURL string
//}
//
//// NewQQLogin 用于创建 QQLogin 实例并进行登录操作
//func NewQQLogin(code string, conf QQConfig) (qqInfo QQInfo, err error) {
//	qqLogin := &QQLogin{
//		appID:       conf.AppID,
//		appKey:      conf.AppKey,
//		redirectURL: conf.RedirectURL,
//		code:        code,
//	}
//
//	err = qqLogin.GetAccessToken()
//	if err != nil {
//		return QQInfo{}, err
//	}
//
//	err = qqLogin.GetOpenID()
//	if err != nil {
//		return QQInfo{}, err
//	}
//
//	qqInfo, err = qqLogin.GetUserInfo()
//	if err != nil {
//		return QQInfo{}, err
//	}
//
//	return qqInfo, nil
//}
//
//type QQLogin struct {
//	appID       string
//	appKey      string
//	redirectURL string
//	code        string
//	AccessToken string
//	OpenID      string
//}
//
//// GetAccessToken 获取Access_token
//func (q *QQLogin) GetAccessToken() error {
//	// 构建请求参数
//	params := url.Values{}
//	params.Add("grant_type", "authorization_code")
//	params.Add("client_id", q.appID)
//	params.Add("client_secret", q.appKey)
//	params.Add("code", q.code)
//	params.Add("redirect_uri", q.redirect)
//
//	// 构建请求URL
//	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
//	if err != nil {
//		return err
//	}
//	u.RawQuery = params.Encode()
//
//	// 发送HTTP GET请求
//	res, err := http.Get(u.String())
//	if err != nil {
//		return err
//	}
//	defer res.Body.Close()
//
//	// 解析响应体为键值对形式
//	qs, err := parseQS(res.Body)
//	if err != nil {
//		return err
//	}
//
//	// 从解析后的响应中获取access_token，并设置到QQLogin结构体中
//	q.AccessToken = qs["access_token"][0]
//
//	return nil
//}
//
//// parseQS 将HTTP响应的正文解析为键值对形式
//func parseQS(r io.Reader) (val map[string][]string, err error) {
//	body := readAll(r)
//	val, err = url.ParseQuery(body)
//	if err != nil {
//		return nil, err
//	}
//	return val, nil
//}
//
//// GetUserInfo 获取用户信息
//func (q *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
//	params := url.Values{}
//	params.Add("access_token", q.AccessToken)
//	params.Add("oauth_consumer_key", q.appID)
//	params.Add("openid", q.OpenID)
//
//	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
//	if err != nil {
//		return qqInfo, err
//	}
//	u.RawQuery = params.Encode()
//
//	res, err := http.Get(u.String())
//	if err != nil {
//		return qqInfo, err
//	}
//	defer res.Body.Close()
//
//	data, err := io.ReadAll(res.Body)
//	if err != nil {
//		return qqInfo, err
//	}
//
//	err = json.Unmarshal(data, &qqInfo)
//	if err != nil {
//		return qqInfo, err
//	}
//
//	return qqInfo, nil
//}
//// GetOpenID 获取openid
//func (q *QQLogin) GetOpenID() error {
//	// 构建请求URL
//	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessTok))
//	if err != nil {
//		return err
//	}
//
//	// 发送HTTP GET请求
//	res, err := http.Get(u.String())
//	if err != nil {
//		return err
//	}
//	defer res.Body.Close()
//
//	// 解析获取openid
//	openID, err := getopenID(res.Body)
//	if err != nil {
//		return err
//	}
//
//	// 将获取的openid设置到QQLogin结构体中
//	q.OpenID = openID
//
//	return nil
//}
//
//
////// getOpenID 从 HTTP 响应的正文中解析出 openid
////func getOpenID(r io.Reader) (string, error) {
////	body := readAll(r)
////	start := strings.Index(body, `"openid":"`)
////	if start == -1 {
////		return "", fmt.Errorf("openid not found")
////	}
////	start += len(`"openid":"`)
////	end := strings.Index(body[start:], `"`)
////	if end == -1 {
////		return "", fmt.Errorf("openid not found")
////	}
////	return body[start : start+end], nil
////}
//
//
//
//
//// readAll 读取所有数据并将其转换为字符串
//func readAll(r io.Reader) string {
//	b, err := ioutil.ReadAll(r)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return string(b)
//}

type QQInfo struct {
	NickName string `json:"nickname"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	OpenID   string `json:"third_party_id"`
}

type QQLoginInfo struct {
	AppID       string
	AppKey      string
	RedirectURL string
	Code        string
	AccessToken string
	OpenID      string
}

type QQConfig struct {
	AppID       string
	AppKey      string
	RedirectURL string
}

// NewQQLogin 用于创建 QQLogin 实例并进行登录操作
func NewQQLogin(code string, conf QQConfig) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		AppID:       conf.AppID,
		AppKey:      conf.AppKey,
		RedirectURL: conf.RedirectURL,
		Code:        code,
	}

	err = qqLogin.GetAccessToken()
	if err != nil {
		return QQInfo{}, err
	}

	err = qqLogin.GetOpenID()
	if err != nil {
		return QQInfo{}, err
	}

	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		return QQInfo{}, err
	}

	return qqInfo, nil
}

type QQLogin struct {
	AppID       string
	AppKey      string
	RedirectURL string
	Code        string
	AccessToken string
	OpenID      string
}

// GetAccessToken 获取Access_token
func (q *QQLogin) GetAccessToken() error {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", q.AppID)
	params.Add("client_secret", q.AppKey)
	params.Add("code", q.Code)
	params.Add("redirect_uri", q.RedirectURL)

	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	qs, err := parseQS(res.Body)
	if err != nil {
		return err
	}

	q.AccessToken = qs["access_token"][0]

	return nil
}

// parseQS 将HTTP响应的正文解析为键值对形式
func parseQS(r io.Reader) (val map[string][]string, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	val, err = url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetUserInfo 获取用户信息
func (q *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.AccessToken)
	params.Add("oauth_consumer_key", q.AppID)
	params.Add("openid", q.OpenID)

	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqInfo, err
	}
	u.RawQuery = params.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return qqInfo, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return qqInfo, err
	}

	err = json.Unmarshal(data, &qqInfo)
	if err != nil {
		return qqInfo, err
	}

	return qqInfo, nil
}

// GetOpenID 获取openid
func (q *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.AccessToken))
	if err != nil {
		return err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data map[string]string
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	q.OpenID = data["openid"]

	return nil
}
