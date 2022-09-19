package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	log *log.Helper
}

func NewData(logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l}, nil
}
