package server

import (
	"Janna-IM/app/auth/service/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegister)

func NewRegister(conf *conf.Registry) registry.Registrar {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}
