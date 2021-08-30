package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Area struct {
	Id        uint32
	AreaInfo  string
	IsDeleted bool
}

type AreaRepo interface {
	GetArea(ctx context.Context, id uint32) (*Area, error)
	ListArea(ctx context.Context) ([]*Area, error)
	CreateArea(ctx context.Context, a *Area) (*Area, error)
	UpdateArea(ctx context.Context, a *Area) (*Area, error)
	DeleteArea(ctx context.Context, id uint32) (bool, error)
}

type AreaUseCase struct {
	repo AreaRepo
	log  *log.Helper
}

func NewAreaUseCase(repo AreaRepo, logger log.Logger) *AreaUseCase {
	return &AreaUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/area")),
	}
}

func (uc *AreaUseCase) Get(ctx context.Context, id uint32) (*Area, error) {
	return uc.repo.GetArea(ctx, id)
}

func (uc *AreaUseCase) List(ctx context.Context) ([]*Area, error) {
	return uc.repo.ListArea(ctx)
}

func (uc *AreaUseCase) Create(ctx context.Context, area *Area) (*Area, error) {
	return uc.repo.CreateArea(ctx, area)
}

func (uc *AreaUseCase) Update(ctx context.Context, area *Area) (*Area, error) {
	return uc.repo.UpdateArea(ctx, area)
}

func (uc *AreaUseCase) Delete(ctx context.Context, id uint32) (bool, error) {
	return uc.repo.DeleteArea(ctx, id)
}
