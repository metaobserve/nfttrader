package req

// nft show on page. select by pageIndex
type NftRequest struct {
	Category  []string `json:category`
	PageIndex int      `json:pageIndex`
	PageSize  int      `json:pageSize`
}
