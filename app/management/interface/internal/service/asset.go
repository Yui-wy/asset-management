package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
)

func (s *ManageMentInterface) ListAsset(ctx context.Context, req *pb.ListAssetReq) (*pb.ListAssetReply, error) {
	return &pb.ListAssetReply{}, nil
}
func (s *ManageMentInterface) CreateAsset(ctx context.Context, req *pb.CreateAssetReq) (*pb.CreateAssetReply, error) {
	return &pb.CreateAssetReply{}, nil
}
func (s *ManageMentInterface) GetAsset(ctx context.Context, req *pb.GetAssetReq) (*pb.GetAssetReply, error) {
	return &pb.GetAssetReply{}, nil
}
func (s *ManageMentInterface) UpdateAsset(ctx context.Context, req *pb.UpdateAssetReq) (*pb.UpdateAssetReply, error) {
	return &pb.UpdateAssetReply{}, nil
}