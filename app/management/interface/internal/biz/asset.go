package biz

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Asset struct {
	Id         uint64
	Classes    string
	Code       string
	AreaId     uint32
	Address    string
	AssetInfo  string
	PicUrl     string
	Price      float32
	OrderAt    int64
	OrderNum   string
	StateNum   int32
	State      string
	AppliedAt  int64
	StorageAt  int64
	ScrappedAt int64
}

type AssetCondition struct {
	Classes      string
	Address      string
	StateNum     int32
	LowStorageAt int64
	UpStorageAt  int64
	OrderBy      string
	SortDesc     bool
	AreaId       []uint32
}

type StorageForm struct {
	Id          int64
	AppliedAt   int64
	ApplicantId uint64
	Applicant   string
	OperatedAt  int64
	OperatorId  uint64
	Operator    string
	StateNum    int32
	State       string
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type StorageCondition struct {
	ApplicantId uint64
	OperatorId  uint64
	StateNum    int32
	AssetId     uint64
	AssetCode   string
	AreaId      []uint32
}

type ScrappedForm struct {
	Id          int64
	AppliedAt   int64
	ApplicantId uint64
	Applicant   string
	OperatedAt  int64
	OperatorId  uint64
	Operator    string
	StateNum    int32
	State       string
	AssetId     uint64
	AssetCode   string
	AreaId      uint32
}

type ScrappedCondition struct {
	ApplicantId uint64
	OperatorId  uint64
	StateNum    int32
	AssetId     uint64
	AssetCode   string
	AreaId      []uint32
}

type AssetRepo interface {
	// asset
	ListAsset(ctx context.Context, condition *AssetCondition, pageNum, pageSize int64) ([]*Asset, error)
	GetAsset(ctx context.Context, assetId uint64) (*Asset, error)
	UpdateAsset(ctx context.Context, asset *Asset) (*Asset, error)
	CreateAsset(ctx context.Context, asset *Asset) (*Asset, error)
	// storage
	ListStorageForm(ctx context.Context, condition *StorageCondition, pageNum, pageSize int64) ([]*StorageForm, error)
	GetStorageForm(ctx context.Context, id int64) (*StorageForm, error)
	CreateStorageForm(ctx context.Context, form *StorageForm) (*StorageForm, error)
	UpdateStorageForm(ctx context.Context, form *StorageForm) (*StorageForm, error)
	// scrapp
	ListScrappedForm(ctx context.Context, condition *ScrappedCondition, pageNum, pageSize int64) ([]*ScrappedForm, error)
	GetScrappedForm(ctx context.Context, id int64) (*ScrappedForm, error)
	CreateScrappedForm(ctx context.Context, form *ScrappedForm) (*ScrappedForm, error)
	UpdateScrappedForm(ctx context.Context, form *ScrappedForm) (*ScrappedForm, error)
}

type AssetUseCase struct {
	repo AssetRepo
	log  *log.Helper
}

func NewAssetUseCase(repo AssetRepo, logger log.Logger) *AssetUseCase {
	return &AssetUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usercase/asset")),
	}
}

func (s *AssetUseCase) List(ctx context.Context, condition *AssetCondition, pageNum, pageSize int64) ([]*Asset, error) {
	return s.repo.ListAsset(ctx, condition, pageNum, pageSize)
}

func (s *AssetUseCase) Get(ctx context.Context, assetId uint64) (*Asset, error) {
	return s.repo.GetAsset(ctx, assetId)
}

func (s *AssetUseCase) Update(ctx context.Context, asset *Asset) (*Asset, error) {
	return s.repo.UpdateAsset(ctx, asset)
}

func (s *AssetUseCase) ListStorageForm(ctx context.Context, condition *StorageCondition, pageNum, pageSize int64) ([]*StorageForm, error) {
	return s.repo.ListStorageForm(ctx, condition, pageNum, pageSize)
}
func (s *AssetUseCase) GetStorageForm(ctx context.Context, id int64) (*StorageForm, error) {
	return s.repo.GetStorageForm(ctx, id)
}
func (s *AssetUseCase) CreateStorageForm(ctx context.Context, asset *Asset, uid uint64, username string) (*StorageForm, error) {
	asset.AppliedAt = time.Now().Unix()
	asset.StateNum = setting.ASSETS_STATE_ST_APPLY
	a, err := s.repo.CreateAsset(ctx, asset)
	if err != nil {
		return nil, err
	}
	form, err := s.repo.CreateStorageForm(ctx, &StorageForm{
		AssetId:     a.Id,
		AssetCode:   a.Code,
		AreaId:      a.AreaId,
		StateNum:    setting.FORM_SUBMITTED,
		ApplicantId: uid,
		Applicant:   username,
		AppliedAt:   a.AppliedAt,
	})
	if err != nil {
		return nil, err
	}
	return form, nil
}
func (s *AssetUseCase) UpdateStorageForm(ctx context.Context, form *StorageForm) (*StorageForm, error) {
	f, err := s.repo.UpdateStorageForm(ctx, form)
	if err != nil {
		return nil, err
	}
	stateNum := setting.ASSETS_STATE_ST
	storageAt := f.OperatedAt
	if (f.StateNum == setting.FORM_CANCELED) || (f.StateNum == setting.FORM_FAIL) {
		stateNum = setting.ASSETS_STATE_ST_FAIL
		storageAt = 0
	}
	_, err = s.repo.UpdateAsset(ctx, &Asset{
		Id:        f.AssetId,
		StateNum:  int32(stateNum),
		StorageAt: storageAt,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *AssetUseCase) ListScrappedForm(ctx context.Context, condition *ScrappedCondition, pageNum, pageSize int64) ([]*ScrappedForm, error) {
	return s.repo.ListScrappedForm(ctx, condition, pageNum, pageSize)
}
func (s *AssetUseCase) GetScrappedForm(ctx context.Context, id int64) (*ScrappedForm, error) {
	return s.repo.GetScrappedForm(ctx, id)
}
func (s *AssetUseCase) CreateScrappedForm(ctx context.Context, form *ScrappedForm) (*ScrappedForm, error) {
	asset, err := s.repo.GetAsset(ctx, form.AssetId)
	if err != nil {
		return nil, err
	}
	if (asset.StateNum == setting.ASSETS_STATE_SP) ||
		(asset.StateNum == setting.ASSETS_STATE_SP_APPLY) ||
		(asset.StateNum == setting.ASSETS_STATE_ST_APPLY) ||
		(asset.StateNum == setting.ASSETS_STATE_ORDER_APPLY) {
		return nil, errors.New(500, "cannot scrapped", "cannot scrapped")
	}
	f, err := s.repo.CreateScrappedForm(ctx, form)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.UpdateAsset(ctx, &Asset{
		Id:       f.AssetId,
		StateNum: setting.ASSETS_STATE_SP_APPLY,
	})
	return f, nil
}
func (s *AssetUseCase) UpdateScrappedForm(ctx context.Context, form *ScrappedForm) (*ScrappedForm, error) {
	f, err := s.repo.UpdateScrappedForm(ctx, form)
	if err != nil {
		return nil, err
	}
	stateNum := setting.ASSETS_STATE_SP
	scrappedAt := f.OperatedAt
	if (f.StateNum == setting.FORM_CANCELED) || (f.StateNum == setting.FORM_FAIL) {
		stateNum = setting.ASSETS_STATE_ST
		scrappedAt = 0
	}
	_, err = s.repo.UpdateAsset(ctx, &Asset{
		Id:         f.AssetId,
		StateNum:   int32(stateNum),
		ScrappedAt: scrappedAt,
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}
