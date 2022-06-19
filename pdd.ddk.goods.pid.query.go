package ddjb

import (
	"encoding/json"
)

type PidQueryResponse struct {
	PidQueryResponse struct {
		PidList []struct {
			CreateTime int    `json:"create_time,omitempty"` // 推广位生成时间
			MediaId    int    `json:"media_id,omitempty"`    // 媒体id
			PidName    string `json:"pid_name,omitempty"`    // 推广位名称
			PId        string `json:"p_id,omitempty"`        // 推广位ID
			Status     int    `json:"status,omitempty"`      // 推广位状态
		} `json:"p_id_list"`
		TotalCount int `json:"total_count"`
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
