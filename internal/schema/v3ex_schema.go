package schema

type SignV3exReq struct {
	ChainId int64  `json:"chain_id" form:"chain_id" binding:"required"`
	To      string `json:"to" form:"to" binding:"required"`
	Date    string `json:"date" form:"date" binding:"required"`
	Amount  int64  `json:"amount" form:"amount" binding:"required"`
}

type SignV3exResp struct {
	SignHash string `json:"sign_hash,omitempty"`
}
