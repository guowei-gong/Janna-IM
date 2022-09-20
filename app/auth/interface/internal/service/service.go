package service

import (
	v1 "Janna-IM/api/auth/interface/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuthService)

type AuthService struct {
	v1.UnimplementedAuthServer

	ac  *biz.AuthUseCase
	log *log.Helper
}

func NewAuthService(ac *biz.AuthUseCase, logger log.Logger) *AuthService {
	return &AuthService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "interface/server-service")),
	}
}
