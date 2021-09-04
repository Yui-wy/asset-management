package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type StorageForm struct {
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

type StConfig struct {
	ApplicantId uint64
	OperatorId  uint64
	StateNum    int64
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type StorageRepo interface {
	GetForm(ctx context.Context, formNum string) (*StorageForm, error)
	ListForm(ctx context.Context, conf *StConfig) ([]*StorageForm, error)
	CreateForm(ctx context.Context, sf *StorageForm) (*StorageForm, error)
	UpdateForm(ctx context.Context, sf *StorageForm) (*StorageForm, error)
}

type StorageUseCase struct {
	repo StorageRepo
	log  *log.Helper
}

func NewStorageUseCase(repo StorageRepo, logger log.Logger) *StorageUseCase {
	return &StorageUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/storageForm")),
	}
}

func (s *StorageUseCase) Get(ctx context.Context, formNum string) (*StorageForm, error) {
	return s.repo.GetForm(ctx, formNum)
}
func (s *StorageUseCase) List(ctx context.Context, conf *StConfig) ([]*StorageForm, error) {
	return s.repo.ListForm(ctx, conf)
}
func (s *StorageUseCase) Create(ctx context.Context, sf *StorageForm) (*StorageForm, error) {
	return s.repo.CreateForm(ctx, sf)
}
func (s *StorageUseCase) Update(ctx context.Context, sf *StorageForm) (*StorageForm, error) {
	return s.repo.UpdateForm(ctx, sf)
}
