package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/form/service/v1"
	"github.com/Yui-wy/asset-management/app/form/service/internal/biz"
)

func (s *FormService) GetScrappedForm(ctx context.Context, req *pb.GetScrappedFormReq) (*pb.GetScrappedFormReply, error) {
	form, err := s.spu.Get(ctx, req.FormNum)
	return &pb.GetScrappedFormReply{
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
func (s *FormService) ListScrappedForm(ctx context.Context, req *pb.ListScrappedFormReq) (*pb.ListScrappedFormReply, error) {
	forms, err := s.spu.List(ctx, &biz.ScConfig{
		ApplicantId: req.Conf.ApplicantId,
		OperatorId:  req.Conf.OperatorId,
		StateNum:    req.Conf.StateNum,
		AssetId:     req.Conf.AssetId,
		AssetCode:   req.Conf.AssetCode,
		AreaId:      req.Conf.AreaId,
	})
	reply := make([]*pb.ListScrappedFormReply_Form, 0)
	for _, form := range forms {
		reply = append(reply, &pb.ListScrappedFormReply_Form{
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
	return &pb.ListScrappedFormReply{
		Forms: reply,
	}, err
}
func (s *FormService) CreateScrappedForm(ctx context.Context, req *pb.CreateScrappedFormReq) (*pb.CreateScrappedFormReply, error) {
	form, err := s.spu.Create(ctx, &biz.ScrappedForm{
		ApplicantId: req.ApplicantId,
		Applicant:   req.Applicant,
		AssetId:     req.AssetId,
		AssetCode:   req.AssetCode,
		AreaId:      req.AreaId,
	})
	return &pb.CreateScrappedFormReply{
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
func (s *FormService) UpdateScrappedForm(ctx context.Context, req *pb.UpdateScrappedFormReq) (*pb.UpdateScrappedFormReply, error) {
	form, err := s.spu.Update(ctx, &biz.ScrappedForm{
		FormNum:   req.FormNum,
		OptAt:     req.OptAt,
		AssetId:   req.OperatorId,
		AssetCode: req.Operator,
		StateNum:  req.StateNum,
	})
	return &pb.UpdateScrappedFormReply{
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
