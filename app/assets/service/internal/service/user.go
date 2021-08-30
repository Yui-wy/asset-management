package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
)

func (s *AssetsService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *AssetsService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *AssetsService) UpdateUserArea(ctx context.Context, req *pb.UpdateUserAreaReq) (*pb.UpdateUserAreaReply, error) {
	return &pb.UpdateUserAreaReply{}, nil
}
func (s *AssetsService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

// func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
// 	user, err := s.uc.Create(ctx, &biz.User{
// 		Username:  req.Username,
// 		Password:  req.Password,
// 		IsDeleted: false,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.CreateUserReply{
// 		Id:       user.Id,
// 		Username: user.Username,
// 	}, err
// }

// func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
// 	user, err := s.uc.Get(ctx, req.Id)
// 	return &pb.GetUserReply{
// 		Id:       user.Id,
// 		Username: user.Username,
// 	}, err
// }

// func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
// 	users, err := s.uc.List(ctx, req.Ids, req.PageNum, req.PageSize)
// 	rs := make([]*pb.ListUserReply_User, 0)
// 	for _, x := range users {
// 		rs = append(rs, &pb.ListUserReply_User{
// 			Id:       x.Id,
// 			Username: x.Username,
// 		})
// 	}
// 	return &pb.ListUserReply{
// 		Results: rs,
// 	}, err
// }

// func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserReply, error) {
// 	user, err := s.uc.Deleted(ctx, &biz.User{
// 		Id: req.Id,
// 	})
// 	return &pb.DeleteUserReply{
// 		Id:       user.Id,
// 		Username: user.Username,
// 	}, err
// }

// func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) (*pb.UpdatePasswordReply, error) {
// 	user, err := s.uc.Update(ctx, &biz.User{
// 		Id:       req.Id,
// 		Password: req.Password,
// 	})
// 	return &pb.UpdatePasswordReply{
// 		Id:       user.Id,
// 		Username: user.Username,
// 	}, err
// }

// func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
// 	rv, err := s.uc.VerifyPassword(ctx, &biz.User{
// 		Username: req.Username,
// 		Password: req.Password,
// 	})
// 	return &pb.VerifyPasswordReply{
// 		Ok: rv,
// 	}, err
// }
