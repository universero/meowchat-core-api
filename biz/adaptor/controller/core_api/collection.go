// Code generated by hertz generator.

package core_api

import (
	"context"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/obs/log"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
	"github.com/xh-polaris/meowchat-core-api/provider"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCatPreviews .
// @router /collection/get_cat_previews [GET]
func GetCatPreviews(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetCatPreviewsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	user := p.Extractor.ExtractUserMeta(ctx, c)
	log.Info(util.JSONF(user))
	resp, err := p.CollectionService.GetCatPreviews(ctx, &req)
	p.Handler.HandlerError(ctx, err, c)

	c.JSON(consts.StatusOK, resp)
}

// GetCatDetail .
// @router /collection/get_cat_detail [GET]
func GetCatDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetCatDetailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.GetCatDetailResp)

	c.JSON(consts.StatusOK, resp)
}

// NewCat .
// @router /collection/new_cat [POST]
func NewCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.NewCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.NewCatResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteCat .
// @router /collection/delete_cat [POST]
func DeleteCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.DeleteCatResp)

	c.JSON(consts.StatusOK, resp)
}

// SearchCat .
// @router /collection/search_cat [GET]
func SearchCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.SearchCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.SearchCatResp)

	c.JSON(consts.StatusOK, resp)
}

// CreateImage .
// @router /collection/create_image [POST]
func CreateImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.CreateImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.CreateImageResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteImage .
// @router /collection/delete_image [POST]
func DeleteImage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.DeleteImageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.DeleteImageResp)

	c.JSON(consts.StatusOK, resp)
}

// GetImageByCat .
// @router /collection/get_image_by_cat [GET]
func GetImageByCat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core_api.GetImageByCatReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core_api.GetImageByCatResp)

	c.JSON(consts.StatusOK, resp)
}
