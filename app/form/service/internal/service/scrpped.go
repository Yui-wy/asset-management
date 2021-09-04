package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
)

func (s *FormService) GetScrappedForm(ctx context.Context, req *pb.GetScrappedFormReq) (*pb.GetScrappedFormReply, error) {
	form, err := s.spu.Get(ctx, req.Id)
	return &pb.GetScrappedFormReply{
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
	}, err
}
func (s *FormService) ListScrappedForm(ctx context.Context, req *pb.ListScrappedFormReq) (*pb.ListScrappedFormReply, error) {
	forms, err := s.spu.List(ctx, &biz.ScConfig{
		BaseConfig: &biz.BaseConfig{
			ApplicantId: req.Conf.ApplicantId,
			OperatorId:  req.Conf.OperatorId,
			StateNum:    req.Conf.StateNum,
			AssetId:     req.Conf.AssetId,
			AssetCode:   req.Conf.AssetCode,
			AreaId:      req.Conf.AreaId,
		},
	}, req.PageNum, req.PageSize)
	reply := make([]*pb.ListScrappedFormReply_Form, 0)
	for _, form := range forms {
		reply = append(reply, &pb.ListScrappedFormReply_Form{
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
	return &pb.ListScrappedFormReply{
		Forms: reply,
	}, err
}
func (s *FormService) CreateScrappedForm(ctx context.Context, req *pb.CreateScrappedFormReq) (*pb.CreateScrappedFormReply, error) {
	form, err := s.spu.Create(ctx, &biz.ScrappedForm{
		BaseForm: &biz.BaseForm{
			AreaId:      req.AreaId,
			ApplicantId: req.ApplicantId,
			Applicant:   req.Applicant,
			AssetId:     req.AssetId,
			AssetCode:   req.AssetCode,
		},
	})
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
	}, err
}
func (s *FormService) UpdateScrappedForm(ctx context.Context, req *pb.UpdateScrappedFormReq) (*pb.UpdateScrappedFormReply, error) {
	form, err := s.spu.Update(ctx, &biz.ScrappedForm{
		BaseForm: &biz.BaseForm{
			Id:         req.Id,
			OperatedAt: req.OperatedAt,
			OperatorId: req.OperatorId,
			Operator:   req.Operator,
			StateNum:   req.StateNum,
		},
	})
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
	}, err
}
