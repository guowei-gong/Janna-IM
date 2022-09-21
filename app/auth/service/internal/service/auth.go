package service

import (
	"context"
	"fmt"

	pb "Janna-IM/api/auth/service/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	fmt.Println("Hello, world!")
	return &pb.UserRegisterResp{}, nil
}
func (s *AuthService) UserToken(ctx context.Context, req *pb.UserTokenReq) (*pb.UserTokenResp, error) {
	return &pb.UserTokenResp{}, nil
}
func (s *AuthService) ForceLogout(ctx context.Context, req *pb.ForceLogoutReq) (*pb.ForceLogoutResp, error) {
	return &pb.ForceLogoutResp{}, nil
}
