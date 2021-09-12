package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewManagementInterface)

type ManageMentInterface struct {
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
) *ManageMentInterface {
	return &ManageMentInterface{
		log:    log.NewHelper(log.With(logger, "module", "service/user")),
		uc:     uc,
		authUc: authUc,
		ac:     ac,
	}
}

func (uc *ManageMentInterface) checkPower(ctx context.Context, power int32, areaId []uint32) (*biz.AuthUser, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	userAuth, ok := uc.authUc.FromContext(ctx)
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

func (uc *ManageMentInterface) getUserDetail(ctx context.Context, areaId []uint32) (*biz.AuthUser, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	userAuth, ok := uc.authUc.FromContext(ctx)
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

func (uc *ManageMentInterface) GetAuthUseCase() *biz.AuthUseCase {
	return uc.authUc
}
