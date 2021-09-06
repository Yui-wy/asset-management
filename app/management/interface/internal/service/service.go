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
	AuthUc *biz.AuthUseCase
}

func NewManagementInterface(
	logger log.Logger,
	uc *biz.UserUseCase,
	authUc *biz.AuthUseCase,
) *ManageMentInterface {
	return &ManageMentInterface{
		log:    log.NewHelper(log.With(logger, "module", "service/user")),
		uc:     uc,
		AuthUc: authUc,
	}
}

func (uc *ManageMentInterface) checkPower(ctx context.Context, power int32, areaId []uint32) (*biz.UserDetails, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrPowerFail
	}
	result := ctx.Value("x-md-global-user").(map[string]interface{})
	userPower := result["power"].(int32)
	if userPower != power {
		return nil, auth.ErrPowerFail
	}
	uaid := result["area_id"].([]uint32)
	for _, i := range areaId {
		notInArray := true
		for _, k := range uaid {
			if k == i {
				notInArray = false
				break
			}
		}
		if notInArray {
			return nil, auth.ErrPowerFail
		}
	}
	return &biz.UserDetails{
		Id:     result["user_id"].(uint64),
		AreaId: uaid,
		Power:  userPower,
	}, nil
}
