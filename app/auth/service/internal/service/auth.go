package service

import (
	pb "Janna-IM/api/auth/service/v1"
	"Janna-IM/app/auth/service/internal/biz"
	"Janna-IM/pkg/constant"
	"Janna-IM/pkg/utils"
	"context"
)

func (s *AuthService) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	s.log.Info(req.OperationID, "rpc args", req.String())

	user := loadUserRegisterReq(req)
	if req.UserInfo.Birth != 0 {
		user.Birth = utils.UnixSecondToTime(int64(req.UserInfo.Birth))
	}
	if err := s.ac.Create(ctx, user); err != nil {
		errMsg := req.OperationID + " imdb.UserRegister failed " + err.Error() + user.UserID
		return &pb.UserRegisterResp{CommonResp: &pb.CommonResp{ErrCode: constant.ErrDB.ErrCode, ErrMsg: errMsg}}, nil
	}
	return &pb.UserRegisterResp{CommonResp: &pb.CommonResp{}}, nil
}

func (s *AuthService) UserToken(ctx context.Context, req *pb.UserTokenReq) (*pb.UserTokenResp, error) {
	return &pb.UserTokenResp{}, nil
}
func (s *AuthService) ForceLogout(ctx context.Context, req *pb.ForceLogoutReq) (*pb.ForceLogoutResp, error) {
	return &pb.ForceLogoutResp{}, nil
}

func loadUserRegisterReq(input *pb.UserRegisterReq) *biz.User {
	var ret biz.User
	userInfo := input.UserInfo

	ret.UserID = userInfo.UserID
	ret.Nickname = userInfo.Nickname
	ret.FaceURL = userInfo.FaceURL
	ret.PhoneNumber = userInfo.PhoneNumber
	ret.Email = userInfo.Email
	ret.Ex = userInfo.Ex
	ret.CreateIp = userInfo.CreateIp
	ret.LastLoginIp = userInfo.LastLoginIp
	ret.InvitationCode = userInfo.InvitationCode
	ret.Gender = userInfo.Gender
	ret.LoginTimes = userInfo.LoginTimes
	ret.LoginLimit = userInfo.LoginLimit
	ret.AppMangerLevel = userInfo.AppMangerLevel
	ret.GlobalRecvMsgOpt = userInfo.GlobalRecvMsgOpt

	return &ret
}
