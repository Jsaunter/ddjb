package ddjb

import (
	"encoding/json"
	"gitee.com/dtapps/gohttp"
)

// App 公共请求参数
type App struct {
	ClientId     string // POP分配给应用的client_id
	ClientSecret string // POP分配给应用的client_secret
	Pid          string // 推广位
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string      `json:"error_msg"`
		SubMsg    string      `json:"sub_msg"`
		SubCode   interface{} `json:"sub_code"`
		ErrorCode int         `json:"error_code"`
		RequestId string      `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func (app *App) request(params map[string]interface{}) (resp []byte, err error) {
	// 签名
	app.Sign(params)
	// 发送请求
	get, err := gohttp.Get("https://gw-api.pinduoduo.com/api/router", params)
	// 检查错误
	var errResp ErrResp
	_ = json.Unmarshal(get.Body, &errResp)
	return get.Body, err
}
