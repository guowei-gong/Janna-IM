package biz

import (
	pbAuth "Janna-IM/api/auth/service/v1"
	ws "Janna-IM/api/ws/v1"
	"Janna-IM/app/auth/interface/internal/conf"
	"context"
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

type UserRegister struct {
	ApiUserInfo

	Secret      string `json:"secret"`
	Platform    int32  `json:"platform"`
	OperationID string `json:"operationID"`
}

type AuthRepo interface {
	RegisterUser(ctx context.Context)
}

type AuthUseCase struct {
	bootstrap *conf.Bootstrap
}

func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

func (s *AuthUseCase) Register(ctx context.Context, params *UserRegister) {
	// TODO 前置校验
	if params.Secret != s.bootstrap.Secret {

	}

	// TODO 请求参数
	req := &pbAuth.UserRegisterReq{
		UserInfo: &ws.UserInfo{},
	}

	req.OperationID = params.OperationID
	// TODO 远程调用
}
