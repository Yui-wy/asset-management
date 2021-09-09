// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/conf"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/data"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/server"
	"github.com/Yui-wy/asset-management/app/management/interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, client *conf.Client, registry *conf.Registry, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(discovery, client)
	assetsClient := data.NewAssetServiceClient(discovery, client)
	formClient := data.NewFormServiceClient(discovery, client)
	dataData, cleanup, err := data.NewData(logger, userClient, assetsClient, formClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	authUseCase := biz.NewAuthUseCase(auth, userRepo)
	assetRepo := data.NewAssetRepo(dataData, logger)
	assetUseCase := biz.NewAssetUseCase(assetRepo, logger)
	manageMentInterface := service.NewManagementInterface(logger, userUseCase, authUseCase, assetUseCase)
	httpServer := server.NewHTTPServer(confServer, logger, manageMentInterface)
	grpcServer := server.NewGRPCServer(confServer, manageMentInterface, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
