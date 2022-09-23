package data

import (
	pbAuth "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "data/server-interface")),
	}
}

func (ap *authRepo) UserToken(ctx context.Context, u *pbAuth.UserTokenReq) (*pbAuth.UserTokenResp, error) {
	_, err := ap.data.ac.UserToken(ctx, u)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (ap *authRepo) UserRegister(ctx context.Context, u *pbAuth.UserRegisterReq) error {
	_, err := ap.data.ac.UserRegister(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
