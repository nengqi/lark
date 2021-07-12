// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_AI_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAIRecognizeBasicImage(func(ctx context.Context, request *lark.RecognizeBasicImageReq, options ...lark.MethodOptionFunc) (*lark.RecognizeBasicImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeBasicImage()

			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAIRecognizeSpeechStream(func(ctx context.Context, request *lark.RecognizeSpeechStreamReq, options ...lark.MethodOptionFunc) (*lark.RecognizeSpeechStreamResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeSpeechStream()

			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAIRecognizeSpeechFile(func(ctx context.Context, request *lark.RecognizeSpeechFileReq, options ...lark.MethodOptionFunc) (*lark.RecognizeSpeechFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIRecognizeSpeechFile()

			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAITranslateText(func(ctx context.Context, request *lark.TranslateTextReq, options ...lark.MethodOptionFunc) (*lark.TranslateTextResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAITranslateText()

			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAIDetectTextLanguage(func(ctx context.Context, request *lark.DetectTextLanguageReq, options ...lark.MethodOptionFunc) (*lark.DetectTextLanguageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAIDetectTextLanguage()

			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.AI

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.AI
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeBasicImage(ctx, &lark.RecognizeBasicImageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechStream(ctx, &lark.RecognizeSpeechStreamReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RecognizeSpeechFile(ctx, &lark.RecognizeSpeechFileReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TranslateText(ctx, &lark.TranslateTextReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DetectTextLanguage(ctx, &lark.DetectTextLanguageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
