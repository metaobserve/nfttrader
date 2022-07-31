package dao

import (
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	NFT_SelectError = errors.New("nft table select error")
)

type nftDao interface {
	//Insert(nft po.NftPO) (bool, error)
	SelectNft(nfts *[]po.NftPO, category []string, pageIndex int, pageSize int) error
	SelectNftBanner(nfts *[]po.NftPO, airdropTokenId string) error
}

type nftDaoImpl struct {
}

func NewNftDao() nftDao {
	return &nftDaoImpl{}
}

//func (dao nftDaoImpl) Insert(nft po.NftPO) (bool, error) {
//	sqlstr := "insert into nft(name,description,token,airdropTokenId,status,createTime,updateTime) values(?,?,?,?,?,?,?)"
//	ret, err := global.MysqlClient.Exec(sqlstr, nft.Name, nft.Description, nft.Token, nft.AirdropTokenId, nft.Status, nft.CreateTime, nft.UpdateTime)
//	if err != nil {
//		global.Logger.
//			WithField("insertNft", "error").
//			WithField("nft", nft).Errorln("insert nft error", err)
//		return false, InsertError
//	}
//	affectRows, err := ret.RowsAffected()
//	if err != nil {
//		global.Logger.
//			WithField("insertNft", "noAffectedError").WithField("nft", nft).Errorln("no affect error")
//		return false, InsertError
//	}
//	if affectRows > 0 {
//		return true, nil
//	} else {
//		return false, NothingInsertError
//	}
//}

func (dao nftDaoImpl) SelectNft(nfts *[]po.NftPO, category []string, pageIndex int, pageSize int) error {

	pageIndex = pageIndex - 1
	pageNum := pageSize * pageIndex
	actualPageSize := pageSize + 1

	allQuery := "select id,name,description,status,theme,category,price,priceUnit,nftAddress,author,authorAddress,createTime,updateTime from nft where  status=1 order by Id asc limit ?,?"
	categoryQuery := "select  id,name,description,status,theme,category,price,priceUnit,nftAddress,author,authorAddress,createTime,updateTime  from nft where  status=1 and category in (?) order by Id asc limit ?,?"

	var err error
	if len(category) > 0 {
		query, args, err := sqlx.In(categoryQuery, category, pageNum, actualPageSize)
		if err != nil {
			global.Logger.
				WithField("selectNft", "argsError").
				WithField("category", category).Errorln("get nfts with args error =>", err)
			return NFT_SelectError
		}
		rebindQuery := global.MysqlClient.Rebind(query)
		err = global.MysqlClient.Select(nfts, rebindQuery, args...)
		println(err)
	} else {
		err = global.MysqlClient.Select(nfts, allQuery, pageNum, actualPageSize)
	}

	if err != nil {
		global.Logger.
			WithField("selectNft", "error").
			WithField("category", category).Errorln("get nfts with category error", err)
		return NFT_SelectError
	}

	return nil
}

func (dao nftDaoImpl) SelectNftBanner(nfts *[]po.NftPO, airdropTokenId string) error {
	err := global.MysqlClient.Select(nfts, "select * from nft where  status=1 and airdropTokenId= ? and isBanner=true", airdropTokenId)
	if err != nil {
		global.Logger.
			WithField("selectNft", "error").
			WithField("airdropTokenId", airdropTokenId).Errorln("get nfts with airdropTokenId", err)
		return SelectError
	} else {
		l := len(*nfts)
		println("nftlen => ", l)
		return nil
	}
}
