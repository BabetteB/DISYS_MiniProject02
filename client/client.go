

package main

import (
	log "github.com/BabetteB/DISYS_MiniProject02/logFile"
	"fmt"
	"io"
	"strings"
	"bufio"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"

)

var (
	ID int32
	user string
	closed bool
)

func main() {
	Output(WelcomeMsg())
	EnterUsername()
	Output("Connecting to server...")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.ErrorLogger("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChittyChatServiceClient(conn)

	response, _ := c.Connect(context.Background(), &chat.UserInfo{
		Name: user, });
	ID = *response.NewId
	Output(fmt.Sprintf("You have id #%v", ID))

	go ServerObserver(c)
	
	Output("Connection to server was successful! Ready to chat!")

	go ServerRequester(c)

	for{
		if(closed) {
			break
		}
	}
}

func ServerObserver(c chat.ChittyChatServiceClient) {
	lastMsg := ""
	for {
		response, err := c.Broadcast(context.Background(), new(google_protobuf.Empty))
		if err != nil {
			log.ErrorLogger("Error when calling Broadcast: %s", err)
		}
		chatLog := response.Msg
		if chatLog != "" && chatLog != lastMsg{
			Output(FormatToChat(response.Username, response.Msg, response.Timestamp))
		}
		lastMsg = chatLog
	}
}

func ServerRequester(c chat.ChittyChatServiceClient) {
	for {
		chatMsg := UserInput()
		var currentId int32 = ID
		_, err := c.Publish(context.Background(), &chat.ClientMessage{
			ClientId: currentId,
			UserName: user,
			Msg: chatMsg, 
		})
		if err != nil {
			log.ErrorLogger("Error when calling Publish: %s", err)
		}
		//log.Printf("Response from server: %s", response.Body)
	}
}



func WelcomeMsg() string {
	return `>>> WELCOME TO CHITTY CHAT <<<
--------------------------------------------------
Please enter an username to begin chatting:
Press Ctrl + C to leave!
			`
}

func LimitReader(s string)  string {
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
	log.InfoLogger("User registred: %v", user)
}

func UserInput() (string){
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
		if err != nil {
			log.ErrorLogger(" Failed to read from console :: %v", err)
		}
	msg = strings.Trim(msg, "\r\n")
	
	return LimitReader(msg)
}

func Welcome(input string) {
	Output("Welcome to the chat " + input)
}

func FormatToChat(user, msg string, timestamp string) string{
	return fmt.Sprintf("%v - %v:  %v", timestamp, user, msg)
}

func Output(input string) {
	fmt.Println(input)
}



