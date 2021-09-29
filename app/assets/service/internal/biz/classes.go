package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Class struct {
	Id      uint64
	Code    string
	Level   uint32
	Pcode   string
	ClzInfo string
}

type ClassRepo interface {
	GetClasses(ctx context.Context) ([]*Class, error)
	CreateClasses(ctx context.Context, clz []*Class) ([]*Class, error)
}

type ClassUseCase struct {
	repo ClassRepo
	log  *log.Helper
}

func NewClassUseCase(repo ClassRepo, logger log.Logger) *ClassUseCase {
	return &ClassUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/class")),
	}
}

func (uc *ClassUseCase) Get(ctx context.Context) ([]*Class, error) {
	return uc.repo.GetClasses(ctx)
}

func (uc *ClassUseCase) Create(ctx context.Context, clz []*Class) ([]*Class, error) {
	return uc.repo.CreateClasses(ctx, clz)
}
