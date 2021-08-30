package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id    uint64
	Uid   uint64
	Power int32
	Areas []*Area
}

type UserRepo interface {
	GetUser(ctx context.Context, id uint64) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	ListUser(ctx context.Context, areaIds []uint32) ([]*User, error)
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

func (uc *UserUseCase) List(ctx context.Context, areaIds []uint32) ([]*User, error) {
	return uc.repo.ListUser(ctx, areaIds)
}
