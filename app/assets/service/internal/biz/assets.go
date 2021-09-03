package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// a. ID
// b. 区域ID
// c. 资产编号 (自动生成)
// 	  i. 组成方式: A类B类C类4为流水码00 00 00 0000
//    流水号
// d. 详细位置
// e. 资产描述
// f. 资产图片(地址)
// g. 资产报价
// h. 购入日期
// i. 订单号
//    申请日期
// j. 入库日期
// l. 报废日期
// n. 资产状态标识(0:入库;1:入库申请中;2:采购中;3:采购申请中;4.报废;5:报废申请中)

type Asset struct {
	Id         uint64
	Classes    string
	Code       string
	AreaId     uint32
	Address    string
	AssetInfo  string
	PicUrl     string
	Price      float32
	OrderAt    int64
	OrderNum   string
	StateNum   int32
	State      string
	AppliedAt  int64
	StorageAt  int64
	ScrappedAt int64
}

type SearchConf struct {
	Classes      string
	Address      string
	StateNum     int32
	LowStorageAt int64
	UpStorageAt  int64
	OrderBy      string
	SortDesc     bool
	AreaId       []uint32
}

type AssetRepo interface {
	GetAsset(ctx context.Context, id uint64) (*Asset, error)
	ListAssets(ctx context.Context, conf *SearchConf, pageNum, pageSize int64) ([]*Asset, error)
	CreatAsset(ctx context.Context, a *Asset) (*Asset, error)
	DeleteAsset(ctx context.Context, id uint64) (bool, error)
	UpdateAsset(ctx context.Context, a *Asset) (*Asset, error)
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

func (ac *AssetUseCase) Get(ctx context.Context, id uint64) (*Asset, error) {
	return ac.repo.GetAsset(ctx, id)
}

func (ac *AssetUseCase) List(ctx context.Context, conf *SearchConf, pageNum, pageSize int64) ([]*Asset, error) {
	return ac.repo.ListAssets(ctx, conf, pageNum, pageSize)
}

func (ac *AssetUseCase) Create(ctx context.Context, a *Asset) (*Asset, error) {
	return ac.repo.CreatAsset(ctx, a)
}

func (ac *AssetUseCase) Delete(ctx context.Context, id uint64) (bool, error) {
	return ac.repo.DeleteAsset(ctx, id)
}

func (ac *AssetUseCase) Update(ctx context.Context, a *Asset) (*Asset, error) {
	return ac.repo.UpdateAsset(ctx, a)
}
