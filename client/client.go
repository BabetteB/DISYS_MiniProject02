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

	Output("Connection to server was successful! Ready to chat!")
	for {
	chatMsg := UserInput()

	response, err := c.SayHello(context.Background(), &chat.Message{Body: chatMsg})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	}

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

func EnterUsername() {
	user := UserInput()
	Welcome(user)
}

func UserInput() (string){
	var input string
	fmt.Scanln(&input)
	return input
}

func Welcome(input string) {
	//more in depth welcome msg 
	//individual welcome msg
	Output("Welcome to the chat " + input)
}

func Output(input string) {
	fmt.Println(input)
}
