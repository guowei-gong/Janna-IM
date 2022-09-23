package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
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

type UserToken struct {
	Platform    int32
	FromUserID  string
	OpUserID    string
	OperationID string
	LoginIp     string
}

type AuthRepo interface {
	CreateUser(ctx context.Context, u *User) error
	GetUserByUserId(ctx context.Context, id string) (*User, error)
	UserExisted(ctx context.Context, id string) (bool, error)
}

type TokenRepo interface {
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

func (ac *AuthUseCase) UserToken(ctx context.Context, u *UserToken) (string, int64, error) {
	existed, err := ac.repo.UserExisted(ctx, u.FromUserID)
	if err != nil {
		errMsg := u.OperationID + " ac.rep.UserExisted failed " + err.Error() + u.FromUserID
		log.Error(u.OperationID, errMsg)
		return "", 0, errors.InternalServer("auth", "internal error")
	}

	if !existed {
		return "", 0, errors.NotFound("auth", "not found user by from_user_id")
	}

	return "", 0, err
}
