package main

import (
	log  "github.com/BabetteB/DISYS_MiniProject02/logfile"
	"fmt"
	"net"
	"os"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
	"google.golang.org/grpc"
)

func main() {
	logFileInit()

	Output("Server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.ErrorLogger.Fatalf("FATAL: Connection unable to establish. Failed to listen: %v", err)
	}
	log.InfoLogger("Connection established through TCP, listening at port 3000", )

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChittyChatServiceServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.ErrorLogger.Fatalf("FATAL: Server connection failed: %s", err)
		}
	}()
	

	Output("Server started. Press any key to stop")

	var o string
	fmt.Scanln(&o)
	log.InfoLogger.Println("Exit successfull. Server closing...")
	os.Exit(3)	
}

func Output(input string) {
	fmt.Println(input)
}


