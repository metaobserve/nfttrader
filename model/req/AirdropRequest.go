package req

// show airdrop on page
type AirdropRequest struct {
	Name        string `json:name`
	Description string `json:description`
	TokenId     string `json:tokenId`
}
