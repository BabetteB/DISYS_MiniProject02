package main

import (
	"fmt"
	"net"
	"os"
	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
	"google.golang.org/grpc"
)

func main() {
	logger.LogFileInit()

	Output("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		logger.ErrorLogger.Fatalf("FATAL: Connection unable to establish. Failed to listen: %v", err)
	}
	logger.InfoLogger.Println("Connection established through TCP, listening at port 3000", )


	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChittyChatServiceServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.ErrorLogger.Fatalf("FATAL: Server connection failed: %s", err)
		}
	}()
	

	Output("Server started. Press any key to stop")

	var o string
	fmt.Scanln(&o)
	logger.InfoLogger.Println("Exit successfull. Server closing...")

	os.Exit(3)	
}

func Output(input string) {
	fmt.Println(input)
}
