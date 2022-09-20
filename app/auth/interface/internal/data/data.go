package data

import (
	authv1 "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/interface/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewAuthServiceClient,
)

type Data struct {
	log *log.Helper
	ac  authv1.AuthClient
}

func NewData(logger log.Logger, ac authv1.AuthClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, ac: ac}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := clientv3.Config{
		Endpoints:   []string{conf.Etcd.EndPoint},
		DialTimeout: time.Second,
	}
	cli, err := clientv3.New(c)
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := clientv3.Config{
		Endpoints:   []string{conf.Etcd.EndPoint},
		DialTimeout: time.Second,
	}
	cli, err := clientv3.New(c)
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}

func NewAuthServiceClient(r registry.Discovery) authv1.AuthClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///janna.auth.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := authv1.NewAuthClient(conn)
	return c
}
