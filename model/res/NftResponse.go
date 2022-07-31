package res

import "github.com/dometa/model"

type NftItem struct {
	Name        string `json:name`
	Description string `json:description`

	Theme    string `json:theme`
	Category string `json:category`

	Author string `json:author"`
	// icon url address
	AuthorAddress string `json:authorAddress`
	// nft url address
	NftAddress string `json:nftAddress`

	Price     string `json:price`
	PriceUnit string `json:priceUnit`
}

type NftResponse struct {
	Nfts        []NftItem        `json:nfts`
	HasMoreNFT  bool             `json:hasMoreNFT`
	Status      model.StatusType `json:status`
	Description string           `json:description`
}
