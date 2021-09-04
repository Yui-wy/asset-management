package data

import (
	"context"

	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type scrappedRepo struct {
	data *Data
	log  *log.Helper
}

type ScrappedForm struct {
}

func NewScrappedRepo(data *Data, logger log.Logger) biz.ScrappedRepo {
	return &scrappedRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/scrapped")),
	}

}

func (repo *scrappedRepo) GetForm(ctx context.Context, formNum string) (*biz.ScrappedForm, error) {

	return nil, nil
}
func (repo *scrappedRepo) ListForm(ctx context.Context, conf *biz.ScConfig) ([]*biz.ScrappedForm, error) {

	return nil, nil
}
func (repo *scrappedRepo) CreateForm(ctx context.Context, sf *biz.ScrappedForm) (*biz.ScrappedForm, error) {

	return nil, nil
}
func (repo *scrappedRepo) UpdateForm(ctx context.Context, sf *biz.ScrappedForm) (*biz.ScrappedForm, error) {

	return nil, nil
}
