package data

import (
	"Janna-IM/app/auth/service/internal/biz"
	"Janna-IM/pkg/constant"
	"Janna-IM/pkg/utils"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (a *authRepo) CreateUser(ctx context.Context, u *biz.User) error {
	if u.AppMangerLevel == 0 {
		u.AppMangerLevel = constant.AppOrdinaryUsers
	}
	if u.Birth.Unix() < 0 {
		u.Birth = utils.UnixSecondToTime(0)
	}
	now := time.Now()
	u.CreateTime = now
	u.LastLoginTime = now
	u.LoginTimes = 0
	u.LastLoginIp = u.CreateIp

	_, err := a.data.db.User.
		Create().
		SetUserID(u.UserID).
		SetName(u.Nickname).
		SetFaceURL(u.FaceURL).
		SetPhoneNumber(u.PhoneNumber).
		SetEmail(u.Email).
		SetEx(u.Ex).
		SetCreateIP(u.CreateIp).
		SetLastLoginIP(u.LastLoginIp).
		SetInvitationCode(u.InvitationCode).
		SetGender(u.Gender).
		SetLoginLimit(u.LoginLimit).
		SetLoginTimes(u.LoginTimes).
		SetAppMangerLevel(u.AppMangerLevel).
		SetGlobalRecvMsgOpt(u.GlobalRecvMsgOpt).
		SetBirth(u.Birth).
		SetCreateTime(u.CreateTime).
		SetLastLoginTime(u.LastLoginTime).
		SetStatus(constant.Normal).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}
