package data

import (
	"context"

	av1 "github.com/Yui-wy/asset-management/api/assets/service/v1"
	fv1 "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type assetRepo struct {
	data *Data
	log  *log.Helper
}

func NewAssetRepo(data *Data, logger log.Logger) biz.AssetRepo {
	return &assetRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

// asset
func (rp *assetRepo) ListAsset(ctx context.Context, c *biz.AssetCondition, pageNum, pageSize int64) ([]*biz.Asset, int64, error) {
	as, err := rp.data.ac.ListAssets(ctx, &av1.ListAssetsReq{
		PageNum:  pageNum,
		PageSize: pageSize,
		Conf: &av1.ListAssetsReq_Condition{
			Classes:      c.Classes,
			Address:      c.Address,
			StateNum:     c.StateNum,
			LowStorageAt: c.LowStorageAt,
			UpStorageAt:  c.UpStorageAt,
			OrderBy:      c.OrderBy,
			AreaId:       c.AreaId,
			SortDesc:     c.SortDesc,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	assets := make([]*biz.Asset, 0)
	for _, a := range as.Results {
		assets = append(assets, &biz.Asset{
			Id:         a.Id,
			Classes:    a.Classes,
			AreaId:     a.AreaId,
			Code:       a.Code,
			Address:    a.Address,
			AssetInfo:  a.AssetInfo,
			PicUrl:     a.PicUrl,
			Price:      a.Price,
			OrderAt:    a.OrderAt,
			OrderNum:   a.OrderNum,
			StateNum:   a.StateNum,
			State:      a.State,
			AppliedAt:  a.AppliedAt,
			StorageAt:  a.StorageAt,
			ScrappedAt: a.ScrappedAt,
		})
	}
	return assets, as.PageTotal, nil
}
func (rp *assetRepo) GetAsset(ctx context.Context, assetId uint64) (*biz.Asset, error) {
	a, err := rp.data.ac.GetAssets(ctx, &av1.GetAssetsReq{Id: assetId})
	if err != nil {
		return nil, err
	}
	return &biz.Asset{
		Id:         a.Id,
		Classes:    a.Classes,
		AreaId:     a.AreaId,
		Code:       a.Code,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}, nil
}
func (rp *assetRepo) UpdateAsset(ctx context.Context, asset *biz.Asset) (*biz.Asset, error) {
	a, err := rp.data.ac.UpdateAssets(ctx, &av1.UpdateAssetsdReq{
		Id:         asset.Id,
		Address:    asset.Address,
		AssetInfo:  asset.AssetInfo,
		PicUrl:     asset.PicUrl,
		Price:      asset.Price,
		OrderAt:    asset.OrderAt,
		OrderNum:   asset.OrderNum,
		StateNum:   asset.StateNum,
		StorageAt:  asset.StorageAt,
		ScrappedAt: asset.ScrappedAt,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Asset{
		Id:         a.Id,
		Classes:    a.Classes,
		AreaId:     a.AreaId,
		Code:       a.Code,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}, nil
}
func (rp *assetRepo) CreateAsset(ctx context.Context, asset *biz.Asset) (*biz.Asset, error) {
	a, err := rp.data.ac.CreateAssets(ctx, &av1.CreateAssetsReq{
		Classes:   asset.Classes,
		AreaId:    asset.AreaId,
		Address:   asset.Address,
		AssetInfo: asset.AssetInfo,
		PicUrl:    asset.PicUrl,
		Price:     asset.Price,
		OrderAt:   asset.OrderAt,
		OrderNum:  asset.OrderNum,
		StateNum:  asset.StateNum,
		AppliedAt: asset.AppliedAt,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Asset{
		Id:         a.Id,
		Classes:    a.Classes,
		AreaId:     a.AreaId,
		Code:       a.Code,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}, nil
}

// storage
func (rp *assetRepo) ListStorageForm(ctx context.Context, c *biz.StorageCondition, pageNum, pageSize int64) ([]*biz.StorageForm, int64, error) {
	fs, err := rp.data.fc.ListStorageForm(ctx, &fv1.ListStorageFormReq{PageNum: pageNum, PageSize: pageSize,
		Conf: &fv1.ListStorageFormReq_Conf{
			ApplicantId: c.ApplicantId,
			OperatorId:  c.OperatorId,
			StateNum:    c.StateNum,
			Operator:    c.Operator,
			Applicant:   c.Applicant,
			AssetCode:   c.AssetCode,
			AreaId:      c.AreaId,
		}})
	if err != nil {
		return nil, 0, err
	}
	results := make([]*biz.StorageForm, 0)
	for _, f := range fs.Forms {
		results = append(results, &biz.StorageForm{
			Id:          f.Id,
			AppliedAt:   f.AppliedAt,
			ApplicantId: f.ApplicantId,
			Applicant:   f.Applicant,
			OperatedAt:  f.OperatedAt,
			OperatorId:  f.OperatorId,
			Operator:    f.Operator,
			StateNum:    f.StateNum,
			State:       f.State,
			AssetId:     f.AssetId,
			AssetCode:   f.AssetCode,
			AreaId:      f.AreaId,
		})
	}
	return results, fs.PageTotal, nil
}
func (rp *assetRepo) GetStorageForm(ctx context.Context, id int64) (*biz.StorageForm, error) {
	f, err := rp.data.fc.GetStorageForm(ctx, &fv1.GetStorageFormReq{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.StorageForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}

func (rp *assetRepo) CreateStorageForm(ctx context.Context, form *biz.StorageForm) (*biz.StorageForm, error) {
	f, err := rp.data.fc.CreateStorageForm(ctx, &fv1.CreateStorageFormReq{
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
		StateNum:    form.StateNum,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		AppliedAt:   form.AppliedAt,
	})
	if err != nil {
		return nil, err
	}
	return &biz.StorageForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}
func (rp *assetRepo) UpdateStorageForm(ctx context.Context, form *biz.StorageForm) (*biz.StorageForm, error) {
	f, err := rp.data.fc.UpdateStorageForm(ctx, &fv1.UpdateStorageFormReq{
		Id:         form.Id,
		OperatedAt: form.OperatedAt,
		OperatorId: form.OperatorId,
		Operator:   form.Operator,
		StateNum:   form.StateNum,
	})
	if err != nil {
		return nil, err
	}
	return &biz.StorageForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}

// scrapp
func (rp *assetRepo) ListScrappedForm(ctx context.Context, c *biz.ScrappedCondition, pageNum, pageSize int64) ([]*biz.ScrappedForm, int64, error) {
	fs, err := rp.data.fc.ListScrappedForm(ctx, &fv1.ListScrappedFormReq{PageNum: pageNum, PageSize: pageSize,
		Conf: &fv1.ListScrappedFormReq_Conf{
			ApplicantId: c.ApplicantId,
			OperatorId:  c.OperatorId,
			StateNum:    c.StateNum,
			Operator:    c.Operator,
			Applicant:   c.Applicant,
			AssetCode:   c.AssetCode,
			AreaId:      c.AreaId,
		}})
	if err != nil {
		return nil, 0, err
	}
	results := make([]*biz.ScrappedForm, 0)
	for _, f := range fs.Forms {
		results = append(results, &biz.ScrappedForm{
			Id:          f.Id,
			AppliedAt:   f.AppliedAt,
			ApplicantId: f.ApplicantId,
			Applicant:   f.Applicant,
			OperatedAt:  f.OperatedAt,
			OperatorId:  f.OperatorId,
			Operator:    f.Operator,
			StateNum:    f.StateNum,
			State:       f.State,
			AssetId:     f.AssetId,
			AssetCode:   f.AssetCode,
			AreaId:      f.AreaId,
		})
	}
	return results, fs.PageTotal, nil
}
func (rp *assetRepo) GetScrappedForm(ctx context.Context, id int64) (*biz.ScrappedForm, error) {
	f, err := rp.data.fc.GetScrappedForm(ctx, &fv1.GetScrappedFormReq{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.ScrappedForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}
func (rp *assetRepo) CreateScrappedForm(ctx context.Context, form *biz.ScrappedForm) (*biz.ScrappedForm, error) {
	f, err := rp.data.fc.CreateScrappedForm(ctx, &fv1.CreateScrappedFormReq{
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
		StateNum:    form.StateNum,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		AppliedAt:   form.AppliedAt,
	})
	if err != nil {
		return nil, err
	}
	return &biz.ScrappedForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}
func (rp *assetRepo) UpdateScrappedForm(ctx context.Context, form *biz.ScrappedForm) (*biz.ScrappedForm, error) {
	f, err := rp.data.fc.UpdateScrappedForm(ctx, &fv1.UpdateScrappedFormReq{
		Id:         form.Id,
		OperatedAt: form.OperatedAt,
		OperatorId: form.OperatorId,
		Operator:   form.Operator,
		StateNum:   form.StateNum,
	})
	if err != nil {
		return nil, err
	}
	return &biz.ScrappedForm{
		Id:          f.Id,
		AppliedAt:   f.AppliedAt,
		ApplicantId: f.ApplicantId,
		Applicant:   f.Applicant,
		OperatedAt:  f.OperatedAt,
		OperatorId:  f.OperatorId,
		Operator:    f.Operator,
		StateNum:    f.StateNum,
		State:       f.State,
		AssetId:     f.AssetId,
		AssetCode:   f.AssetCode,
		AreaId:      f.AreaId,
	}, nil
}

func (rp *assetRepo) GetClasses(ctx context.Context) ([]*biz.Class, error) {
	cs, err := rp.data.ac.GetClasses(ctx, &av1.GetClassesReq{})
	if err != nil {
		rp.log.Errorf("Error: %e", err)
		return nil, err
	}
	results := make([]*biz.Class, 0)
	for _, c := range cs.Clz {
		results = append(results, &biz.Class{
			Id:      c.Id,
			Code:    c.Code,
			Pcode:   c.Pcode,
			ClzInfo: c.ClzInfo,
			Level:   c.Level,
		})
	}
	return results, nil
}

func (rp *assetRepo) CreateClzz(ctx context.Context, clzz []*biz.Class) error {
	clzzReq := make([]*av1.CreateClassesReq_Classes, 0)
	for _, clz := range clzz {
		clzzReq = append(clzzReq, &av1.CreateClassesReq_Classes{
			Code:    clz.Code,
			ClzInfo: clz.ClzInfo,
			Level:   clz.Level,
			Pcode:   clz.Pcode,
		})
	}
	_, err := rp.data.ac.CreateClasses(ctx, &av1.CreateClassesReq{
		Clz: clzzReq,
	})
	return err
}