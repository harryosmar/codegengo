package middlewares

import (
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	HeaderRequestId = "X-Request-Id"
)

func LogMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		reqHeaderKey := strings.ToLower(HeaderRequestId)
		runtime.DefaultHeaderMatcher(reqHeaderKey)
		md, foundMeta := metadata.FromIncomingContext(ctx)
		foundRequestId := false
		if foundMeta {
			_, foundRequestId = md[reqHeaderKey]
		}

		if !foundRequestId {
			md[reqHeaderKey] = []string{
				uuid.New().String(),
			}
		}

		ctx = metadata.NewOutgoingContext(ctx, md)

		return handler(ctx, req)
	}
}
