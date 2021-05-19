// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// IsInChat 判断用户或者机器人是否在群里。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat
func (r *ChatService) IsInChat(ctx context.Context, request *IsInChatReq, options ...MethodOptionFunc) (*IsInChatResp, *Response, error) {
	if r.cli.mock.mockChatIsInChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#IsInChat mock enable")
		return r.cli.mock.mockChatIsInChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "IsInChat",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/is_in_chat",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(isInChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatIsInChat(f func(ctx context.Context, request *IsInChatReq, options ...MethodOptionFunc) (*IsInChatResp, *Response, error)) {
	r.mockChatIsInChat = f
}

func (r *Mock) UnMockChatIsInChat() {
	r.mockChatIsInChat = nil
}

type IsInChatReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群 ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
}

type isInChatResp struct {
	Code int64         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string        `json:"msg,omitempty"`  // 错误描述
	Data *IsInChatResp `json:"data,omitempty"` //
}

type IsInChatResp struct {
	IsInChat bool `json:"is_in_chat,omitempty"` // 用户或者机器人是否在群中
}
