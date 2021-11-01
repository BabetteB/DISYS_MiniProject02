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
	ID             		int32  = 0
	user           		string = ""
	checkingStatus 		bool   = true // true: prints statusmessages
	mockClientTimeStamp string = "XXXX/XX/XX XX:XX:XX"
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

	streamOut, err := c.Publish(context.Background())
	if err != nil {
		//log.Fatalf("Failed to call ChatService :: %v", err)
	}
	streamIn, err := c.Broadcast(context.Background())
	if err != nil {
		//log.Fatalf("Failed to call ChatService :: %v", err)
	}

	// implement communication with gRPC server
	ch1 := clienthandle{
		streamOut: streamOut,
	}
	ch2 := serverhandle{
		streamIn: streamIn,
	}

	//ch.clientConfig()
	go ch1.sendMessage()
	go ch1.recvStatus()
	go ch2.receiveMessage()
	go ch2.sendStatus()

	//blocker
	bl := make(chan bool)
	<-bl
}

type clienthandle struct {
	streamOut  chat.ChittyChatService_PublishClient
	clientName string
}

type serverhandle struct {
	streamIn   chat.ChittyChatService_BroadcastClient
	serverName string
}

func (ch *serverhandle) sendStatus() {
	for {
		clientMessageBox := &chat.StatusMessage{
			Operation: "BroadCast()",
			Status:    chat.Status_SUCCESS,
			NewId: &ID,
		}

		err := ch.streamIn.Send(clientMessageBox)

		if err != nil {
			//log.Printf("Error while sending message to server :: %v", err)
		}

	}
}

func (ch *clienthandle) recvStatus() {
	//create a loop
	for {
		mssg, err := ch.streamOut.Recv()
		if err != nil {
			//log.Printf("Error in receiving message from server :: %v", err)
		}

		// Recieving id from server 
		if(ID == 0) {
			ID = *mssg.NewId
		}

		if checkingStatus {
			Output(fmt.Sprintf("%s : %s \n", mssg.Operation, mssg.Status)) 
		}
	}
}

func (ch *clienthandle) sendMessage() {

	// create a loop
	for {

		clientMessage := UserInput()

		clientMessageBox := &chat.ClientMessage{
			ClientId:         ID,
			UserName:         user,
			Msg:              clientMessage,
			LamportTimestamp: 12345,
		}

		err := ch.streamOut.Send(clientMessageBox)

		if err != nil {
			//log.Printf("Error while sending message to server :: %v", err)
		}

	}

}

func (ch *serverhandle) receiveMessage() {

	//create a loop
	for {
		mssg, err := ch.streamIn.Recv()
		if err != nil {
			//log.Printf("Error in receiving message from server :: %v", err)
		}
		senderUniqueCode := mssg.ClientId

		if (senderUniqueCode != ID) {
			//print message to console - - use FormatToChat() when LamportTimestamp is implemented! CHAR
			Output(fmt.Sprintf("%s %s says: %s \n", mockClientTimeStamp, mssg.Username, mssg.Msg))
		}		
	}
}


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
