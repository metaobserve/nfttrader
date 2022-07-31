package req

// nft show on page. select by pageIndex
type WalletLoginRequest struct {
	Token   string `json:token`
	Address string `json:address`
}
