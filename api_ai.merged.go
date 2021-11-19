// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DetectTextLanguage 机器翻译 (MT)，支持 100 多种语言识别，返回符合 ISO 693-1 标准
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/detect
func (r *AIService) DetectTextLanguage(ctx context.Context, request *DetectTextLanguageReq, options ...MethodOptionFunc) (*DetectTextLanguageResp, *Response, error) {
	if r.cli.mock.mockAIDetectTextLanguage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#DetectTextLanguage mock enable")
		return r.cli.mock.mockAIDetectTextLanguage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "DetectTextLanguage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/translation/v1/text/detect",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(detectTextLanguageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIDetectTextLanguage(f func(ctx context.Context, request *DetectTextLanguageReq, options ...MethodOptionFunc) (*DetectTextLanguageResp, *Response, error)) {
	r.mockAIDetectTextLanguage = f
}

func (r *Mock) UnMockAIDetectTextLanguage() {
	r.mockAIDetectTextLanguage = nil
}

type DetectTextLanguageReq struct {
	Text string `json:"text,omitempty"` // 需要被识别语种的文本, 示例值："你好"
}

type detectTextLanguageResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DetectTextLanguageResp `json:"data,omitempty"`
}

type DetectTextLanguageResp struct {
	Language string `json:"language,omitempty"` // 识别的文本语种，返回符合 ISO 693-1 标准
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// RecognizeBasicImage 可识别图片中的文字，按区域划分返回文本列表
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize
func (r *AIService) RecognizeBasicImage(ctx context.Context, request *RecognizeBasicImageReq, options ...MethodOptionFunc) (*RecognizeBasicImageResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeBasicImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeBasicImage mock enable")
		return r.cli.mock.mockAIRecognizeBasicImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeBasicImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/optical_char_recognition/v1/image/basic_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeBasicImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIRecognizeBasicImage(f func(ctx context.Context, request *RecognizeBasicImageReq, options ...MethodOptionFunc) (*RecognizeBasicImageResp, *Response, error)) {
	r.mockAIRecognizeBasicImage = f
}

func (r *Mock) UnMockAIRecognizeBasicImage() {
	r.mockAIRecognizeBasicImage = nil
}

type RecognizeBasicImageReq struct {
	Image *string `json:"image,omitempty"` // base64 后的图片数据, 示例值："base64后的图片二进制数据"
}

type recognizeBasicImageResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeBasicImageResp `json:"data,omitempty"`
}

type RecognizeBasicImageResp struct {
	TextList []string `json:"text_list,omitempty"` // 按区域识别，返回文本列表
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// RecognizeSpeechFile 语音文件识别接口，上传整段语音文件进行一次性识别。接口适合 60 秒以内音频识别
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/file_recognize
func (r *AIService) RecognizeSpeechFile(ctx context.Context, request *RecognizeSpeechFileReq, options ...MethodOptionFunc) (*RecognizeSpeechFileResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeSpeechFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeSpeechFile mock enable")
		return r.cli.mock.mockAIRecognizeSpeechFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeSpeechFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/speech_to_text/v1/speech/file_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeSpeechFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIRecognizeSpeechFile(f func(ctx context.Context, request *RecognizeSpeechFileReq, options ...MethodOptionFunc) (*RecognizeSpeechFileResp, *Response, error)) {
	r.mockAIRecognizeSpeechFile = f
}

func (r *Mock) UnMockAIRecognizeSpeechFile() {
	r.mockAIRecognizeSpeechFile = nil
}

type RecognizeSpeechFileReq struct {
	Speech *RecognizeSpeechFileReqSpeech `json:"speech,omitempty"` // 语音资源
	Config *RecognizeSpeechFileReqConfig `json:"config,omitempty"` // 配置属性
}

type RecognizeSpeechFileReqSpeech struct {
	Speech *string `json:"speech,omitempty"` // base64 后的音频文件进行, 示例值："base64 后的音频内容"
}

type RecognizeSpeechFileReqConfig struct {
	FileID     string `json:"file_id,omitempty"`     // 仅包含字母数字和下划线的 16 位字符串作为文件的标识，用户生成, 示例值："qwe12dd34567890w"
	Format     string `json:"format,omitempty"`      // 语音格式，目前仅支持：pcm, 示例值："pcm"
	EngineType string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合, 示例值："16k_auto"
}

type recognizeSpeechFileResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeSpeechFileResp `json:"data,omitempty"`
}

type RecognizeSpeechFileResp struct {
	RecognitionText string `json:"recognition_text,omitempty"` // 语音识别后的文本信息
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// RecognizeSpeechStream 语音流式接口，将整个音频文件分片进行传入模型。能够实时返回数据。建议每个音频分片的大小为 100-200ms
//
// 单租户限流：20 路（一个 stream_id 称为一路会话），同租户下的应用没有限流，共享本租户的 20路限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/stream_recognize
func (r *AIService) RecognizeSpeechStream(ctx context.Context, request *RecognizeSpeechStreamReq, options ...MethodOptionFunc) (*RecognizeSpeechStreamResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeSpeechStream != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeSpeechStream mock enable")
		return r.cli.mock.mockAIRecognizeSpeechStream(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeSpeechStream",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/speech_to_text/v1/speech/stream_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeSpeechStreamResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIRecognizeSpeechStream(f func(ctx context.Context, request *RecognizeSpeechStreamReq, options ...MethodOptionFunc) (*RecognizeSpeechStreamResp, *Response, error)) {
	r.mockAIRecognizeSpeechStream = f
}

func (r *Mock) UnMockAIRecognizeSpeechStream() {
	r.mockAIRecognizeSpeechStream = nil
}

type RecognizeSpeechStreamReq struct {
	Speech *RecognizeSpeechStreamReqSpeech `json:"speech,omitempty"` // 语音资源
	Config *RecognizeSpeechStreamReqConfig `json:"config,omitempty"` // 配置属性
}

type RecognizeSpeechStreamReqSpeech struct {
	Speech *string `json:"speech,omitempty"` // base64 后的音频文件进行, 示例值："base64 后的音频内容"
}

type RecognizeSpeechStreamReqConfig struct {
	StreamID   string `json:"stream_id,omitempty"`   // 仅包含字母数字和下划线的 16 位字符串作为同一数据流的标识，用户生成, 示例值："asd1234567890ddd"
	SequenceID int64  `json:"sequence_id,omitempty"` // 数据流分片的序号，序号从 0 开始，每次请求递增 1, 示例值：1
	Action     int64  `json:"action,omitempty"`      // 数据流标记：1 首包，2 正常结束，等待结果返回，3 中断数据流不返回最终结果, 示例值：1
	Format     string `json:"format,omitempty"`      // 语音格式，目前仅支持：pcm, 示例值："pcm"
	EngineType string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合, 示例值："16k_auto"
}

type recognizeSpeechStreamResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeSpeechStreamResp `json:"data,omitempty"`
}

type RecognizeSpeechStreamResp struct {
	StreamID        string `json:"stream_id,omitempty"`        // 16 位 String 随机串作为同一数据流的标识
	SequenceID      int64  `json:"sequence_id,omitempty"`      // 数据流分片的序号，序号从 0 开始，每次请求递增 1
	RecognitionText string `json:"recognition_text,omitempty"` // 语音流识别后的文本信息
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// TranslateText 机器翻译 (MT)，支持中日英（zh、ja、en）三语互译
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate
func (r *AIService) TranslateText(ctx context.Context, request *TranslateTextReq, options ...MethodOptionFunc) (*TranslateTextResp, *Response, error) {
	if r.cli.mock.mockAITranslateText != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#TranslateText mock enable")
		return r.cli.mock.mockAITranslateText(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "TranslateText",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/translation/v1/text/translate",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(translateTextResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAITranslateText(f func(ctx context.Context, request *TranslateTextReq, options ...MethodOptionFunc) (*TranslateTextResp, *Response, error)) {
	r.mockAITranslateText = f
}

func (r *Mock) UnMockAITranslateText() {
	r.mockAITranslateText = nil
}

type TranslateTextReq struct {
	SourceLanguage string                    `json:"source_language,omitempty"` // 源语言, 示例值："zh"
	Text           string                    `json:"text,omitempty"`            // 源文本, 示例值："尝试使用一下飞书吧"
	TargetLanguage string                    `json:"target_language,omitempty"` // 目标语言, 示例值："en"
	Glossary       *TranslateTextReqGlossary `json:"glossary,omitempty"`        // 请求级术语表，携带术语，仅在本次翻译中生效（最多能携带 128个术语词）
}

type TranslateTextReqGlossary struct {
	From string `json:"from,omitempty"` // 原文, 示例值："飞书"
	To   string `json:"to,omitempty"`   // 译文, 示例值："Lark"
}

type translateTextResp struct {
	Code int64              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *TranslateTextResp `json:"data,omitempty"`
}

type TranslateTextResp struct {
	Text string `json:"text,omitempty"` // 翻译后的文本
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// GetApplicationAppAdminUserList
//
// 查询审核应用的管理员列表，返回最新10个管理员账户id列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN
func (r *ApplicationService) GetApplicationAppAdminUserList(ctx context.Context, request *GetApplicationAppAdminUserListReq, options ...MethodOptionFunc) (*GetApplicationAppAdminUserListResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationAppAdminUserList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationAppAdminUserList mock enable")
		return r.cli.mock.mockApplicationGetApplicationAppAdminUserList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationAppAdminUserList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/user/v4/app_admin_user/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationAppAdminUserListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApplicationGetApplicationAppAdminUserList(f func(ctx context.Context, request *GetApplicationAppAdminUserListReq, options ...MethodOptionFunc) (*GetApplicationAppAdminUserListResp, *Response, error)) {
	r.mockApplicationGetApplicationAppAdminUserList = f
}

func (r *Mock) UnMockApplicationGetApplicationAppAdminUserList() {
	r.mockApplicationGetApplicationAppAdminUserList = nil
}

type GetApplicationAppAdminUserListReq struct{}

type getApplicationAppAdminUserListResp struct {
	Code int64                               `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 返回码描述
	Data *GetApplicationAppAdminUserListResp `json:"data,omitempty"` // -
}

type GetApplicationAppAdminUserListResp struct {
	UserList []*GetApplicationAppAdminUserListRespUser `json:"user_list,omitempty"` // 管理员列表
}

type GetApplicationAppAdminUserListRespUser struct {
	OpenID *GetApplicationAppAdminUserListRespUserOpenID `json:"open_id,omitempty"` // 某管理员的open_id
}

type GetApplicationAppAdminUserListRespUserOpenID struct {
	UserID  string `json:"user_id,omitempty"`  // 某管理员的user_id
	UnionID string `json:"union_id,omitempty"` // 某管理员的union_id
}
