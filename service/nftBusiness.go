package service

import (
	"github.com/dometa/dao"
	"github.com/dometa/global"
	bo "github.com/dometa/model/bo"
	po "github.com/dometa/model/po"
	"github.com/dometa/model/res"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	Nfts_QueryFailure = errors.New("ntfs query failure")
)

type nftService interface {
	//GetBanner(banners *[]bo.BannerBO) error
	GetNfts(nfts *[]res.NftItem, category []string, pageIndex int) (bool, error)
	//GetNft(nft *bo.NftBO, token string) error
}

var (
	AIRDROP_STAGE_ACTION = 1
	GetBannerError       = errors.New("get Banner error")
)

type nftServiceImpl struct {
}

func NewNftService() nftService {
	return &nftServiceImpl{}
}

func (service nftServiceImpl) GetBanner(banners *[]bo.BannerBO) error {
	airDropDao := dao.NewAirportDao()
	var airDrop po.AirdropPO
	err := airDropDao.GetAirdropByStage(&airDrop, AIRDROP_STAGE_ACTION)
	if err != nil {
		global.Logger.WithField("nftService", "getBanner").
			WithField("getAirdrop", "error").
			Errorln("getAirdrop for banner error", err)
		return GetBannerError
	}
	nftDaoImpl := dao.NewNftDao()

	var nfts []po.NftPO
	err = nftDaoImpl.SelectNftBanner(&nfts, airDrop.TokenId)
	if err != nil {
		global.Logger.WithField("nftService", "getBanner").
			WithField("getNft", "error").
			Errorln("get Banner Nft error", err)
		return GetBannerError
	}

	return nil
}

func (service nftServiceImpl) GetNfts(nfts *[]res.NftItem, category []string, pageIndex int) (bool, error) {

	if pageIndex < 1 {
		pageIndex = 1
	}
	nftDao := dao.NewNftDao()
	var nftPOs = []po.NftPO{}
	pageSize := viper.GetInt("nft.pageSize")
	err := nftDao.SelectNft(&nftPOs, category, pageIndex, pageSize)

	if err != nil {
		global.Logger.WithField("nftService", "getNtfs").
			Errorln("query nfts error ==>", err)
		return false, Nfts_QueryFailure
	}
	hasMoreNfts := false
	nftReturnSize := len(nftPOs)
	if len(nftPOs) > pageSize {
		hasMoreNfts = true
		nftReturnSize = nftReturnSize - 1
	}

	for i, nftPO := range nftPOs {
		if i == nftReturnSize {
			continue
		}
		var nft res.NftItem
		nft.Name = nftPO.Name
		nft.Description = nftPO.Description

		nft.Theme = nftPO.Theme
		nft.Category = nftPO.Category

		nft.Author = nftPO.Author
		// icon url address
		nft.AuthorAddress = nftPO.AuthorAddress
		// nft url address
		nft.NftAddress = nftPO.NftAddress

		nft.Price = nftPO.Price
		nft.PriceUnit = nftPO.PriceUnit
		*nfts = append(*nfts, nft)
	}

	return hasMoreNfts, nil
}

func (service nftServiceImpl) GetNft(nft *bo.NftBO, token string) error {
	return nil
}
