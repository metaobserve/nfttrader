package service

import (
	"github.com/dometa/dao"
	"github.com/dometa/global"
	bo "github.com/dometa/model/bo"
	po "github.com/dometa/model/po"
	"github.com/pkg/errors"
)

type nftService interface {
	GetBanner(banners *[]bo.BannerBO) error
	GetNfts(nfts *[]bo.NftBO, themes []string) error
	GetNft(nft *bo.NftBO, token string) error
}

var (
	AIRDROP_STAGE_ACTION = 1
	GetBannerError       = errors.New("get Banner error")
)

type nftServiceImpl struct {
}

func NewNftBusiness() nftService {
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
	nftDaoImpl := dao.NewNftDao(global.MysqlClient)

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

func (service nftServiceImpl) GetNfts(nfts *[]bo.NftBO, themes []string) error {
	return nil
}

func (service nftServiceImpl) GetNft(nft *bo.NftBO, token string) error {
	return nil
}
