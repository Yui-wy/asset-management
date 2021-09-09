package jwt

import (
	"context"

	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/Yui-wy/asset-management/pkg/util/inspection"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func NewAuthMiddleware(authUc *biz.AuthUseCase) func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, auth.ErrAuthFail
			}
			// kind := tr.Kind().String()
			head := tr.RequestHeader()
			jwtToken := head.Get("Token")
			// ===============================================
			result := make(map[string]interface{}, 4)
			if !inspection.IsZeros(jwtToken) {
				result, err = authUc.CheckJWT(ctx, jwtToken)
				if err != nil {
					return nil, err
				}
			}
			ctx = context.WithValue(ctx, "x-md-global-user", result)
			reply, err = handler(ctx, req)
			return
		}
	}
}
