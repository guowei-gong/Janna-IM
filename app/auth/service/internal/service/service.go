package service

import (
	v1 "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService)

type AuthService struct {
	v1.UnimplementedAuthServer

	ac  *biz.AuthUseCase
	log *log.Helper
}

func NewAuthService(ac *biz.AuthUseCase, logger log.Logger) *AuthService {
	return &AuthService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
