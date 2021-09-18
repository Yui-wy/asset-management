package data

import (
	"context"

	av1 "github.com/Yui-wy/asset-management/api/assets/service/v1"
	uv1 "github.com/Yui-wy/asset-management/api/user/service/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

// user
func (rp *userRepo) Login(ctx context.Context, username, password string) (*biz.User, error) {
	u, err := rp.data.uc.VerifyPassword(ctx, &uv1.VerifyPasswordReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:         u.Id,
		Username:   u.Username,
		UpdataSign: u.UpdataSign,
	}, nil
}

func (rp *userRepo) Logout(ctx context.Context, id uint64) (bool, error) {
	return true, nil
}
func (rp *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	u, err := rp.data.uc.CreateUser(ctx, &uv1.CreateUserReq{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}
	au, err := rp.data.ac.CreateUser(ctx, &av1.CreateUserReq{
		Uid:     u.Id,
		Power:   user.Power,
		AreaIds: user.AreaIds,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       u.Id,
		Username: u.Username,
		Power:    au.Power,
		AreaIds:  au.AreaIds,
	}, nil
}
func (rp *userRepo) GetUser(ctx context.Context, id uint64) (*biz.User, error) {
	u, err := rp.data.uc.GetUser(ctx, &uv1.GetUserReq{Id: id})
	if err != nil {
		return nil, err
	}
	au, err := rp.data.ac.GetUser(ctx, &av1.GetUserReq{Uid: id})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:         u.Id,
		UpdataSign: u.UpdataSign,
		Username:   u.Username,
		Power:      au.Power,
		AreaIds:    au.AreaIds,
	}, nil
}
func (rp *userRepo) ListUser(ctx context.Context, pageNum, pageSize int64, areaIds []uint32, nextPower int32) ([]*biz.User, int64, error) {
	au, err := rp.data.ac.ListUser(ctx, &av1.ListUserReq{
		AreaIds:   areaIds,
		NextPower: nextPower,
	})
	if err != nil {
		return nil, 0, err
	}
	ids := make([]uint64, 0)
	for _, u := range au.Results {
		ids = append(ids, u.Uid)
	}
	us, err := rp.data.uc.ListUser(ctx, &uv1.ListUserReq{
		Ids:      ids,
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, err
	}
	results := make([]*biz.User, 0)
	for _, u := range us.Results {
		for _, uu := range au.Results {
			if u.Id == uu.Uid {
				results = append(results, &biz.User{
					Id:       u.Id,
					Username: u.Username,
					Power:    uu.Power,
					AreaIds:  uu.AreaIds,
				})
				break
			}
		}
	}
	return results, us.PageTotal, nil
}
func (rp *userRepo) ModifyPd(ctx context.Context, id uint64, password string) (bool, error) {
	_, err := rp.data.uc.UpdatePassword(ctx, &uv1.UpdatePasswordReq{
		Id:       id,
		Password: password,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
func (rp *userRepo) DeleteUser(ctx context.Context, id uint64) (bool, error) {
	_, err := rp.data.ac.UpdateUserArea(ctx, &av1.UpdateUserAreaReq{
		Uid: id,
	})
	if err != nil {
		return false, err
	}
	_, err = rp.data.uc.DeleteUser(ctx, &uv1.DeleteUserReq{
		Id: id,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// area
func (rp *userRepo) ListArea(ctx context.Context, areaIds []uint32, pageNum, pageSize int64) ([]*biz.Area, int64, error) {
	as, err := rp.data.ac.GetAreaByIds(ctx, &av1.GetAreaByIdsReq{
		Ids:      areaIds,
		PageNum:  pageNum,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, err
	}
	results := make([]*biz.Area, 0)
	for _, a := range as.Areas {
		results = append(results, &biz.Area{
			Id:       a.Id,
			AreaInfo: a.AreaInfo,
		})
	}
	return results, as.PageTotal, nil
}
func (rp *userRepo) GetArea(ctx context.Context, areaId uint32) (*biz.Area, error) {
	a, err := rp.data.ac.GetArea(ctx, &av1.GetAreaReq{Id: areaId})
	if err != nil {
		return nil, err
	}
	return &biz.Area{
		Id:       a.Id,
		AreaInfo: a.AreaInfo,
	}, nil
}
