package biz

import (
	pbAuth "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/interface/internal/conf"
	ws "Janna-IM/third_party/ws/v1"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
)

type ApiUserInfo struct {
	UserID         string `json:"userID"`
	Nickname       string `json:"nickname"`
	FaceURL        string `json:"faceURL"`
	Gender         int32  `json:"gender"`
	PhoneNumber    string `json:"phoneNumber"`
	Birth          uint32 `json:"birth"`
	Email          string `json:"email"`
	CreateIp       string `json:"createIp"`
	CreateTime     int64  `json:"createTime"`
	LastLoginIp    string `json:"LastLoginIp"`
	LastLoginTime  int64  `json:"lastLoginTime"`
	LoginTimes     int32  `json:"loginTimes"`
	LoginLimit     int32  `json:"loginLimit"`
	Ex             string `json:"ex"`
	InvitationCode string `json:"invitationCode"`
}

type UserTokenReq struct {
	Secret      string
	Platform    int32
	UserID      string
	LoginIp     string
	OperationID string
}

type UserRegister struct {
	ApiUserInfo

	Secret      string `json:"secret"`
	Platform    int32  `json:"platform"`
	OperationID string `json:"operationID"`
}

type AuthRepo interface {
	UserRegister(ctx context.Context, u *pbAuth.UserRegisterReq) error
	UserToken(ctx context.Context, u *pbAuth.UserTokenReq) (*pbAuth.UserTokenResp, error)
}

type AuthUseCase struct {
	Auth     *conf.Auth
	authRepo AuthRepo
}

func NewAuthUseCase(auth *conf.Auth, authRepo AuthRepo) *AuthUseCase {
	return &AuthUseCase{Auth: auth, authRepo: authRepo}
}

func (s *AuthUseCase) Token(ctx context.Context, params *UserTokenReq) (*pbAuth.UserTokenResp, error) {
	// 前置校验
	if params.Secret != s.Auth.Secret {
		return nil, errors.Unauthorized("auth", "secret error")
	}

	req := &pbAuth.UserTokenReq{
		Platform:    params.Platform,
		FromUserID:  params.UserID,
		OperationID: params.OperationID,
		LoginIp:     params.LoginIp,
	}

	_, err := s.authRepo.UserToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *AuthUseCase) Register(ctx context.Context, params *UserRegister) error {
	// 前置校验
	if params.Secret != s.Auth.Secret {
		return errors.Unauthorized("auth", "secret error")
	}

	req := &pbAuth.UserRegisterReq{
		UserInfo: &ws.UserInfo{
			UserID:        params.UserID,
			Nickname:      params.Nickname,
			Gender:        params.Gender,
			PhoneNumber:   params.PhoneNumber,
			Birth:         params.Birth,
			Email:         params.Email,
			CreateIp:      params.CreateIp,
			CreateTime:    uint32(params.CreateTime),
			LastLoginIp:   params.LastLoginIp,
			LastLoginTime: uint32(params.LastLoginTime),
			LoginTimes:    params.LoginTimes,
		},
		OperationID: params.OperationID,
	}

	// 远程调用
	err := s.authRepo.UserRegister(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
