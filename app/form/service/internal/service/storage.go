package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
)

func (s *FormService) GetStorageForm(ctx context.Context, req *pb.GetStorageFormReq) (*pb.GetStorageFormReply, error) {
	form, err := s.stu.Get(ctx, req.FormNum)
	return &pb.GetStorageFormReply{
		FormNum:     form.FormNum,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OptAt:       form.OptAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, err
}
func (s *FormService) ListStorageForm(ctx context.Context, req *pb.ListStorageFormReq) (*pb.ListStorageFormReply, error) {
	forms, err := s.stu.List(ctx, &biz.StConfig{
		ApplicantId: req.Conf.ApplicantId,
		OperatorId:  req.Conf.OperatorId,
		StateNum:    req.Conf.StateNum,
		AssetId:     req.Conf.AssetId,
		AssetCode:   req.Conf.AssetCode,
		AreaId:      req.Conf.AreaId,
	})
	reply := make([]*pb.ListStorageFormReply_Form, 0)
	for _, form := range forms {
		reply = append(reply, &pb.ListStorageFormReply_Form{
			FormNum:     form.FormNum,
			AppliedAt:   form.AppliedAt,
			ApplicantId: form.ApplicantId,
			Applicant:   form.Applicant,
			OptAt:       form.OptAt,
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
	}, err
}
func (s *FormService) CreateStorageForm(ctx context.Context, req *pb.CreateStorageFormReq) (*pb.CreateStorageFormReply, error) {
	form, err := s.stu.Create(ctx, &biz.StorageForm{
		ApplicantId: req.ApplicantId,
		Applicant:   req.Applicant,
		AssetId:     req.AssetId,
		AssetCode:   req.AssetCode,
		AreaId:      req.AreaId,
	})
	return &pb.CreateStorageFormReply{
		FormNum:     form.FormNum,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OptAt:       form.OptAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, err
}
func (s *FormService) UpdateStorageForm(ctx context.Context, req *pb.UpdateStorageFormReq) (*pb.UpdateStorageFormReply, error) {
	form, err := s.stu.Update(ctx, &biz.StorageForm{
		FormNum:   req.FormNum,
		OptAt:     req.OptAt,
		AssetId:   req.OperatorId,
		AssetCode: req.Operator,
		StateNum:  req.StateNum,
	})
	return &pb.UpdateStorageFormReply{
		FormNum:     form.FormNum,
		AppliedAt:   form.AppliedAt,
		ApplicantId: form.ApplicantId,
		Applicant:   form.Applicant,
		OptAt:       form.OptAt,
		OperatorId:  form.OperatorId,
		Operator:    form.Operator,
		StateNum:    form.StateNum,
		State:       form.State,
		AssetId:     form.AssetId,
		AssetCode:   form.AssetCode,
		AreaId:      form.AreaId,
	}, err
}
