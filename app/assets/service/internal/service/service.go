package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAssetsService)

type AssetsService struct {
	pb.UnimplementedAssetsServer

	log *log.Helper
	arc *biz.AreaUseCase
	uc  *biz.UserUseCase
}

func NewAssetsService(logger log.Logger, arc *biz.AreaUseCase, uc *biz.UserUseCase) *AssetsService {
	return &AssetsService{
		log: log.NewHelper(log.With(logger, "module", "service/assets")),
		arc: arc,
		uc:  uc,
	}
}
