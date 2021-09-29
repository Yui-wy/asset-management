package data

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	pb "github.com/Yui-wy/asset-management/api/user/service/v1"
	"github.com/Yui-wy/asset-management/app/user/service/internal/biz"
	"github.com/Yui-wy/asset-management/app/user/service/internal/pkg/util"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID         uint64 `gorm:"primarykey"`
	Username   string `gorm:"not null;size:64;uniqueIndex:user_name"`
	Nickname   string `gorm:"not null"`
	Password   string
	UpdataSign string `gorm:"not null"`
	IsDeleted  bool   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  int64
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
		Id:         u.ID,
		Username:   u.Username,
		UpdataSign: u.UpdataSign,
		CreatedAt:  u.CreatedAt,
	}, result.Error
}

func (repo *userRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	if match, str := inspection.CheckNameFormat(b.Username); !match {
		repo.log.Error(str)
		return nil, pb.ErrorRegisterFailed(str)
	}
	hashPassword, err := util.HashPassword(b.Password)
	if err != nil {
		repo.log.Errorf("CreateUser1 error. Error:%d", err)
		return nil, err
	}
	if inspection.IsZeros(b.Nickname) {
		b.Nickname = b.Username
	}
	u := User{
		Username:   b.Username,
		Password:   hashPassword,
		Nickname:   b.Nickname,
		UpdataSign: util.CreateMD5Random(b.Username),
		IsDeleted:  false,
	}
	result := repo.data.db.WithContext(ctx).Create(&u)
	if result.Error != nil {
		repo.log.Errorf("CreateUser2 error. Error:%d", result.Error)
		return nil, result.Error
	}
	uu := User{}
	result = repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&uu, u.ID)
	if result.Error != nil {
		repo.log.Errorf("CreateUser3 error. Error:%d", result.Error)
		return nil, result.Error
	}
	return &biz.User{
		Id:         uu.ID,
		Username:   uu.Username,
		Nickname:   uu.Nickname,
		CreatedAt:  uu.CreatedAt,
		UpdataSign: uu.UpdataSign,
	}, result.Error
}

func (repo *userRepo) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Id:         u.ID,
		Username:   u.Username,
		Nickname:   u.Nickname,
		CreatedAt:  u.CreatedAt,
		UpdataSign: u.UpdataSign,
	}, result.Error
}

func (repo *userRepo) ListUser(ctx context.Context, ids []uint64, pageNum, pageSize int64) ([]*biz.User, int64, error) {
	var us []User
	tx := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Where("is_deleted = false").WithContext(ctx)
	result := tx.Find(&us, ids)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64
	result = tx.Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	totalPage := int64(math.Ceil(float64(total) / float64(pageSize)))
	bus := make([]*biz.User, 0)
	for _, u := range us {
		bus = append(bus, &biz.User{
			Id:         u.ID,
			Username:   u.Username,
			Nickname:   u.Nickname,
			CreatedAt:  u.CreatedAt,
			UpdataSign: u.UpdataSign,
		})
	}
	return bus, totalPage, nil
}

func (repo *userRepo) UpdateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, b.Id)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser1 error. Error:%d", result.Error)
		return nil, result.Error
	}
	hp, err := util.HashPassword(b.Password)
	if err != nil {
		repo.log.Errorf("UpdateUser2 error. Error:%d", err)
		return nil, err
	}
	u.ID = b.Id
	u.Password = hp
	u.UpdataSign = util.CreateMD5Random(u.Username)
	result = repo.data.db.WithContext(ctx).Save(&u)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser3 error. Error:%d", result.Error)
		return nil, result.Error
	}
	uu := User{}
	result = repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&uu, u.ID)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser4 error. Error:%d", result.Error)
		return nil, result.Error
	}
	return &biz.User{
		Id:         uu.ID,
		Username:   uu.Username,
		CreatedAt:  uu.CreatedAt,
		UpdataSign: uu.UpdataSign,
	}, nil
}

func (repo *userRepo) UpdateNickname(ctx context.Context, b *biz.User) (*biz.User, error) {
	if inspection.IsZeros(b.Nickname) {
		return nil, errors.Errorf(500, "请输入用户名", "请输入用户名")
	}
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, b.Id)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser1 error. Error:%d", result.Error)
		return nil, result.Error
	}
	u.ID = b.Id
	u.Nickname = b.Nickname
	result = repo.data.db.WithContext(ctx).Save(&u)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser3 error. Error:%d", result.Error)
		return nil, result.Error
	}
	uu := User{}
	result = repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&uu, u.ID)
	if result.Error != nil {
		repo.log.Errorf("UpdateUser4 error. Error:%d", result.Error)
		return nil, result.Error
	}
	return &biz.User{
		Id:         uu.ID,
		Username:   uu.Username,
		CreatedAt:  uu.CreatedAt,
		UpdataSign: uu.UpdataSign,
	}, nil
}

func (repo *userRepo) DeleteUser(ctx context.Context, id uint64) (bool, error) {
	u := User{}
	result := repo.data.db.WithContext(ctx).Where("is_deleted = false").First(&u, id)
	if result.Error != nil {
		repo.log.Errorf("DeleteUser1 error. Error:%d", result.Error)
		return false, result.Error
	}
	result = repo.data.db.WithContext(ctx).Model(&u).Updates(User{
		Username:  u.Username + "+^)-" + fmt.Sprintf("%d%d", rand.Intn(100), u.ID),
		IsDeleted: true,
		DeletedAt: time.Now().Unix(),
	})
	if result.Error != nil {
		repo.log.Errorf("DeleteUser2 error. Error:%d", result.Error)
		return false, result.Error
	}
	return true, nil
}

func (repo *userRepo) VerifyPassword(ctx context.Context, b *biz.User) (*biz.User, error) {
	uu := User{}
	result := repo.data.db.WithContext(ctx).Where("username = ? AND is_deleted = false", b.Username).First(&uu)
	if result.Error != nil {
		return nil, result.Error
	}
	if err := util.CheckPasswordHash(b.Password, uu.Password); err != nil {
		return nil, errors.New(401, "password error", "密码错误")
	}
	return &biz.User{
		Id:         uu.ID,
		Username:   uu.Username,
		CreatedAt:  uu.CreatedAt,
		UpdataSign: uu.UpdataSign,
	}, nil
}
