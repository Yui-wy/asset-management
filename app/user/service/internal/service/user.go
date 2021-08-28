package service

import (
	"context"

	pb "github.com/Yui-wy/material/api/user/service/v1"
	"github.com/Yui-wy/material/app/user/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer

	log *log.Helper
	uc  *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper(log.With(logger, "module", "service/user")),
		uc:  uc,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		IsDelete: false,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, err
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	return &pb.GetUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, err
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	users, err := s.uc.List(ctx, req.Ids)
	rs := make([]*pb.ListUserReply_User, 0)
	for _, x := range users {
		rs = append(rs, &pb.ListUserReply_User{
			Id:       x.Id,
			Username: x.Username,
		})
	}
	return &pb.ListUserReply{
		Results: rs,
	}, err
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserReply, error) {
	user, err := s.uc.Update(ctx, &biz.User{
		Id:       req.Id,
		IsDelete: true,
	})
	return &pb.DeleteUserReply{
		Id:       user.Id,
		Username: user.Username,
	}, err
}

func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) (*pb.UpdatePasswordReply, error) {
	user, err := s.uc.Update(ctx, &biz.User{
		Id:       req.Id,
		Password: req.Password,
	})
	return &pb.UpdatePasswordReply{
		Id:       user.Id,
		Username: user.Username,
	}, err
}

func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	return &pb.VerifyPasswordReply{
		Ok: rv,
	}, err
}
