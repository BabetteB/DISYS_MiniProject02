package chat

import (
	//"log"


	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

const (
	mockTimestamp = 20000000;
) 

var (
	//clients = make(map[int32] string);
)


type Server struct {
}

func (s *Server) Publish(ctx context.Context, in *ClientMessage) (*StatusMessage, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	response := StatusMessage{ Status: Status_SUCCESS};
	println(in.Msg) // Test
	return &response, nil
}

func (s *Server) Broadcast(ctx context.Context, in *google_protobuf.Empty) (*ChatRoomMessages, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	return &ChatRoomMessages{
		Msg: []string {"Hello From the Server!"},
		Timestamp: []int32 {mockTimestamp},
		}, nil
}

func (s *Server) Connect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("Receive message body from client: %s", in.Body)
	var newId int32 = 143234
	response := StatusMessage{
		Operation: "Operation: Connect",
		Status: Status_SUCCESS,
		NewId: &newId,
	}
	return &response, nil
}

func (s *Server) Disconnect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("Receive message body from client: %s", in.Body)
	response := StatusMessage{
		Operation: "Operation: Disconnect",
		Status: Status_SUCCESS,
	}
	return &response, nil
}

