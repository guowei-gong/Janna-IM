// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Janna-IM/app/auth/service/internal/biz"
	"Janna-IM/app/auth/service/internal/conf"
	"Janna-IM/app/auth/service/internal/data"
	"Janna-IM/app/auth/service/internal/server"
	"Janna-IM/app/auth/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, registry *conf.Registry) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	authRepo := data.NewAuthRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(authRepo, logger)
	authService := service.NewAuthService(authUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, authService, logger)
	registrar := server.NewRegister(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
