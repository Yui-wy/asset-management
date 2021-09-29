package data

import (
	"context"
	"math"
	"time"

	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/setting"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/Yui-wy/asset-management/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type scrappedRepo struct {
	data *Data
	log  *log.Helper
}

type ScrappedForm struct {
	ID          int64 `gorm:"primarykey;autoIncrement:false"`
	AreaId      uint32
	AppliedAt   int64
	ApplicantId uint64
	Applicant   string
	OperatedAt  int64
	OperatorId  uint64
	Operator    string
	StateNum    int32
	AssetId     uint64
	AssetCode   string
	// ====================
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewScrappedRepo(data *Data, logger log.Logger) biz.ScrappedRepo {
	return &scrappedRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/scrapped")),
	}

}

func (repo *scrappedRepo) GetForm(ctx context.Context, id int64) (*biz.ScrappedForm, error) {
	form := ScrappedForm{}
	result := repo.data.db.WithContext(ctx).First(&form, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return repo.setbiz(&form), nil
}

func (repo *scrappedRepo) ListForm(ctx context.Context, conf *biz.ScConfig, pageNum, pageSize int64) ([]*biz.ScrappedForm, int64, error) {
	var forms []ScrappedForm
	result := repo.data.db.
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize)))

	if inspection.IsZeros(conf.AreaId) {
		err := errors.New(500, "AreaId is nil", "please set areaId")
		repo.log.Errorf(" ListForm1. Error:%d", err)
		return nil, 0, err
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
	if !inspection.IsZeros(conf.Applicant) {
		result = result.Where("INSTR(applicant, ?) > 0", conf.Applicant)
	}
	if !inspection.IsZeros(conf.Operator) {
		result = result.Where("INSTR(operator, ?) > 0", conf.Operator)
	}
	if !inspection.IsZeros(conf.AssetCode) {
		result = result.Where("INSTR(asset_code, ?) > 0", conf.AssetCode)
	}
	tx := result.WithContext(ctx)
	result = tx.Find(&forms)
	if result.Error != nil {
		repo.log.Errorf(" ListForm2. Error:%d", result.Error)
		return nil, 0, result.Error
	}
	var total int64
	result = tx.Count(&total)
	if result.Error != nil {
		repo.log.Errorf(" ListForm2. Error:%d", result.Error)
		return nil, 0, result.Error
	}
	totalPage := int64(math.Ceil(float64(total) / float64(pageSize)))
	bsps := make([]*biz.ScrappedForm, 0)
	for _, f := range forms {
		bsps = append(bsps, repo.setbiz(&f))
	}
	return bsps, totalPage, nil
}
func (repo *scrappedRepo) CreateForm(ctx context.Context, sf *biz.ScrappedForm) (*biz.ScrappedForm, error) {
	id, err := repo.data.sf.NextVal()
	if err != nil {
		return nil, err
	}
	form := &ScrappedForm{
		ID:          id,
		AreaId:      sf.AreaId,
		ApplicantId: sf.ApplicantId,
		Applicant:   sf.Applicant,
		AppliedAt:   sf.AppliedAt,
		StateNum:    sf.StateNum,
		AssetId:     sf.AssetId,
		AssetCode:   sf.AssetCode,
	}
	result := repo.data.db.WithContext(ctx).Create(&form)
	if result.Error != nil {
		repo.log.Errorf(" CreateForm1. Error:%d", result.Error)
		return nil, result.Error
	}
	spf := ScrappedForm{}
	result = repo.data.db.WithContext(ctx).First(&spf, form.ID)
	if result.Error != nil {
		repo.log.Errorf(" CreateForm2. Error: %d", result.Error)
		return nil, result.Error
	}
	return repo.setbiz(&spf), nil
}
func (repo *scrappedRepo) UpdateForm(ctx context.Context, sf *biz.ScrappedForm) (*biz.ScrappedForm, error) {
	spf := ScrappedForm{}
	result := repo.data.db.WithContext(ctx).First(&spf, sf.Id)
	if result.Error != nil {
		repo.log.Errorf(" UpdateForm1. Error: %d", result.Error)
		return nil, result.Error
	}
	if spf.StateNum != setting.FORM_SUBMITTED {
		return nil, errors.New(500, "form cant be operated", "Form is locked.")
	}
	spf.OperatedAt = sf.OperatedAt
	spf.OperatorId = sf.OperatorId
	spf.Operator = sf.Operator
	spf.StateNum = sf.StateNum
	result = repo.data.db.WithContext(ctx).Save(&spf)
	if result.Error != nil {
		repo.log.Errorf(" UpdateForm2. Error: %d", result.Error)
		return nil, result.Error
	}
	return repo.setbiz(&spf), nil
}

func (repo *scrappedRepo) setbiz(sf *ScrappedForm) *biz.ScrappedForm {
	state, ok := setting.FORM_STATE_MAP[sf.StateNum]
	if !ok {
		state = setting.FORM_STATE_MAP[setting.FORM_UNKNOWN]
	}
	return &biz.ScrappedForm{
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
