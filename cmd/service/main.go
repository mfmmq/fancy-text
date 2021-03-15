package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/mfmmq/fancy-text/pkg/api"

	pb "github.com/mfmmq/fancy-text/pkg/proto"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
			return value
	}
	return fallback
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFancyTextGeneratorServer(grpcServer, &api.Server{})
	fmt.Printf("Grpc server now listening at localhost:8080" + "\n")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
}