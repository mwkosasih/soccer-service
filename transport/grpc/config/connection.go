package config

import (
	"log"
	"net"
	"os"

	server "github.com/mwkosasih/soccer-service/transport/grpc"
	pb "github.com/mwkosasih/soccer-service/transport/grpc/proto/soccer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Connect() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("port_grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port " + os.Getenv("port_grpc"))
	}

	s := grpc.NewServer()
	pb.RegisterSoccerServiceServer(s, server.NewServer())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
