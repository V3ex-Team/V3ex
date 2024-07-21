package controller

import (
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/v3ex"
	"github.com/gin-gonic/gin"
)

//import "github.com/apache/incubator-answer/internal/service/V3ex"

type V3exController struct {
	v3exService *v3ex.V3exService
}

func NewV3exController(
	v3exService *v3ex.V3exService,
) *V3exController {
	return &V3exController{v3exService: v3exService}
}

// CheckInSign sign v3ex checkin data
// @Summary sign v3ex checkin data
// @Description sign v3ex checkin data
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.SignV3exReq true "SignV3exReq"
// @Success 200 {object} handler.RespBody{data=schema.SignV3exResp}
// @Router /answer/api/v1/v3ex/checkin/sign [post]
func (c *V3exController) CheckInSign(ctx *gin.Context) {
	req := &schema.SignV3exReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	resp := &schema.SignV3exResp{}
	sign, err := c.v3exService.CheckSign(ctx, req)
	//3. 返回签名 hash
	resp.SignHash = sign
	handler.HandleResponse(ctx, err, resp)
}

// Sign sign v3ex data
// @Summary sign v3ex data
// @Description sign v3ex data
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.SignReq true "SignReq"
// @Success 200 {object} handler.RespBody{data=schema.SignResp}
// @Router /answer/api/v1/v3ex/sign [post]
func (c *V3exController) Sign(ctx *gin.Context) {
	req := &schema.SignReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	resp := &schema.SignResp{}
	r, s, v, err := c.v3exService.Sign(ctx, req)
	//3. 返回签名 hash
	resp.R = r
	resp.S = s
	resp.V = v
	handler.HandleResponse(ctx, err, resp)
}

// ListToken list v3ex support erc20 token
// @Summary list v3ex support erc20 token
// @Description list v3ex support erc20 token
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ListTokenReq true "ListTokenReq"
// @Success 200 {object} handler.RespBody{data=schema.ListTokenResp}
// @Router /answer/api/v1/v3ex/tokens [get]
func (c *V3exController) ListToken(ctx *gin.Context) {
	req := &schema.ListTokenReq{}
	resp := &schema.ListTokenResp{}
	//1. 获取 token 列表
	tokens, err := c.v3exService.ListToken(ctx, req)
	//2. 返回 token 列表
	if len(tokens) != 0 {
		resp.List = tokens
	}
	handler.HandleResponse(ctx, err, resp)
}

// AddToken add v3ex support erc20 token
// @Summary add v3ex support erc20 token
// @Description add v3ex support erc20 token
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ERC20Token true "ERC20Token"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/v3ex/token [post]
func (c *V3exController) AddToken(ctx *gin.Context) {
	req := &schema.ERC20Token{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := c.v3exService.AddToken(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// DeleteToken delete v3ex support erc20 token
// @Summary delete v3ex support erc20 token
// @Description delete v3ex support erc20 token
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ERC20Token true "ERC20Token"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/v3ex/token [delete]
func (c *V3exController) DeleteToken(ctx *gin.Context) {
	req := &schema.ERC20Token{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := c.v3exService.DeleteToken(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateToken update v3ex support erc20 token
// @Summary update v3ex support erc20 token
// @Description update v3ex support erc20 token
// @Tags V3ex
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ERC20Token true "ERC20Token"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/v3ex/token [put]
func (c *V3exController) UpdateToken(ctx *gin.Context) {
	req := &schema.ERC20Token{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := c.v3exService.UpdateToken(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
