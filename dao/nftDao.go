package dao

//
//import (
//	"github.com/dometa/bootstrap"
//	"github.com/dometa/model/po"
//	"github.com/spf13/viper"
//)
//
//type nftDao interface {
//	Insert(nft po.NftPO) (bool, error)
//	SelectNft(nfts *[]po.NftPO, airdropTokenId string, pageIndex int) error
//}
//
//type nftDaoImpl struct {
//	context *bootstrap.Context
//}
//
//func NewNftDao(context *bootstrap.Context) nftDao {
//	return &nftDaoImpl{
//		context: context,
//	}
//}
//
//func (dao nftDaoImpl) Insert(nft po.NftPO) (bool, error) {
//	sqlstr := "insert into nft(name,description,token,airdropTokenId,status,createTime,updateTime) values(?,?,?,?,?,?,?)"
//	ret, err := dao.context.MysqlClient.Exec(sqlstr, nft.Name, nft.Description, nft.Token, nft.AirdropTokenId, nft.Status, nft.CreateTime, nft.UpdateTime)
//	if err != nil {
//		dao.context.Logger.
//			WithField("insertNft", "error").
//			WithField("nft", nft).Errorln("insert nft error", err)
//		return false, InsertError
//	}
//	affectRows, err := ret.RowsAffected()
//	if err != nil {
//		dao.context.Logger.WithField("insertNft", "noAffectedError").WithField("nft", nft).Errorln("no affect error")
//		return false, InsertError
//	}
//	if affectRows > 0 {
//		return true, nil
//	} else {
//		return false, NothingInsertError
//	}
//}
//
//func (dao nftDaoImpl) SelectNft(nfts *[]po.NftPO, airdropTokenId string, pageIndex int) error {
//	pageSize := viper.GetInt("nft.pageSize")
//	pageNum := pageSize * pageIndex
//	actualPageSize := pageSize + 1
//	err := dao.context.MysqlClient.Select(nfts, "select * from nft where  status=1 and airdropTokenId= ? order by Id asc limits ?,?", airdropTokenId, pageNum, actualPageSize)
//	if err != nil {
//		dao.context.Logger.
//			WithField("selectNft", "error").
//			WithField("airdropTokenId", airdropTokenId).Errorln("get nfts with airdropTokenId", err)
//		return SelectError
//	} else {
//		l := len(*nfts)
//		println("nftlen => ", l)
//		return nil
//	}
//}
