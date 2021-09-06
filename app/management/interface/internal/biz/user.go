package biz

import (
	"context"

	"github.com/Yui-wy/asset-management/pkg/util/rsakey"
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
	Create(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id uint64) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int64, areaIds []uint32, power int32) ([]*User, error)
	ModifyPd(ctx context.Context, id uint64, password string) (bool, error)
	DeleteUser(ctx context.Context, id uint64) (bool, error)
	// area
	ListArea(ctx context.Context, areaIds []uint32,pageNum, pageSize int64) ([]*Area, error)
	GetArea(ctx context.Context, areaId uint32) (*Area, error)
}

type UserUseCase struct {
	repo   UserRepo
	log    *log.Helper
	priKey string
	pubKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	priKey, pubKey := rsakey.GenerateRsaKey(512)

	return &UserUseCase{
		repo:   repo,
		log:    log.NewHelper(log.With(logger, "module", "usecase/user")),
		priKey: priKey,
		pubKey: pubKey,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*User, error) {
	password, err := rsakey.RSADecrypt(password, uc.priKey)
	if err != nil {
		return nil, err
	}
	return uc.repo.Login(ctx, username, password)
}

func (uc *UserUseCase) Logout(ctx context.Context, id uint64) (bool, error) {
	return uc.repo.Logout(ctx, id)
}

func (uc *UserUseCase) GetKey(ctx context.Context) (string, error) {
	return uc.pubKey, nil
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id uint64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64, areaIds []uint32, nextPower int32) ([]*User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize, areaIds, nextPower)
}

func (uc *UserUseCase) ModifyPd(ctx context.Context, id uint64, password string) (bool, error) {
	password, err := rsakey.RSADecrypt(password, uc.priKey)
	if err != nil {
		return false, err
	}
	return uc.repo.ModifyPd(ctx, id, password)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint64) (bool, error) {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) ListArea(ctx context.Context, areaIds []uint32,pageNum, pageSize int64) ([]*Area, error) {
	return uc.repo.ListArea(ctx, areaIds,pageNum, pageSize)
}

func (uc *UserUseCase) GetArea(ctx context.Context, areaId uint32) (*Area, error) {
	return uc.repo.GetArea(ctx, areaId)
}
