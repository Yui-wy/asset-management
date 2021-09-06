package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/management/interface/v1"
)

func (s *ManageMentInterface) ListStorageForm(ctx context.Context, req *pb.ListStorageFormReq) (*pb.ListStorageFormReply, error) {
	return &pb.ListStorageFormReply{}, nil
}
func (s *ManageMentInterface) GetStorageForm(ctx context.Context, req *pb.GetStorageFormReq) (*pb.GetStorageFormReply, error) {
	return &pb.GetStorageFormReply{}, nil
}
func (s *ManageMentInterface) CreateStorageForm(ctx context.Context, req *pb.CreateStorageFormReq) (*pb.CreateStorageFormReply, error) {
	return &pb.CreateStorageFormReply{}, nil
}
func (s *ManageMentInterface) UpdateStorageForm(ctx context.Context, req *pb.UpdateScrappedFormReq) (*pb.UpdateScrappedFormReply, error) {
	return &pb.UpdateScrappedFormReply{}, nil
}
func (s *ManageMentInterface) ListScrappedForm(ctx context.Context, req *pb.ListScrappedFormReq) (*pb.ListScrappedFormReply, error) {
	return &pb.ListScrappedFormReply{}, nil
}
func (s *ManageMentInterface) GetScrappedForm(ctx context.Context, req *pb.GetScrappedFormReq) (*pb.GetScrappedFormReply, error) {
	return &pb.GetScrappedFormReply{}, nil
}
func (s *ManageMentInterface) CreateScrappedForm(ctx context.Context, req *pb.CreateScrappedFormReq) (*pb.CreateScrappedFormReply, error) {
	return &pb.CreateScrappedFormReply{}, nil
}
func (s *ManageMentInterface) UpdateScrappedForm(ctx context.Context, req *pb.UpdateStorageFormReq) (*pb.UpdateStorageFormReply, error) {
	return &pb.UpdateStorageFormReply{}, nil
}
