package service

import (
	v1 "Janna-IM/api/auth/interface/v1"
	"Janna-IM/app/auth/interface/internal/biz"
	"context"
)

func (s *AuthService) UserRegister(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	userInfoReq := req.UserInfo
	userInfo := biz.ApiUserInfo{
		UserID:      userInfoReq.UserId,
		Nickname:    userInfoReq.NickName,
		Gender:      userInfoReq.Gender,
		PhoneNumber: userInfoReq.PhoneNumber,
		Birth:       userInfoReq.Birth,
		Email:       userInfoReq.Email,
		CreateIp:    userInfoReq.CreateIp,
		CreateTime:  userInfoReq.CreateTime,
		LastLoginIp: userInfoReq.LastLoginIp,
	}

	if err := s.ac.Register(ctx, &biz.UserRegister{
		ApiUserInfo: userInfo,
		Secret:      req.UserRegister.Secret,
		Platform:    req.UserRegister.Platform,
		OperationID: req.UserRegister.OperationId,
	}); err != nil {
		return nil, err
	}
	return &v1.UserReply{}, nil
}

func (s *AuthService) UserToken(ctx context.Context, req *v1.UserTokenReq) (*v1.UserTokenResp, error) {
	params := biz.UserTokenReq{
		Secret:      req.Secret,
		Platform:    req.Platform,
		UserID:      req.UserId,
		LoginIp:     req.LoginId,
		OperationID: req.OperationId,
	}

	_, err := s.ac.Token(ctx, &params)
	if err != nil {
		return nil, err
	}
	return &v1.UserTokenResp{}, nil
}
