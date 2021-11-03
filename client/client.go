package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"
	"github.com/BabetteB/DISYS_MiniProject02/protos"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//google_protobuf "github.com/golang/protobuf/ptypes/empty"

var (
	checkingStatus bool
	connected      bool
)

type ChatClient struct {
	clientService protos.ChittyChatServiceClient
	conn          *grpc.ClientConn // conn is the client gRPC connection
	id            int32            // id is the client ID used for subscribing
	clientName    string
	lamport       protos.LamportTimestamp
}

type clienthandle struct {
	streamOut protos.ChittyChatService_PublishClient
}

func main() {
	logger.LogFileInit()
	Output(WelcomeMsg())

	rand.Seed(time.Now().UnixNano())

	client, err := makeClient(int32(rand.Intn(1e6)))
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to make Client: %v", err)
	}

	client.EnterUsername()

	streamOut, err := client.clientService.Publish(context.Background())
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to call ChatService :: %v", err)
	}

	// implement communication with gRPC server
	ch1 := clienthandle{
		streamOut: streamOut,
	}

	go client.receiveMessage()
	go ch1.sendMessage(*client)
	go ch1.recvStatus()

	//blocker
	bl := make(chan bool)
	<-bl
}

func (cc *ChatClient) receiveMessage() {
	var err error
	var stream protos.ChittyChatService_BroadcastClient

	for {
		if stream == nil {
			if stream, err = cc.subscribe(); err != nil {
				UserInput() // wait for user to close
				Output("Closing client")
				logger.ErrorLogger.Fatalf("Failed to subscribe: %v", err)
				cc.sleep()
				// Retry on failure SHould do something more
				continue
			}
		}
		response, err := stream.Recv()

		if err != nil {
			logger.WarningLogger.Printf("Failed to receive message: %v", err)
			stream = nil
			cc.sleep()
			continue
		}
		if response.ClientId != cc.id {
			result := protos.RecievingCompareToLamport(&cc.lamport, response.LamportTimestamp)
			msgCode := response.Code
			switch {
			case msgCode == 1:
				Output(fmt.Sprintf("Logical Timestamp:%d, %s joined the server\n", result, response.Username))
			case msgCode == 2:
				// det går galt her
				Output(fmt.Sprintf("Logical Timestamp:%d, %s says: %s \n", result, response.Username, response.Msg))
			case msgCode == 3:
				Output(fmt.Sprintf("Logical Timestamp:%d, %s left the server\n", result, response.Username))
			case msgCode == 4:
				Output(fmt.Sprintf("Logical Timestamp:%d, server closed. Press ctrl + c to exit.\n", result))
			}
		}
	}
}

func (c *ChatClient) subscribe() (protos.ChittyChatService_BroadcastClient, error) {
	logger.InfoLogger.Printf("Subscribing client ID: %d", c.id)
	return c.clientService.Broadcast(context.Background(), &protos.Subscription{
		ClientId:         c.id,
		UserName:         c.clientName,
		LamportTimestamp: c.lamport.Timestamp,
	})
}

func makeClient(idn int32) (*ChatClient, error) {
	conn, err := makeConnection()
	if err != nil {
		return nil, err
	}
	return &ChatClient{
		clientService: protos.NewChittyChatServiceClient(conn),
		conn:          conn,
		id:            idn,
	}, nil
}

func makeConnection() (*grpc.ClientConn, error) {
	logger.InfoLogger.Print("Connecting to server...")
	return grpc.Dial(":3000", []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
}

func (ch *clienthandle) recvStatus() {
	//create a loop

	for !connected {
		mssg, err := ch.streamOut.Recv()
		if err != nil {
			logger.ErrorLogger.Fatalf("Error in receiving message from server :: %v", err)
		}

		if checkingStatus {
			Output(fmt.Sprintf("%s : %s \n", mssg.Operation, mssg.Status))
		}
		connected = true
	}
}

func (ch *clienthandle) sendMessage(client ChatClient) {
	// create a loop
	for {
		clientMessage := UserInput()
		client.lamport.Tick()

		if strings.Contains(clientMessage, "-- quit") {
			Output(fmt.Sprintf("Logical Timestamp:%d, connection to server closed. Press any key to exit.\n", client.lamport.Timestamp))
			clientMessageBox := &protos.ClientMessage{
				ClientId:         client.id,
				UserName:         client.clientName,
				Msg:              "",
				LamportTimestamp: client.lamport.Timestamp,
				Code:             2,
			}

			err := ch.streamOut.Send(clientMessageBox)
			if err != nil {
				logger.WarningLogger.Printf("Error while sending message to server :: %v", err)
			}
			UserInput()
			os.Exit(3)

		} else {

			clientMessageBox := &protos.ClientMessage{
				ClientId:         client.id,
				UserName:         client.clientName,
				Msg:              clientMessage,
				LamportTimestamp: client.lamport.Timestamp,
				Code:             1,
			}

			err := ch.streamOut.Send(clientMessageBox)
			if err != nil {
				logger.WarningLogger.Printf("Error while sending message to server :: %v", err)
			}
		}

	}

}

func (c *ChatClient) sleep() {
	time.Sleep(time.Second * 2)
}

func WelcomeMsg() string {
	return `>>> WELCOME TO CHITTY CHAT <<<
--------------------------------------------------
Please enter an username to begin chatting:
			`
}

func LimitReader(s string) string {
	limit := 128

	reader := strings.NewReader(s)

	buff := make([]byte, limit)

	n, _ := io.ReadAtLeast(reader, buff, limit)

	if n != 0 {
		return string(buff)
	} else {
		return s
	}
}

func (s *ChatClient) EnterUsername() {
	s.clientName = UserInput()
	Welcome(s.clientName)
	s.lamport.Tick()
	logger.InfoLogger.Printf("User registred: %s", s.clientName)
}

func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		logger.ErrorLogger.Printf(" Failed to read from console :: %v", err)
	}
	msg = strings.Trim(msg, "\r\n")

	return LimitReader(msg)
}

func Welcome(input string) {
	Output("Welcome to the chat " + input)
	Output("Type: '-- quit' to exit")
}

func FormatToChat(user, msg string, timestamp int32) string {
	return fmt.Sprintf("%d - %v:  %v", timestamp, user, msg)
}

func Output(input string) {
	fmt.Println(input)
	fmt.Println("-------------------")
}
