package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	u, err := s.uc.Get(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Uid:     u.Uid,
		Power:   u.Power,
		AreaIds: u.AreaIds,
	}, nil
}
func (s *AssetsService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	if len(req.AreaIds) == 0 {
		return nil, pb.ErrorNoAreaError("Please set correct area")
	}
	u, err := s.uc.Create(ctx, &biz.User{
		Uid:     req.Uid,
		Power:   req.Power,
		AreaIds: req.AreaIds,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Uid:     u.Uid,
		Power:   u.Power,
		AreaIds: u.AreaIds,
	}, nil
}
func (s *AssetsService) UpdateUserArea(ctx context.Context, req *pb.UpdateUserAreaReq) (*pb.UpdateUserAreaReply, error) {
	u, err := s.uc.Update(ctx, &biz.User{
		Uid:     req.Uid,
		AreaIds: req.AreaIds,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserAreaReply{
		Uid:     u.Uid,
		Power:   u.Power,
		AreaIds: u.AreaIds,
	}, nil
}

func (s *AssetsService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	us, err := s.uc.List(ctx, req.NextPower, req.AreaIds)
	if err != nil {
		return nil, err
	}
	users := make([]*pb.ListUserReply_User, 0)
	for _, user := range us {
		users = append(users, &pb.ListUserReply_User{
			Uid:     user.Uid,
			Power:   user.Power,
			AreaIds: user.AreaIds,
		})
	}
	return &pb.ListUserReply{
		Results: users,
	}, nil
}
