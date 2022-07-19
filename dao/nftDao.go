package dao

import (
	"github.com/dometa/global"
	"github.com/dometa/model/po"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type nftDao interface {
	Insert(nft po.NftPO) (bool, error)
	SelectNft(nfts *[]po.NftPO, airdropTokenId string, pageIndex int) error
	SelectNftBanner(nfts *[]po.NftPO, airdropTokenId string) error
}

type nftDaoImpl struct {
	sqlClient *sqlx.DB
}

func NewNftDao(sqlClient *sqlx.DB) nftDao {
	return &nftDaoImpl{
		sqlClient: sqlClient,
	}
}

func (dao nftDaoImpl) Insert(nft po.NftPO) (bool, error) {
	sqlstr := "insert into nft(name,description,token,airdropTokenId,status,createTime,updateTime) values(?,?,?,?,?,?,?)"
	ret, err := dao.sqlClient.Exec(sqlstr, nft.Name, nft.Description, nft.Token, nft.AirdropTokenId, nft.Status, nft.CreateTime, nft.UpdateTime)
	if err != nil {
		global.Logger.
			WithField("insertNft", "error").
			WithField("nft", nft).Errorln("insert nft error", err)
		return false, InsertError
	}
	affectRows, err := ret.RowsAffected()
	if err != nil {
		global.Logger.
			WithField("insertNft", "noAffectedError").WithField("nft", nft).Errorln("no affect error")
		return false, InsertError
	}
	if affectRows > 0 {
		return true, nil
	} else {
		return false, NothingInsertError
	}
}

func (dao nftDaoImpl) SelectNft(nfts *[]po.NftPO, airdropTokenId string, pageIndex int) error {
	pageSize := viper.GetInt("nft.pageSize")
	pageNum := pageSize * pageIndex
	actualPageSize := pageSize + 1
	err := dao.sqlClient.Select(nfts, "select * from nft where  status=1 and airdropTokenId= ? order by Id asc limits ?,?", airdropTokenId, pageNum, actualPageSize)
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

func (dao nftDaoImpl) SelectNftBanner(nfts *[]po.NftPO, airdropTokenId string) error {
	err := dao.sqlClient.Select(nfts, "select * from nft where  status=1 and airdropTokenId= ? and isBanner=true", airdropTokenId)
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
