// Code generated by hertz generator.

package core_api

import (
	"context"
	"github.com/xh-polaris/meowchat-core-api/biz/adaptor"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	core_api "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
)

// GetMomentPreviews .
// @router /moment/get_moment_previews [GET]
func GetMomentPreviews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetMomentPreviewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.MomentService.GetMomentPreviews(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// GetMomentDetail .
// @router /moment/get_moment_detail [GET]
func GetMomentDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetMomentDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.MomentService.GetMomentDetail(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// NewMoment .
// @router /moment/new_moment [POST]
func NewMoment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewMomentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.MomentService.NewMoment(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// DeleteMoment .
// @router /moment/delete_moment [POST]
func DeleteMoment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteMomentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.MomentService.DeleteMoment(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}

// SearchMoment .
// @router /moment/search_moment [GET]
func SearchMoment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SearchMomentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.MomentService.SearchMoment(ctx, &req)
	adaptor.Return(ctx, c, &req, resp, err)
}
