package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
	"google.golang.org/grpc"
)

func main() {

	Output("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChittyChatServiceServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
	

	Output("Server started. Press any key to stop")

	var o string
	fmt.Scanln(&o)
	// Log exit
	os.Exit(3)	
}

func Output(input string) {
	fmt.Println(input)
}


