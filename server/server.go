package main

import (
	"fmt"
	"log"
	"net"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChittyChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
