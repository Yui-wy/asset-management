package data

import (
	"context"

	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type storageRepo struct {
	data *Data
	log  *log.Helper
}

type StorageForm struct {
}

func NewStorageRepo(data *Data, logger log.Logger) biz.StorageRepo {
	return &storageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/storage")),
	}

}

func (repo *storageRepo) GetForm(ctx context.Context, formNum string) (*biz.StorageForm, error) {

	return nil, nil
}
func (repo *storageRepo) ListForm(ctx context.Context, conf *biz.StConfig) ([]*biz.StorageForm, error) {

	return nil, nil
}
func (repo *storageRepo) CreateForm(ctx context.Context, sf *biz.StorageForm) (*biz.StorageForm, error) {

	return nil, nil
}
func (repo *storageRepo) UpdateForm(ctx context.Context, sf *biz.StorageForm) (*biz.StorageForm, error) {

	return nil, nil
}
