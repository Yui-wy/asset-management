package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// a. ID
// b. 区域ID
// c. 资产编号 (自动生成)
// 	i. 组成方式: A类B类C类4为流水码00 00 00 0000
// d. 详细位置
// e. 资产描述
// f. 资产图片(地址)
// g. 资产报价
// h. 购入日期
// i. 订单号
// j. 入库日期
// k. 入库申请流水号（申请单号）
// l. 报废日期
// m. 报废申请流水号（报废单号）
// n. 资产状态标识(0:入库;1:入库申请中;2:采购中;3:采购申请中;4.报废;5:报废申请中)

type Asset struct {
	Id          uint64
	AreaId      uint
	Code        string
	Address     string
	AssetInfo   string
	Price       float64
	orderAt     int64
	orderNum    string
	StorageAt   int64
	StorageNum  string
	ScrappedAt  int64
	ScrappedNum string
	stateNum    int32
	state       string
}

type AssetRepo interface {
}

type AssetUseCase struct {
	repo AssetRepo
	log  *log.Helper
}

func NewAssetUseCase(repo AssetRepo, logger log.Logger) *AssetUseCase {
	return &AssetUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/aassets")),
	}
}

