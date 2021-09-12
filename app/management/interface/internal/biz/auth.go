package biz

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Yui-wy/asset-management/app/management/interface/internal/conf"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

type authKey struct{}

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request header.
	authorizationKey string = "Authorization"
)

type AuthUseCase struct {
	key  string
	repo UserRepo
}

type AuthUser struct {
	Uid      uint64
	Username string
	Power    int32
	AreaIds  []uint32
}

type AuthClaims struct {
	UserId uint64 `json:"uid"`
	Sign   string `json:"sign"`
	jwt.StandardClaims
}

func NewAuthUseCase(conf *conf.Auth, repo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		key:  conf.Key,
		repo: repo,
	}
}

func (r AuthUseCase) Auth(userId uint64, sign string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		UserId: userId,
		Sign:   sign,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 864000,
		},
	})
	tokenStr, err := token.SignedString([]byte(r.key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bearerFormat, tokenStr), err
}

// TODO: 搜索用户验证密码
func (r AuthUseCase) CheckJWT(ctx context.Context) (context.Context, error) {
	header, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, auth.ErrWrongContext
	}
	auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
		return nil, auth.ErrMissingJwtToken
	}
	jwtToken := auths[1]
	tokenInfo, err := jwt.Parse(jwtToken, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(r.key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, auth.ErrTokenInvalid
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, auth.ErrTokenExpired
			} else {
				return nil, auth.ErrTokenParseFail
			}
		}
	} else if !tokenInfo.Valid {
		return nil, auth.ErrTokenInvalid
	} else if tokenInfo.Method != jwt.SigningMethodHS256 {
		return nil, auth.ErrUnSupportSigningMethod
	}
	claims, ok := tokenInfo.Claims.(AuthClaims)
	u, err := r.repo.GetUser(ctx, claims.UserId)
	if strings.Compare(u.UpdataSign, claims.Sign) == -1 {
		return nil, auth.ErrTokenExpired
	}
	info := &AuthUser{
		Uid:      u.Id,
		Username: u.Username,
		Power:    u.Power,
		AreaIds:  u.AreaIds,
	}
	return r.NewContext(ctx, info), nil
}

// NewContext put auth info into context
func (r AuthUseCase) NewContext(ctx context.Context, info *AuthUser) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

// FromContext extract auth info from context
func (r AuthUseCase) FromContext(ctx context.Context) (*AuthUser, bool) {
	authUser, ok := ctx.Value(authKey{}).(*AuthUser)
	return authUser, ok
}
