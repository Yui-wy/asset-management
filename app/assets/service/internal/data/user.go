package data

import (
	"context"
	"fmt"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	Uid           uint64 `gorm:"primarykey;;autoIncrement:false"`
	Power         int32  `gorm:"not null"`
	AreaTableName string `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type AdminArea struct {
	ID        uint64 `gorm:"primarykey"`
	Uid       uint64 `gorm:"not null;index"`
	Aid       uint32 `gorm:"not null;index"`
	CreatedAt time.Time
}

func (AdminArea) TableName() string {
	return setting.TABLE_MAP[setting.AREA_ADMIN_USER]
}

type UserArea struct {
	ID        uint64 `gorm:"primarykey"`
	Uid       uint64 `gorm:"not null;uniqueIndex"`
	Aid       uint32 `gorm:"not null;index"`
	CreatedAt time.Time
}

func (UserArea) TableName() string {
	return setting.TABLE_MAP[setting.AREA_USER]
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (repo *userRepo) GetUser(ctx context.Context, uid uint64) (*biz.User, error) {
	/* 先得到user, 得到areaId */
	u := User{}
	result := repo.data.db.WithContext(ctx).First(&u, uid)
	if result.Error != nil {
		repo.log.Errorf(" GetUser1. Error:%d", result.Error)
		return nil, result.Error
	}
	if u.Power == setting.SUPER_ADMIN_USER {
		return &biz.User{
			Uid:     u.Uid,
			Power:   u.Power,
			AreaIds: nil,
		}, nil
	}
	var areaIdMaps []map[string]interface{}
	result = repo.data.db.WithContext(ctx).
		Table(setting.TABLE_MAP[u.Power]).
		Where("uid = ?", uid).
		Find(&areaIdMaps)
	if result.Error != nil {
		repo.log.Errorf(" GetUser2. Error:%d", result.Error)
		return nil, result.Error
	}
	areaIds := make([]uint32, 0)
	for _, areaIdMap := range areaIdMaps {
		areaIds = append(areaIds, areaIdMap["aid"].(uint32))
	}
	return &biz.User{
		Uid:     u.Uid,
		Power:   u.Power,
		AreaIds: areaIds,
	}, nil
}

func (repo *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	/* 先创建user, 在关联areaId */
	uc := User{
		Uid:           u.Uid,
		Power:         u.Power,
		AreaTableName: setting.TABLE_MAP[u.Power],
	}
	tx := repo.data.db.Begin()
	result := tx.WithContext(ctx).Create(&uc)
	if result.Error != nil {
		tx.Rollback()
		repo.log.Errorf(" CreateUser1. Error:%d", result.Error)
		return nil, result.Error
	}
	// 权限认证
	if u.Power == setting.SUPER_ADMIN_USER {
		repo.log.Debug(" CreateUser2. Debug: create super admin")
		return &biz.User{
			Uid:     uc.Uid,
			Power:   uc.Power,
			AreaIds: nil,
		}, nil
	} else if u.Power == setting.AREA_USER {
		repo.log.Debug(" CreateUser3. Debug: create user")
		if len(u.AreaIds) > 1 {
			u.AreaIds = u.AreaIds[0:1]
		}
	}
	var umap = []map[string]interface{}{}
	for _, aid := range u.AreaIds {
		umap = append(umap, map[string]interface{}{
			"uid":        uc.Uid,
			"aid":        aid,
			"created_at": time.Now().Local(),
		})
	}
	result = tx.WithContext(ctx).
		Table(setting.TABLE_MAP[u.Power]).
		Create(umap)
	if result.Error != nil {
		tx.Rollback()
		repo.log.Errorf(" CreateUser4. Error:%d", result.Error)
		return nil, result.Error
	}
	tx.Commit()
	return &biz.User{
		Uid:     uc.Uid,
		Power:   uc.Power,
		AreaIds: u.AreaIds,
	}, nil
}

func (repo *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	/*
		如果传入areaIds为空, 则删除全部区域。
	*/
	uu := User{}
	tx := repo.data.db.Begin()
	result := tx.WithContext(ctx).First(&uu, u.Uid)
	if result.Error != nil {
		tx.Rollback()
		repo.log.Errorf(" UpdateUser1. Error:%d", result.Error)
		return nil, result.Error
	}
	if uu.Power == setting.SUPER_ADMIN_USER {
		return nil, errors.New(500, "Super admin", "super admin can not be updated.")
	}
	result = tx.WithContext(ctx).
		Exec(fmt.Sprintf("DELETE FROM %s WHERE uid=?", setting.TABLE_MAP[uu.Power]), uu.Uid)
	if result.Error != nil {
		tx.Rollback()
		repo.log.Errorf(" UpdateUser2. Error:%d", result.Error)
		return nil, result.Error
	}
	if len(u.AreaIds) == 0 {
		return &biz.User{
			Uid:     uu.Uid,
			Power:   uu.Power,
			AreaIds: u.AreaIds,
		}, nil
	}
	var umap = []map[string]interface{}{}
	for _, aid := range u.AreaIds {
		umap = append(umap, map[string]interface{}{
			"uid":        uu.Uid,
			"aid":        aid,
			"created_at": time.Now().Local(),
		})
	}
	result = tx.WithContext(ctx).
		Table(setting.TABLE_MAP[uu.Power]).
		Create(umap)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return &biz.User{
		Uid:     uu.Uid,
		Power:   uu.Power,
		AreaIds: u.AreaIds,
	}, nil
}

func (repo *userRepo) ListUser(ctx context.Context, power int32, areaIds []uint32) ([]*biz.User, error) {
	/*
		TODO: 有点傻, 要改  group_concat
		通过areaId, 得到user组
		若没有areaIds,返回全部user
	*/
	var us []User
	if len(areaIds) == 0 {
		// 搜索全部
		result := repo.data.db.WithContext(ctx).Where("power = ?", setting.AREA_ADMIN_USER).Find(&us)
		if result.Error != nil {
			repo.log.Errorf(" ListUser1. Error:%d", result.Error)
			return nil, result.Error
		}
	} else {
		// 按Areaids搜索
		results := []map[string]interface{}{}
		result := repo.data.db.WithContext(ctx).
			Table(setting.TABLE_MAP[power]).
			Where("aid = ?", areaIds[0]).Find(&results)
		if result.Error != nil {
			repo.log.Errorf(" ListUser2. Error:%d", result.Error)
			return nil, result.Error
		}
		uids := make([]uint64, 0)
		for _, v := range results {
			uids = append(uids, v["uid"].(uint64))
		}
		for i := 1; i < len(areaIds); i++ {
			results = []map[string]interface{}{}
			result = repo.data.db.WithContext(ctx).
				Table(setting.TABLE_MAP[power]).
				Where("aid = ?", areaIds[i]).
				Where("uid IN ?", uids).
				Find(&results)
			if result.Error != nil {
				repo.log.Errorf(" ListUser3. Error:%d", result.Error)
				return nil, result.Error
			}
			uids = make([]uint64, 0)
			for _, v := range results {
				uids = append(uids, v["uid"].(uint64))
			}
			if len(uids) == 0 {
				return []*biz.User{}, nil
			}
		}
		// ======================================================
		result = repo.data.db.WithContext(ctx).Find(&us, uids)
		if result.Error != nil {
			repo.log.Errorf(" ListUser4. Error:%d", result.Error)
			return nil, result.Error
		}
	}
	bu := make([]*biz.User, 0)
	for _, u := range us {
		var areaIdMaps []map[string]interface{}
		result := repo.data.db.WithContext(ctx).
			Table(setting.TABLE_MAP[u.Power]).
			Where("uid = ?", u.Uid).
			Find(&areaIdMaps)
		if result.Error != nil {
			repo.log.Errorf(" GetUser2. Error:%d", result.Error)
			return nil, result.Error
		}
		areaIds := make([]uint32, 0)
		for _, areaIdMap := range areaIdMaps {
			areaIds = append(areaIds, areaIdMap["aid"].(uint32))
		}
		bu = append(bu, &biz.User{
			Uid:     u.Uid,
			Power:   u.Power,
			AreaIds: areaIds,
		})
	}
	return bu, nil
}
