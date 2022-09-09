package service

import (
	v1 "Janna-IM/api/auth/interface/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuthService)

func NewAuthService() *AuthService {
	return &AuthService{}
}

type AuthService struct {
	v1.UnimplementedAuthServer

	ac *biz.AuthUseCase
}
