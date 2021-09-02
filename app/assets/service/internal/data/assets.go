package data

import (
	"context"
	"fmt"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AssetRepo = (*assetRepo)(nil)

var STATE_MAP = map[int32]string{
	0: "库内",
	1: "入库申请中",
	2: "采购中",
	3: "采购申请中",
	4: "报废",
	5: "报废申请中",
	6: "",
	7: "",
	8: "",
	9: "",
}

type assetRepo struct {
	data *Data
	log  *log.Helper
}

type Asset struct {
	ID         uint64 `gorm:"primarykey"`
	Classes    string `gorm:"not null"`
	AreaId     uint32 `gorm:"not null"`
	SuffCode   uint32 `gorm:"not null"`
	Code       string `gorm:"not null"`
	Address    string
	AssetInfo  string
	PicUrl     string
	Price      float32
	OrderAt    int64
	OrderNum   string
	StateNum   int32
	AppliedAt  int64
	StorageAt  int64
	ScrappedAt int64
	// ====================
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAssetRepo(data *Data, logger log.Logger) biz.AssetRepo {
	return &assetRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/assets")),
	}
}

func (repo *assetRepo) GetAsset(ctx context.Context, id uint64) (*biz.Asset, error) {
	a := Asset{}
	result := repo.data.db.WithContext(ctx).First(&a, id)
	if result.Error != nil {
		repo.log.Errorf("GetAsset error. Error:%d", result.Error)
		return nil, result.Error
	}
	return repo.setbizAsset(&a), nil
}

func (repo *assetRepo) ListAssets(ctx context.Context,
	pageNum int64,
	pageSize int64,
	code string,
	address string,
	stateNum int32,
	lowStorageAt int64,
	upStorageAt int64,
	orderBy string,
	areaId uint32,
) ([]*biz.Asset, error) {
	return nil, nil
}

func (repo *assetRepo) CreatAsset(ctx context.Context, ba *biz.Asset) (*biz.Asset, error) {
	// code 自动生成
	// 生成方式: areaId(3位) + class + 搜索同类型最大+1
	// var code string = xxx
	su := map[string]interface{}{}
	result := repo.data.db.WithContext(ctx).
		Model(&Asset{}).
		Select("MAX(suff_code) as maxcode").
		Where("classes = ? AND area_id = ?").
		Take(&su)
	if result.Error != nil {
		repo.log.Errorf("CreatAsset error. Error:%d", result.Error)
		return nil, result.Error
	}
	suffCode := su["maxcode"].(uint32)
	a := Asset{
		Classes:    ba.Classes,
		AreaId:     ba.AreaId,
		Code:       fmt.Sprintf("%03d-%s-%05d", ba.AreaId, ba.Classes, suffCode+1),
		Address:    ba.Address,
		AssetInfo:  ba.AssetInfo,
		PicUrl:     ba.PicUrl,
		Price:      ba.Price,
		OrderAt:    ba.OrderAt,
		OrderNum:   ba.OrderNum,
		StateNum:   ba.StateNum,
		AppliedAt:  ba.AppliedAt,
		StorageAt:  ba.StorageAt,
		ScrappedAt: ba.ScrappedAt,
	}
	result = repo.data.db.WithContext(ctx).Create(&a)
	if result.Error != nil {
		repo.log.Errorf("CreatAsset error. Error:%d", result.Error)
		return nil, result.Error
	}
	return repo.setbizAsset(&a), nil
}

func (repo *assetRepo) DeleteAsset(ctx context.Context, id uint64) (bool, error) {
	result := repo.data.db.WithContext(ctx).Delete(&Asset{}, id)
	if result.Error != nil {
		repo.log.Errorf("DeleteAsset error. Error:%d", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (repo *assetRepo) UpdateAsset(ctx context.Context, ba *biz.Asset) (*biz.Asset, error) {
	return nil, nil
}

func (repo *assetRepo) setbizAsset(a *Asset) *biz.Asset {
	state, ok := STATE_MAP[a.StateNum]
	if !ok {
		a.CreatedAt.Unix()
		state = "未知问题"
	}
	return &biz.Asset{
		Code:       a.Code,
		AreaId:     a.AreaId,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      state,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}
}
