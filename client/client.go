package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/BabetteB/DISYS_MiniProject02/chat"
)

var (
	ID *int32
	User string
)

func main() {
	Output(WelcomeMsg())
	EnterUsername()
	Output("Connecting to server...")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChittyChatServiceClient(conn)

	response, _ := c.Connect(context.Background(), &chat.UserInfo{
		Name: User, });
	ID = response.NewId
	Output(fmt.Sprintf("You have id %v", ID))

	
	Output("Connection to server was successful! Ready to chat!")

	for {
	chatMsg := UserInput()
	var currentId int32 = *ID
	_, err := c.Publish(context.Background(), &chat.ClientMessage{
		ClientId: currentId,
		Msg: chatMsg, 
	})
	if err != nil {
		log.Fatalf("Error when calling Publish: %s", err)
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
	User = UserInput()
	Welcome(User)
}

func UserInput() (string){
	var input string
	fmt.Scanln(&input)
	return LimitReader(input)
}

func Welcome(input string) {
	//more in depth welcome msg 
	//individual welcome msg
	Output("Welcome to the chat " + input)
}

func Output(input string) {
	fmt.Println(input)
}
