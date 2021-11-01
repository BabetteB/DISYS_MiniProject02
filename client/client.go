package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	logger "github.com/BabetteB/DISYS_MiniProject02/logFile"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	chat "github.com/BabetteB/DISYS_MiniProject02/chat"
)

//google_protobuf "github.com/golang/protobuf/ptypes/empty"

var (
	ID     int32
	user   string
	closed bool
)

func main() {
	Output(WelcomeMsg())
	EnterUsername()
	Output("Connecting to server...")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		logger.ErrorLogger.Printf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChittyChatServiceClient(conn)

	/* 	response, _ := c.Connect(context.Background(), &chat.UserInfo{
	   		Name: user})
	   	ID = *response.NewId
	   	Output(fmt.Sprintf("You have id #%v", ID)) */

	stream, err := c.Publish(context.Background())
	if err != nil {
		//log.Fatalf("Failed to call ChatService :: %v", err)
	}

	// implement communication with gRPC server
	ch := clienthandle{stream: stream}
	//ch.clientConfig()
	go ch.sendMessage()
	//go ch.receiveMessage()

	//blocker
	bl := make(chan bool)
	<-bl
}

type clienthandle struct {
	stream     chat.ChittyChatService_PublishClient
	clientName string
}

func (ch *clienthandle) sendMessage() {

	// create a loop
	for {

		clientMessage := UserInput()

		clientMessageBox := &chat.ClientMessage{
			ClientId:         123,
			UserName:         "Babse",
			Msg:              clientMessage,
			LamportTimestamp: 1234,
		}

		err := ch.stream.Send(clientMessageBox)

		if err != nil {
			//log.Printf("Error while sending message to server :: %v", err)
		}

	}

}

/* func (ch *clienthandle) receiveMessage() {

	//create a loop
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			//log.Printf("Error in receiving message from server :: %v", err)
		}

		//print message to console
		fmt.Printf("%s : %s \n",mssg.Name,mssg.Body)

	}
} */

/* func ServerObserver(c chat.ChittyChatServiceClient) {
	lastMsg := ""
	for {
		response, err := c.Broadcast(context.Background(), new(google_protobuf.Empty))
		if err != nil {
			logger.ErrorLogger.Printf("Error when calling Broadcast: %s", err)
		}
		chatLog := response.Msg
		if chatLog != "" && chatLog != lastMsg {
			Output(FormatToChat(response.Username, response.Msg, response.LamportTimestamp))
		}
		lastMsg = chatLog
	}
} */

/* func ServerRequester(c chat.ChittyChatServiceClient) {
	for {
		chatMsg := UserInput()
		var currentId int32 = ID
		_, err := c.Publish(context.Background(), &chat.ClientMessage{
			ClientId: currentId,
			UserName: user,
			Msg:      chatMsg,
		})
		if err != nil {
			logger.ErrorLogger.Printf("Error when calling Publish: %s", err)
		}
		//log.Printf("Response from server: %s", response.Body)
	}
} */

func WelcomeMsg() string {
	return `>>> WELCOME TO CHITTY CHAT <<<
--------------------------------------------------
Please enter an username to begin chatting:
Press Ctrl + C to leave!
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

func EnterUsername() {
	user = UserInput()
	Welcome(user)
	//logger.InfoLogger.Printf("User registred: %v", user) /// BAAAAARBETSE:P
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
}

func FormatToChat(user, msg string, timestamp int32) string {
	return fmt.Sprintf("%d - %v:  %v", timestamp, user, msg)
}

func Output(input string) {
	fmt.Println(input)
}
