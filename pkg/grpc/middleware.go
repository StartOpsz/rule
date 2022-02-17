package grpc

import (
	"context"
	"errors"
	
	"github.com/go-kratos/kratos/v2/middleware"
)

// 如果在 grpc metadata 中未发现 userUUID 则认为未通过鉴权
func KratosAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			
			reply, err = handler(ctx, req)
			
			userUUID, _ := GetStringFromGrpc(ctx, "userUUID")
			
			if userUUID == "" {
				return nil, errors.New("not found UserUUID, in grpc metadata")
			}
			return
		}
	}
}
