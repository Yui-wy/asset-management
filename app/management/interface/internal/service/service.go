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
	ac     *biz.AssetUseCase
}

type UserDetails struct {
	Id       uint64
	Username string
	Power    int32
	AreaId   []uint32
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
		AuthUc: authUc,
		ac:     ac,
	}
}

func (uc *ManageMentInterface) checkPower(ctx context.Context, power int32, areaId []uint32) (*UserDetails, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	result := ctx.Value("x-md-global-user").(map[string]interface{})
	userPower, ok := result["power"].(int32)
	if !ok || (userPower != power) {
		return nil, auth.ErrPowerFail
	}
	// 确保在区域范围内的权限正确
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
			return nil, auth.ErrAreaFail
		}
	}
	return &UserDetails{
		Id:       result["user_id"].(uint64),
		Username: result["user_name"].(string),
		AreaId:   uaid,
		Power:    userPower,
	}, nil
}

func (uc *ManageMentInterface) getUserDetail(ctx context.Context, areaId []uint32) (*UserDetails, error) {
	if inspection.IsZeros(areaId) {
		return nil, auth.ErrAreaFail
	}
	result := ctx.Value("x-md-global-user").(map[string]interface{})
	userPower := result["power"].(int32)
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
			return nil, auth.ErrAreaFail
		}
	}
	return &UserDetails{
		Id:       result["user_id"].(uint64),
		Username: result["user_name"].(string),
		AreaId:   uaid,
		Power:    userPower,
	}, nil
}
