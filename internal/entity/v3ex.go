package entity

import "time"

type ERC20Token struct {
	ID int64 `json:"id,omitempty" xorm:"not null pk autoincr INT(11) id"`
	// Name 代币的名称
	Name string `json:"name,omitempty" xorm:"not null VARCHAR(255) name"`
	// Symbol 代币的简称
	Symbol string `json:"symbol,omitempty" xorm:"not null VARCHAR(255) symbol"`
	// Address 代币合约的地址
	Address string `json:"address,omitempty" xorm:"not null VARCHAR(255) address"`
	// Decimals 代币的小数位数
	Decimals int `json:"decimals,omitempty" xomr:"not null INT(11) decimals"`
	// LogoURI 代币的图标 URI
	LogoURI string `json:"logoURI,omitempty" xorm:"VARCHAR(255) logo_uri"`
	// Description 对代币的描述
	Description string `json:"description,omitempty" xorm:"VARCHAR(255) description"`
	// UpdatedAt 更新时间
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	// CreatedAt 创建时间
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
}

// TableName user role rel table name
func (ERC20Token) TableName() string {
	return "erc20_token"
}
