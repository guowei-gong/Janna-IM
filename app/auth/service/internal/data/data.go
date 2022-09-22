package data

import (
	"Janna-IM/app/auth/service/internal/conf"
	"Janna-IM/app/auth/service/internal/data/ent"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(NewData, NewEntClient, NewAuthRepo)

type Data struct {
	db *ent.Client
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "auth-server/data/ent"))

	client, err := ent.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data"))

	d := &Data{db: entClient}

	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
