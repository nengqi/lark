package test

import (
	"context"
	"strings"
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_ContactUserUpdatedV3(t *testing.T) {
	as := assert.New(t)
	cli := AppAllPermission.Ins()

	t.Run("", func(t *testing.T) {
		cli.EventCallback.HandlerEventV2ContactUserUpdatedV3(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2ContactUserUpdatedV3) (string, error) {
			panic("")
		})

		req := ``
		resp := newFakeHTTPWriter()
		cli.EventCallback.ListenCallback(ctx, strings.NewReader(req), resp)

		as.Equal(500, resp.code)
		as.Equal(`{"err":"lark event unmarshal event_req  failed"}`, resp.str())
	})

	t.Run("", func(t *testing.T) {
		cli.EventCallback.HandlerEventV2ContactUserUpdatedV3(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2ContactUserUpdatedV3) (string, error) {
			panic("")
		})

		req := `{}`
		resp := newFakeHTTPWriter()
		cli.EventCallback.ListenCallback(ctx, strings.NewReader(req), resp)

		as.Equal(500, resp.code)
		as.Equal(`{"err":"must set verification token"}`, resp.str())
	})
}
