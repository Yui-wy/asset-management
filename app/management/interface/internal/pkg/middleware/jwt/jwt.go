package jwt

import (
	"context"

	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/middleware"
)

func NewAuthMiddleware(authUc *biz.AuthUseCase) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ctx, err = authUc.CheckJWT(ctx)
			if err != nil {
				return nil, err
			}
			reply, err = handler(ctx, req)
			return
		}
	}
}
