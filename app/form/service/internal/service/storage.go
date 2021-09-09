package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
)

func (s *FormService) GetStorageForm(ctx context.Context, req *pb.GetStorageFormReq) (*pb.GetStorageFormReply, error) {
	form, err := s.stu.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStorageFormReply{
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
func (s *FormService) ListStorageForm(ctx context.Context, req *pb.ListStorageFormReq) (*pb.ListStorageFormReply, error) {
	forms, err := s.stu.List(ctx,
		&biz.StConfig{
			BaseConfig: &biz.BaseConfig{
				AreaId:      req.Conf.AreaId,
				ApplicantId: req.Conf.ApplicantId,
				OperatorId:  req.Conf.OperatorId,
				StateNum:    req.Conf.StateNum,
				AssetId:     req.Conf.AssetId,
				AssetCode:   req.Conf.AssetCode,
			},
		}, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	reply := make([]*pb.ListStorageFormReply_Form, 0)
	for _, form := range forms {
		reply = append(reply, &pb.ListStorageFormReply_Form{
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
		})
	}
	return &pb.ListStorageFormReply{
		Forms: reply,
	}, nil
}
func (s *FormService) CreateStorageForm(ctx context.Context, req *pb.CreateStorageFormReq) (*pb.CreateStorageFormReply, error) {
	form, err := s.stu.Create(ctx, &biz.StorageForm{
		BaseForm: &biz.BaseForm{
			ApplicantId: req.ApplicantId,
			Applicant:   req.Applicant,
			AppliedAt:   req.AppliedAt,
			AssetId:     req.AssetId,
			AssetCode:   req.AssetCode,
			AreaId:      req.AreaId,
			StateNum:    req.StateNum,
		},
	})
	if err != nil {
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
func (s *FormService) UpdateStorageForm(ctx context.Context, req *pb.UpdateStorageFormReq) (*pb.UpdateStorageFormReply, error) {
	form, err := s.stu.Update(ctx, &biz.StorageForm{
		BaseForm: &biz.BaseForm{
			Id:         req.Id,
			OperatedAt: req.OperatedAt,
			OperatorId: req.OperatorId,
			Operator:   req.Operator,
			StateNum:   req.StateNum,
		},
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
