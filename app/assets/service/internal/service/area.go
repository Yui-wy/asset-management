package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) GetArea(ctx context.Context, req *pb.GetAreaReq) (*pb.GetAreaReply, error) {
	area, err := s.arc.Get(ctx, req.Id)
	return &pb.GetAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, err
}

func (s *AssetsService) ListAreas(ctx context.Context, req *pb.ListAreasReq) (*pb.ListAreasReply, error) {
	areas, err := s.arc.List(ctx)
	rs := make([]*pb.ListAreasReply_Areas, 0)
	for _, x := range areas {
		rs = append(rs, &pb.ListAreasReply_Areas{
			Id:       x.Id,
			AreaInfo: x.AreaInfo,
		})
	}
	return &pb.ListAreasReply{
		Areas: rs,
	}, err
}

func (s *AssetsService) CreateArea(ctx context.Context, req *pb.CreateAreaReq) (*pb.CreateAreaReply, error) {
	area, err := s.arc.Create(ctx, &biz.Area{
		AreaInfo:  req.AreaInfo,
		IsDeleted: false,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, err
}

func (s *AssetsService) UpdateArea(ctx context.Context, req *pb.UpdateAreaReq) (*pb.UpdateAreaReply, error) {
	area, err := s.arc.Update(ctx, &biz.Area{
		Id:       req.Id,
		AreaInfo: req.AreaInfo,
	})
	return &pb.UpdateAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, err
}

func (s *AssetsService) DeleteArea(ctx context.Context, req *pb.DeleteAreaReq) (*pb.DeleteAreaReply, error) {
	ok, err := s.arc.Delete(ctx, req.Id)
	return &pb.DeleteAreaReply{
		Ok: ok,
	}, err
}
