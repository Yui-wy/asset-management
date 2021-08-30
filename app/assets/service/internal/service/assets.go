package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
)

func (s *AssetsService) GetAssets(ctx context.Context, req *pb.GetAssetsReq) (*pb.GetAssetsReply, error) {
	return &pb.GetAssetsReply{}, nil
}
func (s *AssetsService) ListAssets(ctx context.Context, req *pb.ListAssetsReq) (*pb.ListAssetsReply, error) {
	return &pb.ListAssetsReply{}, nil
}
func (s *AssetsService) CreateAssets(ctx context.Context, req *pb.CreateAssetsReq) (*pb.CreateAssetsReply, error) {
	return &pb.CreateAssetsReply{}, nil
}
func (s *AssetsService) DeleteAssets(ctx context.Context, req *pb.DeleteAssetsReq) (*pb.DeleteAssetsReply, error) {
	return &pb.DeleteAssetsReply{}, nil
}
func (s *AssetsService) UpdateAssets(ctx context.Context, req *pb.UpdateAssetsdReq) (*pb.UpdateAssetsReply, error) {
	return &pb.UpdateAssetsReply{}, nil
}

