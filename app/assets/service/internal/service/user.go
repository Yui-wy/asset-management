package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	u, err := s.uc.Get(ctx, req.Uid)
	areas := make([]*pb.GetUserReply_Areas, 0)
	for _, area := range u.Areas {
		areas = append(areas, &pb.GetUserReply_Areas{
			AreaId:   area.Id,
			AreaInfo: area.AreaInfo,
		})
	}
	return &pb.GetUserReply{
		Uid:   u.Uid,
		Power: u.Power,
		Areas: areas,
	}, err
}
func (s *AssetsService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	a := make([]*biz.Area, 0)
	for _, area := range req.Areas {
		a = append(a, &biz.Area{
			Id: area.AreaId,
		})
	}
	u, err := s.uc.Create(ctx, &biz.User{
		Uid:   req.Uid,
		Power: req.Power,
		Areas: a,
	})
	areas := make([]*pb.CreateUserReply_Areas, 0)
	for _, area := range u.Areas {
		areas = append(areas, &pb.CreateUserReply_Areas{
			AreaId:   area.Id,
			AreaInfo: area.AreaInfo,
		})
	}
	return &pb.CreateUserReply{
		Uid:   u.Uid,
		Power: u.Power,
		Areas: areas,
	}, err
}
func (s *AssetsService) UpdateUserArea(ctx context.Context, req *pb.UpdateUserAreaReq) (*pb.UpdateUserAreaReply, error) {
	a := make([]*biz.Area, 0)
	for _, area := range req.Areas {
		a = append(a, &biz.Area{
			Id: area.AreaId,
		})
	}
	u, err := s.uc.Update(ctx, &biz.User{
		Uid:   req.Uid,
		Areas: a,
	})
	areas := make([]*pb.UpdateUserAreaReply_Areas, 0)
	for _, area := range u.Areas {
		areas = append(areas, &pb.UpdateUserAreaReply_Areas{
			AreaId:   area.Id,
			AreaInfo: area.AreaInfo,
		})
	}
	return &pb.UpdateUserAreaReply{
		Uid:   u.Uid,
		Power: u.Power,
		Areas: areas,
	}, err
}

func (s *AssetsService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	us, err := s.uc.List(ctx, req.AreaIds)
	users := make([]*pb.ListUserReply_User, 0)
	for _, user := range us {
		areas := make([]*pb.ListUserReply_User_Areas, 0)
		for _, area := range user.Areas {
			areas = append(areas, &pb.ListUserReply_User_Areas{
				AreaId:   area.Id,
				AreaInfo: area.AreaInfo,
			})
		}
		users = append(users, &pb.ListUserReply_User{
			Uid:   user.Uid,
			Power: user.Power,
			Areas: areas,
		})
	}
	return &pb.ListUserReply{
		Results: users,
	}, err
}
