package grpc

import (
	"context"
	"errors"
	
	gmetadata "google.golang.org/grpc/metadata"
	
	kmetadata "github.com/go-kratos/kratos/v2/metadata"
)

// grpc 返回 []string, 只获取第一个 String
func GetStringFromGrpc(ctx context.Context,key string) (string, error) {
	if md, ok := gmetadata.FromIncomingContext(ctx); ok {
		values := md.Get(key)
		if len(values) == 0 {
			return "", nil
		} else {
			return values[0], nil
		}
	} else {
		return "", errors.New("ctx metadata not exists")
	}
}

// kratos 只返回 string
func GetStringFromKratos(ctx context.Context, key string) (string, error) {
	if md, ok := kmetadata.FromServerContext(ctx); ok {
		value := md.Get(key)
		return value, nil
	} else {
		return "", errors.New("ctx metadata not exists")
	}
}