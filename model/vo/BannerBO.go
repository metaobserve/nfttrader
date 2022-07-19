package vo

type BannerVO struct {
	Name       string `json:name`
	Descrption string `json:descrption`
	// nft token
	Token          string `json:token`
	AirdropTokenId string `json:airdropTokenId`
	Address        string `json:address`
	Price          string `json:price`
	Author         string `json:author`
	Theme          string `json:theme`
}
