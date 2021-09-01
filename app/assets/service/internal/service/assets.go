package service

import (
	"context"

	pb "github.com/Yui-wy/asset-management/api/assets/service/v1"
	"github.com/Yui-wy/asset-management/app/assets/service/internal/biz"
)

func (s *AssetsService) GetAssets(ctx context.Context, req *pb.GetAssetsReq) (*pb.GetAssetsReply, error) {
	a, err := s.ac.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAssetsReply{
		Id:          a.Id,
		Code:        a.Code,
		CodeInfo:    a.CodeInfo,
		AreaId:      a.AreaId,
		Address:     a.Address,
		AssetInfo:   a.AssetInfo,
		PicUrl:      a.PicUrl,
		Price:       a.Price,
		OrderAt:     a.OrderAt,
		OrderNum:    a.OrderNum,
		StorageAt:   a.StorageAt,
		StorageNum:  a.StorageNum,
		ScrappedAt:  a.ScrappedAt,
		ScrappedNum: a.ScrappedNum,
		StateNum:    a.StateNum,
		State:       a.State,
	}, nil
}
func (s *AssetsService) ListAssets(ctx context.Context, req *pb.ListAssetsReq) (*pb.ListAssetsReply, error) {
	as, err := s.ac.List(
		ctx,
		req.PageNum, req.PageSize,
		req.Con.AssetCode,
		req.Con.Address,
		req.Con.StateNum,
		req.Con.LowStorageAt,
		req.Con.UpStorageAt,
	)
	if err != nil {
		return nil, err
	}
	ars := make([]*pb.ListAssetsReply_Assets, 0)
	for _, a := range as {
		ars = append(ars, &pb.ListAssetsReply_Assets{
			Id:          a.Id,
			Code:        a.Code,
			CodeInfo:    a.CodeInfo,
			AreaId:      a.AreaId,
			Address:     a.Address,
			AssetInfo:   a.AssetInfo,
			PicUrl:      a.PicUrl,
			Price:       a.Price,
			OrderAt:     a.OrderAt,
			OrderNum:    a.OrderNum,
			StorageAt:   a.StorageAt,
			StorageNum:  a.StorageNum,
			ScrappedAt:  a.ScrappedAt,
			ScrappedNum: a.ScrappedNum,
			StateNum:    a.StateNum,
			State:       a.State,
		})
	}
	return &pb.ListAssetsReply{
		Results: ars,
	}, nil
}
func (s *AssetsService) CreateAssets(ctx context.Context, req *pb.CreateAssetsReq) (*pb.CreateAssetsReply, error) {
	a, err := s.ac.Create(ctx, &biz.Asset{
		Code:      req.Code,
		CodeInfo:  req.CodeInfo,
		AreaId:    req.AreaId,
		Address:   req.Address,
		AssetInfo: req.AssetInfo,
		PicUrl:    req.PicUrl,
		Price:     req.Price,
		OrderAt:   req.OrderAt,
		OrderNum:  req.OrderNum,
		StateNum:  req.StateNum,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAssetsReply{
		Id:          a.Id,
		Code:        a.Code,
		CodeInfo:    a.CodeInfo,
		AreaId:      a.AreaId,
		Address:     a.Address,
		AssetInfo:   a.AssetInfo,
		PicUrl:      a.PicUrl,
		Price:       a.Price,
		OrderAt:     a.OrderAt,
		OrderNum:    a.OrderNum,
		StorageAt:   a.StorageAt,
		StorageNum:  a.StorageNum,
		ScrappedAt:  a.ScrappedAt,
		ScrappedNum: a.ScrappedNum,
		StateNum:    a.StateNum,
		State:       a.State,
	}, nil
}
func (s *AssetsService) DeleteAssets(ctx context.Context, req *pb.DeleteAssetsReq) (*pb.DeleteAssetsReply, error) {
	ok, err := s.ac.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAssetsReply{
		Ok: ok,
	}, nil
}
func (s *AssetsService) UpdateAssets(ctx context.Context, req *pb.UpdateAssetsdReq) (*pb.UpdateAssetsReply, error) {
	a, err := s.ac.Create(ctx, &biz.Asset{
		Id:          req.Id,
		Address:     req.Address,
		AssetInfo:   req.AssetInfo,
		PicUrl:      req.PicUrl,
		Price:       req.Price,
		OrderAt:     req.OrderAt,
		OrderNum:    req.OrderNum,
		StorageAt:   req.StorageAt,
		StorageNum:  req.StorageNum,
		ScrappedAt:  req.ScrappedAt,
		ScrappedNum: req.ScrappedNum,
		StateNum:    req.StateNum,
		State:       req.State,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAssetsReply{
		Id:          a.Id,
		Code:        a.Code,
		CodeInfo:    a.CodeInfo,
		AreaId:      a.AreaId,
		Address:     a.Address,
		AssetInfo:   a.AssetInfo,
		PicUrl:      a.PicUrl,
		Price:       a.Price,
		OrderAt:     a.OrderAt,
		OrderNum:    a.OrderNum,
		StorageAt:   a.StorageAt,
		StorageNum:  a.StorageNum,
		ScrappedAt:  a.ScrappedAt,
		ScrappedNum: a.ScrappedNum,
		StateNum:    a.StateNum,
		State:       a.State,
	}, nil
}
