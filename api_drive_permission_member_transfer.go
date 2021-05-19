// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// TransferDriveMemberPermission 该接口用于根据文档信息和用户信息转移文档的所有者。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQzNzUjL0czM14CN3MTN
func (r *DriveService) TransferDriveMemberPermission(ctx context.Context, request *TransferDriveMemberPermissionReq, options ...MethodOptionFunc) (*TransferDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveTransferDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#TransferDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveTransferDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "TransferDriveMemberPermission",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/transfer",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(transferDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveTransferDriveMemberPermission(f func(ctx context.Context, request *TransferDriveMemberPermissionReq, options ...MethodOptionFunc) (*TransferDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveTransferDriveMemberPermission = f
}

func (r *Mock) UnMockDriveTransferDriveMemberPermission() {
	r.mockDriveTransferDriveMemberPermission = nil
}

type TransferDriveMemberPermissionReq struct {
	Token          string                                 `json:"token,omitempty"`            // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type           string                                 `json:"type,omitempty"`             // 文档类型  "doc"  or  "sheet" or "file"
	Owner          *TransferDriveMemberPermissionReqOwner `json:"owner,omitempty"`            // 要转移到的新的文档所有者
	RemoveOldOwner *bool                                  `json:"remove_old_owner,omitempty"` // true 为转移后删除旧 owner 的权限，默认为false
	CancelNotify   *bool                                  `json:"cancel_notify,omitempty"`    // true为不通知新owner，默认为false
}

type TransferDriveMemberPermissionReqOwner struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型，可选 **email、openid、userid**
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
}

type transferDriveMemberPermissionResp struct {
	Code int64                              `json:"code,omitempty"`
	Msg  string                             `json:"msg,omitempty"`
	Data *TransferDriveMemberPermissionResp `json:"data,omitempty"`
}

type TransferDriveMemberPermissionResp struct {
	IsSuccess bool                                    `json:"is_success,omitempty"` // 请求是否成功
	Type      string                                  `json:"type,omitempty"`       // 文档类型 "doc" or "sheet" or "file"
	Token     string                                  `json:"token,omitempty"`      // 文档的 token
	Owner     *TransferDriveMemberPermissionRespOwner `json:"owner,omitempty"`      // 文档当前所有者
}

type TransferDriveMemberPermissionRespOwner struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型，有 email 、openid、userid
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
}
