package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ScrappedForm struct {
	*BaseForm
}

type ScConfig struct {
	*BaseConfig
}

type ScrappedRepo interface {
	GetForm(ctx context.Context, id int64) (*ScrappedForm, error)
	ListForm(ctx context.Context, conf *ScConfig,pageNum, pageSize int64) ([]*ScrappedForm, error)
	CreateForm(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error)
	UpdateForm(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error)
}

type ScrappedUseCase struct {
	repo ScrappedRepo
	log  *log.Helper
}

func NewScrappedUseCase(repo ScrappedRepo, logger log.Logger) *ScrappedUseCase {
	return &ScrappedUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/scrappedForm")),
	}
}

func (s *ScrappedUseCase) Get(ctx context.Context, id int64) (*ScrappedForm, error) {
	return s.repo.GetForm(ctx, id)
}
func (s *ScrappedUseCase) List(ctx context.Context, conf *ScConfig,pageNum, pageSize int64) ([]*ScrappedForm, error) {
	return s.repo.ListForm(ctx, conf,pageNum, pageSize)
}
func (s *ScrappedUseCase) Create(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error) {
	return s.repo.CreateForm(ctx, sf)
}
func (s *ScrappedUseCase) Update(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error) {
	return s.repo.UpdateForm(ctx, sf)
}
