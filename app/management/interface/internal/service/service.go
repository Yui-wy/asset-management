package service

import (
	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
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
