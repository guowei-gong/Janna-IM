package service

import (
	v1 "Janna-IM/api/auth/interface/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"context"
)

func (s *AuthService) UserRegister(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	s.ac.Register(ctx, &biz.UserRegister{
		ApiUserInfo: biz.ApiUserInfo{},
		Secret:      ctx.Value("secret").(string),
		Platform:    ctx.Value("platform").(int32),
		OperationID: ctx.Value("operation_id").(string),
	})
	return &v1.UserReply{}, nil
}
