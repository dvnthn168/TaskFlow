package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandler() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		resp, err = handler(ctx, req)
		if err != nil {
			log.Printf("Error occurred in API %s: %v", info.FullMethod, err)
			return nil, handleError(err)
		}

		return resp, nil
	}
}

func handleError(err error) error {
	if st, ok := status.FromError(err); ok {
		return status.New(st.Code(), st.Message()).Err()
	}

	log.Printf("Converting unknown error to Internal: %v", err)
	return status.Error(codes.Internal, "Internal server error")
}
