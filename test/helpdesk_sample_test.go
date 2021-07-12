// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Helpdesk_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Helpdesk

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicket(ctx, &lark.GetHelpdeskTicketReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DownloadHelpdeskTicketImage(ctx, &lark.DownloadHelpdeskTicketImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AnswerHelpdeskTicketUserQuery(ctx, &lark.AnswerHelpdeskTicketUserQueryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendHelpdeskTicketMessage(ctx, &lark.SendHelpdeskTicketMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedFieldList(ctx, &lark.GetHelpdeskTicketCustomizedFieldListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedField(ctx, &lark.GetHelpdeskTicketCustomizedFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategory(ctx, &lark.GetHelpdeskCategoryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategoryList(ctx, &lark.GetHelpdeskCategoryListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQ(ctx, &lark.GetHelpdeskFAQReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQList(ctx, &lark.GetHelpdeskFAQListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQImage(ctx, &lark.GetHelpdeskFAQImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchHelpdeskFAQ(ctx, &lark.SearchHelpdeskFAQReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeHelpdeskEvent(ctx, &lark.SubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UnsubscribeHelpdeskEvent(ctx, &lark.UnsubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Helpdesk

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskStartHelpdeskService(func(ctx context.Context, request *lark.StartHelpdeskServiceReq, options ...lark.MethodOptionFunc) (*lark.StartHelpdeskServiceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskStartHelpdeskService()

			_, _, err := moduleCli.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskTicket(func(ctx context.Context, request *lark.GetHelpdeskTicketReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskTicketResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskTicket()

			_, _, err := moduleCli.GetHelpdeskTicket(ctx, &lark.GetHelpdeskTicketReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskTicketList(func(ctx context.Context, request *lark.GetHelpdeskTicketListReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskTicketListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskTicketList()

			_, _, err := moduleCli.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskDownloadHelpdeskTicketImage(func(ctx context.Context, request *lark.DownloadHelpdeskTicketImageReq, options ...lark.MethodOptionFunc) (*lark.DownloadHelpdeskTicketImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskDownloadHelpdeskTicketImage()

			_, _, err := moduleCli.DownloadHelpdeskTicketImage(ctx, &lark.DownloadHelpdeskTicketImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskAnswerHelpdeskTicketUserQuery(func(ctx context.Context, request *lark.AnswerHelpdeskTicketUserQueryReq, options ...lark.MethodOptionFunc) (*lark.AnswerHelpdeskTicketUserQueryResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskAnswerHelpdeskTicketUserQuery()

			_, _, err := moduleCli.AnswerHelpdeskTicketUserQuery(ctx, &lark.AnswerHelpdeskTicketUserQueryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskTicketMessageList(func(ctx context.Context, request *lark.GetHelpdeskTicketMessageListReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskTicketMessageListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskTicketMessageList()

			_, _, err := moduleCli.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskSendHelpdeskTicketMessage(func(ctx context.Context, request *lark.SendHelpdeskTicketMessageReq, options ...lark.MethodOptionFunc) (*lark.SendHelpdeskTicketMessageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskSendHelpdeskTicketMessage()

			_, _, err := moduleCli.SendHelpdeskTicketMessage(ctx, &lark.SendHelpdeskTicketMessageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskTicketCustomizedFieldList(func(ctx context.Context, request *lark.GetHelpdeskTicketCustomizedFieldListReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskTicketCustomizedFieldListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskTicketCustomizedFieldList()

			_, _, err := moduleCli.GetHelpdeskTicketCustomizedFieldList(ctx, &lark.GetHelpdeskTicketCustomizedFieldListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskTicketCustomizedField(func(ctx context.Context, request *lark.GetHelpdeskTicketCustomizedFieldReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskTicketCustomizedFieldResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskTicketCustomizedField()

			_, _, err := moduleCli.GetHelpdeskTicketCustomizedField(ctx, &lark.GetHelpdeskTicketCustomizedFieldReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskCategory(func(ctx context.Context, request *lark.GetHelpdeskCategoryReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskCategoryResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskCategory()

			_, _, err := moduleCli.GetHelpdeskCategory(ctx, &lark.GetHelpdeskCategoryReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskCategoryList(func(ctx context.Context, request *lark.GetHelpdeskCategoryListReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskCategoryListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskCategoryList()

			_, _, err := moduleCli.GetHelpdeskCategoryList(ctx, &lark.GetHelpdeskCategoryListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskFAQ(func(ctx context.Context, request *lark.GetHelpdeskFAQReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskFAQResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskFAQ()

			_, _, err := moduleCli.GetHelpdeskFAQ(ctx, &lark.GetHelpdeskFAQReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskFAQList(func(ctx context.Context, request *lark.GetHelpdeskFAQListReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskFAQListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskFAQList()

			_, _, err := moduleCli.GetHelpdeskFAQList(ctx, &lark.GetHelpdeskFAQListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskGetHelpdeskFAQImage(func(ctx context.Context, request *lark.GetHelpdeskFAQImageReq, options ...lark.MethodOptionFunc) (*lark.GetHelpdeskFAQImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskGetHelpdeskFAQImage()

			_, _, err := moduleCli.GetHelpdeskFAQImage(ctx, &lark.GetHelpdeskFAQImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskSearchHelpdeskFAQ(func(ctx context.Context, request *lark.SearchHelpdeskFAQReq, options ...lark.MethodOptionFunc) (*lark.SearchHelpdeskFAQResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskSearchHelpdeskFAQ()

			_, _, err := moduleCli.SearchHelpdeskFAQ(ctx, &lark.SearchHelpdeskFAQReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskSubscribeHelpdeskEvent(func(ctx context.Context, request *lark.SubscribeHelpdeskEventReq, options ...lark.MethodOptionFunc) (*lark.SubscribeHelpdeskEventResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskSubscribeHelpdeskEvent()

			_, _, err := moduleCli.SubscribeHelpdeskEvent(ctx, &lark.SubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockHelpdeskUnsubscribeHelpdeskEvent(func(ctx context.Context, request *lark.UnsubscribeHelpdeskEventReq, options ...lark.MethodOptionFunc) (*lark.UnsubscribeHelpdeskEventResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockHelpdeskUnsubscribeHelpdeskEvent()

			_, _, err := moduleCli.UnsubscribeHelpdeskEvent(ctx, &lark.UnsubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Helpdesk

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicket(ctx, &lark.GetHelpdeskTicketReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DownloadHelpdeskTicketImage(ctx, &lark.DownloadHelpdeskTicketImageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AnswerHelpdeskTicketUserQuery(ctx, &lark.AnswerHelpdeskTicketUserQueryReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendHelpdeskTicketMessage(ctx, &lark.SendHelpdeskTicketMessageReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedFieldList(ctx, &lark.GetHelpdeskTicketCustomizedFieldListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedField(ctx, &lark.GetHelpdeskTicketCustomizedFieldReq{
				TicketCustomizedFieldID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategory(ctx, &lark.GetHelpdeskCategoryReq{
				ID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategoryList(ctx, &lark.GetHelpdeskCategoryListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQ(ctx, &lark.GetHelpdeskFAQReq{
				ID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQList(ctx, &lark.GetHelpdeskFAQListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQImage(ctx, &lark.GetHelpdeskFAQImageReq{
				ID:       "x",
				ImageKey: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchHelpdeskFAQ(ctx, &lark.SearchHelpdeskFAQReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeHelpdeskEvent(ctx, &lark.SubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UnsubscribeHelpdeskEvent(ctx, &lark.UnsubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Helpdesk
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.StartHelpdeskService(ctx, &lark.StartHelpdeskServiceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicket(ctx, &lark.GetHelpdeskTicketReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketList(ctx, &lark.GetHelpdeskTicketListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DownloadHelpdeskTicketImage(ctx, &lark.DownloadHelpdeskTicketImageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.AnswerHelpdeskTicketUserQuery(ctx, &lark.AnswerHelpdeskTicketUserQueryReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketMessageList(ctx, &lark.GetHelpdeskTicketMessageListReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SendHelpdeskTicketMessage(ctx, &lark.SendHelpdeskTicketMessageReq{
				TicketID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedFieldList(ctx, &lark.GetHelpdeskTicketCustomizedFieldListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskTicketCustomizedField(ctx, &lark.GetHelpdeskTicketCustomizedFieldReq{
				TicketCustomizedFieldID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategory(ctx, &lark.GetHelpdeskCategoryReq{
				ID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskCategoryList(ctx, &lark.GetHelpdeskCategoryListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQ(ctx, &lark.GetHelpdeskFAQReq{
				ID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQList(ctx, &lark.GetHelpdeskFAQListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetHelpdeskFAQImage(ctx, &lark.GetHelpdeskFAQImageReq{
				ID:       "x",
				ImageKey: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchHelpdeskFAQ(ctx, &lark.SearchHelpdeskFAQReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeHelpdeskEvent(ctx, &lark.SubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UnsubscribeHelpdeskEvent(ctx, &lark.UnsubscribeHelpdeskEventReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
