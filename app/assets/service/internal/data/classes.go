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
	ID        uint64 `gorm:"primarykey"`
	Code      string `gorm:"not null;uniqueIndex"`
	Level     uint32 `gorm:"not null;autoIncrement:false"`
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
	var clz []Class
	result := repo.data.db.WithContext(ctx).Find(&clz)
	if result.Error != nil {
		repo.log.Errorf("GetClasses error. Error:%d", result.Error)
		return nil, result.Error
	}
	bc := make([]*biz.Class, 0)
	for _, c := range clz {
		bc = append(bc, &biz.Class{
			Id:      c.ID,
			Code:    c.Code,
			Level:   c.Level,
			Pcode:   c.Pcode,
			ClzInfo: c.ClzInfo,
		})
	}
	return bc, nil
}
func (repo *classRepo) CreateClasses(ctx context.Context, clz []*biz.Class) ([]*biz.Class, error) {
	// 删除全部的类型重新导入
	result := repo.data.db.WithContext(ctx).Where("1 = 1").Delete(&Asset{})
	if result.Error != nil {
		repo.log.Errorf("CreateClasses error. Error:%d", result.Error)
		return nil, result.Error
	}
	cs := make([]*Class, 0)
	for _, c := range clz {
		cs = append(cs, &Class{
			Code:    c.Code,
			Level:   c.Level,
			Pcode:   c.Pcode,
			ClzInfo: c.ClzInfo,
		})
	}
	result = repo.data.db.WithContext(ctx).Create(cs)
	if result.Error != nil {
		repo.log.Errorf("CreateClasses error. Error:%d", result.Error)
		return nil, result.Error
	}
	bcs := make([]*biz.Class, 0)
	for _, c := range cs {
		bcs = append(bcs, &biz.Class{
			Id:      c.ID,
			Code:    c.Code,
			Pcode:   c.Pcode,
			Level:   c.Level,
			ClzInfo: c.ClzInfo,
		})
	}
	return bcs, nil
}
