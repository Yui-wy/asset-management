package data

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

const (
	SUPER_ADMIN_USER = 0
	AREA_ADMIN_USER  = 1
	AREA_USER        = 2
)

var TABLE_MAP = map[int32]string{
	AREA_ADMIN_USER: "admin_areas",
	AREA_USER:       "user_areas",
}

type User struct {
	Uid           uint64 `gorm:"primarykey"`
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
	return TABLE_MAP[AREA_ADMIN_USER]
}

type UserArea struct {
	ID        uint64 `gorm:"primarykey"`
	Uid       uint64 `gorm:"not null;uniqueIndex"`
	Aid       uint32 `gorm:"not null;index"`
	CreatedAt time.Time
}

func (UserArea) TableName() string {
	return TABLE_MAP[AREA_USER]
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
		repo.log.Errorf("GetUser1 error. Error:%d", result.Error)
		return nil, result.Error
	}
	if u.Power == SUPER_ADMIN_USER {
		return &biz.User{
			Uid:     u.Uid,
			Power:   u.Power,
			AreaIds: nil,
		}, nil
	}
	var areaIdMaps []map[string]interface{}
	result = repo.data.db.WithContext(ctx).
		Table(TABLE_MAP[u.Power]).
		Where("uid = ?", uid).
		Find(&areaIdMaps)
	if result.Error != nil {
		repo.log.Errorf("GetUser2 error. Error:%d", result.Error)
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
		AreaTableName: TABLE_MAP[u.Power],
	}
	result := repo.data.db.WithContext(ctx).Create(&uc)
	if result.Error != nil {
		repo.log.Errorf("CreateUser1 error. Error:%d", result.Error)
		return nil, result.Error
	}
	// 权限认证
	if u.Power == SUPER_ADMIN_USER {
		repo.log.Debug("CreateUser2 info. create super admin")
		return &biz.User{
			Uid:     uc.Uid,
			Power:   uc.Power,
			AreaIds: nil,
		}, nil
	} else if u.Power == AREA_USER {
		repo.log.Debug("CreateUser3 info. create user")
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
	result = repo.data.db.WithContext(ctx).
		Table(TABLE_MAP[u.Power]).
		Create(umap)
	if result.Error != nil {
		repo.log.Errorf("CreateUser4 error. Error:%d", result.Error)
		return nil, result.Error
	}
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
	result := repo.data.db.WithContext(ctx).First(&uu, u.Uid)
	if result != nil {
		repo.log.Errorf("UpdateUser1 error. Error:%d", result.Error)
		return nil, result.Error
	}
	result = repo.data.db.WithContext(ctx).
		Exec("DELETE FROM `?` WHERE uid=?", TABLE_MAP[u.Power], uu.Uid)
	if result != nil {
		repo.log.Errorf("UpdateUser2 error. Error:%d", result.Error)
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
	result = repo.data.db.WithContext(ctx).
		Table(TABLE_MAP[u.Power]).
		Create(umap)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Uid:     uu.Uid,
		Power:   uu.Power,
		AreaIds: u.AreaIds,
	}, nil
}

func (repo *userRepo) ListUser(ctx context.Context, power int32, areaIds []uint32) ([]*biz.User, error) {
	/*
		通过areaId, 得到user组
		若没有areaIds,返回全部user
	*/
	var us []User
	if len(areaIds) == 0 {
		// 搜索全部
		result := repo.data.db.WithContext(ctx).Find(&us)
		if result.Error != nil {
			repo.log.Errorf("ListUser1 error. Error:%d", result.Error)
			return nil, result.Error
		}
	} else {
		// 按Areaids搜索
		results := []map[string]interface{}{}
		result := repo.data.db.WithContext(ctx).
			Table(TABLE_MAP[power]).
			Where("aid = ?", areaIds[0]).Find(&results)
		if result.Error != nil {
			repo.log.Errorf("ListUser2 error. Error:%d", result.Error)
			return nil, result.Error
		}
		uids := make([]uint64, 0)
		for _, v := range results {
			uids = append(uids, v["uid"].(uint64))
		}
		for i := 1; i < len(areaIds); i++ {
			results = []map[string]interface{}{}
			result = repo.data.db.WithContext(ctx).
				Table(TABLE_MAP[power]).
				Where("aid = ?", areaIds[i]).Where("uid IN ?", uids).Find(&results)
			if result.Error != nil {
				repo.log.Errorf("ListUser3 error. Error:%d", result.Error)
				return nil, result.Error
			}
			uids = make([]uint64, 0)
			for _, v := range results {
				uids = append(uids, v["uid"].(uint64))
			}
		}
		// ======================================================
		result = repo.data.db.WithContext(ctx).Find(&us, uids)
		if result.Error != nil {
			repo.log.Errorf("ListUser4 error. Error:%d", result.Error)
			return nil, result.Error
		}
	}
	bu := make([]*biz.User, 0)
	for _, u := range us {
		bu = append(bu, &biz.User{
			Uid:     u.Uid,
			Power:   u.Power,
			AreaIds: areaIds,
		})
	}
	return bu, nil
}
