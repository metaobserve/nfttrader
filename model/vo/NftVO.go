package vo

// nft show on page. select by pageIndex
type NftVO struct {
	Name           string `json:name`
	Description    string `json:description`
	Address        string `json:address`
	AirdropTokenId string `json:airdropTokenId`
}
