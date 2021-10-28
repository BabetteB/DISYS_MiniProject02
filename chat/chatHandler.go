package chat

import (
	//"log"


	"golang.org/x/net/context"
)

var (
	clients = make(map[int32] string);
)


type Server struct {
}

func (s *Server) Publish(ctx context.Context, in *ClientMessage) (*StatusMessage, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	response := StatusMessage{ Status: Status_SUCCESS};
	println(in.Msg)
	return &response, nil
}

func (s *Server) BroadCast(ctx context.Context, in *StatusMessage) (*ServerMessage, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &ServerMessage{Msg: "Hello From the Server!"}, nil
}

func (s *Server) JoinedTheChat(ctx context.Context, in *UserInfo) (*JoinedServer, error) {
	// log.Printf("Receive message body from client: %s", in.Body)
	id := len(clients)
	response := JoinedServer{
		Id: int32(id),
	}
	return &response, nil
}

