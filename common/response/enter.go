package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 函数用于返回 HTTP 响应
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 如果没有错误，返回成功的 JSON 响应
		response := &Body{
			Code: 0,
			Msg:  "成功",
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, response)
		return
	}

	// 如果有错误，返回失败的 JSON 响应
	errCode := uint32(10086) // 自定义的错误码
	errMsg := "服务器错误"   // 自定义的错误消息

	httpx.WriteJson(w, http.StatusBadRequest, &Body{
		Code: errCode,
		Msg:  errMsg,
		Data: nil,
	})
}
