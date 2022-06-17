package ddjb

import (
	"encoding/json"
)

type PidQueryResponse struct {
	PidQueryResponse struct {
		PidList []struct {
			CreateTime int    `json:"create_time,omitempty"` // 活动佣金比例，千分比（特定活动期间的佣金比例）
			PidName    string `json:"pid_name,omitempty"`    // 商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
			PId        string `json:"pid,omitempty"`         // 商品品牌词信息，如“苹果”、“阿迪达斯”、“李宁”等
			Status     int    `json:"status,omitempty"`      // 全局礼金金额，单位分
		} `json:"p_id_list"`
	} `json:"p_id_query_response"`
}

type PidQueryResult struct {
	Result PidQueryResponse // 结果
	Body   []byte           // 内容
	Err    error            // 错误
}

func NewPidResult(result PidQueryResponse, body []byte, err error) *PidQueryResult {
	return &PidQueryResult{Result: result, Body: body, Err: err}
}

// PidQuery 查询已经生成的推广位信息
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.query
func (app *App) PidQuery(notMustParams ...Params) *PidQueryResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.goods.pid.query", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response PidQueryResponse
	err = json.Unmarshal(body, &response)
	return NewPidResult(response, body, err)
}
