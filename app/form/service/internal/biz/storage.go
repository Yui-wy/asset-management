package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type StorageForm struct {
	*BaseForm
}

type StConfig struct {
	*BaseConfig
}

type StorageRepo interface {
	GetForm(ctx context.Context, id int64) (*StorageForm, error)
	ListForm(ctx context.Context, conf *StConfig,pageNum, pageSize int64) ([]*StorageForm, error)
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

func (s *StorageUseCase) Get(ctx context.Context, id int64) (*StorageForm, error) {
	return s.repo.GetForm(ctx, id)
}
func (s *StorageUseCase) List(ctx context.Context, conf *StConfig, pageNum, pageSize int64) ([]*StorageForm, error) {
	return s.repo.ListForm(ctx, conf,pageNum, pageSize)
}
func (s *StorageUseCase) Create(ctx context.Context, sf *StorageForm) (*StorageForm, error) {
	return s.repo.CreateForm(ctx, sf)
}
func (s *StorageUseCase) Update(ctx context.Context, sf *StorageForm) (*StorageForm, error) {
	return s.repo.UpdateForm(ctx, sf)
}
