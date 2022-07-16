package dao

import (
	"github.com/dometa/bootstrap"
	"github.com/dometa/model/po"
)

type nftDao interface {
	Insert(nft po.NftPO) (bool, error)
	SelectNft(nfts []po.NftPO) error
}

type nftDaoImpl struct {
	context *bootstrap.Context
}

func (dao nftDaoImpl) Insert(nft po.NftPO) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (dao nftDaoImpl) SelectNft(nfts []po.NftPO) error {
	//TODO implement me
	panic("implement me")
}
