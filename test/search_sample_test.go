// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Search_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchCreateSearchDataSourceItem(func(ctx context.Context, request *lark.CreateSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.CreateSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchCreateSearchDataSourceItem()

			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchGetSearchDataSourceItem(func(ctx context.Context, request *lark.GetSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSourceItem()

			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchDeleteSearchDataSourceItem(func(ctx context.Context, request *lark.DeleteSearchDataSourceItemReq, options ...lark.MethodOptionFunc) (*lark.DeleteSearchDataSourceItemResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchDeleteSearchDataSourceItem()

			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchCreateSearchDataSource(func(ctx context.Context, request *lark.CreateSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.CreateSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchCreateSearchDataSource()

			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchGetSearchDataSource(func(ctx context.Context, request *lark.GetSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSource()

			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchUpdateSearchDataSource(func(ctx context.Context, request *lark.UpdateSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.UpdateSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchUpdateSearchDataSource()

			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchGetSearchDataSourceList(func(ctx context.Context, request *lark.GetSearchDataSourceListReq, options ...lark.MethodOptionFunc) (*lark.GetSearchDataSourceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchGetSearchDataSourceList()

			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockSearchDeleteSearchDataSource(func(ctx context.Context, request *lark.DeleteSearchDataSourceReq, options ...lark.MethodOptionFunc) (*lark.DeleteSearchDataSourceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockSearchDeleteSearchDataSource()

			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Search

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Search
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSourceItem(ctx, &lark.CreateSearchDataSourceItemReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceItem(ctx, &lark.GetSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSourceItem(ctx, &lark.DeleteSearchDataSourceItemReq{
				DataSourceID: "x",
				ItemID:       "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateSearchDataSource(ctx, &lark.CreateSearchDataSourceReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSource(ctx, &lark.GetSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateSearchDataSource(ctx, &lark.UpdateSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetSearchDataSourceList(ctx, &lark.GetSearchDataSourceListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteSearchDataSource(ctx, &lark.DeleteSearchDataSourceReq{
				DataSourceID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
