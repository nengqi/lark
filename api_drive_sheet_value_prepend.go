// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// PrependSheetValue 该接口用于根据 spreadsheetToken 和 range 向范围之前增加相应数据的行和相应的数据，相当于数组的插入操作；单次写入不超过5000行，100列，每个格子大小为0.5M。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIjMzUjLyIzM14iMyMTN
func (r *DriveService) PrependSheetValue(ctx context.Context, request *PrependSheetValueReq, options ...MethodOptionFunc) (*PrependSheetValueResp, *Response, error) {
	if r.cli.mock.mockDrivePrependSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#PrependSheetValue mock enable")
		return r.cli.mock.mockDrivePrependSheetValue(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "PrependSheetValue",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_prepend",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(prependSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDrivePrependSheetValue(f func(ctx context.Context, request *PrependSheetValueReq, options ...MethodOptionFunc) (*PrependSheetValueResp, *Response, error)) {
	r.mockDrivePrependSheetValue = f
}

func (r *Mock) UnMockDrivePrependSheetValue() {
	r.mockDrivePrependSheetValue = nil
}

type PrependSheetValueReq struct {
	SpreadSheetToken string                          `path:"spreadsheetToken" json:"-"` // sheet的token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	ValueRange       *PrependSheetValueReqValueRange `json:"valueRange,omitempty"`      // 值与范围
}

type PrependSheetValueReqValueRange struct {
	Range  string           `json:"range,omitempty"`  // ⁣查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见 ⁣[对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)  的第 5 项
	Values [][]SheetContent `json:"values,omitempty"` // 需要写入的值，如要写入公式、超链接、email、@人等，可详看附录[sheet 支持写入数据类型](/ssl:ttdoc/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN)
}

type prependSheetValueResp struct {
	Code int64                  `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data *PrependSheetValueResp `json:"data,omitempty"`
}

type PrependSheetValueResp struct {
	SpreadSheetToken string                        `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	TableRange       string                        `json:"tableRange,omitempty"`       // 写入的范围
	Revision         int64                         `json:"revision,omitempty"`         // sheet 的版本号
	Updates          *PrependSheetValueRespUpdates `json:"updates,omitempty"`          // 插入数据的范围、行列数等
}

type PrependSheetValueRespUpdates struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 写入的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 写入的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 写入的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 写入的单元格总数
	Revision         int64  `json:"revision,omitempty"`         // sheet 的版本号
}
