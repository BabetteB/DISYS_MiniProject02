package chat

import (
	//"log"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

//const (
//	mockTimestamp = "2021-10-29 00:00:00";
//)

var (
	broadcastMessage = ""
	broadcaster      = ""
	isNewMessage     bool
	clients          = make(map[int32]string)
	//timestamp        int32
)

type Server struct {
	timestamp lamport_timestamp
}

// used in the client as well - maybe move to seperate file? + its functions?
type lamport_timestamp struct {
	id        int32
	timestamp int32
}

func (s *Server) setServer() {
	s.timestamp.id = 0
}

func (s *Server) Publish(ctx context.Context, in *ClientMessage) (*StatusMessage, error) {
	//log.Printf("Receive message body from client: %s", in.Body)
	response := StatusMessage{
		Operation: "Operation: Publish",
		Status:    Status_SUCCESS}
	broadcastMessage = in.Msg
	broadcaster = in.UserName
	isNewMessage = true
	s.timestamp.timestamp += 1
	// println(broadcastMessage) // Test
	return &response, nil
}

func (s *Server) Broadcast(ctx context.Context, in *google_protobuf.Empty) (*ChatRoomMessages, error) {
	// log.Printf("")
	receivers := 0
	s.timestamp.timestamp += 1
	if isNewMessage {
		receivers++
		if receivers == len(clients) {
			isNewMessage = false
			receivers = 0
		}
		return &ChatRoomMessages{
			Msg:              broadcastMessage,
			LamportTimestamp: s.timestamp.timestamp,
			Username:         broadcaster,
		}, nil
	} else {
		return &ChatRoomMessages{
			Msg:              "",
			LamportTimestamp: s.timestamp.timestamp,
			Username:         "",
		}, nil
	}
}

func (s *Server) Connect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("")
	var newId int32 = int32(len(clients) + 1)
	clients[newId] = in.Name
	response := StatusMessage{
		Operation: "Operation: Connect",
		Status:    Status_SUCCESS,
		NewId:     &newId,
	}
	return &response, nil
}

func (s *Server) Disconnect(ctx context.Context, in *UserInfo) (*StatusMessage, error) {
	// log.Printf("")
	response := StatusMessage{
		Operation: "Operation: Disconnect",
		Status:    Status_SUCCESS,
	}
	return &response, nil
}
