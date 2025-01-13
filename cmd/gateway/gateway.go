package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dvnthn168/TaskFlow/gen/userpb"
	"github.com/dvnthn168/TaskFlow/pkg/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	mux := runtime.NewServeMux()

	err = userpb.RegisterUsersHandlerFromEndpoint(
		context.Background(),
		mux,
		config.HTTPServerAddr,
		[]grpc.Dial{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to register gRPC-Gateway handler: %v", err)
	}

	log.Printf("Gateway is running on %s", config.GateWayServerAddr)
	if err := http.ListenAndServe(config.GateWayServerAddr, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}