package data

import (
	"context"
	"time"

	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type storageRepo struct {
	data *Data
	log  *log.Helper
}

type StorageForm struct {
	ID          int64  `gorm:"primarykey;autoIncrement:false"`
	AreaId      uint32 `gorm:"not null"`
	AppliedAt   int64  `gorm:"not null"`
	ApplicantId uint64 `gorm:"not null"`
	Applicant   string `gorm:"not null"`
	OperatedAt  int64
	OperatorId  uint64
	Operator    string
	StateNum    int32  `gorm:"not null"`
	AssetId     uint64 `gorm:"not null"`
	AssetCode   string `gorm:"not null"`
	// ====================
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStorageRepo(data *Data, logger log.Logger) biz.StorageRepo {
	return &storageRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/storage")),
	}
}

func (repo *storageRepo) GetForm(ctx context.Context, id int64) (*biz.StorageForm, error) {
	form := &StorageForm{}
	result := repo.data.db.WithContext(ctx).First(form, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.setbiz(form), nil
}

func (repo *storageRepo) ListForm(ctx context.Context, conf *biz.StConfig, pageNum, pageSize int64) ([]*biz.StorageForm, error) {
	var forms []StorageForm
	result := repo.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize)))

	if inspection.IsZeros(conf.AreaId) {
		err := errors.New(500, "AreaId is nil", "please set areaId")
		repo.log.Errorf(" ListForm1. Error:%d", err)
		return nil, err
	}
	result = result.Where("area_id IN ?", conf.AreaId)
	if !inspection.IsZeros(conf.ApplicantId) {
		result = result.Where("applicant_id = ?", conf.ApplicantId)
	}
	if !inspection.IsZeros(conf.OperatorId) {
		result = result.Where("operator_id = ?", conf.OperatorId)
	}
	if !inspection.IsZeros(conf.StateNum) {
		result = result.Where("state_num = ?", conf.StateNum)
	}
	if !inspection.IsZeros(conf.AssetId) {
		result = result.Where("asset_id = ?", conf.AssetId)
	}
	if !inspection.IsZeros(conf.AssetCode) {
		result = result.Where("asset_code like ?", conf.AssetCode)
	}
	result = result.Find(&forms)
	if result.Error != nil {
		repo.log.Errorf(" ListForm2. Error:%d", result.Error)
		return nil, result.Error
	}
	bsfs := make([]*biz.StorageForm, 0)
	for _, f := range forms {
		bsfs = append(bsfs, repo.setbiz(&f))
	}
	return bsfs, nil
}

func (repo *storageRepo) CreateForm(ctx context.Context, sf *biz.StorageForm) (*biz.StorageForm, error) {
	id, err := repo.data.sf.NextVal()
	if err != nil {
		return nil, err
	}
	form := &StorageForm{
		ID:          id,
		AreaId:      sf.AreaId,
		AppliedAt:   time.Now().Unix(),
		ApplicantId: sf.ApplicantId,
		Applicant:   sf.Applicant,
		StateNum:    setting.FORM_SUBMITTED,
		AssetId:     sf.AssetId,
		AssetCode:   sf.AssetCode,
	}
	result := repo.data.db.WithContext(ctx).Create(&form)
	if result.Error != nil {
		repo.log.Errorf(" CreateForm1. Error:%d", result.Error)
		return nil, result.Error
	}
	sff := StorageForm{}
	result = repo.data.db.WithContext(ctx).First(&sff, form.ID)
	if result.Error != nil {
		repo.log.Errorf(" CreateForm2. Error: %d", result.Error)
		return nil, result.Error
	}
	return repo.setbiz(&sff), nil
}
func (repo *storageRepo) UpdateForm(ctx context.Context, sf *biz.StorageForm) (*biz.StorageForm, error) {
	s := StorageForm{
		ID: sf.Id,
	}
	result := repo.data.db.WithContext(ctx).Model(&s).Updates(StorageForm{
		OperatedAt: sf.OperatedAt,
		OperatorId: sf.OperatorId,
		Operator:   sf.Operator,
		StateNum:   sf.StateNum,
	})
	if result.Error != nil {
		repo.log.Errorf(" UpdateForm1. Error: %d", result.Error)
		return nil, result.Error
	}
	sff := StorageForm{}
	result = repo.data.db.WithContext(ctx).First(&sff, s.ID)
	if result.Error != nil {
		repo.log.Errorf(" UpdateForm2. Error: %d", result.Error)
		return nil, result.Error
	}
	return repo.setbiz(&sff), nil
}

func (repo *storageRepo) setbiz(sf *StorageForm) *biz.StorageForm {
	state, ok := setting.FORM_STATE_MAP[sf.StateNum]
	if !ok {
		state = setting.FORM_STATE_MAP[setting.FORM_UNKNOWN]
	}
	return &biz.StorageForm{
		BaseForm: &biz.BaseForm{
			Id:          sf.ID,
			AppliedAt:   sf.AppliedAt,
			ApplicantId: sf.ApplicantId,
			Applicant:   sf.Applicant,
			OperatedAt:  sf.OperatedAt,
			OperatorId:  sf.OperatorId,
			Operator:    sf.Operator,
			StateNum:    sf.StateNum,
			State:       state,
			AssetId:     sf.AssetId,
			AssetCode:   sf.AssetCode,
			AreaId:      sf.AreaId,
		},
	}
}
