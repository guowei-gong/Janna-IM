package data

import (
	"Janna-IM/app/auth/service/internal/biz"
	"Janna-IM/pkg/constant"
	"Janna-IM/pkg/utils"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var _ biz.TokenRepo = (*tokenRepo)(nil)

type tokenRepo struct {
	data *Data
	log  *log.Helper
}

func NewTokenRepo(data *Data, logger log.Logger) biz.TokenRepo {
	return &tokenRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

type Claims struct {
	UID      string
	Platform string //login platform
	jwt.RegisteredClaims
}

func (t *tokenRepo) BuildClaims(uid, platform string, ttl int64) Claims {
	now := time.Now()
	before := now.Add(-time.Minute * 5)
	return Claims{
		UID:      uid,
		Platform: platform,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl*24) * time.Hour)), //Expiration time
			IssuedAt:  jwt.NewNumericDate(now),                                        //Issuing time
			NotBefore: jwt.NewNumericDate(before),                                     //Begin Effective time
		}}
}

func (t *tokenRepo) CreateToken(secret, userID string, platformID int, accessExpire int64) (string, int64, error) {
	claims := BuildClaims(userID, constant.PlatformIDToName(platformID), accessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	//remove Invalid token
	m, err := t.GetTokenMapByUidPid(userID, constant.PlatformIDToName(platformID))
	if err != nil && err != go_redis.Nil {
		return "", 0, err
	}
	var deleteTokenKey []string
	for k, v := range m {
		_, err = GetClaimFromToken(k)
		if err != nil || v != constant.NormalToken {
			deleteTokenKey = append(deleteTokenKey, k)
		}
	}
	if len(deleteTokenKey) != 0 {
		err = t.DeleteTokenByUidPid(userID, platformID, deleteTokenKey)
		if err != nil {
			return "", 0, err
		}
	}
	err = t.AddTokenFlag(userID, platformID, tokenString, constant.NormalToken)
	if err != nil {
		return "", 0, err
	}
	return tokenString, claims.ExpiresAt.Time.Unix(), err
}

func (t *tokenRepo) GetTokenMapByUidPidFromCache(userID, platformID string) (map[string]int, error) {
	key := constant.UidPidToken + userID + ":" + platformID
	m, err := t.data.redisCli.HGetAll(context.Background(), key).Result()
	mm := make(map[string]int)
	for k, v := range m {
		mm[k] = utils.StringToInt(v)
	}
	return mm, err
}

func (t *tokenRepo) GetClaimFromToken(tokensString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, utils.Wrap(constant.ErrTokenMalformed, "")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, utils.Wrap(constant.ErrTokenExpired, "")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
			} else {
				return nil, utils.Wrap(constant.ErrTokenUnknown, "")
			}
		} else {
			return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
		}
	} else {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, utils.Wrap(constant.ErrTokenNotValidYet, "")
	}
}
