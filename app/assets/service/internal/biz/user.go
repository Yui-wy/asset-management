package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Uid     uint64
	Power   int32
	AreaIds []uint32
}

type UserRepo interface {
	// user
	GetUser(ctx context.Context, uid uint64) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	// 通过区域得到用户
	ListUser(ctx context.Context, power int32, areaIds []uint32) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (uc *UserUseCase) Get(ctx context.Context, uid uint64) (*User, error) {
	return uc.repo.GetUser(ctx, uid)
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) List(ctx context.Context, power int32, areaIds []uint32) ([]*User, error) {
	return uc.repo.ListUser(ctx, power, areaIds)
}
