package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id        uint32
	Username  string
	Password  string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id uint32) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	ListUser(ctx context.Context, ids []uint32) ([]*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
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

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, ids []uint32) ([]*User, error) {
	return uc.repo.ListUser(ctx, ids)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, user *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, user)
}
