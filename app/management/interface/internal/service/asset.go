package service

import (
	"context"
	"time"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/Yui-wy/asset-management/pkg/setting"
)

func (s *ManageMentInterface) ListAsset(ctx context.Context, req *pb.ListAssetReq) (*pb.ListAssetReply, error) {
	_, err := s.getUserDetail(ctx, req.Con.AreaId)
	if err != nil {
		return nil, err
	}
	as, err := s.ac.List(ctx, &biz.AssetCondition{
		Classes:      req.Con.Classes,
		Address:      req.Con.Address,
		StateNum:     req.Con.StateNum,
		LowStorageAt: req.Con.LowStorageAt,
		UpStorageAt:  req.Con.UpStorageAt,
		OrderBy:      req.Con.OrderBy,
		SortDesc:     req.Con.SortDesc,
		AreaId:       req.Con.AreaId,
	},
		req.PageNum,
		req.PageSize)
	results := make([]*pb.ListAssetReply_Assets, 0)
	for _, r := range as {
		results = append(results, &pb.ListAssetReply_Assets{
			Id:         r.Id,
			Classes:    r.Classes,
			Code:       r.Code,
			AreaId:     r.AreaId,
			Address:    r.Address,
			AssetInfo:  r.AssetInfo,
			PicUrl:     r.PicUrl,
			Price:      r.Price,
			OrderAt:    r.OrderAt,
			OrderNum:   r.OrderNum,
			StateNum:   r.StateNum,
			State:      r.State,
			AppliedAt:  r.AppliedAt,
			StorageAt:  r.StorageAt,
			ScrappedAt: r.ScrappedAt,
		})
	}
	return &pb.ListAssetReply{
		Results: results,
	}, nil
}

func (s *ManageMentInterface) GetAsset(ctx context.Context, req *pb.GetAssetReq) (*pb.GetAssetReply, error) {
	r, err := s.ac.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.getUserDetail(ctx, []uint32{r.AreaId})
	if err != nil {
		return nil, err
	}
	return &pb.GetAssetReply{
		Id:         r.Id,
		Classes:    r.Classes,
		Code:       r.Code,
		AreaId:     r.AreaId,
		Address:    r.Address,
		AssetInfo:  r.AssetInfo,
		PicUrl:     r.PicUrl,
		Price:      r.Price,
		OrderAt:    r.OrderAt,
		OrderNum:   r.OrderNum,
		StateNum:   r.StateNum,
		State:      r.State,
		AppliedAt:  r.AppliedAt,
		StorageAt:  r.StorageAt,
		ScrappedAt: r.ScrappedAt,
	}, nil
}

func (s *ManageMentInterface) UpdateAsset(ctx context.Context, req *pb.UpdateAssetReq) (*pb.UpdateAssetReply, error) {
	r, err := s.ac.Update(ctx, &biz.Asset{
		Id:        req.Id,
		Address:   req.Address,
		AssetInfo: req.AssetInfo,
		PicUrl:    req.PicUrl,
		Price:     req.Price,
		OrderAt:   req.OrderAt,
		OrderNum:  req.OrderNum,
	})
	if err != nil {
		return nil, err
	}
	_, err = s.getUserDetail(ctx, []uint32{r.AreaId})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAssetReply{
		Id:         r.Id,
		Classes:    r.Classes,
		Code:       r.Code,
		AreaId:     r.AreaId,
		Address:    r.Address,
		AssetInfo:  r.AssetInfo,
		PicUrl:     r.PicUrl,
		Price:      r.Price,
		OrderAt:    r.OrderAt,
		OrderNum:   r.OrderNum,
		StateNum:   r.StateNum,
		State:      r.State,
		AppliedAt:  r.AppliedAt,
		StorageAt:  r.StorageAt,
		ScrappedAt: r.ScrappedAt}, nil
}

