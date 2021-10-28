package main
/*
import (
	"fmt"
	"io"
	"strings"
)

func main()  {
	fmt.Println(
		`>>> WELCOME TO CHITTY CHAT <<<
--------------------------------------------------
Please enter an username to begin chatting:
		`)
	var input string
	fmt.Scanln(&input)
	welcome(input)

	for{
		publish()
	}

}

func welcome(input string) {
	//more in depth welcome msg 
	//individual welcome msg
	fmt.Println("Welcome to the chat " + input)
}

func publish() {
	//Find way to send to server
	//Find way to limit 
	var msg string
	fmt.Scan(&msg)

	fmt.Println(limitReader(msg))
}


func leave() {
// client wants to leave the server - promts with a confirmation
}

func limitReader(s string) string {
	limit:= 128

	reader:= strings.NewReader(s)

	buff := make([]byte, limit)

	n, _ := io.ReadAtLeast(reader, buff, limit)

	if n !=0 {
		return string(buff)
	} else {
		return s
	}
}

type msgInfo struct {
	//msg string - timestamp - int (msg type)
}
 */