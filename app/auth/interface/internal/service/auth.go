package service

import (
	v1 "Janna-IM/api/auth/interface/v1"
	"context"
)

func (s *AuthService) UserRegister(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {

	return &v1.UserReply{}, nil
}
