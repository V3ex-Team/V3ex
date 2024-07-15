package v3ex

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

const (
	_PrivateKey = "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e"
)

type V3exRepo interface {
}

type V3exService struct {
	v3exRepo V3exRepo
}

func NewV3exService(
	v3exRepo V3exRepo,
) *V3exService {
	return &V3exService{
		v3exRepo: v3exRepo,
	}
}

func Sign() {
	// 获取链 ID
	chainId := uint64(1) // 假设在主网上操作

	// 生成交易 ID
	to := common.HexToAddress("0x1234567890123456789012345678901234567890")
	date := big.NewInt(1618304400) // 2021年4月13日
	transId := crypto.Keccak256Hash(
		new(big.Int).SetUint64(chainId).Bytes(),
		to.Bytes(),
		date.Bytes(),
	).Hex()

	// 生成哈希值
	newAmount := big.NewInt(100)
	hash := crypto.Keccak256Hash(
		new(big.Int).SetUint64(chainId).Bytes(),
		[]byte("CHECKIN"),
		common.HexToHash(transId).Bytes(),
		newAmount.Bytes(),
		to.Bytes(),
	).Hex()

	privateKey, err := crypto.HexToECDSA(_PrivateKey)

	signatureBytes, err := crypto.Sign(common.HexToHash(hash).Bytes(), privateKey)
	if err != nil {
		fmt.Println("Error signing hash:", err)
		return
	}

	// 将签名结果拼接成完整的签名字符串
	signature := "0x" + common.Bytes2Hex(signatureBytes[:64]) + common.Bytes2Hex([]byte{signatureBytes[64]})
	fmt.Println("Signature:", signature)
}
