package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewManagementInterface)

type ManagementInterface struct {
	pb.UnimplementedManagementInterfaceServer

	log    *log.Helper
	uc     *biz.UserUseCase
	authUc *biz.AuthUseCase
	ac     *biz.AssetUseCase
}

func NewManagementInterface(
	logger log.Logger,
	uc *biz.UserUseCase,
	authUc *biz.AuthUseCase,
	ac *biz.AssetUseCase,
) *ManagementInterface {
	s := &ManagementInterface{
		log:    log.NewHelper(log.With(logger, "module", "service/user")),
		uc:     uc,
		authUc: authUc,
		ac:     ac,
	}
	s.init(context.Background())
	return s
}

func (s *ManagementInterface) init(ctx context.Context) error {
	// 创建一系列初始化数据
	// 超级管理员
	_, err := s.uc.Create(ctx, &biz.User{
		Username: "SuperAdmin",
		Password: "admin12345",
		Power:    setting.SUPER_ADMIN_USER,
		Nickname: "超级管理员",
	})
	if err != nil {
		return err
	}
	// 创建区域
	area, err := s.uc.CreateArea(ctx, "复旦上雅园物业")
	if err != nil {
		return err
	}
	// 创建分类
	_, err = s.ac.CreateClasses(ctx, biz.Clzzz)
	if err != nil {
		return err
	}
	// 创建区域管理员
	_, err = s.uc.Create(ctx, &biz.User{
		Username: "fdsyyAdmin",
		Password: "fdsyy12345",
		Power:    setting.AREA_ADMIN_USER,
		Nickname: "复旦上雅园管理员",
		AreaIds:  []uint32{area.Id},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *ManagementInterface) checkPower(ctx context.Context, power int32, areaId []uint32) (*biz.AuthUser, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	userAuth, ok := s.authUc.FromContext(ctx)
	if !ok {
		return nil, auth.ErrWrongContext
	}
	if userAuth.Power != power {
		return nil, auth.ErrPowerFail
	}
	// 确保在区域范围内的权限正确
	for _, i := range areaId {
		notInArray := true
		for _, k := range userAuth.AreaIds {
			if k == i {
				notInArray = false
				break
			}
		}
		if notInArray {
			return nil, auth.ErrAreaFail
		}
	}
	return userAuth, nil
}

func (s *ManagementInterface) getUserDetail(ctx context.Context, areaId []uint32) (*biz.AuthUser, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	userAuth, ok := s.authUc.FromContext(ctx)
	if !ok {
		return nil, auth.ErrWrongContext
	}
	for _, i := range areaId {
		notInArray := true
		for _, k := range userAuth.AreaIds {
			if k == i {
				notInArray = false
				break
			}
		}
		if notInArray {
			return nil, auth.ErrAreaFail
		}
	}
	return userAuth, nil
}

func (s *ManagementInterface) GetAuthUseCase() *biz.AuthUseCase {
	return s.authUc
}
