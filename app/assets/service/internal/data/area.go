package data

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AreaRepo = (*areaRepo)(nil)

type areaRepo struct {
	data *Data
	log  *log.Helper
}

type Area struct {
	ID        uint32 `gorm:"primarykey;autoIncrement:true"`
	AreaInfo  string `gorm:"not null;"`
	IsDeleted bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewAreaRepo(data *Data, logger log.Logger) biz.AreaRepo {
	return &areaRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/area")),
	}
}

func (repo *areaRepo) GetArea(ctx context.Context, id uint32) (*biz.Area, error) {
	a := Area{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&a, id)
	repo.log.Debugf(result.Statement.SQL.String())
	if result.Error != nil {
		repo.log.Errorf(" GetArea. Error:%d", result.Error)
		return nil, result.Error
	}
	return &biz.Area{
		Id:       a.ID,
		AreaInfo: a.AreaInfo,
	}, nil
}

func (repo *areaRepo) GetAreasByIds(ctx context.Context, ids []uint32, pageNum, pageSize int64) ([]*biz.Area, error) {
	var as []Area
	result := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Where("is_deleted = false").
		Where("id IN ?", ids).
		Find(&as)
	if result.Error != nil {
		repo.log.Errorf(" GetAreasByIds. Error:%d", result.Error)
		return nil, result.Error
	}
	bas := make([]*biz.Area, 0)
	for _, a := range as {
		bas = append(bas, &biz.Area{
			Id:       a.ID,
			AreaInfo: a.AreaInfo,
		})
	}
	return bas, nil
}

func (repo *areaRepo) ListArea(ctx context.Context, pageNum, pageSize int64) ([]*biz.Area, error) {
	var as []Area
	result := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Where("is_deleted = false").
		Find(&as)
	if result.Error != nil {
		repo.log.Errorf(" ListArea. %d", result.Error)
		return nil, result.Error
	}
	bas := make([]*biz.Area, 0)
	for _, a := range as {
		bas = append(bas, &biz.Area{
			Id:       a.ID,
			AreaInfo: a.AreaInfo,
		})
	}
	return bas, nil
}

func (repo *areaRepo) CreateArea(ctx context.Context, a *biz.Area) (*biz.Area, error) {
	ac := Area{
		AreaInfo:  a.AreaInfo,
		IsDeleted: false,
	}
	result := repo.data.db.WithContext(ctx).Create(&ac)
	if result.Error != nil {
		repo.log.Errorf(" CreateArea. %d", result.Error)
		return nil, result.Error
	}
	return &biz.Area{
		Id:       ac.ID,
		AreaInfo: ac.AreaInfo,
	}, result.Error
}

func (repo *areaRepo) UpdateArea(ctx context.Context, a *biz.Area) (*biz.Area, error) {
	au := Area{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&au, a.Id)
	if result.Error != nil {
		repo.log.Errorf(" UpdateArea1. Error:%d", result.Error)
		return nil, result.Error
	}
	au.AreaInfo = a.AreaInfo
	result = repo.data.db.WithContext(ctx).Save(&au)
	if result.Error != nil {
		repo.log.Errorf(" UpdateArea2. Error:%d", result.Error)
		return nil, result.Error
	}
	return &biz.Area{
		Id:       au.ID,
		AreaInfo: au.AreaInfo,
	}, nil
}

func (repo *areaRepo) DeleteArea(ctx context.Context, id uint32) (bool, error) {
	au := Area{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&au, id)
	if result.Error != nil {
		repo.log.Errorf(" DeleteArea1. Error:%d", result.Error)
		return false, result.Error
	}
	au.IsDeleted = true
	au.DeletedAt = time.Now()
	result = repo.data.db.WithContext(ctx).Save(&au)
	if result.Error != nil {
		repo.log.Errorf(" DeleteArea2. Error:%d", result.Error)
		return false, result.Error
	}
	return true, nil
}
