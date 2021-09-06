package jwt

import (
	"context"
	"net/http"

	"github.com/Yui-wy/asset-management/app/management/interface/internal/biz"
	"github.com/Yui-wy/asset-management/pkg/errors/auth"
	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc/metadata"
)

func NewAuthMiddleware(authUc *biz.AuthUseCase) func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// if tr, ok := transport.FromServerContext(ctx); ok {
			// 	kind = tr.Kind().String()
			// 	operation = tr.Operation()
			// 	// 断言成HTTP的Transport可以拿到特殊信息
			// 	if ht, ok := tr.(*http.Transport); ok {
			// 	}
			// }
			var jwtToken string
			if request, ok := req.(http.Request); ok {
				jwtToken = request.Header.Get("Token")
			} else if md, ok := metadata.FromIncomingContext(ctx); ok {
				jwtToken = md.Get("Token")[0]
			} else {
				// 缺少可认证的token，返回错误
				return nil, auth.ErrAuthFail
			}
			token, err := authUc.CheckJWT(ctx, jwtToken)
			if err != nil {
				// 缺少合法的token，返回错误
				return nil, auth.ErrAuthFail
			}
			ctx = context.WithValue(ctx, "x-md-global-user", token)
			reply, err = handler(ctx, req)
			return
		}
	}
}
