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
		Id:         a.Id,
		Classes:    a.Classes,
		Code:       a.Code,
		AreaId:     a.AreaId,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}, nil
}
func (s *AssetsService) ListAssets(ctx context.Context, req *pb.ListAssetsReq) (*pb.ListAssetsReply, error) {
	conf := &biz.SearchConf{
		Classes:      req.Conf.Classes,
		Address:      req.Conf.Address,
		StateNum:     req.Conf.StateNum,
		LowStorageAt: req.Conf.LowStorageAt,
		UpStorageAt:  req.Conf.UpStorageAt,
		OrderBy:      req.Conf.OrderBy,
		SortDesc:     req.Conf.SortDesc,
		AreaId:       req.Conf.AreaId,
	}
	as, err := s.ac.List(ctx, conf, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	ars := make([]*pb.ListAssetsReply_Assets, 0)
	for _, a := range as {
		ars = append(ars, &pb.ListAssetsReply_Assets{
			Id:         a.Id,
			Classes:    a.Classes,
			Code:       a.Code,
			AreaId:     a.AreaId,
			Address:    a.Address,
			AssetInfo:  a.AssetInfo,
			PicUrl:     a.PicUrl,
			Price:      a.Price,
			OrderAt:    a.OrderAt,
			OrderNum:   a.OrderNum,
			StateNum:   a.StateNum,
			State:      a.State,
			AppliedAt:  a.AppliedAt,
			StorageAt:  a.StorageAt,
			ScrappedAt: a.ScrappedAt,
		})
	}
	return &pb.ListAssetsReply{
		Results: ars,
	}, nil
}
func (s *AssetsService) CreateAssets(ctx context.Context, req *pb.CreateAssetsReq) (*pb.CreateAssetsReply, error) {
	a, err := s.ac.Create(ctx, &biz.Asset{
		Classes:    req.Classes,
		AreaId:     req.AreaId,
		Address:    req.Address,
		AssetInfo:  req.AssetInfo,
		PicUrl:     req.PicUrl,
		Price:      req.Price,
		OrderAt:    req.OrderAt,
		OrderNum:   req.OrderNum,
		StateNum:   req.StateNum,
		AppliedAt:  req.AppliedAt,
		StorageAt:  req.StorageAt,
		ScrappedAt: req.ScrappedAt,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAssetsReply{
		Id:         a.Id,
		Classes:    a.Classes,
		Code:       a.Code,
		AreaId:     a.AreaId,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
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
	a, err := s.ac.Update(ctx, &biz.Asset{
		Id:         req.Id,
		Address:    req.Address,
		AssetInfo:  req.AssetInfo,
		PicUrl:     req.PicUrl,
		Price:      req.Price,
		OrderAt:    req.OrderAt,
		OrderNum:   req.OrderNum,
		StateNum:   req.StateNum,
		AppliedAt:  req.AppliedAt,
		StorageAt:  req.StorageAt,
		ScrappedAt: req.ScrappedAt,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAssetsReply{
		Id:         a.Id,
		Classes:    a.Classes,
		Code:       a.Code,
		AreaId:     a.AreaId,
		Address:    a.Address,
		AssetInfo:  a.AssetInfo,
		PicUrl:     a.PicUrl,
		Price:      a.Price,
		OrderAt:    a.OrderAt,
		OrderNum:   a.OrderNum,
		StateNum:   a.StateNum,
		State:      a.State,
		AppliedAt:  a.AppliedAt,
		StorageAt:  a.StorageAt,
		ScrappedAt: a.ScrappedAt,
	}, nil
}
