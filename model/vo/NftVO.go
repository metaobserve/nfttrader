package vo

// nft show on page. select by pageIndex
type NftVO struct {
	Name string `json:name`
	// nft token
	Token          string `json:token`
	AirdropTokenId string `json:airdropTokenId`
	Address        string `json:address`
	Price          string `json:price`
	Author         string `json:author`
	Theme          string `json:theme`
}
