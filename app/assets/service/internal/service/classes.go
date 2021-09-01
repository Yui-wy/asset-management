package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) CreateClasses(ctx context.Context, req *pb.CreateClassesReq) (*pb.CreateClassesReply, error) {
	clz := make([]*biz.Class, 0)
	for _, c := range req.Clz {
		clz = append(clz, &biz.Class{
			Code:    c.Code,
			ClzInfo: c.ClzInfo,
			Level:   c.Level,
			Pcode:   c.Pcode,
		})
	}
	cs, err := s.cc.Create(ctx, clz)
	if err != nil {
		return nil, err
	}
	rc := make([]*pb.CreateClassesReply_Classes, 0)
	for _, c := range cs {
		rc = append(rc, &pb.CreateClassesReply_Classes{
			Id:      c.Id,
			Code:    c.Code,
			ClzInfo: c.ClzInfo,
			Level:   c.Level,
			Pcode:   c.Pcode,
		})
	}
	return &pb.CreateClassesReply{
		Clz: rc,
	}, nil
}

func (s *AssetsService) GetClasses(ctx context.Context, req *pb.GetClassesReq) (*pb.GetClassesReply, error) {
	cs, err := s.cc.Get(ctx)
	if err != nil {
		return nil, err
	}
	rc := make([]*pb.GetClassesReply_Classes, 0)
	for _, c := range cs {
		rc = append(rc, &pb.GetClassesReply_Classes{
			Id:      c.Id,
			Code:    c.Code,
			ClzInfo: c.ClzInfo,
			Level:   c.Level,
			Pcode:   c.Pcode,
		})
	}
	return &pb.GetClassesReply{
		Clz: rc,
	}, nil
}
