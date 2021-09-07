package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	assetv1 "github.com/Yui-wy/asset-management/api/assets/service/v1"
	formv1 "github.com/Yui-wy/asset-management/api/form/service/v1"
	userv1 "github.com/Yui-wy/asset-management/api/user/service/v1"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewDiscovery,
	NewUserServiceClient,
	NewAssetServiceClient,
	NewFormServiceClient,
	NewUserRepo,
	NewAssetRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
	ac  assetv1.AssetsClient
	fc  formv1.FormClient
}

// NewData .
func NewData(
	logger log.Logger,
	uc userv1.UserClient,
	ac assetv1.AssetsClient,
	fc formv1.FormClient,
) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{log: log, uc: uc, ac: ac, fc: fc}
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return d, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///asmg.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(conn)
}

func NewAssetServiceClient(r registry.Discovery) assetv1.AssetsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///asmg.asset.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return assetv1.NewAssetsClient(conn)
}

func NewFormServiceClient(r registry.Discovery) formv1.FormClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///asmg.form.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return formv1.NewFormClient(conn)
}
