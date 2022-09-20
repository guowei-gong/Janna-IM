package data

import (
	pbAuth "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"context"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
}

func (a authRepo) RegisterUser(ctx context.Context, u *pbAuth.UserRegisterReq) error {
	//TODO implement me
	panic("implement me")
}
