// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateDriveComment 往云文档添加一条评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/create
func (r *DriveService) CreateDriveComment(ctx context.Context, request *CreateDriveCommentReq, options ...MethodOptionFunc) (*CreateDriveCommentResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveComment mock enable")
		return r.cli.mock.mockDriveCreateDriveComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveComment",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCreateDriveComment(f func(ctx context.Context, request *CreateDriveCommentReq, options ...MethodOptionFunc) (*CreateDriveCommentResp, *Response, error)) {
	r.mockDriveCreateDriveComment = f
}

func (r *Mock) UnMockDriveCreateDriveComment() {
	r.mockDriveCreateDriveComment = nil
}

type CreateDriveCommentReq struct {
	FileType     FileType                        `query:"file_type" json:"-"`      // 文档类型, 示例值："doc", 可选值有: `doc`：文档, `sheet`：表格, `file`：文件
	UserIDType   *IDType                         `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	FileToken    string                          `path:"file_token" json:"-"`      // 文档token, 示例值："doccnGp4UK1UskrOEJwBXd3****"
	CommentID    *string                         `json:"comment_id,omitempty"`     // 评论ID, 示例值："6916106822734578184"
	UserID       *string                         `json:"user_id,omitempty"`        // 用户ID, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab*****"
	CreateTime   *int64                          `json:"create_time,omitempty"`    // 创建时间, 示例值：1610281603
	UpdateTime   *int64                          `json:"update_time,omitempty"`    // 更新时间, 示例值：1610281603
	IsSolved     *bool                           `json:"is_solved,omitempty"`      // 是否已解决, 示例值：false
	SolvedTime   *int64                          `json:"solved_time,omitempty"`    // 解决评论时间, 示例值：1610281603
	SolverUserID *string                         `json:"solver_user_id,omitempty"` // 解决评论者的用户ID, 示例值："null"
	ReplyList    *CreateDriveCommentReqReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

type CreateDriveCommentReqReplyList struct {
	Replies []*CreateDriveCommentReqReplyListReplie `json:"replies,omitempty"` // 回复列表
}

type CreateDriveCommentReqReplyListReplie struct {
	ReplyID    *string                                      `json:"reply_id,omitempty"`    // 回复ID, 示例值："6916106822734594568"
	UserID     *string                                      `json:"user_id,omitempty"`     // 用户ID, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab2*****"
	CreateTime *int64                                       `json:"create_time,omitempty"` // 创建时间, 示例值：1610281603
	UpdateTime *int64                                       `json:"update_time,omitempty"` // 更新时间, 示例值：1610281603
	Content    *CreateDriveCommentReqReplyListReplieContent `json:"content,omitempty"`     // 回复内容
}

type CreateDriveCommentReqReplyListReplieContent struct {
	Elements []*CreateDriveCommentReqReplyListReplieContentElement `json:"elements,omitempty"` // 回复的内容
}

type CreateDriveCommentReqReplyListReplieContentElement struct {
	Type     string                                                      `json:"type,omitempty"`      // 回复的内容元素, 示例值："text_run", 可选值有: `text_run`：普通文本, `docs_link`：at 云文档链接, `person`：at 联系人
	TextRun  *CreateDriveCommentReqReplyListReplieContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *CreateDriveCommentReqReplyListReplieContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *CreateDriveCommentReqReplyListReplieContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

type CreateDriveCommentReqReplyListReplieContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本, 示例值："comment text"
}

type CreateDriveCommentReqReplyListReplieContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档, 示例值："https://bytedance.feishu.cn/docs/doccnHh7U87HOFpii5u5G*****"
}

type CreateDriveCommentReqReplyListReplieContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人, 示例值："ou_cc19b2bfb93f8a44db4b4d6eab*****"
}

type createDriveCommentResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveCommentResp `json:"data,omitempty"` //
}

type CreateDriveCommentResp struct {
	CommentID    string                           `json:"comment_id,omitempty"`     // 评论ID
	UserID       string                           `json:"user_id,omitempty"`        // 用户ID
	CreateTime   int64                            `json:"create_time,omitempty"`    // 创建时间
	UpdateTime   int64                            `json:"update_time,omitempty"`    // 更新时间
	IsSolved     bool                             `json:"is_solved,omitempty"`      // 是否已解决
	SolvedTime   int64                            `json:"solved_time,omitempty"`    // 解决评论时间
	SolverUserID string                           `json:"solver_user_id,omitempty"` // 解决评论者的用户ID
	ReplyList    *CreateDriveCommentRespReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

type CreateDriveCommentRespReplyList struct {
	Replies []*CreateDriveCommentRespReplyListReplie `json:"replies,omitempty"` // 回复列表
}

type CreateDriveCommentRespReplyListReplie struct {
	ReplyID    string                                        `json:"reply_id,omitempty"`    // 回复ID
	UserID     string                                        `json:"user_id,omitempty"`     // 用户ID
	CreateTime int64                                         `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64                                         `json:"update_time,omitempty"` // 更新时间
	Content    *CreateDriveCommentRespReplyListReplieContent `json:"content,omitempty"`     // 回复内容
}

type CreateDriveCommentRespReplyListReplieContent struct {
	Elements []*CreateDriveCommentRespReplyListReplieContentElement `json:"elements,omitempty"` // 回复的内容
}

type CreateDriveCommentRespReplyListReplieContentElement struct {
	Type     string                                                       `json:"type,omitempty"`      // 回复的内容元素, 可选值有: `text_run`：普通文本, `docs_link`：at 云文档链接, `person`：at 联系人
	TextRun  *CreateDriveCommentRespReplyListReplieContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *CreateDriveCommentRespReplyListReplieContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *CreateDriveCommentRespReplyListReplieContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

type CreateDriveCommentRespReplyListReplieContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本
}

type CreateDriveCommentRespReplyListReplieContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档
}

type CreateDriveCommentRespReplyListReplieContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人
}
