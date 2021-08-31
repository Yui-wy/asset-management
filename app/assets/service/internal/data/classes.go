package data

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ClassRepo = (*classRepo)(nil)

type classRepo struct {
	data *Data
	log  *log.Helper
}

type Class struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"not null;uniqueIndex"`
	Pcode     string
	ClzInfo   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClassRepo(data *Data, logger log.Logger) biz.ClassRepo {
	return &classRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/classes")),
	}
}

func (repo *classRepo) GetClasses(ctx context.Context) ([]*biz.Class, error) {
	return nil, nil
}
func (repo *classRepo) CreateClasses(ctx context.Context, clz []*biz.Class) ([]*biz.Class, error) {
	return nil, nil
}
