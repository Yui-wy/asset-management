package data

import (
	"context"
	"time"

	pb "github.com/Yui-wy/asset-management/api/user/service/v1"
	"github.com/Yui-wy/asset-management/app/user/service/internal/biz"
	"github.com/Yui-wy/asset-management/app/user/service/internal/pkg/util"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

// var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID         uint   `gorm:"primarykey"`
	Username   string `gorm:"not null;uniqueIndex:username"`
	Password   string
	UpdataSign string `gorm:"not null"`
	IsDeleted  bool   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (repo *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("username = ? AND is_deleted = false", username).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:        uint64(u.ID),
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}, result.Error
}

func (repo *userRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	if match, str := util.CheckNameFormat(b.Username); !match {
		repo.log.Error(str)
		return nil, pb.ErrorRegisterFailed(str)
	}
	hashPassword, err := util.HashPassword(b.Password)
	if err != nil {
		return nil, err
	}
	u := User{
		Username:   b.Username,
		Password:   hashPassword,
		UpdataSign: util.CreateMD5Random(b.Username),
		IsDeleted:  false,
	}
	result := repo.data.db.WithContext(ctx).Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:         uint64(u.ID),
		Username:   u.Username,
		CreatedAt:  u.CreatedAt,
		UpdataSign: util.CreateMD5Random(b.Username),
	}, result.Error
}

func (repo *userRepo) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:         uint64(u.ID),
		Username:   u.Username,
		CreatedAt:  u.CreatedAt,
		UpdataSign: u.UpdataSign,
	}, result.Error
}

func (repo *userRepo) ListUser(ctx context.Context, ids []uint64, pageNum, pageSize int64) ([]*biz.User, error) {
	var us []User
	result := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Where("is_deleted = false").
		Find(&us, ids)
	if result.Error != nil {
		return nil, result.Error
	}
	bus := make([]*biz.User, 0)
	for _, u := range us {
		bus = append(bus, &biz.User{
			Id:         uint64(u.ID),
			Username:   u.Username,
			CreatedAt:  u.CreatedAt,
			UpdataSign: u.UpdataSign,
		})
	}
	return bus, nil
}

func (repo *userRepo) UpdateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	hp, err := util.HashPassword(b.Password)
	if err != nil {
		return nil, err
	}
	result = repo.data.db.WithContext(ctx).Model(&u).Updates(User{
		Password:   hp,
		UpdataSign: util.CreateMD5Random(u.Username),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:         uint64(u.ID),
		Username:   u.Username,
		CreatedAt:  u.CreatedAt,
		UpdataSign: util.CreateMD5Random(u.Username),
	}, nil
}

func (repo *userRepo) DeleteUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	result = repo.data.db.WithContext(ctx).Model(&u).Updates(User{
		IsDeleted: true,
		DeletedAt: time.Now(),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:        uint64(u.ID),
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		IsDeleted: u.IsDeleted,
		DeletedAt: u.DeletedAt,
	}, nil
}

func (repo *userRepo) VerifyPassword(ctx context.Context, b *biz.User) (bool, error) {
	uu, err := repo.GetUserByUsername(ctx, b.Username)
	if err != nil {
		return false, err
	}
	result := util.CheckPasswordHash(b.Password, uu.Password)
	return result, nil
}
