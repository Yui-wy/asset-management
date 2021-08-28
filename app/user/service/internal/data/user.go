package data

import (
	"context"
	"time"

	pb "github.com/Yui-wy/material/api/user/service/v1"
	"github.com/Yui-wy/material/app/user/service/internal/biz"
	"github.com/Yui-wy/material/app/user/service/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

// var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"not null"`
	Password  string
	IsDeleted bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (repo *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("username = ?", username).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	if u.IsDeleted {
		return nil, pb.ErrorUnknownError("User is not found")
	}
	return &biz.User{
		Id:        uint32(u.ID),
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}, result.Error
}

func (repo *userRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	if match, str := util.CheckNameFormat(b.Username); !match {
		repo.log.Error(str)
		return nil, pb.ErrorRegisterFailed(str)
	}
	uu, _ := repo.GetUserByUsername(ctx, b.Username)
	if uu != nil {
		repo.log.Errorf("User has existed. name: %s", b.Username)
		return nil, pb.ErrorRegisterFailed("User has existed. name: %s", b.Username)
	}
	hashPassword, err := util.HashPassword(b.Password)
	if err != nil {
		return nil, err
	}
	u := User{
		Username:  b.Username,
		Password:  hashPassword,
		IsDeleted: false,
	}
	result := repo.data.db.WithContext(ctx).Create(&u)
	return &biz.User{
		Id:        uint32(u.ID),
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}, result.Error
}

func (repo *userRepo) GetUser(ctx context.Context, id uint32) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).First(&u, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if u.IsDeleted {
		repo.log.Errorf("User is not found.")
		return nil, pb.ErrorUnknownError("User is not found.")
	}
	return &biz.User{
		Id:        uint32(u.ID),
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}, result.Error
}

func (repo *userRepo) ListUser(ctx context.Context, ids []uint32) ([]*biz.User, error) {
	return nil, nil
}

func (repo *userRepo) UpdateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	return nil, nil
}

func (repo *userRepo) VerifyPassword(ctx context.Context, b *biz.User) (bool, error) {
	uu, err := repo.GetUserByUsername(ctx, b.Username)
	if err != nil {
		return false, err
	}
	result := util.CheckPasswordHash(b.Password, uu.Password)
	return result, nil
}
