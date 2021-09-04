package service

import (
	pb "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFormService)

type FormService struct {
	pb.UnimplementedFormServer

	log *log.Helper
	spu *biz.ScrappedUseCase
	stu *biz.StorageUseCase
}

func NewFormService(logger log.Logger, spu *biz.ScrappedUseCase, stu *biz.StorageUseCase) *FormService {
	return &FormService{
		log: log.NewHelper(log.With(logger, "module", "service/form")),
		spu: spu,
		stu: stu,
	}
}
