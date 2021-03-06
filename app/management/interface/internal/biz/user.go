package biz

import (
	"context"

	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/Yui-wy/asset-management/pkg/util/rsakey"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         uint64
	Username   string
	Password   string
	Nickname   string
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
	ListUser(ctx context.Context, pageNum, pageSize int64, areaIds []uint32, power int32) ([]*User, int64, error)
	ModifyPd(ctx context.Context, id uint64, password string) (bool, error)
	DeleteUser(ctx context.Context, id uint64) (bool, error)
	// area
	ListArea(ctx context.Context, areaIds []uint32, pageNum, pageSize int64) ([]*Area, int64, error)
	GetArea(ctx context.Context, areaId uint32) (*Area, error)
	// superAdmin
	CreateArea(ctx context.Context, areaInfo string) (*Area, error)
}

type UserUseCase struct {
	repo   UserRepo
	log    *log.Helper
	priKey string
	pubKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	priKey, pubKey := rsakey.GenerateRsaKey(1024)
	uc := &UserUseCase{
		repo:   repo,
		log:    log.NewHelper(log.With(logger, "module", "usecase/user")),
		priKey: priKey,
		pubKey: pubKey,
	}
	err := uc.init(context.Background())
	if err != nil {
	}
	return uc
}

func (uc *UserUseCase) init(ctx context.Context) error {
	// 创建一系列数据
	// 超级管理员
	_, err := uc.Create(ctx, &User{
		Username: "SuperAdmin",
		Password: "admin12345",
		Power:    setting.SUPER_ADMIN_USER,
		Nickname: "超级管理员",
	})
	if err != nil {
		return err
	}
	// 创建区域
	area, err := uc.CreateArea(ctx, "复旦上雅园物业")
	if err != nil {
		return err
	}
	// 创建分类

	// 创建区域管理员
	_, err = uc.Create(ctx, &User{
		Username: "fdsyyAdmin",
		Password: "fdsyy12345",
		Power:    setting.AREA_ADMIN_USER,
		Nickname: "复旦上雅园管理员",
		AreaIds:  []uint32{area.Id},
	})
	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*User, error) {
	decodePd, err := rsakey.Base64Decrypt(password)
	if err != nil {
		return nil, err
	}
	password, err = rsakey.RSADecrypt(decodePd, uc.priKey)
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

func (uc *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64, areaIds []uint32, nextPower int32) ([]*User, int64, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize, areaIds, nextPower)
}

func (uc *UserUseCase) ModifyPd(ctx context.Context, id uint64, password string) (bool, error) {
	decodePd, err := rsakey.Base64Decrypt(password)
	if err != nil {
		return false, err
	}
	password, err = rsakey.RSADecrypt(decodePd, uc.priKey)
	if err != nil {
		return false, err
	}
	return uc.repo.ModifyPd(ctx, id, password)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint64) (bool, error) {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) ListArea(ctx context.Context, areaIds []uint32, pageNum, pageSize int64) ([]*Area, int64, error) {
	return uc.repo.ListArea(ctx, areaIds, pageNum, pageSize)
}

func (uc *UserUseCase) GetArea(ctx context.Context, areaId uint32) (*Area, error) {
	return uc.repo.GetArea(ctx, areaId)
}

func (uc *UserUseCase) CreateArea(ctx context.Context, areaInfo string) (*Area, error) {
	return uc.repo.CreateArea(ctx, areaInfo)
}
