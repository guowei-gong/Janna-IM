package data

import (
	"Janna-IM/app/auth/interface/internal/biz"
	"context"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
}

func (a authRepo) RegisterUser(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
