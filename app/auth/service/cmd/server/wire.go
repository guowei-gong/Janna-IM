//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Janna-IM/app/auth/service/internal/biz"
	"Janna-IM/app/auth/service/internal/conf"
	"Janna-IM/app/auth/service/internal/data"
	"Janna-IM/app/auth/service/internal/server"
	"Janna-IM/app/auth/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
