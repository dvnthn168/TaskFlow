package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()
		log.Printf("Started call: %s", info.FullMethod)

		resp, err := handler(ctx, req)

		duration := time.Since(start)
		if err != nil {
			log.Printf("Completed call: %s | Error: %v | Duration: %v", info.FullMethod, err, duration)
		} else {
			log.Printf("Completed call: %s | Success | Duration: %v", info.FullMethod, duration)
		}

		return resp, err
	}
}
