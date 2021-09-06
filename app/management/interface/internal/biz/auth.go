package biz

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Yui-wy/asset-management/app/management/interface/internal/conf"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	key  string
	repo UserRepo
}

func NewAuthUseCase(conf *conf.Auth, repo UserRepo) *AuthUseCase {
	return &AuthUseCase{
		key:  conf.Key,
		repo: repo,
	}
}

func (r AuthUseCase) Auth(userId uint64, areaId []uint32, power int32, sign string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"area_id": areaId,
		"power":   power,
		"sign":    sign,
		"exp":     time.Now().Unix() + 864000,
	})
	return claims.SignedString(r.key)
}

// TODO: 搜索用户验证密码
func (r AuthUseCase) CheckJWT(ctx context.Context, jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(jwtToken *jwt.Token) (interface{}, error) {
		return r.key, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token type error")
	}
	if claims["exp"].(int64) < time.Now().Unix() {
		return nil, errors.New("token overtime.")
	}
	u, err := r.repo.GetUser(ctx, claims["user_id"].(uint64))
	if strings.Compare(u.UpdataSign, claims["sign"].(string)) == -1 {
		return nil, errors.New("token overtime.")
	}
	result := make(map[string]interface{}, 3)
	result["user_id"] = claims["user_id"]
	result["area_id"] = claims["area_id"]
	result["power"] = claims["power"]
	return result, nil
}
