package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	log *log.Helper
}

func NewCatalogService(logger log.Logger) *CatalogService {
	return &CatalogService{
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
