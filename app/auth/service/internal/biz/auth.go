package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	UserID           string
	Nickname         string
	FaceURL          string
	PhoneNumber      string
	Email            string
	Ex               string
	CreateIp         string
	LastLoginIp      string
	InvitationCode   string
	status           int32
	Gender           int32
	LoginTimes       int32
	LoginLimit       int32
	AppMangerLevel   int32
	GlobalRecvMsgOpt int32
	Birth            time.Time
	CreateTime       time.Time
	LastLoginTime    time.Time
}

type AuthRepo interface {
	CreateUser(ctx context.Context, u *User) error
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/auth"))}
}

func (ac *AuthUseCase) Create(ctx context.Context, u *User) error {
	//TODO 是否限制 IP

	err := ac.repo.CreateUser(ctx, u)
	return err
}
