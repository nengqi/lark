// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateSearchDataSource 创建一个数据源
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create
func (r *SearchService) CreateSearchDataSource(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchCreateSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#CreateSearchDataSource mock enable")
		return r.cli.mock.mockSearchCreateSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "CreateSearchDataSource",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/search/v2/data_sources",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSearchCreateSearchDataSource(f func(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error)) {
	r.mockSearchCreateSearchDataSource = f
}

func (r *Mock) UnMockSearchCreateSearchDataSource() {
	r.mockSearchCreateSearchDataSource = nil
}

type CreateSearchDataSourceReq struct {
	Name        *string `json:"name,omitempty"`        // data_source的展示名称, 示例值："客服工单"
	State       int64   `json:"state,omitempty"`       // 数据源状态，0-未上线，1-已上线, 示例值：0, 可选值有: `0`：未上线, `1`：已上线
	Description *string `json:"description,omitempty"` // 对于数据源的描述, 示例值："搜索客服工单数据"
}

type createSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateSearchDataSourceResp `json:"data,omitempty"` //
}

type CreateSearchDataSourceResp struct {
	DataSource *CreateSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源实例
}

type CreateSearchDataSourceRespDataSource struct {
	ID            string `json:"id,omitempty"`              // 数据源的唯一标识
	Name          string `json:"name,omitempty"`            // data_source的展示名称
	State         int64  `json:"state,omitempty"`           // 数据源状态，0-未上线，1-已上线, 可选值有: `0`：未上线, `1`：已上线
	Description   string `json:"description,omitempty"`     // 对于数据源的描述
	CreateTime    string `json:"create_time,omitempty"`     // 创建时间，使用Unix时间戳，秒
	UpdateTime    string `json:"update_time,omitempty"`     // 更新时间，使用Unix时间戳，秒
	IsExceedQuota bool   `json:"is_exceed_quota,omitempty"` // 是否超限
}
