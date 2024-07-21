package schema

type SignV3exReq struct {
	ChainId int64  `json:"chain_id" form:"chain_id" binding:"required"`
	To      string `json:"to" form:"to" binding:"required"`
	Date    string `json:"date" form:"date" binding:"required"`
	Amount  int64  `json:"amount" form:"amount" binding:"required"`
	Method  string `json:"method" form:"method" binding:"required"`
}

type SignV3exResp struct {
	SignHash string `json:"sign_hash,omitempty"`
}

type (
	SignReq struct {
		Data string `json:"data,omitempty" form:"data" binding:"required"`
	}
	SignResp struct {
		R string `json:"r,omitempty"`
		S string `json:"s,omitempty"`
		V string `json:"v,omitempty"`
	}
)

type (
	ListTokenReq struct {
	}
	ListTokenResp struct {
		List []*ERC20Token `json:"list,omitempty"`
	}
	ERC20Token struct {
		ID int64 `json:"id,omitempty"`
		// Name 代币的名称
		Name string `json:"name,omitempty"`
		// Symbol 代币的简称
		Symbol string `json:"symbol,omitempty"`
		// Address 代币合约的地址
		Address string `json:"address,omitempty"`
		// Decimals 代币的小数位数
		Decimals int `json:"decimals,omitempty"`
		// LogoURI 代币的图标 URI
		LogoURI string `json:"logoURI,omitempty"`
		// Description 对代币的描述
		Description string `json:"description,omitempty"`
	}
)
