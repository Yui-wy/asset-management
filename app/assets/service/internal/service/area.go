package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) GetArea(ctx context.Context, req *pb.GetAreaReq) (*pb.GetAreaReply, error) {
	area, err := s.arc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, nil
}

func (s *AssetsService) ListAreas(ctx context.Context, req *pb.ListAreasReq) (*pb.ListAreasReply, error) {
	areas, pageTotal, err := s.arc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.ListAreasReply_Areas, 0)
	for _, x := range areas {
		rs = append(rs, &pb.ListAreasReply_Areas{
			Id:       x.Id,
			AreaInfo: x.AreaInfo,
		})
	}
	return &pb.ListAreasReply{
		Areas:     rs,
		PageTotal: pageTotal,
	}, nil
}

func (s *AssetsService) GetAreaByIds(ctx context.Context, req *pb.GetAreaByIdsReq) (*pb.GetAreaByIdsReply, error) {
	areas, pageTotal, err := s.arc.GetByIds(ctx, req.Ids, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.GetAreaByIdsReply_Areas, 0)
	for _, x := range areas {
		rs = append(rs, &pb.GetAreaByIdsReply_Areas{
			Id:       x.Id,
			AreaInfo: x.AreaInfo,
		})
	}
	return &pb.GetAreaByIdsReply{
		Areas:     rs,
		PageTotal: pageTotal,
	}, nil
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
	}, nil
}

func (s *AssetsService) UpdateArea(ctx context.Context, req *pb.UpdateAreaReq) (*pb.UpdateAreaReply, error) {
	area, err := s.arc.Update(ctx, &biz.Area{
		Id:       req.Id,
		AreaInfo: req.AreaInfo,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, nil
}

func (s *AssetsService) DeleteArea(ctx context.Context, req *pb.DeleteAreaReq) (*pb.DeleteAreaReply, error) {
	ok, err := s.arc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAreaReply{
		Ok: ok,
	}, nil
}
