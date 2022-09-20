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
