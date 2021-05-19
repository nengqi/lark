// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableMeta 通过 app_token 获取多维表格元数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/get
func (r *BitableService) GetBitableMeta(ctx context.Context, request *GetBitableMetaReq, options ...MethodOptionFunc) (*GetBitableMetaResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableMeta != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableMeta mock enable")
		return r.cli.mock.mockBitableGetBitableMeta(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Bitable",
		API:                 "GetBitableMeta",
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getBitableMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableMeta(f func(ctx context.Context, request *GetBitableMetaReq, options ...MethodOptionFunc) (*GetBitableMetaResp, *Response, error)) {
	r.mockBitableGetBitableMeta = f
}

func (r *Mock) UnMockBitableGetBitableMeta() {
	r.mockBitableGetBitableMeta = nil
}

type GetBitableMetaReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
}

type getBitableMetaResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableMetaResp `json:"data,omitempty"` //
}

type GetBitableMetaResp struct {
	App *GetBitableMetaRespApp `json:"app,omitempty"` // 多维表格元数据
}

type GetBitableMetaRespApp struct {
	AppToken string `json:"app_token,omitempty"` // 多维表格的 app_token
	Revision int64  `json:"revision,omitempty"`  // 多维表格的版本号
}
