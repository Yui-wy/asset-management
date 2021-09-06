package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         uint64
	Username   string
	Password   string
	UpdataSign string
	Power      int32
	AreaIds    []uint32
}

type Area struct {
	Id       uint32
	AreaInfo string
}

type UserRepo interface {
	// user
	Login(ctx context.Context, username, password string) (*User, error)
	Logout(ctx context.Context, id uint64) (bool, error)
	Create(ctx context.Context, user User) (*User, error)
	GetUser(ctx context.Context, id uint64) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*User, error)
	ModifyPd(ctx context.Context, id uint64, pd string) (bool, error)
	DeleteUser(ctx context.Context, id uint64) (bool, error)
	// area
	ListArea(ctx context.Context, areaIds []uint32) ([]*Area, error)
	GetArea(ctx context.Context, areaId uint32) (*Area, error)
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

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*User, error) {
	// TODO: 解密传入
	return uc.repo.Login(ctx, username, password)
}

func (uc *UserUseCase) Logout(ctx context.Context, id uint64) (bool, error) {
	return uc.repo.Logout(ctx, id)
}

func (uc *UserUseCase) GetKey(ctx context.Context) (string, error) {
	// TODO: 不作为接口RSA
	return "", nil
}

func (uc *UserUseCase) Create(ctx context.Context, user User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id uint64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) ModifyPd(ctx context.Context, id uint64, pd string) (bool, error) {
	// TODO: 解密传入
	return uc.repo.ModifyPd(ctx, id, pd)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint64) (bool, error) {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) ListArea(ctx context.Context, areaIds []uint32) ([]*Area, error) {
	return uc.repo.ListArea(ctx, areaIds)
}

func (uc *UserUseCase) GetArea(ctx context.Context, areaId uint32) (*Area, error) {
	return uc.repo.GetArea(ctx, areaId)
}
