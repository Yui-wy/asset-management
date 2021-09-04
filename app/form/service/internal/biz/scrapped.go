package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ScrappedForm struct {
	FormNum     string
	AppliedAt   int64
	ApplicantId uint64
	Applicant   string
	OptAt       int64
	OperatorId  uint64
	Operator    string
	StateNum    int64
	State       string
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type ScConfig struct {
	ApplicantId uint64
	OperatorId  uint64
	StateNum    int64
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type ScrappedRepo interface {
	GetForm(ctx context.Context, formNum string) (*ScrappedForm, error)
	ListForm(ctx context.Context, conf *ScConfig) ([]*ScrappedForm, error)
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

func (s *ScrappedUseCase) Get(ctx context.Context, formNum string) (*ScrappedForm, error) {
	return s.repo.GetForm(ctx, formNum)
}
func (s *ScrappedUseCase) List(ctx context.Context, conf *ScConfig) ([]*ScrappedForm, error) {
	return s.repo.ListForm(ctx, conf)
}
func (s *ScrappedUseCase) Create(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error) {
	return s.repo.CreateForm(ctx, sf)
}
func (s *ScrappedUseCase) Update(ctx context.Context, sf *ScrappedForm) (*ScrappedForm, error) {
	return s.repo.UpdateForm(ctx, sf)
}
