// Code generated by hertz generator.

package core_api

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
)

// ApplySignedUrl .
// @router /sts/apply_signed_url [POST]
func ApplySignedUrl(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ApplySignedUrlReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.StsService.ApplySignedUrl(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}

// ApplySignedUrlAsCommunity .
// @router /sts/apply_signed_url_as_community [POST]
func ApplySignedUrlAsCommunity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.ApplySignedUrlAsCommunityReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.StsService.ApplySignedUrlAsCommunity(ctx, &req, adaptor.ExtractUserMeta(ctx, c))
	adaptor.PostProcess(ctx, c, &req, resp, err)
}
