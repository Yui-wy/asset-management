package data

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type areaRepo struct {
	data *Data
	log  *log.Helper
}

type Area struct {
	ID        uint   `gorm:"primarykey"`
	AreaInfo  string `gorm:"not null;uniqueIndex:info"`
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
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false")
	return nil, nil
}

func (repo *areaRepo) ListArea(ctx context.Context) ([]*biz.Area, error) {
	return nil, nil
}

func (repo *areaRepo) CreateArea(ctx context.Context, a *biz.Area) (*biz.Area, error) {
	return nil, nil
}

func (repo *areaRepo) UpdateArea(ctx context.Context, a *biz.Area) (*biz.Area, error) {
	return nil, nil
}

func (repo *areaRepo) DeleteArea(ctx context.Context, id uint32) (bool, error) {
	return false, nil
}
