package data

import (
	"context"
	"fmt"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AssetRepo = (*assetRepo)(nil)

type assetRepo struct {
	data *Data
	log  *log.Helper
}

type Asset struct {
	ID         uint64 `gorm:"primarykey"`
	Classes    string `gorm:"not null"`
	AreaId     uint32 `gorm:"not null"`
	SuffCode   int64  `gorm:"not null"`
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
		repo.log.Errorf(" GetAsset. Error:%d", result.Error)
		return nil, result.Error
	}
	return repo.setbizAsset(&a), nil
}

func (repo *assetRepo) ListAssets(ctx context.Context, conf *biz.SearchConf, pageNum, pageSize int64) ([]*biz.Asset, error) {
	var as []Asset
	result := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize)))
	if !inspection.IsZeros(conf.AreaId) {
		err := errors.New(500, "AreaId is nil", "please set areaId")
		repo.log.Errorf(" ListForm1. Error:%d", err)
		return nil, err
	}
	result = result.Where("area_id IN ?", conf.AreaId)
	if !inspection.IsZeros(conf.Classes) {
		result = result.Where("INSTR(classes, ?) > 0", conf.Classes)
	}
	if !inspection.IsZeros(conf.Address) {
		result = result.Where("INSTR(address, ?) > 0", conf.Address)
	}
	if !inspection.IsZeros(conf.StateNum) {
		result = result.Where("state_num = ?", conf.StateNum)
	}
	if !inspection.IsZeros(conf.LowStorageAt) {
		result = result.Where("storage_at >= ?", conf.LowStorageAt)
	}
	if !inspection.IsZeros(conf.UpStorageAt) {
		result = result.Where("storage_at =< ?", conf.UpStorageAt)
	}
	if !inspection.IsZeros(conf.OrderBy) {
		if conf.SortDesc {
			result.Order(fmt.Sprintf("%s desc", conf.OrderBy))
		} else {
			result.Order(fmt.Sprintf("%s asc", conf.OrderBy))
		}
	} else {
		result.Order("suff_code asc")
	}
	result = result.Find(&as)
	if result.Error != nil {
		repo.log.Errorf(" ListAssets. Error:%d", result.Error)
		return nil, result.Error
	}
	bas := make([]*biz.Asset, 0)
	for _, a := range as {
		bas = append(bas, repo.setbizAsset(&a))
	}
	return bas, nil
}

func (repo *assetRepo) CreatAsset(ctx context.Context, ba *biz.Asset) (*biz.Asset, error) {
	// code 自动生成
	// 生成方式: areaId(3位) + class + 搜索同类型最大+1
	// var code string = xxx
	su := map[string]interface{}{}
	result := repo.data.db.WithContext(ctx).
		Model(&Asset{}).
		Select("MAX(suff_code) as maxcode").
		Where("classes = ? AND area_id = ?", ba.Classes, ba.AreaId).
		Take(&su)
	if result.Error != nil {
		repo.log.Errorf(" CreatAsset1. Error:%d", result.Error)
		return nil, result.Error
	}
	var suffCode int64
	if su["maxcode"] == nil {
		suffCode = 0
	} else {
		suffCode = su["maxcode"].(int64)
	}
	suffCode = suffCode + 1
	a := Asset{
		Classes:    ba.Classes,
		AreaId:     ba.AreaId,
		SuffCode:   suffCode,
		Code:       fmt.Sprintf("%03d-%s-%05d", ba.AreaId, ba.Classes, suffCode),
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
		repo.log.Errorf(" CreatAsset2. Error:%d", result.Error)
		return nil, result.Error
	}
	aa := Asset{}
	result = repo.data.db.WithContext(ctx).First(&aa, a.ID)
	if result.Error != nil {
		repo.log.Errorf(" CreatAsset3. Error:%d", result.Error)
		return nil, result.Error
	}
	return repo.setbizAsset(&aa), nil
}

func (repo *assetRepo) DeleteAsset(ctx context.Context, id uint64) (bool, error) {
	result := repo.data.db.WithContext(ctx).Delete(&Asset{}, id)
	if result.Error != nil {
		repo.log.Errorf(" DeleteAsset. Error:%d", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (repo *assetRepo) UpdateAsset(ctx context.Context, ba *biz.Asset) (*biz.Asset, error) {
	a := Asset{
		ID: ba.Id,
	}
	result := repo.data.db.WithContext(ctx).Model(&a).Updates(Asset{
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
	})
	if result.Error != nil {
		repo.log.Errorf(" UpdateAsset1. Error:%d", result.Error)
		return nil, result.Error
	}
	result = repo.data.db.WithContext(ctx).First(&a)
	if result.Error != nil {
		repo.log.Errorf(" UpdateAsset2. Error:%d", result.Error)
		return nil, result.Error
	}
	return repo.setbizAsset(&a), nil
}

func (repo *assetRepo) setbizAsset(a *Asset) *biz.Asset {
	state, ok := setting.ASSETS_STATE_MAP[a.StateNum]
	if !ok {
		state = "未知问题"
	}
	return &biz.Asset{
		Id:         a.ID,
		Classes:    a.Classes,
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
