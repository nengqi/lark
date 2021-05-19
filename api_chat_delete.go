// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteChat 解散群组
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 仅有 群主 或 创建群组且具备[更新应用所创建群的群信息]权限的机器人，可解散群
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete
func (r *ChatService) DeleteChat(ctx context.Context, request *DeleteChatReq, options ...MethodOptionFunc) (*DeleteChatResp, *Response, error) {
	if r.cli.mock.mockChatDeleteChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#DeleteChat mock enable")
		return r.cli.mock.mockChatDeleteChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "DeleteChat",
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatDeleteChat(f func(ctx context.Context, request *DeleteChatReq, options ...MethodOptionFunc) (*DeleteChatResp, *Response, error)) {
	r.mockChatDeleteChat = f
}

func (r *Mock) UnMockChatDeleteChat() {
	r.mockChatDeleteChat = nil
}

type DeleteChatReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群 ID, 示例值："oc_a0553eda9014c201e6969b478895c230"
}

type deleteChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteChatResp `json:"data,omitempty"`
}

type DeleteChatResp struct{}
