package controller

import (
	"crypto/ecdsa"
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"math/big"
)

//import "github.com/apache/incubator-answer/internal/service/V3ex"

type V3exController struct {
	//V3exService *V3ex.V3exService
}

func NewV3exController(
// V3exService *V3ex.V3exService,
) *V3exController {
	return &V3exController{
		//V3exService: V3exService,
	}
}

// /user/tx/sign
// @Summary get tx sign
// @Description get tx sign
// @Tags V3ex
// @Accept json
// @Produce json
func (c *V3exController) Sign(ctx *gin.Context) {
	req := &schema.SignV3exReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	var (
		toAddress = common.HexToAddress(req.To)
		value     = big.NewInt(req.Amount)
		nonce     = uint64(0)
		gasLimit  = uint64(0)
		gasPrice  = big.NewInt(0)
		data      []byte
		//data          = []byte("we are v3ex!")
		privateKeyStr = "0x85d4d57a07fe43ee1a1df3a640f425ccfbc04bd170df21900b6e17793d2ee0b7"
		resp          = &schema.SignV3exResp{}
	)
	privateKey, err := StringToPrivateKey(privateKeyStr)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return

	}
	// TODO 对交易数据进行签名
	//1. 校验 buff
	//2. 校验金额是否正确
	// 生成交易信息
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	// 签名
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(req.ChainId)), privateKey)
	//3. 返回签名 hash
	resp.SignHash = signTx.Hash().String()
	handler.HandleResponse(ctx, err, resp)
}
func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func (c *V3exController) Verify(ctx *gin.Context) {

}

func (c *V3exController) ListToken(ctx *gin.Context) {

}
func (c *V3exController) AddToken(ctx *gin.Context) {
	//合约名、简称、地址、图标

}

func (c *V3exController) UpdateToken(ctx *gin.Context) {

}

func (c *V3exController) DeleteToken(ctx *gin.Context) {

}
