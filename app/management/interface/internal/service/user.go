package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/setting"
)

func (s *ManageMentInterface) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	user, err := s.uc.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := s.AuthUc.Auth(user.Id, user.AreaIds, user.Power, user.UpdataSign)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: token,
	}, nil
}
func (s *ManageMentInterface) GetKey(ctx context.Context, req *pb.GetKeyReq) (*pb.GetKeyReply, error) {
	key, err := s.uc.GetKey(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetKeyReply{
		Key: key,
	}, nil
}

func (s *ManageMentInterface) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	ok, err := s.uc.Logout(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutReply{
		Ok: ok,
	}, nil
}

func (s *ManageMentInterface) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, req.AreaId)
	if err != nil {
		return nil, err
	}
	_, err = s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Power:    setting.AREA_USER,
		AreaIds:  req.AreaId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Ok: true,
	}, nil
}
func (s *ManageMentInterface) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:         user.Id,
		Username:   user.Username,
		AssetPower: user.Power,
		AreaIds:    user.AreaIds,
	}, nil
}
func (s *ManageMentInterface) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, req.AreaIds)
	if err != nil {
		return nil, err
	}
	users, err := s.uc.ListUser(ctx, req.PageNum, req.PageSize, req.AreaIds, setting.AREA_USER)
	if err != nil {
		return nil, err
	}
	r := make([]*pb.ListUserReply_Users, 0)
	for _, user := range users {
		r = append(r, &pb.ListUserReply_Users{
			Id:         user.Id,
			Username:   user.Username,
			AssetPower: user.Power,
			AreaIds:    user.AreaIds,
		})
	}
	return &pb.ListUserReply{
		Users: r,
	}, nil
}
func (s *ManageMentInterface) ModifyUserPd(ctx context.Context, req *pb.ModifyUserPdReq) (*pb.ModifyUserPdReply, error) {
	_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, req.AreaId)
	if err != nil {
		return nil, err
	}
	ok, err := s.uc.ModifyPd(ctx, req.Id, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.ModifyUserPdReply{Ok: ok}, nil
}
func (s *ManageMentInterface) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserReply, error) {
	_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, req.AreaId)
	if err != nil {
		return nil, err
	}
	ok, err := s.uc.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{Ok: ok}, nil
}

func (s *ManageMentInterface) ListArea(ctx context.Context, req *pb.ListAreaReq) (*pb.ListAreaReply, error) {
	_, err := s.checkPower(ctx, setting.AREA_ADMIN_USER, req.Ids)
	if err != nil {
		return nil, err
	}
	areas, err := s.uc.ListArea(ctx, req.Ids, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	r := make([]*pb.ListAreaReply_Areas, 0)
	for _, area := range areas {
		r = append(r, &pb.ListAreaReply_Areas{
			Id:       area.Id,
			AreaInfo: area.AreaInfo,
		})
	}
	return &pb.ListAreaReply{Areas: r}, nil
}

func (s *ManageMentInterface) GetArea(ctx context.Context, req *pb.GetAreaReq) (*pb.GetAreaReply, error) {
	area, err := s.uc.GetArea(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAreaReply{
		Id:       area.Id,
		AreaInfo: area.AreaInfo,
	}, nil
}
