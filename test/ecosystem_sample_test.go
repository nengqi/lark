// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Ecosystem_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Ecosystem

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetEcosystemBindAwemeUser(ctx, &lark.GetEcosystemBindAwemeUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Ecosystem

		t.Run("", func(t *testing.T) {
			cli.Mock().MockEcosystemGetEcosystemBindAwemeUser(func(ctx context.Context, request *lark.GetEcosystemBindAwemeUserReq, options ...lark.MethodOptionFunc) (*lark.GetEcosystemBindAwemeUserResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockEcosystemGetEcosystemBindAwemeUser()

			_, _, err := moduleCli.GetEcosystemBindAwemeUser(ctx, &lark.GetEcosystemBindAwemeUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Ecosystem

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetEcosystemBindAwemeUser(ctx, &lark.GetEcosystemBindAwemeUserReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Ecosystem
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetEcosystemBindAwemeUser(ctx, &lark.GetEcosystemBindAwemeUserReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
