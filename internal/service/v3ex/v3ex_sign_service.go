package v3ex

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/segmentfault/pacman/log"
	"math/big"
)

const (
	_PrivateKey = "92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e"
)

type V3exRepo interface {
	ListToken(ctx context.Context) (tokens []*entity.ERC20Token, err error)
	AddToken(ctx context.Context, req *entity.ERC20Token) (err error)
	DeleteToken(ctx context.Context, id int64) (err error)
	UpdateToken(ctx context.Context, req *entity.ERC20Token) (err error)
}

type V3exService struct {
	v3exRepo   V3exRepo
	privateKey *ecdsa.PrivateKey
}

func NewV3exService(
	v3exRepo V3exRepo,
) *V3exService {
	privateKey, err := crypto.HexToECDSA(_PrivateKey)
	if err != nil {
		log.Errorf("[v3ex] Sign error: %v", err)
		panic(err)
	}
	return &V3exService{
		v3exRepo:   v3exRepo,
		privateKey: privateKey,
	}
}

func (svc *V3exService) Sign(ctx context.Context, req *schema.SignReq) (r, s, v string, err error) {
	// 生成哈希值
	signatureBytes, err := crypto.Sign(common.HexToHash(req.Data).Bytes(), svc.privateKey)
	if err != nil {
		fmt.Println("Error signing hash:", err)
		log.Errorf("[v3ex] Sign error: %v", err)
		return
	}
	// 将签名结果拼接成完整的签名字符串
	r = common.Bytes2Hex(signatureBytes[:32])
	s = common.Bytes2Hex(signatureBytes[32:64])
	v = common.Bytes2Hex([]byte{signatureBytes[64] + 27})
	log.Debugf("[v3ex] Sign success: r=%s, s=%s, v=%s", r, s, v)
	return r, s, v, nil
}

/**
 * CheckSign 签名 V3ex 签到数据
 * @Description: 签名 V3ex 数据
 * @receiver svc
 * @param ctx
 * @param req
 * @return signature
 */
func (svc *V3exService) CheckSign(ctx context.Context, req *schema.SignV3exReq) (signature string, err error) {
	// 获取链 ID
	var (
		chainId   = new(big.Int).SetInt64(req.ChainId).Bytes()
		newAmount = big.NewInt(req.Amount)
		to        = common.HexToAddress(req.To)
		date      = req.Date
	)
	// 根据交易数据生成交易 ID
	transId := crypto.Keccak256Hash(
		chainId,
		to.Bytes(),
		[]byte(date),
	).Hex()
	// 生成哈希值
	hashBytes := crypto.Keccak256Hash(
		chainId,
		[]byte(req.Method),
		common.HexToHash(transId).Bytes(),
		newAmount.Bytes(),
		to.Bytes(),
	).Bytes()
	signatureBytes, err := crypto.Sign(hashBytes, svc.privateKey)
	if err != nil {
		log.Errorf("[v3ex] Sign error: %v", err)
		return
	}
	// 将签名结果拼接成完整的签名字符串
	signature = "0x" + common.Bytes2Hex(signatureBytes[:64]) + common.Bytes2Hex([]byte{signatureBytes[64]})
	log.Debugf("[v3ex] Sign success: %s", signature)
	return signature, nil
}

/**
 * ListToken 获取 V3ex 支持的 ERC20 代币列表
 * @Description: 获取 V3ex 支持的 ERC20 代币列表 从文件中加载
 */
func (svc *V3exService) ListToken(ctx context.Context, req *schema.ListTokenReq) (tokens []*schema.ERC20Token, err error) {
	tokens = make([]*schema.ERC20Token, 0)
	tks, err := svc.v3exRepo.ListToken(ctx)
	if err != nil {
		return tokens, err
	}
	for _, v := range tks {
		tokens = append(tokens, &schema.ERC20Token{
			ID:          v.ID,
			Name:        v.Name,
			Symbol:      v.Symbol,
			Address:     v.Address,
			Decimals:    v.Decimals,
			LogoURI:     v.LogoURI,
			Description: v.Description,
		})
	}
	return tokens, nil
}

func (svc *V3exService) AddToken(ctx context.Context, req *schema.ERC20Token) (err error) {
	return svc.v3exRepo.AddToken(ctx, &entity.ERC20Token{
		Name:        req.Name,
		Symbol:      req.Symbol,
		Address:     req.Address,
		Decimals:    req.Decimals,
		LogoURI:     req.LogoURI,
		Description: req.Description,
	})
}

func (svc *V3exService) DeleteToken(ctx context.Context, req *schema.ERC20Token) (err error) {
	return svc.v3exRepo.DeleteToken(ctx, req.ID)
}

func (svc *V3exService) UpdateToken(ctx context.Context, req *schema.ERC20Token) (err error) {
	return svc.v3exRepo.UpdateToken(ctx, &entity.ERC20Token{
		ID:          req.ID,
		Name:        req.Name,
		Symbol:      req.Symbol,
		Address:     req.Address,
		Decimals:    req.Decimals,
		LogoURI:     req.LogoURI,
		Description: req.Description,
	})
}
