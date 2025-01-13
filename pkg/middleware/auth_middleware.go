package middleware

import (
	"context"
	"log"
	"strings"

	"github.com/dvnthn168/TaskFlow/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
	payloadKey          = contextKey("payload")
)

func AuthInterceptor(tokenMaker token.Maker) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		if isUnauthenticatedAPI(ctx) {
			log.Printf("API không yêu cầu xác thực: %s", info.FullMethod)
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		values := md.Get(authorizationHeader)
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		authHeader := values[0]
		fields := strings.Fields(authHeader)
		if len(authHeader) != 2 {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header format")
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationBearer {
			return nil, status.Errorf(codes.Unauthenticated, "unsupported authorization type: %v", authType)
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		const contextKey = "payload"
		ctx = context.WithValue(ctx, contextKey, payload)

		return handler(ctx, req)
	}
}

func isUnauthenticatedAPI(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	if val, exists := md["X-Unauthenticated"]; exists && val[0] == "true" {
		return true
	}

	return false
}