func (s *ManageMentInterface) ListStorageForm(ctx context.Context, req *pb.ListStorageFormReq) (*pb.ListStorageFormReply, error) {
	_, err := s.getUserDetail(ctx, req.Conf.AreaId)
	sfs, err := s.ac.ListStorageForm(ctx, &biz.StorageCondition{
		ApplicantId: req.Conf.ApplicantId,
		OperatorId:  req.Conf.OperatorId,
		StateNum:    req.Conf.StateNum,
		AssetId:     req.Conf.AssetId,
		AssetCode:   req.Conf.AssetCode,
		AreaId:      req.Conf.AreaId,
	}, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.ListStorageFormReply_Form, 0)
	for _, sf := range sfs {
		results = append(results, &pb.ListStorageFormReply_Form{
			Id:          sf.Id,
			AppliedAt:   sf.AppliedAt,
			ApplicantId: sf.ApplicantId,
			Applicant:   sf.Applicant,
			OperatedAt:  sf.OperatedAt,
			OperatorId:  sf.OperatorId,
			Operator:    sf.Operator,
			StateNum:    sf.StateNum,
			State:       sf.State,
			AssetId:     sf.AssetId,
			AssetCode:   sf.AssetCode,
			AreaId:      sf.AreaId,
		})
	}
	return &pb.ListStorageFormReply{
		Forms: results,
	}, nil
}
func (s *ManageMentInterface) GetStorageForm(ctx context.Context, req *pb.GetStorageFormReq) (*pb.GetStorageFormReply, error) {
	sf, err := s.ac.GetStorageForm(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.getUserDetail(ctx, []uint32{sf.AreaId})
	if err != nil {
		return nil, err
	}
	return &pb.GetStorageFormReply{
		Id:          sf.Id,
		AppliedAt:   sf.AppliedAt,
		ApplicantId: sf.ApplicantId,
		Applicant:   sf.Applicant,
		OperatedAt:  sf.OperatedAt,
		OperatorId:  sf.OperatorId,
		Operator:    sf.Operator,
		StateNum:    sf.StateNum,
		State:       sf.State,
		AssetId:     sf.AssetId,
		AssetCode:   sf.AssetCode,
		AreaId:      sf.AreaId,
	}, nil
}
func (s *ManageMentInterface) CreateStorageForm(ctx context.Context, req *pb.CreateStorageFormReq) (*pb.CreateStorageFormReply, error) {
	u, err := s.getUserDetail(ctx, []uint32{req.AreaId})
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	form, err := s.ac.CreateStorageForm(ctx, &biz.Asset{
		Classes:   req.Classes,
		AreaId:    req.AreaId,
		Address:   req.Address,
		AssetInfo: req.AssetInfo,
		PicUrl:    req.PicUrl,
		Price:     req.Price,
		OrderAt:   req.OrderAt,
		OrderNum:  req.OrderNum,
	}, u.Id, u.Username)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &pb.CreateStorageFormReply{
		Id:          form.Id,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OperatedAt:  form.OperatedAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, nil
}

// 重写
func (s *ManageMentInterface) CreateStoragesForm(ctx context.Context, req *pb.CreateStorageFormsReq) (*pb.CreateStorageFormsReply, error) {
	for _, v := range req.Assets {
		_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, []uint32{v.AreaId})
		if err != nil {
			s.log.Error(err)
			continue
		}
		_, err = s.CreateStorageForm(ctx, &pb.CreateStorageFormReq{
			Classes:   v.Classes,
			AreaId:    v.AreaId,
			Address:   v.Address,
			AssetInfo: v.AssetInfo,
			PicUrl:    v.PicUrl,
			Price:     v.Price,
			OrderAt:   v.OrderAt,
			OrderNum:  v.OrderNum,
		})
		if err != nil {
			s.log.Error(err)
			continue
		}
	}
	return &pb.CreateStorageFormsReply{
		Ok: true,
	}, nil
}

func (s *ManageMentInterface) UpdateStorageForm(ctx context.Context, req *pb.UpdateStorageFormReq) (*pb.UpdateStorageFormReply, error) {
	u, err := s.getUserDetail(ctx, []uint32{req.AreaId})
	if err != nil {
		return nil, err
	}
	if ((req.StateNum == setting.FORM_CONFIRMED) || (req.StateNum == setting.FORM_FAIL)) && (u.Power != setting.AREA_ADMIN_USER) {
		return nil, auth.ErrPowerFail
	}
	form, err := s.ac.UpdateStorageForm(ctx, &biz.StorageForm{
		Id:         req.Id,
		StateNum:   req.StateNum,
		OperatorId: u.Id,
		Operator:   u.Username,
		OperatedAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStorageFormReply{
		Id:          form.Id,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OperatedAt:  form.OperatedAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, nil
}

func (s *ManageMentInterface) ListScrappedForm(ctx context.Context, req *pb.ListScrappedFormReq) (*pb.ListScrappedFormReply, error) {
	_, err := s.getUserDetail(ctx, req.Conf.AreaId)
	sps, err := s.ac.ListScrappedForm(ctx, &biz.ScrappedCondition{
		ApplicantId: req.Conf.ApplicantId,
		OperatorId:  req.Conf.OperatorId,
		StateNum:    req.Conf.StateNum,
		AssetId:     req.Conf.AssetId,
		AssetCode:   req.Conf.AssetCode,
		AreaId:      req.Conf.AreaId,
	}, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.ListScrappedFormReply_Form, 0)
	for _, sp := range sps {
		results = append(results, &pb.ListScrappedFormReply_Form{
			Id:          sp.Id,
			AppliedAt:   sp.AppliedAt,
			ApplicantId: sp.ApplicantId,
			Applicant:   sp.Applicant,
			OperatedAt:  sp.OperatedAt,
			OperatorId:  sp.OperatorId,
			Operator:    sp.Operator,
			StateNum:    sp.StateNum,
			State:       sp.State,
			AssetId:     sp.AssetId,
			AssetCode:   sp.AssetCode,
			AreaId:      sp.AreaId,
		})
	}
	return &pb.ListScrappedFormReply{
		Forms: results,
	}, nil
}
func (s *ManageMentInterface) GetScrappedForm(ctx context.Context, req *pb.GetScrappedFormReq) (*pb.GetScrappedFormReply, error) {
	sp, err := s.ac.GetScrappedForm(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.getUserDetail(ctx, []uint32{sp.AreaId})
	if err != nil {
		return nil, err
	}
	return &pb.GetScrappedFormReply{
		Id:          sp.Id,
		AppliedAt:   sp.AppliedAt,
		ApplicantId: sp.ApplicantId,
		Applicant:   sp.Applicant,
		OperatedAt:  sp.OperatedAt,
		OperatorId:  sp.OperatorId,
		Operator:    sp.Operator,
		StateNum:    sp.StateNum,
		State:       sp.State,
		AssetId:     sp.AssetId,
		AssetCode:   sp.AssetCode,
		AreaId:      sp.AreaId,
	}, nil
}
func (s *ManageMentInterface) CreateScrappedForm(ctx context.Context, req *pb.CreateScrappedFormReq) (*pb.CreateScrappedFormReply, error) {
	u, err := s.getUserDetail(ctx, []uint32{req.AreaId})
	if err != nil {
		return nil, err
	}
	form, err := s.ac.CreateScrappedForm(ctx, &biz.ScrappedForm{
		AssetId:     req.AssetId,
		AssetCode:   req.AssetCode,
		AreaId:      req.AreaId,
		ApplicantId: u.Id,
		Applicant:   u.Username,
		AppliedAt:   time.Now().Unix(),
		StateNum:    setting.FORM_SUBMITTED,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateScrappedFormReply{
		Id:          form.Id,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OperatedAt:  form.OperatedAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, nil
}
func (s *ManageMentInterface) UpdateScrappedForm(ctx context.Context, req *pb.UpdateScrappedFormReq) (*pb.UpdateScrappedFormReply, error) {
	u, err := s.getUserDetail(ctx, []uint32{req.AreaId})
	if err != nil {
		return nil, err
	}
	if ((req.StateNum == setting.FORM_CONFIRMED) || (req.StateNum == setting.FORM_FAIL)) && (u.Power != setting.AREA_ADMIN_USER) {
		return nil, auth.ErrPowerFail
	}
	form, err := s.ac.UpdateScrappedForm(ctx, &biz.ScrappedForm{
		Id:         req.Id,
		StateNum:   req.StateNum,
		OperatorId: u.Id,
		Operator:   u.Username,
		OperatedAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateScrappedFormReply{
		Id:          form.Id,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OperatedAt:  form.OperatedAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, nil
}
